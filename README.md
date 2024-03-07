```
docker run --rm \
  -v $(pwd):/workspace \
  -v ~/.docker/config.json:/kaniko/.docker/config.json \
  gcr.io/kaniko-project/executor:v1.21.1 \
  --dockerfile=/workspace/Dockerfile \
  --context=/workspace \
  --destination=serg3091/app:kaniko

```
