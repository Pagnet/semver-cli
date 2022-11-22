GOCMD:=go
GOBUILD:=$(GOCMD) build
OUTDIR:=out
VERSION:=$(shell git ls-remote --refs --sort="version:refname" --tags | cut -d/ -f3- |tail -n1 | xargs -I{} echo '{} + 1' | bc)
TEXT:=$(shell git log -1 --pretty=%B)
NAME:=semver-cli
MAIN:=main.go

ifeq ($(VERSION),)
VERSION=0.0.0
endif

LDFLAGS=-ldflags "-X main.version=$(VERSION)"

build:
	@rm -fr $(OUTDIR)
	@mkdir -p $(OUTDIR)
	@echo "Building $(NAME).rpi version $(VERSION)"
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).rpi $(MAIN)
	@echo "Building $(NAME).lin version $(VERSION)"
	GOOS=linux $(GOBUILD)  ${LDFLAGS} -o $(OUTDIR)/$(NAME).lin $(MAIN)
	@echo "Building $(NAME).mac version $(VERSION)"
	GOOS=darwin $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).mac $(MAIN)
	zip $(OUTDIR)/$(NAME).rpi.zip $(OUTDIR)/$(NAME).rpi 
	zip $(OUTDIR)/$(NAME).lin.zip $(OUTDIR)/$(NAME).lin
	zip $(OUTDIR)/$(NAME).mac.zip $(OUTDIR)/$(NAME).mac

markdown:
	@echo "Building markdown docs"
	@rm -fr docs
	@mkdir -p docs
	@go run main.go markdown > docs/README.md

release:
	 go run main.go ghrelease --repos=semver-cli --branch=main --tag=$(VERSION) --orgname=Pagnet --text="$(TEXT)" --files="$(shell ls -m out/*.zip)"

all: build release

.PHONY: build markdown release