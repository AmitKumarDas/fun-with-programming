package shellx_carvel

var appValuesYML = `
#!
#! Note: Do not EDIT. This file is GENERATED
#!
#@data/values-schema
---
#@schema/desc "Port for metrics"
metrics_port: 8080
#@schema/desc "Port for health"
health_port: 9440
`

var appDeploymentYML = `
#!
#! Note: Do not EDIT. This file is GENERATED
#!
#@ load("@ytt:data", "data")

---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: ${K8S_NAMESPACE}
  name: ${APP_DEPLOYMENT_NAME}
spec:
  selector:
    matchLabels:
      ${APP_DEPLOYMENT_LABEL_KEY}: ${APP_DEPLOYMENT_LABEL_VAL}
  template:
    metadata:
      labels:
        ${APP_DEPLOYMENT_LABEL_KEY}: ${APP_DEPLOYMENT_LABEL_VAL}
    annotations:
      prometheus.io/scrape: "true"
      prometheus.io/port: "8080"
    spec:
      serviceAccountName: tkg-remediator
      terminationGracePeriodSeconds: 60
      #! Required for AWS IAM Role bindings
      #! https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts-technical-overview.html
      securityContext:
        fsGroup: 1337
      containers:
        - name: manager
          image: ${APP_IMAGE_NAME}:${APP_IMAGE_VERSION}
          imagePullPolicy: IfNotPresent
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsNonRoot: true
            capabilities:
              drop: [ "ALL" ]
            seccompProfile:
              type: RuntimeDefault
          ports:
            - containerPort: 8080
              name: http-prom
              protocol: TCP
            - containerPort: #@ data.values.health_port
              name: healthz
              protocol: TCP
          env:
            - name: RUNTIME_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          args:
            - --log-level=info
            - --log-encoding=json
            - --enable-leader-election
            - --watch-all-namespaces
          readinessProbe:
            httpGet:
              path: /readyz
              port: healthz
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
          resources:
            limits:
              cpu: 1000m
              memory: 1Gi
            requests:
              cpu: 100m
              memory: 64Mi
          volumeMounts:
            - name: temp
              mountPath: /tmp
      volumes:
        - name: temp
          emptyDir: {}
`

var packageMetadataYML = `
#!
#! Note: Do not EDIT. This file is GENERATED
#!
apiVersion: data.packaging.carvel.dev/v1alpha1
kind: PackageMetadata
metadata:
  #! This will be the name of our package
  name: ${PACKAGE_NAME}
spec:
  displayName: "K8s remediator"
  longDescription: "Package consists of a K8s controller as a K8s deployment "
  shortDescription: ""
  categories:
  - release
`

var packageTemplateYML = `
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
        Initial release of the tkg-remediator package
  valuesSchema: #! configurable properties that exist for the version
    openAPIv3: #@ yaml.decode(data.values.openapi)["components"]["schemas"]["dataValues"]
  template:
    spec:
      fetch:
      - imgpkgBundle: #! fetch workload imgpkg bundle
          image: #@ "${REGISTRY_NAME}:${REGISTRY_PORT}/packages/${APP_BUNDLE_NAME}:" + data.values.version
      template:
      - ytt: #! run the templates through ytt
          paths:
          - "config/"
      - kbld: #! kbld transformations
          paths:
          - ".imgpkg/images.yml"
          - "-"
      deploy:
      - kapp: {} #! deploy the resulting manifests through kapp
`

var kindClusterLocalRegistryYML = `
#!
#! Note: Do Not EDIT. This file is GENERATED
#!
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${REGISTRY_PORT}"]
    endpoint = ["http://${REGISTRY_NAME}:${REGISTRY_PORT}"]
`

var kindConfigLocalRegistryHostingYML = `
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
    host: "localhost:${REGISTRY_PORT}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
`

var etcHostsUpdateMsg = `
# ---------------------------
# Ensure below text is appended to /etc/hosts
# ---------------------------

# Added to set up kind & carvel
# This forces imgpkg to use http instead of https
127.0.0.1 ${REGISTRY_NAME}
# End of section
`

var packageRepositoryYML = `
#!
#! Note: Do not EDIT. This file is GENERATED
#!
apiVersion: packaging.carvel.dev/v1alpha1
kind: PackageRepository
metadata:
  name: ${PACKAGE_REPO_NAME}
  namespace: ${K8S_NAMESPACE}
spec:
  fetch:
    imgpkgBundle:
      image: ${REGISTRY_NAME}:${REGISTRY_PORT}/packages/${PACKAGE_REPO_NAME}:${PACKAGE_REPO_VERSION}
`
