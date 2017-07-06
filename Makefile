.PHONY: setup install upgrade test release ci

%:
	@:

# setup project
setup:
	scripts/setup

# install deps
install:
	glide install

# upgrade deps
upgrade:
	glide up

# test package
test:
	go test $(glide nv)

ci: test
	${HOME}/gopath/bin/goveralls -service=travis-ci

# release package
release:
	scripts/release $(filter-out $@, $(MAKECMDGOALS))
