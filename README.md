# testlinkerd
Simple testing of [linkerd](https://linkerd.io/2/getting-started/) with [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) and [gRPC](https://grpc.io/), other examples seemed a bit too bloated. Should work on linux. Has been tested on ubuntu bionic.

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

- [go 1.13](go)
- [kubectl]()
- [minikube]()
- [kustomize]()
- [linkerd]()

[go]: nrk.no