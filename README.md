otus-ma-lesson-6
---

### Build a docker image:

```shell
REPO=kvmayer/otus-ma-lesson-6
TAG=0.1.3
ARCH=linux/amd64,linux/arm64
docker buildx build --platform $ARCH -t $REPO:$TAG --push .
```

### Deploy database with helm

```shell
helm install postgres bitnami/postgresql --values .helm-postgres/postgres-values.yaml
```

### Deploy app
```shell
kubectl apply -f .kube/secret.yaml
kubectl apply -f .kube/configmap.yaml
kubectl apply -f .kube/db-migrations-job.yaml
kubectl apply -f .kube/deployment.yaml
kubectl apply -f .kube/service.yaml
kubectl apply -f .kube/ingress.yaml
```
or
```shell
kubectl apply -f .kube
```

### Run tests
```shell
newman run user_service.postman_collection.json
```