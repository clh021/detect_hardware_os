GO_ARCH:= $(shell go env GOARCH)
gitTime:= $(shell git log -1 --format="%cd" --date=format:'%Y-%m-%d %H:%M:%S')
buildTime=$(shell date +'%Y-%m-%d %H:%M:%S')
gitHash=$(shell git rev-parse HEAD | cut -c1-19)
gitTag=$(shell git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(shell git rev-list --all --count)
# uiTime=$(shell cat resource/public/dist/env | grep UI_BUILD_TIME | cut -d ':' -f 2-)
# 清理可能存在的空格
buildStr=${gitTime}.${gitHash}.${gitTag}.${gitCount}
# 获取当前用户的 UID 和 GID
UID=$(shell id -u)
GID=$(shell id -g)
fileTime=$(shell date +%Y%m%d%H%M)
programName=detect_hardware_os
ImageFullName=leehom/detect:centos7.go1.21.7

.PHONY: build
build:buildgf

.PHONY: buildgo
buildgo:
	docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -u ${UID}:${GID} -v $(shell pwd)/:/app -w /app \
	${ImageFullName} \
	go build \
	-mod vendor \
	-o ./bin/0.0.${gitCount}/linux_amd64/${programName} \
	-ldflags '\
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.BuildTime=${buildTime}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitTime=${gitTime}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitHash=${gitHash}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitCount=${gitCount}\" \
	' \
	cmd/v3/main.go
#	scp bin/0.0.${gitCount}/linux_amd64/${programName} vboxV10:~/lianghong/
#	scp bin/0.0.${gitCount}/linux_arm64/${programName} companyft2:~/lianghong/
# V10 安全中心全部禁用,重启,检查全部禁用

.PHONY: buildgf
buildgf:
	docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -u ${UID}:${GID} -v $(shell pwd)/:/app -w /app \
	-v ${HOME}/.bin/:/usr/local/hostbin/ \
	${ImageFullName} \
	/usr/local/hostbin/gf build \
	cmd/v3/main.go \
	-mod vendor \
	-v 0.0.${gitCount} \
	-n ${programName} \
	-a amd64,arm64 -s linux \
	-p ./bin \
	-e "-trimpath -ldflags '\
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.BuildTime=${buildTime}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitTime=${gitTime}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitHash=${gitHash}\" \
	-X \"github.com/clh021/detect_hardware_os/service/cmd/version.GitCount=${gitCount}\" \
	'"
#	scp bin/0.0.${gitCount}/linux_amd64/${programName} vboxV10:~/lianghong/
#	scp bin/0.0.${gitCount}/linux_arm64/${programName} companyft2:~/lianghong/
# V10 安全中心全部禁用,重启,检查全部禁用

.PHONY: test
test:
	@docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -v $(shell pwd)/:/app -w /app ${ImageFullName} ./0.0.${gitCount}/linux_${GO_ARCH}/${programName}

.PHONY: test_ver
test_ver:
	@docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -v $(shell pwd)/bin:/app -w /app ${ImageFullName} ./0.0.${gitCount}/linux_${GO_ARCH}/${programName} version

.PHONY: test_euler
test_euler:
	@docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -v $(shell pwd)/bin:/app -w /app openeuler/openeuler:20.03-lts ./0.0.${gitCount}/linux_${GO_ARCH}/${programName}

.PHONY: test_debian
test_debian:
	@docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -v $(shell pwd)/bin:/app -w /app debian:10 ./0.0.${gitCount}/linux_${GO_ARCH}/${programName}

.PHONY: test_kylinv10
test_kylinv10:
	@docker run --rm -it -e CGO_ENABLED=0 -e GF_DEBUG=1 -v $(shell pwd)/bin:/app -w /app hoytluo/kylinv10sp2:0608 ./0.0.${gitCount}/linux_${GO_ARCH}/${programName}
