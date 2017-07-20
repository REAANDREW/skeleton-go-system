USERNAME="reaandrew"
PROJECT="skeleton-go-system"
GITHUB_TOKEN=$$GITHUB_TOKEN
VERSION=`cat VERSION`
BUILD_TIME=`date +%FT%T%z`
COMMIT_HASH=`git rev-parse HEAD`
DIST_NAME_CONVENTION="dist/{{.OS}}_{{.Arch}}_{{.Dir}}"

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
SOURCES += VERSION
# This is how we want to name the binary output
BINARY=${PROJECT}

# These are the values we want to pass for Version and BuildTime

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.CommitHash=${COMMIT_HASH} -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): deps $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} 

.PHONY: deps 
deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -t ./...

.PHONY: deploy-deps
deploy-deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/tcnksm/ghr
	go get -u github.com/mattn/goveralls

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: lint
lint:
	gometalinter --concurrency=4

.PHONY: test
test:
	go test -cover -coverprofile=coverage.out

.PHONY: coverage-report
coverage-report: 
	go tool cover -html=c.out -o coverage.html

.PHONY: cross-platform-compile
cross-platform-compile: deploy-deps
	gox -output ${DIST_NAME_CONVENTION} ${LDFLAGS}

.PHONY: upload-release
upload-release:
	ghr -username ${USERNAME} -token ${GITHUB_TOKEN} --delete ${VERSION} dist/

.PHONY: ensure-version-increment
ensure-version-increment:
	echo "Check the version is not the same or lower than the last tag"
	echo "If there is no latest tag then this should return no error code"
