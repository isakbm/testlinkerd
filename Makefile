
.PHONY: all hello world clean
all: hello world

NAMESPACE := testlinkerd
CA_DIR := .ca

# hello

hello: services/hello/.image.was.built services/hello/tls.crt

services/hello/.image.was.built: services/hello/Dockerfile services/hello/service
	./minidocker build -t testlinkerd/hello $(@D)
	touch $@

services/hello/service: pkg/world/* pkg/world/world.pb.go services/hello/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "-s -w" \
		-pkgdir .pkg \
		-i \
		-o ./$@ \
		./$(@D)/*.go 

# world

world: services/world/.image.was.built services/world/tls.crt

services/world/.image.was.built: services/world/Dockerfile services/world/service
	./minidocker build -t testlinkerd/world $(@D)
	touch $@

services/world/service: pkg/world/* pkg/world/world.pb.go services/world/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-tags netgo \
		-installsuffix netgo \
		-ldflags "-s -w" \
		-pkgdir .pkg \
		-i \
		-o ./$@ \
		./$(@D)/*.go 

pkg/world/world.pb.go: pkg/world/world.proto
	protoc $< --go_out=plugins=grpc:.

# certificates

$(CA_DIR)/tls.crt: $(CA_DIR)/tls.key
	openssl req -x509 -new -nodes -key $< -days 3650 -out $@ -subj "/CN=$(NAMESPACE) Certificate Authority"

%.key:
	mkdir -p $(@D)
	openssl genrsa -out $@ 2048

%.csr: %.key $(CA_DIR)/tls.cnf
	SAN="DNS:$(lastword $(subst /, ,$(<D)))" openssl req -new -key $< -out $@ -subj "/CN=$(lastword $(subst /, ,$(<D)))" -extensions san_env -config $(word 2, $^)

%.crt: %.key $(CA_DIR)/tls.cnf %.csr $(CA_DIR)/tls.crt $(CA_DIR)/tls.key
	SAN="DNS:$(lastword $(subst /, ,$(<D)))" openssl x509 -req -in $(word 3,$^) -CA $(word 4,$^) -CAkey $(word 5,$^) -CAcreateserial -out $@ -days 365 -extensions v3_req -extensions san_env -extfile $(word 2, $^)

# clean

clean:
	rm services/*/service
	rm services/*/tls.key
	rm services/*/tls.crt
	rm services/*/.image.was.built
	rm $(CA_DIR)/tls.crt
	rm $(CA_DIR)/tls.key
	rm $(CA_DIR)/tls.srl

