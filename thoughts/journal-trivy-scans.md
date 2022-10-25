- Start Date: 17-Oct-2022
- Last Updated Date: 17-Oct-2022
- Author(s): @amitd2

### Machine Details
```shell
uname
Darwin
```

### Docker Version
```shell
docker version
Error response from daemon: dial unix docker.raw.sock: connect: no such file or directory
Client:
 Cloud integration: v1.0.28
 Version:           20.10.17
 API version:       1.41
 Go version:        go1.17.11
 Git commit:        100c701
 Built:             Mon Jun  6 23:04:45 2022
 OS/Arch:           darwin/arm64
 Context:           default
 Experimental:      true
```

### Trivy Version
```shell
trivy --version
Version: 0.32.1
Vulnerability DB:
  Version: 2
  UpdatedAt: 2022-10-14 00:18:41.356360643 +0000 UTC
  NextUpdate: 2022-10-14 06:18:41.356360043 +0000 UTC
  DownloadedAt: 2022-10-14 05:28:59.196476 +0000 UTC
```

### OCI Registry Details
```yaml
- projects-stg.registry.vmware.com/tkg/grafana/k8s-sidecar:v1.12.1_vmware.3 # user facing
- projects.registry.vmware.com/tkg/grafana/k8s-sidecar:v1.12.1_vmware.3 # staging
```

### Source Code
```shell
git config -l  
credential.helper=osxkeychain
init.defaultbranch=main
rerere.enabled=true
user.name=AmitKumarDas
user.email=amitd2@vmware.com
core.editor=vim
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
core.ignorecase=true
core.precomposeunicode=true
remote.origin.url=git@gitlab.eng.vmware.com:core-build/cayman_k8s-sidecar.git
remote.origin.fetch=+refs/heads/*:refs/remotes/origin/*
branch.master.remote=origin
branch.master.merge=refs/heads/master
branch.topic/amitd/v1.15.6+vmware.2.remote=origin
branch.topic/amitd/v1.15.6+vmware.2.merge=refs/heads/vmware-1.15.6
submodule.k8s-sidecar/src.active=true
submodule.k8s-sidecar/src.url=git@gitlab.eng.vmware.com:core-build/mirrors_github_k8s-sidecar
```

#### Verify Image from Registry
```shell
imgpkg tag list -i projects.registry.vmware.com/tkg/grafana/k8s-sidecar
```

```shell
docker pull projects.registry.vmware.com/tkg/grafana/k8s-sidecar:v1.12.1_vmware.2

REPOSITORY                                                      TAG                IMAGE ID       CREATED          SIZE
projects.registry.vmware.com/tkg/grafana/k8s-sidecar            v1.12.1_vmware.2   74e38e23550d   10 months ago    517MB
```

```shell
trivy image -q projects.registry.vmware.com/tkg/grafana/k8s-sidecar:v1.12.1_vmware.2

projects.registry.vmware.com/tkg/grafana/k8s-sidecar:v1.12.1_vmware.2 (debian 10.11)
====================================================================================
Total: 946 (UNKNOWN: 8, LOW: 493, MEDIUM: 190, HIGH: 199, CRITICAL: 56)

Python (python-pkg)

Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)
```

#### Verify from Source Code @ **vmware-1.12.1**
```shell
git checkout origin/vmware-1.12.1
git submodule init
git submodule update
````

```shell
git status
HEAD detached at origin/vmware-1.12.1
```

```shell
% git diff
diff --git a/k8s-sidecar/src b/k8s-sidecar/src
--- a/k8s-sidecar/src
+++ b/k8s-sidecar/src
@@ -1 +1 @@
-Subproject commit fbcf38f0bc1841490f3c66e32e2bfce4e282438d
+Subproject commit fbcf38f0bc1841490f3c66e32e2bfce4e282438d-dirty
```

```shell
cd k8s-sidecar/src

git status
HEAD detached at fbcf38f
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   Dockerfile
```

```shell
cat Dockerfile

FROM        python:3.8-alpine
ENV         PYTHONUNBUFFERED=1
WORKDIR     /app

COPY        requirements.txt .
RUN         apk add --no-cache gcc && \
	        pip install -r requirements.txt && \
	        apk del -r gcc && \
            rm -rf /var/cache/apk/* requirements.txt

COPY        sidecar/* ./

# Use the nobody user's numeric UID/GID to satisfy MustRunAsNonRoot PodSecurityPolicies
# https://kubernetes.io/docs/concepts/policy/pod-security-policy/#users-and-groups
USER        65534:65534

CMD         [ "python", "-u", "/app/sidecar.py" ]
```

```shell
docker build -t cayman-k8s-sidecar:1.12.1-orig .
```

```shell
docker images
REPOSITORY                                                      TAG                IMAGE ID       CREATED         SIZE
cayman-k8s-sidecar                                              1.12.1-orig        08e368589aac   2 days ago      94.1MB
```

```shell
trivy image cayman-k8s-sidecar:1.12.1-orig

cayman-k8s-sidecar:1.12.1-orig (alpine 3.16.2)

Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)

2022-10-17T17:22:47.371+0530  INFO  Table result includes only package filenames. Use '--format json' option to get the full path to the package file.

Python (python-pkg)

Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 1, CRITICAL: 0)

┌────────────────────┬────────────────┬──────────┬───────────────────┬───────────────┬───────────────────────────────────────────────────────────┐
│      Library       │ Vulnerability  │ Severity │ Installed Version │ Fixed Version │                           Title                           │
├────────────────────┼────────────────┼──────────┼───────────────────┼───────────────┼───────────────────────────────────────────────────────────┤
│ urllib3 (METADATA) │ CVE-2021-33503 │ HIGH     │ 1.25.11           │ 1.26.5        │ python-urllib3: ReDoS in the parsing of authority part of │
│                    │                │          │                   │               │ URL                                                       │
│                    │                │          │                   │               │ https://avd.aquasec.com/nvd/cve-2021-33503                │
└────────────────────┴────────────────┴──────────┴───────────────────┴───────────────┴───────────────────────────────────────────────────────────┘
```


#### Attempt 1 - Scan after updating to Distroless (python3-debian11)


```shell
git diff
```

```diff
diff --git a/Dockerfile b/Dockerfile
index ed6e1b9..6ceebe9 100644
--- a/Dockerfile
+++ b/Dockerfile
@@ -1,17 +1,24 @@
-FROM        python:3.8-alpine
+FROM harbor-repo.vmware.com/dockerhub-proxy-cache/library/python:3-slim as builder
 ENV         PYTHONUNBUFFERED=1
 WORKDIR     /app
 
 COPY        requirements.txt .
-RUN         apk add --no-cache gcc && \
-               pip install -r requirements.txt && \
-               apk del -r gcc && \
-            rm -rf /var/cache/apk/* requirements.txt
+
+RUN         pip3 install --upgrade pip
+RUN            pip install -r requirements.txt && \
+            rm -rf requirements.txt
 
 COPY        sidecar/* ./
 
+FROM gcr.io/distroless/python3-debian11:nonroot
+
 # Use the nobody user's numeric UID/GID to satisfy MustRunAsNonRoot PodSecurityPolicies
 # https://kubernetes.io/docs/concepts/policy/pod-security-policy/#users-and-groups
 USER        65534:65534
 
+ENV         PYTHONUNBUFFERED=1
+WORKDIR     /app
+
+COPY        --from=builder /app /app
+
 CMD         [ "python", "-u", "/app/sidecar.py" ]
```


```shell
docker build -t cayman-k8s-sidecar:1.12.1-attempt.1 .
```

```shell
docker images
REPOSITORY                                                      TAG                IMAGE ID       CREATED         SIZE
cayman-k8s-sidecar                                              1.12.1-attempt.1   3a6078f3c502   6 minutes ago   50.2MB
cayman-k8s-sidecar                                              1.12.1-orig        08e368589aac   2 days ago      94.1MB
```


```shell
trivy image cayman-k8s-sidecar:1.12.1-attempt.1

cayman-k8s-sidecar:1.12.1-attempt.1 (debian 11.5)
Total: 49 (UNKNOWN: 0, LOW: 34, MEDIUM: 6, HIGH: 6, CRITICAL: 3)
```

#### Attempt 2 - Scan after updating requirements.txt

```diff
diff --git a/Dockerfile b/Dockerfile
index ed6e1b9..54d3739 100644
--- a/Dockerfile
+++ b/Dockerfile
@@ -1,4 +1,4 @@
-FROM        python:3.8-alpine
+FROM        harbor-repo.vmware.com/dockerhub-proxy-cache/library/python:alpine3.16
 ENV         PYTHONUNBUFFERED=1
 WORKDIR     /app
 
diff --git a/requirements.txt b/requirements.txt
index 714ee01..3d7f0e9 100644
--- a/requirements.txt
+++ b/requirements.txt
@@ -1,2 +1,3 @@
 kubernetes==12.0.0
-requests==2.24.0
+requests>=2.24.0
+urllib3>=1.26.5
```


```shell
docker build -t cayman-k8s-sidecar:1.12.1-attempt.2 .
```

```shell
% docker images

REPOSITORY                                                      TAG                IMAGE ID       CREATED         SIZE
cayman-k8s-sidecar                                              1.12.1-attempt.2   cf74e367422c   2 minutes ago   95.3MB
```

```shell
trivy image cayman-k8s-sidecar:1.12.1-attempt.2

2022-10-17T17:29:28.516+0530  INFO  Vulnerability scanning is enabled
2022-10-17T17:29:28.516+0530  INFO  Secret scanning is enabled
2022-10-17T17:29:28.516+0530  INFO  If your scanning is slow, please try '--security-checks vuln' to disable secret scanning
2022-10-17T17:29:28.516+0530  INFO  Please see also https://aquasecurity.github.io/trivy/v0.32/docs/secret/scanning/#recommendation for faster secret detection
2022-10-17T17:29:30.917+0530  INFO  Detected OS: alpine
2022-10-17T17:29:30.917+0530  INFO  Detecting Alpine vulnerabilities...
2022-10-17T17:29:30.921+0530  INFO  Number of language-specific files: 1
2022-10-17T17:29:30.921+0530  INFO  Detecting python-pkg vulnerabilities...

cayman-k8s-sidecar:1.12.1-attempt.2 (alpine 3.16.2)

Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)
```

### Merge Request

```shell
git checkout -b topic/amitd/v1.12.1+vmware.4 origin/vmware-1.12.1
```

```shell
% git status
On branch topic/amitd/v1.12.1+vmware.4
Your branch is up to date with 'origin/vmware-1.12.1'.
```

```diff
% git diff
diff --git a/k8s-sidecar/k8s-sidecar_build.sh b/k8s-sidecar/k8s-sidecar_build.sh
index 641f1e0..18f2545 100644
--- a/k8s-sidecar/k8s-sidecar_build.sh
+++ b/k8s-sidecar/k8s-sidecar_build.sh
@@ -69,9 +69,16 @@ mkdir -p "${STAGING_OUTPUT_DIR}"
 mkdir -p "${CONCRETE_IMAGES_OUTPUT_DIR}"
 mkdir -p "${CONCRETE_EXECUTABLES_OUTPUT_DIR}"
 
+# create requirements.txt
+cat << EOF > requirements.txt
+kubernetes==12.0.0
+requests>=2.24.0
+urllib3>=1.26.5
+EOF
+
 # create dockerfile
 cat << EOF > ${DOCKERFILE}
-FROM        harbor-repo.vmware.com/dockerhub-proxy-cache/library/python:3.8-alpine 
+FROM        harbor-repo.vmware.com/dockerhub-proxy-cache/library/python:alpine3.16
 ENV         PYTHONUNBUFFERED=1
 WORKDIR     /app
```

```shell
git add .
git commit -s -v
git push origin HEAD
```

```shell
ssh amitd2@blr-dbc2103

/build/mts/apps/bin/gobuild sandbox queue --branch "topic/amitd/v1.12.1+vmware.4" --no-store-trees --buildtype release --changeset False --no-send-email  --bootstrap "tkg-remediator=git-eng:core-build/cayman_k8s-sidecar.git;%(branch);"  --syncto latest cayman_k8s-sidecar
```

```
https://buildweb.eng.vmware.com/sb/60264252/
```




