# Kubernetes Samples

This project contains the test images for Kubernetes

## Build

These images are intended to be built with multi-arch. Therefore, [buildx](https://docs.docker.com/buildx/working-with-buildx/) is required to build them.
```
docker buildx build -t "IMAGE_TAG" --platform PLATFORMS IMAGE_DIRECTORY
```
For instance, to build the hello-world image:
```
docker buildx build -t "charmed-kubernetes/hello-world:1.0" --platform linux/amd64,linux/arm64,linux/s390x hello-world
```
