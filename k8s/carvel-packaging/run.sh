#!/bin/sh
set -o errexit

# -----------------------
# Learn Carvel Packaging
# -----------------------

# This helps understand the inner workings of carvel kapp-controller

# This file remains the **Single Source of Truth**. All other folders & files found
# in the parent folder are generated while running this file. Note that all files
# should be git versioned even if some of them are generated.

# References:
# - https://github.com/vmware-tanzu/carvel-kapp-controller/tree/develop/examples/local-kind-environment
# - https://carvel.dev/kapp-controller/docs/v0.40.0/packaging-tutorial/

echo "++ this script works only for M1 / ARM Macs \_(^^)_/"

echo "++ ------------------------------"
echo "++ configuration"
echo "++ ------------------------------"

echo "++ configuring local image registry"
REG_NAME="kind-registry.local"
REG_PORT="5000"

echo "++ configuring kubernetes rbac"
K8S_NAMESPACE="default"
K8S_SERVICE_ACCOUNT="default-ns-sa"
K8S_ROLE="default-ns-role"
K8S_ROLE_BINDING="default-ns-role-binding"

echo "++ configuring carvel versions"
KAPP_CTRL_VERSION="v0.40.0"

echo "++ configuring location for carvel binaries"
LOCAL_BIN="local-bin"
KBLD="${LOCAL_BIN}/kbld"
IMGPKG="${LOCAL_BIN}/imgpkg"
YTT="${LOCAL_BIN}/ytt"
KAPP="${LOCAL_BIN}/kapp"

echo "++ configuring location for kind binary"
KIND="${LOCAL_BIN}/kind"

echo "++ configuring workload bundle"
WORKLOAD_BUNDLE_NAME="simple-app"
WORKLOAD_BUNDLE_VERSION="1.0.0"

echo "++ configuring package"
PACKAGE_NAME="simple-app.corp.com"
PACKAGE_VERSION="1.0.0"

echo "++ configuring package repository"
PACKAGE_REPOSITORY_NAME="pkg-repo"
PACKAGE_REPOSITORY_VERSION="1.0.0"

echo "++ configuring package install"
PACKAGE_INSTALL_NAME="pkg-demo"

echo "++ configuring folder names"
FOLDER_KIND="kind"
FOLDER_PKG_SOURCE="pkg-source"
FOLDER_PKG_SOURCE_CONFIG="config"
FOLDER_PKG_RELEASE="pkg-release"
FOLDER_PKG_K8S_DEPLOY="pkg-deploy"

echo "++ configuring file names"
FILE_RESULTING_PACKAGE_TEMPLATE="package-template.yml"
FILE_RESULTING_PACKAGE_VALUES="schema-openapi.yml"

echo "configuring workload names"
WORKLOAD_DEPLOYMENT_NAME="simple-app"
WORKLOAD_SERVICE_NAME="simple-app"

echo "++ ------------------------------"
echo "++ prerequisites"
echo "++ ------------------------------"

echo "++ installing carvel tools locally"
mkdir -p ${LOCAL_BIN}/
ls ${KBLD} && ls ${IMGPKG} && ls ${YTT}
[ $? == '0' ] || curl -L https://carvel.dev/install.sh | K14SIO_INSTALL_BIN_DIR=${LOCAL_BIN} bash

echo "++ installing kind locally"
IS_KIND_INSTALLED=$(ls ${KIND} || echo "no")
[ ${IS_KIND_INSTALLED} == 'no' ] && curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-darwin-arm64
[ ${IS_KIND_INSTALLED} == 'no' ] && chmod +x ./kind
[ ${IS_KIND_INSTALLED} == 'no' ] && mv ./kind ${KIND}

echo "++ list of locally installed binaries"
ls -ltra ${LOCAL_BIN}

echo "++ installing jq (system wide)"
which jq || brew install jq

echo "++ ------------------------------"
echo "++ bundle workload artifacts as an OCI image"
echo "++ ------------------------------"

echo "++ creating folder '${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}'"
mkdir -p ${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}

echo "++ creating workload template at '${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/config.yml'"
cat > ${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/config.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
#@ load("@ytt:data", "data")

#@ def labels():
simple-app: ""
#@ end

---
apiVersion: v1
kind: Service
metadata:
  namespace: default
  name: ${WORKLOAD_SERVICE_NAME}
spec:
  ports:
  - port: #@ data.values.svc_port
    targetPort: #@ data.values.app_port
  selector: #@ labels()
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: default
  name: ${WORKLOAD_DEPLOYMENT_NAME}
spec:
  selector:
    matchLabels: #@ labels()
  template:
    metadata:
      labels: #@ labels()
    spec:
      containers:
      - name: simple-app
        image: docker.io/dkalinin/k8s-simple-app@sha256:4c8b96d4fffdfae29258d94a22ae4ad1fe36139d47288b8960d9958d1e63a9d0
        env:
        - name: HELLO_MSG
          value: #@ data.values.hello_msg
EOF

echo "++ creating workload values at '${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/values.yml'"
cat > ${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/values.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
#@data/values-schema
---
#@schema/desc "Port number for the service."
svc_port: 80
#@schema/desc "Target port for the application."
app_port: 80
#@schema/desc "Name used in hello message from app when app is pinged."
hello_msg: stranger
EOF

echo "++ creating folder for workload bundle at '${FOLDER_PKG_SOURCE}/.imgpkg'"
mkdir -p ${FOLDER_PKG_SOURCE}/.imgpkg

echo "++ using ${KBLD} to populate workload bundle"
${KBLD} -f ${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/ --imgpkg-lock-output ${FOLDER_PKG_SOURCE}/.imgpkg/images.yml

echo "++ using ${IMGPKG} to publish workload bundle ${WORKLOAD_BUNDLE_NAME}:${WORKLOAD_BUNDLE_VERSION} to registry "
${IMGPKG} push -b ${REG_NAME}:${REG_PORT}/packages/${WORKLOAD_BUNDLE_NAME}:${WORKLOAD_BUNDLE_VERSION} -f ${FOLDER_PKG_SOURCE}/

echo "++ verify docker registry catalog for presence of workload bundle"
curl ${REG_NAME}:${REG_PORT}/v2/_catalog

echo "++ ------------------------------"
echo "++ make a release"
echo "++ ------------------------------"

echo "++ creating folders for releasing package"
mkdir -p ${FOLDER_PKG_RELEASE}/.imgpkg ${FOLDER_PKG_RELEASE}/packages/${PACKAGE_NAME}

echo "++ creating 'PackageMetadata' custom resource"
cat > ${FOLDER_PKG_RELEASE}/packages/${PACKAGE_NAME}/package-metadata.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
apiVersion: data.packaging.carvel.dev/v1alpha1
kind: PackageMetadata
metadata:
  #! This will be the name of our package
  name: ${PACKAGE_NAME}
spec:
  displayName: "Simple App"
  longDescription: "A package consisting of a k8s deployment and service"
  shortDescription: "Packaging for demoing"
  categories:
  - demo
EOF

echo "++ using ${YTT} to generate OpenAPI schema from '${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/values.yml'"
cat > ${FOLDER_PKG_RELEASE}/${FILE_RESULTING_PACKAGE_VALUES} << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
EOF
${YTT} -f ${FOLDER_PKG_SOURCE}/${FOLDER_PKG_SOURCE_CONFIG}/values.yml \
  --data-values-schema-inspect -o openapi-v3 >> ${FOLDER_PKG_RELEASE}/${FILE_RESULTING_PACKAGE_VALUES}

echo "++ creating 'Package' custom resource with generated OpenAPI schema for values"
cat > ${FOLDER_PKG_RELEASE}/${FILE_RESULTING_PACKAGE_TEMPLATE} << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
#@ load("@ytt:data", "data")  #! read data values (generated via ytt's data-values-schema-inspect mode)
#@ load("@ytt:yaml", "yaml")  #! dynamically decode the output of ytt's data-values-schema-inspect
---
apiVersion: data.packaging.carvel.dev/v1alpha1
kind: Package
metadata:
  name: #@ "${PACKAGE_NAME}." + data.values.version
spec:
  refName: ${PACKAGE_NAME}
  version: #@ data.values.version
  releaseNotes: |
        Initial release of the simple app package
  valuesSchema: #! configurable properties that exist for the version
    openAPIv3: #@ yaml.decode(data.values.openapi)["components"]["schemas"]["dataValues"]
  template:
    spec:
      fetch:
      - imgpkgBundle: #! fetch workload imgpkg bundle
          image: #@ "${REG_NAME}:${REG_PORT}/packages/${WORKLOAD_BUNDLE_NAME}:" + data.values.version
      template:
      - ytt: #! run the templates through ytt
          paths:
          - "${FOLDER_PKG_SOURCE_CONFIG}/"
      - kbld: #! kbld transformations
          paths:
          - ".imgpkg/images.yml"
          - "-"
      deploy:
      - kapp: {} #! deploy the resulting manifests through kapp
EOF

echo "++ using ${YTT} to create versioned package artifact ${PACKAGE_NAME}:${PACKAGE_VERSION}"
${YTT} -f ${FOLDER_PKG_RELEASE}/${FILE_RESULTING_PACKAGE_TEMPLATE}  \
  --data-value-file openapi=${FOLDER_PKG_RELEASE}/${FILE_RESULTING_PACKAGE_VALUES} \
  -v version="${PACKAGE_VERSION}" > ${FOLDER_PKG_RELEASE}/packages/${PACKAGE_NAME}/${PACKAGE_VERSION}.yml

echo "++ using ${KBLD} to create package repository bundle"
${KBLD} -f ${FOLDER_PKG_RELEASE}/packages/ --imgpkg-lock-output ${FOLDER_PKG_RELEASE}/.imgpkg/images.yml

echo "++ using ${IMGPKG} to publish package repository bundle '${PACKAGE_REPOSITORY_NAME}:${PACKAGE_REPOSITORY_VERSION}'"
${IMGPKG} push -b ${REG_NAME}:${REG_PORT}/packages/${PACKAGE_REPOSITORY_NAME}:${PACKAGE_REPOSITORY_VERSION} -f ${FOLDER_PKG_RELEASE}

echo "++ verifying registry catalog for presence of package repository bundle"
curl ${REG_NAME}:${REG_PORT}/v2/_catalog

echo "++ ------------------------------"
echo "++ set up kind cluster"
echo "++ ------------------------------"

echo "++ [manual] turn on the docker runtime"

echo "++ creating registry container for kind"
if [ "$(docker inspect -f '{{.State.Running}}' "${REG_NAME}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${REG_PORT}:5000" --name "${REG_NAME}" \
    registry:2
fi

echo "++ creating folder for kind artifacts"
mkdir -p ${FOLDER_KIND}

echo "++ creating config for kind cluster with the local registry"
cat > ${FOLDER_KIND}/kind-cluster.yml << EOF
#!
#! Note: Do Not EDIT. This file is GENERATED
#!
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${REG_PORT}"]
    endpoint = ["http://${REG_NAME}:${REG_PORT}"]
EOF

kubectl cluster-info --context kind-kind
[ $? == '0' ] || ${KIND} create cluster --config=${FOLDER_KIND}/kind-cluster.yml

echo "++ connecting registry to kind cluster network"
if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${REG_NAME}")" = 'null' ]; then
  docker network connect "kind" "${REG_NAME}"
fi

echo "++ documenting local registry as a config map"
cat > ${FOLDER_KIND}/local-registry-hosting.yml << EOF
#!
#! Note: Do Not EDIT. This file is GENERATED
#!
#! https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${REG_PORT}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF

echo "++ applying local registry config map"
kubectl apply -f ${FOLDER_KIND}/local-registry-hosting.yml

echo "++ creating etchosts.txt"
cat > ${FOLDER_KIND}/etchosts.txt << EOF
# Added while setting up kind & carvel
# To force imgpkg to use http instead of https
127.0.0.1 ${REG_NAME}
# End of section
EOF

echo "++ [manual] append /etc/hosts with contents from ${FOLDER_KIND}/etchosts.txt \_(^^)_/"

echo "++ deploying carvel kapp controller ${KAPP_CTRL_VERSION}"
kubectl apply -f \
  https://github.com/vmware-tanzu/carvel-kapp-controller/releases/download/${KAPP_CTRL_VERSION}/release.yml

echo "++ verifying kapp controller installations"
kubectl get all -n kapp-controller
kubectl api-resources --api-group packaging.carvel.dev
kubectl api-resources --api-group data.packaging.carvel.dev
kubectl api-resources --api-group kappctrl.k14s.io

echo "++ ------------------------------"
echo "++ test the release by deploying it in the cluster"
echo "++ ------------------------------"

echo "++ creating folder for k8s artifacts"
mkdir -p ${FOLDER_PKG_K8S_DEPLOY}

echo "++ creating PackageRepository custom resource"
cat > ${FOLDER_PKG_K8S_DEPLOY}/package-repo.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
apiVersion: packaging.carvel.dev/v1alpha1
kind: PackageRepository
metadata:
  name: ${PACKAGE_REPOSITORY_NAME}
  namespace: ${K8S_NAMESPACE}
spec:
  fetch:
    imgpkgBundle:
      image: ${REG_NAME}:${REG_PORT}/packages/${PACKAGE_REPOSITORY_NAME}:${PACKAGE_REPOSITORY_VERSION}
EOF

echo "++ applying PackageRepository CR to the cluster"
kubectl delete -f ${FOLDER_PKG_K8S_DEPLOY}/package-repo.yml || true
kubectl apply -f ${FOLDER_PKG_K8S_DEPLOY}/package-repo.yml

echo "++ waiting for availability of package repository"
sleep 10
kubectl get packagerepository -n ${K8S_NAMESPACE}

echo "++ listing PackageMetadata"
kubectl get packagemetadatas -n ${K8S_NAMESPACE}

echo "++ listing Package"
kubectl get packages -n ${K8S_NAMESPACE} --field-selector spec.refName=${PACKAGE_NAME}

echo "++ creating rbac for workload"
cat > ${FOLDER_PKG_K8S_DEPLOY}/workload-rbac.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: ${K8S_SERVICE_ACCOUNT}
  namespace: ${K8S_NAMESPACE}
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ${K8S_ROLE}
  namespace: ${K8S_NAMESPACE}
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ${K8S_ROLE_BINDING}
  namespace: ${K8S_NAMESPACE}
subjects:
- kind: ServiceAccount
  name: ${K8S_SERVICE_ACCOUNT}
  namespace: ${K8S_NAMESPACE}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: ${K8S_ROLE}
EOF

echo "++ applying workload rbac to the cluster"
kubectl delete -f ${FOLDER_PKG_K8S_DEPLOY}/workload-rbac.yml || true
kubectl apply -f ${FOLDER_PKG_K8S_DEPLOY}/workload-rbac.yml

echo "++ creating PackageInstall CR with a secret resource (to provide customized values)"
cat > ${FOLDER_PKG_K8S_DEPLOY}/package-install.yml << EOF
#!
#! Note: Do not EDIT. This file is GENERATED
#!
apiVersion: packaging.carvel.dev/v1alpha1
kind: PackageInstall
metadata:
  name: ${PACKAGE_INSTALL_NAME}
  namespace: ${K8S_NAMESPACE}
spec:
  serviceAccountName: ${K8S_SERVICE_ACCOUNT}
  packageRef:
    refName: ${PACKAGE_NAME}
    versionSelection: #! provides granular control over version selection
      constraints: ${PACKAGE_VERSION}
  values:
  - secretRef:
      name: pkg-install-secret
---
apiVersion: v1
kind: Secret
metadata:
  name: pkg-install-secret
  namespace: ${K8S_NAMESPACE}
stringData:
  #! This provides customized values to package installation template
  #! Users can discover more details on the configurable properties of a
  #! package by inspecting the Package CR’s valuesSchema
  values.yml: |
    ---
    hello_msg: "to all my internet friends from carvel packaging demo"
EOF

echo "++ deploying PackageInstall CR"
kubectl delete -f ${FOLDER_PKG_K8S_DEPLOY}/package-install.yml || true
kubectl delete app ${PACKAGE_INSTALL_NAME} -n ${K8S_NAMESPACE} || true
kubectl delete deployment ${WORKLOAD_DEPLOYMENT_NAME} -n ${K8S_NAMESPACE} || true
kubectl delete service ${WORKLOAD_SERVICE_NAME} -n ${K8S_NAMESPACE} || true
kubectl apply -f ${FOLDER_PKG_K8S_DEPLOY}/package-install.yml

echo "++ verify if workload is running"
sleep 10
kubectl get app -n ${K8S_NAMESPACE}
kubectl get service -n ${K8S_NAMESPACE}
kubectl get pod -n ${K8S_NAMESPACE}
