apiVersion: batch/v1
kind: Job
metadata:
  name: {{ .Name }}
  namespace: {{ .Namespace }}
  labels: {{ .Labels | toJson }}
  annotations: {{ .Annotations | toJson }}

spec:
  backoffLimit: 0
  suspend: false
  template:
    metadata:
      name: {{ .Name }}
    spec:
      shareProcessNamespace: true
      volumes:
      - name: docker-socket
        emptyDir: {}

      - name: hostpath-volume
        hostPath:
          path: /var/docker-data/{{ .AccountName }}
          type: DirectoryOrCreate


      containers:
      - name: dind-server
        volumeMounts:
        - name: docker-socket
          mountPath: /var/run
        - name: hostpath-volume
          mountPath: /var/lib/docker

        image: ghcr.io/kloudlite/platform/apis/docker:dind
        securityContext:
          privileged: true
        resources:
          requests:
            memory: "1024Mi"
            cpu: "0.5"
          limits:
            memory: "2048Mi"
            cpu: "1"


      - name: build-and-push
        volumeMounts:
        - name: docker-socket
          mountPath: /var/run
        image: ghcr.io/kloudlite/image-builder:v1.0.5-nightly
        env:
        - name: DOCKER_PSW
          value: {{ .DockerPassword }}

        command: ["bash", "-c"]
        args:
        - |
          set -o errexit
          set -o pipefail

          trap 'pkill dockerd' SIGINT SIGTERM EXIT
          while ! docker info > /dev/null 2>&1 ; do sleep 1; done

          tag={{ .Registry }}/{{ .RegistryRepoName }}:{{ .Tag }}

          echo $DOCKER_PSW | docker login -u {{ .KlAdmin }} --password-stdin {{ .Registry }}

          docker buildx build -t $tag -o type=registry,oci-mediatypes=true,compression=estargz,force-compression=true {{ .PullUrl }}  2>&1 | grep -v '\[internal\]'
      restartPolicy: Never
