# testlinkerd
Simple testing of [linkerd] with [minikube] and [gRPC], other examples seemed a bit too bloated. Should work on linux. Has been tested on ubuntu bionic.

---

### Clone

You can either clone manually like so:

```git clone https://github.com/isakbm/testlinkerd.git github.com/isakbm/testlinkerd```

or use go get:

```go get github.com/isakbm/testlinkerd```

### Deploy

To deploy simply run

```./deploy```

### Requirements

- [go]
- [kubectl]
- [minikube]
- [kustomize]
- [linkerd]
- [protoc]
- [protoc-gen-go]

[go]: https://golang.org/
[gRPC]: https://grpc.io/
[linkerd]: https://linkerd.io/2/getting-started/
[minikube]: https://kubernetes.io/docs/tasks/tools/install-minikube/
[kustomize]: https://kustomize.io/
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[protoc]:  https://developers.google.com/protocol-buffers
[protoc-gen-go]: https://github.com/golang/protobuf
