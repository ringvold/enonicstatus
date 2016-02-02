# Adds build information from git repo
#
# Taken from Hugo https://github.com/spf13/hugo/blob/master/Makefile

COMMIT_HASH=`git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE=`date +%FT%T%z`
LDFLAGS=-ldflags "-X github.com/haraldringvold/enonicstatus/cmd.CommitHash=${COMMIT_HASH} -X github.com/haraldringvold/enonicstatus/cmd.BuildDate=${BUILD_DATE}"

all: gitinfo

install: install-gitinfo

help:
	echo ${COMMIT_HASH}
	echo ${BUILD_DATE}

gitinfo:
	go build ${LDFLAGS} -o enonicstatus main.go

install-gitinfo:
	go install ${LDFLAGS} ./...

no-git-info:
	go build -o enonicstatus main.go
