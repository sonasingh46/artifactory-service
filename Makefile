ifeq (${IMAGE_ORG}, )
  IMAGE_ORG="sonasingh46"
  export IMAGE_ORG
endif

# Specify the docker arg for repository url
ifeq (${DBUILD_REPO_URL}, )
  DBUILD_REPO_URL="https://github.com/sonasingh46/artifactory-service"
  export DBUILD_REPO_URL
endif

# Specify the date of build
DBUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')


# Determine the arch/os
ifeq (${XC_OS}, )
  XC_OS:=$(shell go env GOOS)
endif
export XC_OS

ifeq (${XC_ARCH}, )
  XC_ARCH:=$(shell go env GOARCH)
endif
export XC_ARCH

ARCH:=${XC_OS}_${XC_ARCH}
export ARCH

ifeq (${IMAGE_TAG}, )
  IMAGE_TAG = ci
  export IMAGE_TAG
endif

export DBUILD_ARGS=--build-arg DBUILD_DATE=${DBUILD_DATE} --build-arg DBUILD_REPO_URL=${DBUILD_REPO_URL} --build-arg ARCH=${ARCH}

# Specify the name of the docker repo for amd64
ART_SERVICE_REPO_NAME=artifactory-service-amd64


# Specify the directory location of main package after bin directory
# e.g. bin/{DIRECTORY_NAME_OF_APP}
ART_SERVICE=artifactory-service

# list only the source code directories
PACKAGES = $(shell go list ./... | grep -v 'vendor\|pkg/client/generated\|tests')

# deps ensures fresh go.mod and go.sum.
.PHONY: deps
deps:
	@go mod tidy
	@go mod verify

.PHONY: test
test:
	go fmt ./...
	go test $(PACKAGES)


.PHONY: artifactory-service.amd64
artifactory-service.amd64 :
	@echo -n "--> Building artifactory service image <--"
	@echo "${IMAGE_ORG}/${ART_SERVICE_REPO_NAME}:${IMAGE_TAG}"
	@echo "----------------------------"
	@PNAME=${ART_SERVICE} CTLNAME=${ART_SERVICE} sh -c "'$(PWD)/build/build.sh'"
	@cp bin/${ART_SERVICE}/${ART_SERVICE} build/artifactory-service/
	@cd build/${ART_SERVICE} && sudo docker build -t ${IMAGE_ORG}/${ART_SERVICE_REPO_NAME}:${IMAGE_TAG} ${DBUILD_ARGS} .
	@rm build/${ART_SERVICE}/${ART_SERVICE}

.PHONY: all.amd64
all.amd64: artifactory-service.amd64