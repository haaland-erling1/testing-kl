{{/*Generated by referencing from output of kustomize build config/default*/}}

{{- $name := .Name | pipefail "no .Name provided" -}}
{{- $namespace := .Namespace | pipefail "no .Namespace provided"  -}}
{{- $env := .Env | default "[]" | mustFromJson -}}
{{- $envFrom := .EnvFrom | default "[]" | mustFromJson -}}
{{- $image := .Image | pipefail "no .Image provided" -}}
{{- $imagePullPolicy := .ImagePullPolicy | pipefail "no .ImagePullPolicy provided" -}}
{{- $svcAccountName := .SvcAccountName | pipefail "no .SvcAccountName provided" -}}
{{- $tolerations := .Tolerations | default "[]" | mustFromJson -}}
{{- $nodeSelectors := .NodeSelectors | default "{}" | mustFromJson -}}

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kl-{{$name}}
  namespace: {{$namespace}}
  labels:
    app: {{$name}}
    control-plane: {{$name}}
spec:
  selector:
    matchLabels:
      app: {{$name}}
      control-plane: {{$name}}
  replicas: 1
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: manager
      labels:
        app: {{$name}}
        control-plane: {{$name}}
    spec:
      securityContext:
        runAsNonRoot: true
      {{ printf "{{- if .Tolerations }}" }}
      {{ printf "tolerations: {{ .Tolerations | default list | mustFromJson | toYAML | nindent 8 }}" }}
      {{ printf "{{- end }}" }}

      {{ printf "{{- if .NodeSelector }}" }}
      {{ printf "nodeSelector: {{ .NodeSelector | default dict | mustFromJson | toYAML | nindent 8 }}" }}
      {{ printf "{{- end }}" }}
      containers:
        - args:
            - --secure-listen-address=0.0.0.0:8443
            - --upstream=http://127.0.0.1:8080/
            - --logtostderr=true
            - --v=0
          image: gcr.io/kubebuilder/kube-rbac-proxy:v0.8.0
          name: kube-rbac-proxy
          ports:
            - containerPort: 8443
              name: https
              protocol: TCP
          resources:
            limits:
              cpu: 20m
              memory: 20Mi
            requests:
              cpu: 5m
              memory: 10Mi

        - command:
            - /manager
          args:
            - --health-probe-bind-address=:8081
            - --metrics-bind-address=127.0.0.1:8080
            - --leader-elect
          env:
            - name: RECONCILE_PERIOD
              value: "30s"
            - name: MAX_CONCURRENT_RECONCILES
              value: "2"
            {{- if $env }}
            {{$env | toYAML | nindent 12}}
            {{- end}}
          {{if $envFrom }}
          envFrom: {{$envFrom |toYAML| nindent 13}}
          {{- end}}
{{/*            - secretRef:*/}}
{{/*                name: operator-env*/}}
          image: {{$image}}
          imagePullPolicy: {{$imagePullPolicy}}
          name: manager
          securityContext:
            allowPrivilegeEscalation: false
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8081
            initialDelaySeconds: 15
            periodSeconds: 20
          readinessProbe:
            httpGet:
              path: /readyz
              port: 8081
            initialDelaySeconds: 5
            periodSeconds: 10
          resources:
            limits:
              memory: 200Mi
            requests:
              cpu: 64m
              memory: 160Mi
      serviceAccountName: {{$svcAccountName}}
      terminationGracePeriodSeconds: 10
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{$name}}
    control-plane: {{$name}}
  name: {{$name}}-metrics-service
  namespace: {{$namespace}}
spec:
  ports:
    - name: https
      port: 8443
      protocol: TCP
      targetPort: https
  selector:
    app: {{$name}}
    control-plane: {{$name}}
