# testlinkerd

Minial example used to test [linkerd] with [gRPC] services on a [kubernetes] cluster deployed using [minikube].

There are two services

> **hello**: sends HelloReq ([gRPC]) to the world service

> **world**: serves HelloReq ([gRPC])

Thanks to [linkerd], the **world** service can be scaled up, and requests from the **hello** service will be load balanced accross them automatically. This is nontrivial since [gRPC] uses http/2, and therefore requires request based load balancing, meaning that traditional connect based load balancing will not suffice. Yay for [linkerd].

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
[kubernetes]: https://kubernetes.io/docs/reference/
