.PHONY: default init install test release ci
.DEFAULT_GOAL=default

default:
	@ mmake help

# init project
init:
	@ scripts/setup

# install deps
install:
	@ go mod vendor

# test package
test:
	go test ./...

# release package
release: test
		$(eval VERSION=$(filter-out $@, $(MAKECMDGOALS)))
		$(if ${VERSION},@true,$(error "VERSION is required"))
		git commit --allow-empty -am ${VERSION}
		git push
		hub release create -m ${VERSION} -e ${VERSION}

ci: test
	${HOME}/gopath/bin/goveralls -service=travis-ci

%:
	@ true
