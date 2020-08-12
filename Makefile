build: check-env
	GITHUB_ORG=$(GITHUB_ORG) go generate -x repos/gen.go
	go build -trimpath

check-env:
ifndef GITHUB_ORG
	$(error GITHUB_ORG is not set)
endif
