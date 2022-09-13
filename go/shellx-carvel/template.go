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
