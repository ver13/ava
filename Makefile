# ------------------------------------------------------------------------------
# MAKEFILE
#
# @author      Valentin Encinas <vencinasrojas@gmail.com>
# @link        https://www.linkedin.com/in/ver13/
# ------------------------------------------------------------------------------

# Use bash as shell (Note: Ubuntu now uses dash which doesn't support PIPESTATUS).
SHELL=/bin/bash

GO                ?= go
GOFMT             ?= $(GO)fmt
GO_VERSION        ?= $(shell $(GO) version)
GO_VERSION_NUMBER ?= $(word 3, $(GO_VERSION))
PRE_GO_111        ?= $(shell echo $(GO_VERSION_NUMBER) | grep -E 'go1\.(10|[0-9])\.')

GOVENDOR :=
GO111MODULE :=
ifeq (, $(PRE_GO_111))
	ifneq (,$(wildcard go.mod))
		# Enforce Go modules support just in case the directory is inside GOPATH (and for Travis CI).
		GO111MODULE := on

		ifneq (,$(wildcard vendor))
			# Always use the local vendor/ directory to satisfy the dependencies.
			GOOPTS := $(GOOPTS) -mod=vendor
		endif
	endif
else
	ifneq (,$(wildcard go.mod))
		ifneq (,$(wildcard vendor))
$(warning This repository requires Go >= 1.11 because of Go modules)
$(warning Some recipes may not work as expected as the current Go runtime is '$(GO_VERSION_NUMBER)')
		endif
	else
		# This repository isn't using Go modules (yet).
		GOVENDOR := $(FIRST_GOPATH)/bin/govendor
	endif
endif

# Project owner
OWNER=ver13

# Project vendor
VENDOR=ver13

# Project name
PROJECT=ava

# Project version
VERSION=$(shell cat VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat RELEASE)

# Name of RPM or DEB package
PKGNAME=${VENDOR}-${PROJECT}

# Current directory
CURRENTDIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# GO lang path
ifneq ($(GOPATH),)
	ifeq ($(findstring $(GOPATH),$(CURRENTDIR)),)
		# the defined GOPATH is not valid
		GOPATH=
	endif
endif
ifeq ($(GOPATH),)
	# extract the GOPATH
	GOPATH=$(firstword $(subst /./, ,$(CURRENTDIR)))
endif

# Add the GO binary dir in the PATH
export PATH := $(GOPATH)/bin:$(PATH)

# Path for binary files (where the executable files will be installed)
BINPATH=bin

# Path for configuration files
CONFIGPATH=configs/$(PROJECT)/

# Path for ssl root certs
#SSLCONFIGPATH=etc/ssl/
SSLCONFIGPATH=

# Path for init script
INITPATH=etc/init.d/

# Path path for documentation
DOCPATH=docs/$(PKGNAME)/

# Path path for man pages
MANPATH=usr/share/man/man1/

# Installation path for the binary files
PATHINSTBIN=./$(BINPATH)

# Installation path for the configuration files
PATHINSTCFG=./$(CONFIGPATH)

# Installation path for the ssl root certs
PATHINSTSSLCFG=./$(SSLCONFIGPATH)

# Installation path for the init file
PATHINSTINIT=./$(INITPATH)

# Installation path for documentation
PATHINSTDOC=./$(DOCPATH)

# Installation path for man pages
PATHINSTMAN=./$(MANPATH)

# RPM Packaging path (where RPMs will be stored)
PATHRPMPKG=./target/bin/RPM

# DEB Packaging path (where DEBs will be stored)
PATHDEBPKG=./target/bin/DEB

# BZ2 Packaging path (where BZ2s will be stored)
PATHBZ2PKG=./target/bin/BZ2

# DOCKER Packaging path (where BZ2s will be stored)
PATHDOCKERPKG=./target/DOCKER

# Cross compilation targets
CCTARGETS=darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm openbsd/386 openbsd/amd64 windows/386 windows/amd64

# docker image name for consul (used during testing)
CONSUL_DOCKER_IMAGE_NAME=consul_$(VENDOR)_$(PROJECT)$(DOCKERSUFFIX)

# STATIC is a flag to indicate whether to build using static or dynamic linking
ifeq ($(STATIC),0)
	STATIC_TAG=dynamic
	STATIC_FLAG=
else
	STATIC_TAG=static
	STATIC_FLAG=-static
endif

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "$(PROJECT) Makefile."
	@echo "GOPATH=$(GOPATH)"
	@echo "The following commands are available:"
	@echo ""
	@echo "    make qa          : Run all the tests and static analysis reports"
	@echo "    make test        : Run the unit tests"
	@echo ""
	@echo "    make format      : Format the source code"
	@echo "    make fmtcheck    : Check if the source code has been formatted"
	@echo "    make confcheck   : Check the JSON configuration files against the schema"
	@echo "    make vet         : Check for suspicious constructs"
	@echo "    make lint        : Check for style errors"
	@echo "    make coverage    : Generate the coverage report"
	@echo "    make cyclo       : Generate the cyclomatic complexity report"
	@echo "    make ineffassign : Detect ineffectual assignments"
	@echo "    make misspell    : Detect commonly misspelled words in source files"
	@echo "    make structcheck : Find unused struct fields"
	@echo "    make varcheck    : Find unused global variables and constants"
	@echo "    make errcheck    : Check that error return values are used"
	@echo "    make staticcheck : Suggest code simplifications"
	@echo "    make astscan     : GO AST scanner"
	@echo ""
	@echo "    make docs        : Generate source code documentation"
	@echo ""
	@echo "    make deps        : Get the dependencies"
	@echo "    make build       : Compile the application"
	@echo "    make clean       : Remove any build artifact"
	@echo "    make nuke        : Deletes any intermediate file"
	@echo "    make install     : Install this application"
	@echo ""
	@echo "    make rpm         : Build an RPM package"
	@echo "    make deb         : Build a DEB package"
	@echo "    make bz2         : Build a tar bz2 (tbz2) compressed archive"
	@echo "    make docker      : Build a scratch docker container to run this service"
	@echo "    make dockertest  : Test the newly built docker container"
	@echo ""
	@echo "    make buildall    : Full build and test sequence"
	@echo "    make dbuild      : Build everything inside a Docker container"
	@echo ""

# Alias for help target
all: help

prepare: tools
	@echo "Installing Tools..."
ifeq (,$(wildcard ${GOPATH}/bin/go-ethereum))
	@echo "Installing go-ethereum..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/ethereum/go-ethereum
else
	@echo "  >	go-ethereum is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/dep))
	@echo "Installing dep..."
	@curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-${GOHOSTOS}-${GOHOSTARCH} -o ${GOPATH}/bin/dep
	@chmod a+x ${GOPATH}/bin/dep
else
	@echo "  >	dep is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/go-enum))
	@echo "Installing go-enum..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/abice/go-enum
else
	@echo "  >	go-enum is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/swagger))
	@echo "Installing swagger..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/go-swagger/go-swagger/cmd/swagger
else
	@echo "  >	swagger is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/swag))
	@echo "Installing swag..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/swaggo/swag/cmd/swag
else
	@echo "  >	swag is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/go-callvis))
	@echo "Installing go-callvis..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/TrueFurby/go-callvis
else
	@echo "  >	go-callvis is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/go-junit-report))
	@echo "Installing go-junit-report..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/jstemmer/go-junit-report
else
	@echo "  >	go-junit-report is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/go-bindata))
	@echo "Installing go-bindata..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/kevinburke/go-bindata/go-bindata
else
	@echo "  >	go-bindata is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/goimports))
	@echo "Installing goimports..."
	@GO111MODULE=$(GO111MODULE) $(GO) install golang.org/x/tools/cmd/goimports
else
	@echo "  >	goimports is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/goveralls))
	@echo "Installing goveralls..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/mattn/goveralls
else
	@echo "  >	goveralls is already installed"
endif
ifeq (,$(wildcard ${GOPATH}/bin/mockgen))
	@echo "Installing mockgen..."
	@GO111MODULE=$(GO111MODULE) $(GO) install github.com/golang/mock/mockgen
else
	@echo "  >	mockgen is already installed"
endif
	@echo "Installing Tools Finished..."

# Run the unit tests
.PHONY: test
test:
	@mkdir -p target/test
	@mkdir -p target/report
	GO111MODULE=$(GO111MODULE) $(GO) test \
	-tags ${STATIC_TAG} \
	-covermode=atomic \
	-bench=. \
	-race \
	-cpuprofile=target/report/cpu.out \
	-memprofile=target/report/mem.out \
	-mutexprofile=target/report/mutex.out \
	-coverprofile=target/report/coverage.out \
	-v ./... | \
	tee >(PATH=$(GOPATH)/bin:$(PATH) go-junit-report > target/test/report.xml); \
	test $${PIPESTATUS[0]} -eq 0

# Format the source code
.PHONY: format
format:
	@echo Running GOFMT
	@for package in $(shell go list ./...); do \
		echo "Checking "$$package; \
		files=$$($(GO) list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(gofmt -d -s $$files 2>&1); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failure"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "gofmt success"; \

.PHONY: style
style:
	@echo ">> checking code style"
	@fmtRes=$$($(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print)); \
	if [ -n "$${fmtRes}" ]; then \
		echo "gofmt checking failed!"; echo "$${fmtRes}"; echo; \
		echo "Please ensure you are using $$($(GO) version) for formatting code."; \
		exit 1; \
	fi

# Check if the source code has been formatted
.PHONY: fmtcheck
fmtcheck:
	@mkdir -p target
	@find . -type f -name "*.go" -exec gofmt -s -d {} \; | tee target/format.diff
	@test ! -s target/format.diff || { echo "ERROR: the source code has not been formatted - please use 'make format' or 'gofmt'"; exit 1; }

# Validate JSON configuration files against the JSON schema
.PHONY: confcheck
confcheck:
	json validate --schema-file=resources/etc/${PROJECT}/config.schema.json --document-file=resources/test/etc/${PROJECT}/config.json
	json validate --schema-file=resources/etc/${PROJECT}/config.schema.json --document-file=resources/etc/${PROJECT}/config.json

# Check for syntax errors
.PHONY: vet
vet:
	@GO111MODULE=$(GO111MODULE) $(GO) vet ./...

# Check for style errors
.PHONY: lint
lint:
	GOPATH=$(GOPATH) PATH=$(GOPATH)/bin:$(PATH) golint ./...

# Generate the coverage report
.PHONY: coverage
coverage:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) $(GO) tool cover -html=target/report/coverage.out -o target/report/coverage.html

# Report cyclomatic complexity
.PHONY: cyclo
cyclo:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) gocyclo -avg . | tee target/report/cyclo.txt ; test $${PIPESTATUS[0]} -eq 0

# Detect ineffectual assignments
.PHONY: ineffassign
ineffassign:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) ineffassign . tee | tee ./target/report/ineffassign.txt ; test $${PIPESTATUS[0]} -eq 0

# Detect commonly misspelled words in source files
.PHONY: misspell
misspell:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) misspell -error ./... | tee ./target/report/misspell.txt ; test $${PIPESTATUS[0]} -eq 0

# Find unused struct fields.
.PHONY: structcheck
structcheck:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) structcheck -a ./... | tee ./target/report/structcheck.txt

# Find unused global variables and constants.
.PHONY: varcheck
varcheck:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) varcheck -e ./... | target/report/varcheck.txt

# Check that error return values are used.
.PHONY: errcheck
errcheck:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) errcheck ./... | tee target/report/errcheck.txt

# Suggest code simplifications"
.PHONY: staticcheck
staticcheck:
	@mkdir -p target/report
	@GO111MODULE=$(GO111MODULE) staticcheck ./... | tee target/report/staticcheck.txt

# AST scanner
.PHONY: astscan
astscan:
	@mkdir -p ./target/report
	@GO111MODULE=$(GO111MODULE) gosec ./... | tee ./target/report/astscan.txt

# Generate source docs
.PHONY: docs
docs:
	@mkdir -p target/docs
	nohup sh -c 'GOPATH=$(GOPATH) godoc -http=127.0.0.1:6060' > target/godoc_server.log 2>&1 &
	wget --directory-prefix=target/docs/ --execute robots=off --retry-connrefused --recursive --no-parent --adjust-extension --page-requisites --convert-links http://127.0.0.1:6060/pkg/github.com/${VENDOR}/${PROJECT}/ ; kill -9 `lsof -ti :6060`
	@echo '<html><head><meta http-equiv="refresh" content="0;./127.0.0.1:6060/pkg/'${CVSPATH}'/'${PROJECT}'/index.html"/></head><a href="./127.0.0.1:6060/pkg/'${CVSPATH}'/'${PROJECT}'/index.html">'${PKGNAME}' Documentation ...</a></html>' > target/docs/index.html

# Alias to run targets: fmtcheck test vet lint coverage
.PHONY: qa
qa: fmtcheck confcheck test vet lint coverage cyclo ineffassign misspell structcheck varcheck errcheck staticcheck astscan

# --- INSTALL ---

# Get the dependencies
.PHONY: deps
deps:
	@echo ">> getting dependencies"
ifdef GO111MODULE
	GO111MODULE=$(GO111MODULE) $(GO) mod download
else
	$(GO) get $(GOOPTS) -t ./...
endif

# Install this application
.PHONY: install
install: uninstall
	mkdir -p ./target/${BINPATH}
	cp -r ./target/${BINPATH}* $(PATHINSTBIN)
	find $(PATHINSTBIN) -type d -exec chmod 755 {} \;
	find $(PATHINSTBIN) -type f -exec chmod 755 {} \;
	mkdir -p $(PATHINSTDOC)
	cp -f ./LICENSE $(PATHINSTDOC)
	cp -f ./README.md $(PATHINSTDOC)
	cp -f ./CONFIG.md $(PATHINSTDOC)
	cp -f ./VERSION $(PATHINSTDOC)
	cp -f ./RELEASE $(PATHINSTDOC)
	chmod -R 644 $(PATHINSTDOC)*
ifneq ($(strip $(INITPATH)),)
	mkdir -p $(PATHINSTINIT)
	cp -ru ./resources/${INITPATH}* $(PATHINSTINIT)
	find $(PATHINSTINIT) -type d -exec chmod 755 {} \;
	find $(PATHINSTINIT) -type f -exec chmod 755 {} \;
endif
ifneq ($(strip $(CONFIGPATH)),)
	mkdir -p $(PATHINSTCFG)
	touch -c $(PATHINSTCFG)*
	cp -ru ./resources/${CONFIGPATH}* $(PATHINSTCFG)
	find $(PATHINSTCFG) -type d -exec chmod 755 {} \;
	find $(PATHINSTCFG) -type f -exec chmod 644 {} \;
endif
ifneq ($(strip $(MANPATH)),)
	mkdir -p $(PATHINSTMAN)
	cat ./resources/${MANPATH}${PROJECT}.1 | gzip -9 > $(PATHINSTMAN)${PROJECT}.1.gz
	find $(PATHINSTMAN) -type f -exec chmod 644 {} \;
endif

# Install SSL certificates
.PHONY: installssl
installssl:
ifneq ($(strip $(SSLCONFIGPATH)),)
	mkdir -p $(PATHINSTSSLCFG)
	touch -c $(PATHINSTSSLCFG)*
	cp -ru ./resources/${SSLCONFIGPATH}* $(PATHINSTSSLCFG)
	find $(PATHINSTSSLCFG) -type d -exec chmod 755 {} \;
	find $(PATHINSTSSLCFG) -type f -exec chmod 644 {} \;
endif

# Remove all installed files (excluding configuration files)
.PHONY: uninstall
uninstall:
	rm -rf ./target/$(BINPATH)/$(PROJECT)
	rm -rf $(PATHINSTDOC)

# Remove any build artifact
.PHONY: clean
clean:
	@GO111MODULE=$(GO111MODULE) $(GO) clean ./...

# Deletes any intermediate file
.PHONY: nuke
nuke:
	rm -rf ./target
	@GO111MODULE=$(GO111MODULE) $(GO) clean -i ./...

# Compile the application
.PHONY: build
build: deps
	GOPATH=$(GOPATH) \
	GO111MODULE=$(GO111MODULE) $(GO) build \
	-tags ${STATIC_TAG} \
	-ldflags '-linkmode external -extldflags ${STATIC_FLAG} -w -s -X main.ProgramVersion=${VERSION} -X main.ProgramRelease=${RELEASE}' \
	-o ./target/${BINPATH}$(PROJECT) .
ifneq (${UPXENABLED},)
	upx --brute ./target/${BINPATH}$(PROJECT)
endif

# Cross-compile the application for several platforms
.PHONY: crossbuild
crossbuild: deps
	@echo "" > target/ccfailures.txt
	$(foreach TARGET,$(CCTARGETS), \
		$(eval GOOS = $(word 1,$(subst /, ,$(TARGET)))) \
		$(eval GOARCH = $(word 2,$(subst /, ,$(TARGET)))) \
		$(shell which mkdir) -p target/$(TARGET) && \
		GOOS=${GOOS} \
		GOARCH=${GOARCH} \
		GOPATH=$(GOPATH) \
		GO111MODULE=$(GO111MODULE) $(GO) build \
		-tags ${STATIC_TAG} \
		-ldflags '-s -extldflags ${STATIC_FLAG} -w -s -X main.ProgramVersion=${VERSION} -X main.ProgramRelease=${RELEASE}' \
		-o ./target/${GOOS}/${GOARCH}/$(PROJECT) . \
		|| echo $(TARGET) >> target/ccfailures.txt ; \
	)
ifneq ($(strip $(cat target/ccfailures.txt)),)
	echo target/ccfailures.txt
	exit 1
endif

# --- PACKAGING ---

# Build the RPM package for RedHat-like Linux distributions
.PHONY: rpm
rpm:
	rm -rf $(PATHRPMPKG)
	rpmbuild \
	--define "_topdir $(PATHRPMPKG)" \
	--define "_vendor $(VENDOR)" \
	--define "_owner $(OWNER)" \
	--define "_project $(PROJECT)" \
	--define "_package $(PKGNAME)" \
	--define "_version $(VERSION)" \
	--define "_release $(RELEASE)" \
	--define "_current_directory $(CURRENTDIR)" \
	--define "_binpath /$(BINPATH)" \
	--define "_docpath /$(DOCPATH)" \
	--define "_configpath /$(CONFIGPATH)" \
	--define "_initpath /$(INITPATH)" \
	--define "_manpath /$(MANPATH)" \
	-bb resources/rpm/rpm.spec

# Build the DEB package for Debian-like Linux distributions
.PHONY: deb
deb:
	rm -rf $(PATHDEBPKG)
	make install DESTDIR=$(PATHDEBPKG)/$(PKGNAME)-$(VERSION)
	rm -f $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/$(DOCPATH)LICENSE
	tar -zcvf $(PATHDEBPKG)/$(PKGNAME)_$(VERSION).orig.tar.gz -C $(PATHDEBPKG)/ $(PKGNAME)-$(VERSION)
	cp -rf ./resources/debian $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian
	mkdir -p $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/missing-sources
	echo "// fake source for lintian" > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/missing-sources/$(PROJECT).c
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#DATE#~/`date -R`/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#PKGNAME#~/$(PKGNAME)/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#VERSION#~/$(VERSION)/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#RELEASE#~/$(RELEASE)/" {} \;
	echo $(BINPATH) > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(BINPATH)* $(BINPATH)" > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
	echo $(DOCPATH) >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(DOCPATH)* $(DOCPATH)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
ifneq ($(strip $(INITPATH)),)
	echo $(INITPATH) >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(INITPATH)* $(INITPATH)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
endif
ifneq ($(strip $(CONFIGPATH)),)
	echo $(CONFIGPATH) >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(CONFIGPATH)* $(CONFIGPATH)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
endif
ifneq ($(strip $(MANPATH)),)
	echo $(MANPATH) >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(MANPATH)* $(MANPATH)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
endif
	echo "statically-linked-binary usr/bin/$(PROJECT)" > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).lintian-overrides
	echo "new-package-should-close-itp-bug" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).lintian-overrides
	echo "hardening-no-relro $(BINPATH)$(PROJECT)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).lintian-overrides
	echo "embedded-library $(BINPATH)$(PROJECT): libyaml" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).lintian-overrides
	cd $(PATHDEBPKG)/$(PKGNAME)-$(VERSION) && debuild -us -uc

# Build a compressed bz2 archive
.PHONY: bz2
bz2:
	rm -rf $(PATHBZ2PKG)
	make install DESTDIR=$(PATHBZ2PKG)
	tar -jcvf $(PATHBZ2PKG)/$(PKGNAME)-$(VERSION)-$(RELEASE).tbz2 -C $(PATHBZ2PKG) usr/ etc/

# Build a docker container to run this service
.PHONY: docker
docker:
	rm -rf $(PATHDOCKERPKG)
	make install DESTDIR=$(PATHDOCKERPKG)
	make installssl DESTDIR=$(PATHDOCKERPKG)
	cp resources/DockerDeploy/Dockerfile $(PATHDOCKERPKG)/Dockerfile
	docker build --no-cache --tag=${VENDOR}/${PROJECT}$(DOCKERSUFFIX):latest $(PATHDOCKERPKG)
	docker save --output=$(PATHDOCKERPKG)/${VENDOR}-${PROJECT}$(DOCKERSUFFIX)-$(VERSION)-$(RELEASE) ${VENDOR}/${PROJECT}$(DOCKERSUFFIX):latest

# Check if the deployment container starts
.PHONY: dockertest
dockertest:
	# clean any previous container (if any)
	rm -f target/old_docker_containers.id
	docker ps -a | grep $(CONSUL_DOCKER_IMAGE_NAME) | awk '{print $$1}' >> target/old_docker_containers.id
	docker ps -a | grep ${VENDOR}/${PROJECT}$(DOCKERSUFFIX) | awk '{print $$1}' >> target/old_docker_containers.id
	docker stop `cat target/old_docker_containers.id` 2> /dev/null || true
	docker rm `cat target/old_docker_containers.id` 2> /dev/null || true
	# start Consul service inside a Docker container
	docker run --detach=true --name=$(CONSUL_DOCKER_IMAGE_NAME) --publish=8500 --hostname=test.consul progrium/consul -server -bootstrap > target/consul_docker_container.id
	sleep 3
	# push Consul configuration
	docker inspect --format='{{(index (index .NetworkSettings.Ports "8500/tcp") 0).HostPort}}' `cat target/consul_docker_container.id` > target/consul_docker_container.port
	curl --request PUT --data @resources/test/etc/rndpwd/config.json http://127.0.0.1:`cat target/consul_docker_container.port`/v1/kv/config/rndpwd
	# Run the program container
	docker run --detach=true --net="host" --tty=true \
	--env="RNDPWD_REMOTECONFIGPROVIDER=consul" \
	--env="RNDPWD_REMOTECONFIGENDPOINT=127.0.0.1:`cat target/consul_docker_container.port`" \
	--env="RNDPWD_REMOTECONFIGPATH=/config/rndpwd" \
	--env="RNDPWD_REMOTECONFIGSECRETKEYRING=" \
	${VENDOR}/${PROJECT}$(DOCKERSUFFIX):latest > target/project_docker_container.id || true
	sleep 3
	# check if the container is working
	docker inspect -f {{.State.Running}} `cat target/project_docker_container.id` > target/project_docker_container.run || true
	# remove containers
	docker stop `cat target/project_docker_container.id` || true
	docker logs `cat target/project_docker_container.id` || true
	docker rm `cat target/project_docker_container.id` || true
	docker stop `cat target/consul_docker_container.id` || true
	docker rm `cat target/consul_docker_container.id` || true
	@exit `grep -ic "false" target/project_docker_container.run`

# Full build and test sequence
# You may want to change this and remove the options you don't need
#buildall: deps qa rpm deb bz2 crossbuild
.PHONY: buildall
buildall: build qa rpm deb

# Build everything inside a Docker container
.PHONY: dbuild
dbuild:
	@mkdir -p target
	@rm -rf target/*
	@echo 0 > target/make.exit
	CVSPATH=$(CVSPATH) VENDOR=$(VENDOR) PROJECT=$(PROJECT) MAKETARGET='$(MAKETARGET)' ./dockerbuild.sh
	@exit `cat target/make.exit`

# Upload linux packages to bintray
.PHONY: bintray
bintray: rpm deb
	@curl -T target/RPM/RPMS/x86_64/${VENDOR}-${PROJECT}-${VERSION}-${RELEASE}.x86_64.rpm -u${APIUSER}:${APIKEY} -H "X-Bintray-Package:${PROJECT}" -H "X-Bintray-Version:${VERSION}" -H "X-Bintray-Publish:1" -H "X-Bintray-Override:1" https://api.bintray.com/content/tecnickcom/rpm/~#VENDOR#~-${PROJECT}-${VERSION}-${RELEASE}.x86_64.rpm
	@curl -T target/DEB/${VENDOR}-${PROJECT}_${VERSION}-${RELEASE}_amd64.deb -u${APIUSER}:${APIKEY} -H "X-Bintray-Package:${PROJECT}" -H "X-Bintray-Version:${VERSION}" -H "X-Bintray-Debian-Distribution:amd64" -H "X-Bintray-Debian-Component:main" -H "X-Bintray-Debian-Architecture:amd64" -H "X-Bintray-Publish:1" -H "X-Bintray-Override:1" https://api.bintray.com/content/tecnickcom/deb/~#VENDOR#~-${PROJECT}_${VERSION}-${RELEASE}_amd64.deb
