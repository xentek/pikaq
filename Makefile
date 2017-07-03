.PHONY: setup install upgrade test release

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

# release package
release:
	scripts/release $(filter-out $@, $(MAKECMDGOALS))
