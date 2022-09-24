# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
BINARIES  := codegen
BINARIES_DOCKER = $(patsubst %,docker-%, $(BINARIES))
BINARIES_PUSH = $(patsubst %,push-%, $(BINARIES))
BINARIES_BUMP = $(patsubst %,bump-%, $(BINARIES))

# tools
CP := cp -u -v

# Required for globs to work correctly
SHELL=/usr/bin/env bash

.PHONY: release
release:
	$(eval VERSION_FILE := version.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	git tag -a ${VERSION} -m ${VERSION}
	git push origin ${VERSION}


.PHONY: bump
BUMP := patch
bump:
	@echo "(${BUMP})ing, target: $(patsubst bump-%,%, $@)"

	$(eval VERSION_FILE := version.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))

	@echo "Bumping version.txt from $(VERSION) to $(NEW_VERSION)"
	@echo $(NEW_VERSION) > ${VERSION_FILE}
	@git add ${VERSION_FILE}
	@git commit -vsam "Bump '$(patsubst bump-%,%, $@)' version to $(NEW_VERSION)"
