# testlinkerd
Simple testing of linkerd with minikube and gRPC, other examples seemed a bit too bloated. Should work on linux. Has been tested on ubuntu bionic.

---

### Clone

You can either clone manually like so:

```git clone https://github.com/isakbm/testlinkerd.git github.com/isakbm/testlinkerd```

or you use go get:

```go get github.com/isakbm/testlinkerd```

### Deploy

To deploy simply run

```./deploy```

### Requirements

- go 1.13
- kubectl
- minikube
- kustomize
- linkerd

