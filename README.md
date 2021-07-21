otus-ma-lesson-6
---

Build a docker image:

```shell
REPO=kvmayer/otus-ma-lesson-6
TAG=0.1.0
ARCH=linux/amd64,linux/arm64
docker buildx build --platform $ARCH -t $REPO:$TAG --push .
```
