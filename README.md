# Toolbox Image

Toolbox is a docker image that contains debug tool for network issues.

## Docker
```console
docker run -it anfernee/network-toolbox
```

## Kubernetes
```console
kubectl run toolbox --tty --rm -i --image anfernee/network-toolbox -- sh
```
