gitTime=$(shell date +%Y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD | cut -c1-8)
gitTag=$(shell git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(shell git log --pretty=format:'' | wc -l)/$(shell git rev-list --all --count)
buildStr=${gitTime}.${gitCID}.${gitTag}.${gitCount}
# 获取当前用户的 UID 和 GID
UID=$(shell id -u)
GID=$(shell id -g)
fileTime=$(shell date +%Y%m%d%H%M)
ImageFullName=leehom/detect:centos7.go1.21.7

.PHONY: build
build:
	@docker run --rm -it -e CGO_ENABLED=1 -e GF_DEBUG=1 -u ${UID}:${GID} -v $(shell pwd)/:/app -w /app ${ImageFullName} gf build -mod vendor -o bin/dho.amd64 cmd/v3/main.go

.PHONY: buildarm
buildarm:
	@docker run --rm -it -e CGO_ENABLED=1 -e GF_DEBUG=1 -u ${UID}:${GID} -v $(shell pwd)/:/app -w /app ${ImageFullName} gf build -a arm64 -mod vendor -o bin/dho.arm64 cmd/v3/main.go
	scp bin/dho.arm64 companyft2:~/lianghong/

.PHONY: test
test:
	@docker run --rm -it -e CGO_ENABLED=1 -e GF_DEBUG=1 -v $(shell pwd)/:/app -w /app ${ImageFullName} ./bin/dho.amd64
