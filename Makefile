
all: hello world

# hello

hello: services/hello/.image.was.built

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

world: services/world/.image.was.built

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
