## Motivation
A journal that records the issues I faced while trying Carvel tools with KIND & local
docker registry. This log book gives enough hints & ideas for someone to understand
& get comfortable to contribute to shellx-carvel codebase.

## Journal
### Start 1/
- **GIVEN**: `/etc/hosts`
```shell
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost
# Added by Docker Desktop
# To allow the same kube context to work on the host and the container:
127.0.0.1 kubernetes.docker.internal
```

### Attempt 2/
- **WHEN**: `make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (83.65s)
    --- FAIL: TestCarvelReleaseE2E/publishing_the_app_bundle (4.98s)
        testutil_test.go:27: expected no err got: running "tmp/imgpkg push -b kind-registry.local:5000/packages/k8s-remediator-app:1.0.1 -f artifacts/packaging/source" failed with exit code 1 : imgpkg: Error: Writing 'kind-registry.local:5000/packages/k8s-remediator-app:1.0.1':
              Error while preparing a transport to talk with the registry:
                Unable to create round tripper:
                  Get "https://kind-registry.local:5000/v2/": dial tcp: lookup kind-registry.local on 10.112.16.144:53: no such host; Get "http://kind-registry.local:5000/v2/": dial tcp: lookup kind-registry.local on 10.112.16.144:53: no such host : 
    --- FAIL: TestCarvelReleaseE2E/publishing_the_package (0.79s)
        testutil_test.go:27: expected no err got: running "tmp/kbld -f artifacts/packaging/release/packages --imgpkg-lock-output artifacts/packaging/release/.imgpkg/images.yml" failed with exit code 1 : kbld: Error: 
            - Resolving image 'kind-registry.local:5000/packages/k8s-remediator-app:1.0.1': Get "https://kind-registry.local:5000/v2/": dial tcp: lookup kind-registry.local on 10.112.16.144:53: no such host; Get "http://kind-registry.local:5000/v2/": dial tcp: lookup kind-registry.local on 10.112.16.144:53: no such host : 
    --- FAIL: TestCarvelReleaseE2E/verifying_app_deployment (56.40s)
        testutil_test.go:27: expected no err got: running "kubectl get package -n k8s-remediator-system k8s-remediator.experiment.dev.com.1.0.1" failed with exit code 1 : error: the server doesn't have a resource type "package" : : func timed out after 47.010980958s
```

### Attempt 3/
- **IDEA**: Use `kubernetes.docker.internal` as the registry name since it is set in `/etc/hosts`
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.internal make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (52.92s)
    --- FAIL: TestCarvelReleaseE2E/setting_up_local_docker_container_as_an_OCI_registry (0.32s)
        testutil_test.go:27: expected no err got: running "docker run -d --restart always -p 127.0.0.1:5000:5000 --name kubernetes.docker.internal registry:2" failed with exit code 125 : docker: Error response from daemon: driver failed programming external connectivity on endpoint kubernetes.docker.internal (27a91978c3d15f609375b8e5c40528d4309f23b84d5205a022356369cafcb604): Bind for 127.0.0.1:5000 failed: port is already allocated. : 7be7d003c82b9e4c8855ba16b78ef07b80d0f894db27d6ee14528bb664e8a37d
    --- FAIL: TestCarvelReleaseE2E/publishing_the_app_bundle (4.17s)
        testutil_test.go:27: expected no err got: running "tmp/imgpkg push -b kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1 -f artifacts/packaging/source" failed with exit code 1 : imgpkg: Error: Writing 'kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1':
              Error while preparing a transport to talk with the registry:
                Unable to create round tripper:
                  Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client : 
    --- FAIL: TestCarvelReleaseE2E/publishing_the_package (0.19s)
        testutil_test.go:27: expected no err got: running "tmp/kbld -f artifacts/packaging/release/packages --imgpkg-lock-output artifacts/packaging/release/.imgpkg/images.yml" failed with exit code 1 : kbld: Error: 
            - Resolving image 'kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1': Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client : 
    --- FAIL: TestCarvelReleaseE2E/verifying_app_deployment (47.34s)
        testutil_test.go:27: expected no err got: running "kubectl get package -n k8s-remediator-system k8s-remediator.experiment.dev.com.1.0.1" failed with exit code 1 : Error from server (NotFound): internalpackages.internal.packaging.carvel.dev "dcs76bbiclmmap39c5q6ushecls70pbid5mmarjk5pi6athecdnmqbhh5oo2sc8" not found : : func timed out after 31.087228208s
```
- **CAUSE**: 127.0.0.1:5000 failed: port is already allocated

### Attempt 4/
- **WHY**: If a test fails then subsequent tests should not run & clutter the whole output
- **CHANGE**: Remove subtests to exit test in case of failure
```diff
 func TestCarvelReleaseE2E(t *testing.T) {
-       t.Run("installing carvel CLIs", func(t *testing.T) {
-               tryInstallCarvelCLIs(t)
-       })
-       t.Run("installing KIND cli", func(t *testing.T) {
-               tryInstallKindCLI(t)
-       })
-       t.Run("setting up local docker container as an OCI registry", func(t *testing.T) {
-               trySetupRegistryAsLocalDockerContainer(t)
-       })
-       t.Run("publishing the app bundle", func(t *testing.T) {
-               tryCutAppBundle(t)
-       })
-       t.Run("publishing the package", func(t *testing.T) {
-               tryCutPackageRelease(t)
-       })
-       t.Run("setting up KIND cluster", func(t *testing.T) {
-               trySetupKindCluster(t)
-       })
-       t.Run("verifying app deployment", func(t *testing.T) {
-               tryVerifyApplication(t)
-       })
+       tryInstallCarvelCLIs(t)
+       tryInstallKindCLI(t)
+       trySetupRegistryAsLocalDockerContainer(t)
+       tryCutAppBundle(t)
+       tryCutPackageRelease(t)
+       trySetupKindCluster(t)
+       tryVerifyApplication(t)
 }
```
- **MANUAL**: Delete docker containers that serve as image registry & KIND cluster
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.internal make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (6.36s)
    testutil_test.go:27: expected no err got: running "tmp/imgpkg push -b kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1 -f artifacts/packaging/source" failed with exit code 1 : imgpkg: Error: Writing 'kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1':
          Error while preparing a transport to talk with the registry:
            Unable to create round tripper:
              Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client : 
FAIL
```
- **CAUSE**: Error says `http: server gave HTTP response to HTTPS client`
- **CAUSE**: imgpkg is the CLI that fails

### Attempt 5/
- **NOTE**: Automate to always start a test run on a clean test setup 
- **CHANGE**: Ensure test runs are idiomatic by first deleting & then creating local registry
```diff
func setupRegistryAsLocalDockerContainer() error {
        if shx.IsNotEq(EnvSetupLocalRegistry, "true") {
                return nil
        }
-       if err := docker("inspect", "-f", "{{.State.Running}}", EnvRegistryName); err != nil {
-               var envErr *shx.InvalidEnvError // Note: Must be a pointer
-               if errors.As(err, &envErr) {
-                       return err
-               }
-               // Note: This error is mostly due to docker in NOT RUNNING state
-               // Start a local registry as docker service
-               return docker("run", "-d", "--restart", "always", "-p", format("127.0.0.1:%s:5000", EnvRegistryPort), "--name", EnvRegistryName, "registry:2")
-       }
-       return nil
+       _ = docker("stop", EnvRegistryName)
+       _ = docker("rm", EnvRegistryName)
+       return docker("run", "-d", "--restart", "always", "-p", format("127.0.0.1:%s:5000", EnvRegistryPort), "--name", EnvRegistryName, "registry:2")
 }
```
- **CHANGE**: Ensure test runs are idiomatic by first deleting & then creating KIND cluster
```diff
-       if err := kubectl("cluster-info", "--context", "kind-kind"); err != nil {
-               // Create kind cluster on error
-               // Note: error is swallowed
-               return kind("create", "cluster", "--config", kindClusterFilePath)
-       }
-       return nil
+       _ = kind("delete", "cluster")
+       return kind("create", "cluster", "--config", kindClusterFilePath)
```
- **FIX**: Set imgpkg to run with insecure registry to solve above https transport issue
```diff
-var imgpkg = shx.RunCmd(binPathCarvel + "/imgpkg")
+var insecureRegistry = shx.MaybeSetEnv("${INSECURE_REGISTRY}", "true")
+var isInsecureRegistry, _ = strconv.ParseBool(insecureRegistry)
+var imgpkg = func(args ...string) error {
+       binImgpkg := binPathCarvel + "/imgpkg"
+       if !isInsecureRegistry {
+               return shx.Run(binImgpkg, args...)
+       }
+       newArgs := append(args, "--registry-insecure")
+       return shx.Run(binImgpkg, newArgs...)
+}
```
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.internal make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (5.13s)
    testutil_test.go:27: expected no err got: running "tmp/kbld -f artifacts/packaging/release/packages --imgpkg-lock-output artifacts/packaging/release/.imgpkg/images.yml" failed with exit code 1 : kbld: Error: 
        - Resolving image 'kubernetes.docker.internal:5000/packages/k8s-remediator-app:1.0.1': Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client : 
FAIL
FAIL	carvel.shellx.dev	5.370s
ok  	carvel.shellx.dev/internal/sh	0.294s [no tests to run]
FAIL
```
- **CAUSE**: Error says `http: server gave HTTP response to HTTPS client`
- **CAUSE**: kbld is the CLI that fails

### Attempt 6/
- **FIX**: Set kbld to run with insecure registry to solve above https transport issue
```diff
-var kbld = shx.RunCmd(binPathCarvel + "/kbld")
-var whichKbld = shx.RunCmd("ls", binPathCarvel+"/kbld")
 var insecureRegistry = shx.MaybeSetEnv("${INSECURE_REGISTRY}", "true")
 var isInsecureRegistry, _ = strconv.ParseBool(insecureRegistry)
+var kbld = func(args ...string) error {
+       binKbld := binPathCarvel + "/kbld"
+       if !isInsecureRegistry {
+               return shx.Run(binKbld, args...)
+       }
+       return shx.Run(binKbld, append(args, "--registry-insecure")...)
+}
+var whichKbld = shx.RunCmd("ls", binPathCarvel+"/kbld")
```
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.internal make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (80.05s)
    testutil_test.go:27: expected no err got: running "kubectl get package -n k8s-remediator-system k8s-remediator.experiment.dev.com.1.0.1" failed with exit code 1 : error: the server doesn't have a resource type "package" : : func timed out after 46.052309208s
FAIL
```
- **DEBUG**: `kubectl describe pkgr -A`
```shell
Name:         k8s-remediator-repo.experiment.dev.com
Namespace:    k8s-remediator-system
Labels:       <none>
Annotations:  <none>
API Version:  packaging.carvel.dev/v1alpha1
Kind:         PackageRepository
Metadata:
  Creation Timestamp:  2022-10-03T11:49:18Z
Spec:
  Fetch:
    Imgpkg Bundle:
      Image:  kubernetes.docker.internal:5000/packages/k8s-remediator-repo.experiment.dev.com:1.0.1
Status:
  Conditions:
    Message:                       Fetching resources: Error (see .status.usefulErrorMessage for details)
    Status:                        True
    Type:                          ReconcileFailed
  Consecutive Reconcile Failures:  7
  Fetch:
    Error:       Fetching resources: Error (see .status.usefulErrorMessage for details)
    Exit Code:   1
    Started At:  2022-10-03T11:52:48Z
    Stderr:      vendir: Error: Syncing directory '0':
  Syncing directory '.' with imgpkgBundle contents:
    Imgpkg: exit status 1 (stderr: imgpkg: Error: Checking if image is bundle:
  Fetching image:
    Error while preparing a transport to talk with the registry:
      Unable to create round tripper:
        Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client
)

    Updated At:          2022-10-03T11:52:54Z
  Friendly Description:  Reconcile failed: Fetching resources: Error (see .status.usefulErrorMessage for details)
  Observed Generation:   1
  Useful Error Message:  vendir: Error: Syncing directory '0':
  Syncing directory '.' with imgpkgBundle contents:
    Imgpkg: exit status 1 (stderr: imgpkg: Error: Checking if image is bundle:
  Fetching image:
    Error while preparing a transport to talk with the registry:
      Unable to create round tripper:
        Get "https://kubernetes.docker.internal:5000/v2/": http: server gave HTTP response to HTTPS client
)

Events:  <none>
```
- **CAUSE**: Error says `http: server gave HTTP response to HTTPS client`
- **CAUSE**: kapp-controller reconciliation seems to fail

### Attempt 7/
- **IDEA**: Try `*.localhost` & check if host is automatically resolved to 127.0.0.1 
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.localhost make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (4.71s)
    testutil_test.go:27: expected no err got: running "tmp/imgpkg push -b kubernetes.docker.localhost:5000/packages/k8s-remediator-app:1.0.1 -f artifacts/packaging/source --registry-insecure" failed with exit code 1 : imgpkg: Error: Writing 'kubernetes.docker.localhost:5000/packages/k8s-remediator-app:1.0.1':
          Error while preparing a transport to talk with the registry:
            Unable to create round tripper:
              Get "https://kubernetes.docker.localhost:5000/v2/": dial tcp: lookup kubernetes.docker.localhost on 10.112.16.144:53: no such host; Get "http://kubernetes.docker.localhost:5000/v2/": dial tcp: lookup kubernetes.docker.localhost on 10.112.16.144:53: no such host : 
FAIL
```

### Attempt 8/
- **CHANGE**: Add `127.0.0.1 kubernetes.docker.localhost` entry in /etc/hosts
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.localhost make`
- **THEN**:
```shell
--- FAIL: TestCarvelReleaseE2E (78.53s)
    testutil_test.go:27: expected no err got: running "kubectl get package -n k8s-remediator-system k8s-remediator.experiment.dev.com.1.0.1" failed with exit code 1 : Error from server (NotFound): internalpackages.internal.packaging.carvel.dev "dcs76bbiclmmap39c5q6ushecls70pbid5mmarjk5pi6athecdnmqbhh5oo2sc8" not found : : func timed out after 44.744728542s
FAIL
```
- **DEBUG**: `kubectl get pkgr -A`
```shell
NAMESPACE               NAME                                     AGE     DESCRIPTION
k8s-remediator-system   k8s-remediator-repo.experiment.dev.com   3m11s   Reconcile succeeded
```
- **DEBUG**: `kubectl get package -A`
```shell
kubectl get package -A
NAMESPACE               NAME                                      PACKAGEMETADATA NAME                VERSION   AGE
k8s-remediator-system   k8s-remediator.experiment.dev.com.1.0.1   k8s-remediator.experiment.dev.com   1.0.1     2m35s
```
- **DEBUG**: `kubectl get pkgi -A`
```shell
No resources found
```
- **CAUSE**: Running "kubectl get package -n k8s-remediator-system k8s-remediator.experiment.dev.com.1.0.1": func timed out after 44.744728542s
- **CAUSE**: Logic might be timing out early without giving a chance for reconciliation to succeed

### Attempt 9/
- **FIX**: Increase the timeout to verify availability of package
```diff
 func verifyPresenceOfPackage() error {
-       return eventually(func() error {
+       return eventuallyWith(func() error {
                return kubectl("get", "package", "-n", EnvK8sNamespace, EnvPackageName+"."+EnvPackageVersion)
+       }, eventuallyConfig{
+               Attempts: ptrInt(20),
+               Interval: ptrDuration(3 * time.Second),
        })
 }
```
- **WHEN**: `REGISTRY_NAME=kubernetes.docker.localhost make`
- **THEN**:
```shell
go clean -testcache
SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -run TestCarvelReleaseE2E
ok  	carvel.shellx.dev	99.869s
ok  	carvel.shellx.dev/internal/sh	0.540s [no tests to run]
```
- **PROOF**: `kubectl get po -A`
```shell
NAMESPACE               NAME                                         READY   STATUS    RESTARTS   AGE
k8s-remediator-system   k8s-remediator-7fbb84fff6-5f5vl              1/1     Running   0          65s
```
- **PROOF**: `kubectl logs -n k8s-remediator-system   k8s-remediator-7fbb84fff6-5f5vl`
```shell
{"level":"info","ts":"2022-10-04T06:52:31.313Z","logger":"controller-runtime.metrics","msg":"Metrics server is starting to listen","addr":":8080"}
{"level":"info","ts":"2022-10-04T06:52:31.318Z","msg":"controller-vagc-va is enabled"}
{"level":"info","ts":"2022-10-04T06:52:31.319Z","logger":"setup","msg":"starting tkg-remediator controller manager"}
{"level":"info","ts":"2022-10-04T06:52:31.326Z","msg":"Starting server","path":"/metrics","kind":"metrics","addr":"[::]:8080"}
{"level":"info","ts":"2022-10-04T06:52:31.326Z","msg":"Starting server","kind":"health probe","addr":"[::]:9440"}
I1004 06:52:31.328836      11 leaderelection.go:248] attempting to acquire leader lease k8s-remediator-system/tkg-remediator-controller-leader-election...
I1004 06:52:31.355137      11 leaderelection.go:258] successfully acquired lease k8s-remediator-system/tkg-remediator-controller-leader-election
{"level":"info","ts":"2022-10-04T06:52:31.404Z","msg":"Starting EventSource","controller":"volumeattachment","controllerGroup":"storage.k8s.io","controllerKind":"VolumeAttachment","source":"kind source: *v1.VolumeAttachment"}
{"level":"info","ts":"2022-10-04T06:52:31.405Z","msg":"Starting Controller","controller":"volumeattachment","controllerGroup":"storage.k8s.io","controllerKind":"VolumeAttachment"}
{"level":"info","ts":"2022-10-04T06:52:31.512Z","msg":"Starting workers","controller":"volumeattachment","controllerGroup":"storage.k8s.io","controllerKind":"VolumeAttachment","worker count":1}
```
