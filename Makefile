.PHONY: setup install test release ci

%:
	@:

# setup project
setup:
	@ scripts/setup

# install deps
install:
	@ go mod vendor

# test package
test:
	go test ./...

ci: test
	${HOME}/gopath/bin/goveralls -service=travis-ci

# release package
release:
	scripts/release $(filter-out $@, $(MAKECMDGOALS))
