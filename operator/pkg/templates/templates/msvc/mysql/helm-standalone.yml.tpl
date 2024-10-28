{{- $ownerRefs := get . "owner-refs" }}
{{- $storageClass := get . "storage-class"}}
{{- $obj := get . "obj"}}
{{- $rootPassword := get . "root-password"  -}}
{{- $mysqlUserPassword := get . "mysql-user-password"  -}}
{{- $existingSecret := get . "existing-secret"  -}}
{{- $priorityClassName := get . "priority-class-name"  | default "stateful" -}}

{{- with $obj }}
{{- /*gotype: github.com/kloudlite/operator/apis/mysql.msvc/v1.StandaloneService*/ -}}
apiVersion: msvc.kloudlite.io/v1
kind: HelmMySqlDB
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
  ownerReferences: {{ $ownerRefs | toYAML | nindent 4 }}
  labels: {{ .Labels | default dict | toYAML | nindent 4 }}
spec:
  global:
    storageClass: {{$storageClass}}

  fullnameOverride: {{.Name}}
  image:
    tag: 8.0.30-debian-11-r6

  {{- if .Labels }}
  commonLabels: {{.Labels | toYAML | nindent 4}}
  {{- end }}

  architecture: standalone
  auth:
    enabled: true
    existingSecret: {{$existingSecret}}
  primary:
    persistence:
      enabled: true
      size: {{.Spec.Resources.Storage.Size}}
    podLabels:
      {{.Labels  | default dict | toYAML | nindent 6 }}
      kloudlite.io/region: {{.Spec.Region}}
      kloudlite.io/stateful-node: "true"

    priorityClassName: {{$priorityClassName}}
    affinity: {{include "NodeAffinity" (dict) | toYAML | nindent 6}}
    tolerations: {{include "RegionToleration" (dict "region" .Spec.Region) | nindent 6}}
    nodeSelector: {{include "RegionNodeSelector" (dict "region" .Spec.Region) | nindent 6}}

    resources:
      limits:
        cpu: {{.Spec.Resources.Cpu.Max}}
        memory: {{.Spec.Resources.Memory}}
      requests:
        cpu: {{.Spec.Resources.Cpu.Min}}
        memory: {{.Spec.Resources.Memory}}
{{- end}}
