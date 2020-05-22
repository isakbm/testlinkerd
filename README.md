# testlinkerd

Minimal example to test [linkerd] with [gRPC] services on a [kubernetes] cluster deployed using [minikube].

There are two services

> **hello**: sends HelloReq ([gRPC]) to the world service

> **world**: serves HelloReq ([gRPC])

Thanks to [linkerd], the **world** service can be scaled up, and requests from the **hello** service will be load balanced accross them automatically. This is nontrivial since [gRPC] uses http/2, and therefore requires request based load balancing, meaning that traditional connect based load balancing will not suffice. Yay for [linkerd].

Don't forget to view the linkerd dashboard once you're done

```linkerd dashboard```

---

### Clone

You can either clone manually like so:

```git clone https://github.com/isakbm/testlinkerd.git github.com/isakbm/testlinkerd```

or use go get:

```go get github.com/isakbm/testlinkerd```

### Requirements

- [go]
- [kubectl]
- [minikube]
- [kustomize]
- [linkerd]
- [protoc]
- [protoc-gen-go]

### Deploy

To deploy simply run

```./depl```


[go]: https://golang.org/
[gRPC]: https://grpc.io/
[linkerd]: https://linkerd.io/2/getting-started/
[minikube]: https://kubernetes.io/docs/tasks/tools/install-minikube/
[kustomize]: https://kustomize.io/
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[protoc]:  https://developers.google.com/protocol-buffers
[protoc-gen-go]: https://github.com/golang/protobuf
[kubernetes]: https://kubernetes.io/docs/reference/

### Issues

If you do not already have a minikube cluster running, or if otherwise
you are having issues with minikube, do

```minikube delete```

```./start-minikube```

Note that running certain minikube commands fail if there is a file or directory
named deploy. Therefore the deploy script is called depl.