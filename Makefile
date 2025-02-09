# Copyright 2022 The CubeFS Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied. See the License for the specific language governing
# permissions and limitations under the License.

PROJECTDIR=$(shell pwd)
OS=$(shell uname -s)
BINDIR=$(PROJECTDIR)/bin
GCFLAGS=all=-trimpath=$(PROJECTDIR)
ASMFLAGS=all=-trimpath=$(PROJECTDIR)
BRANCHNAME=${branchName}
COMMITID=${commitId}

ifeq ($(BRANCHNAME),)
	BRANCHNAME=$(shell git symbolic-ref --short -q HEAD)
	COMMITID=$(shell git rev-parse --short HEAD)
endif

LDFLAGS=-w -s
ifneq ($(BRANCHNAME),)
    LDFLAGS+= -X "github.com/cubefs/blobstore/common/util/version.version=$(BRANCHNAME)/$(COMMITID)"
endif

BUILD=go build -v -gcflags=$(GCFLAGS) -asmflags=$(ASMFLAGS) -ldflags='$(LDFLAGS)' -o $(BINDIR)
INSTALL=CGO_ENABLED=0 $(BUILD)
CGOINSTALL=CGO_ENABLED=1 $(BUILD)
CMDDIR=github.com/cubefs/blobstore/cmd
TARGETS=blobnode cm allocator access mqproxy scheduler tinker worker cli

.PHONY: clean all $(TARGETS)
all:$(TARGETS)

cm:
	@echo "building clustermgr"
	@$(CGOINSTALL) $(CMDDIR)/clustermgr

blobnode:
	@echo "building blobnode"
	@$(CGOINSTALL) $(CMDDIR)/blobnode

allocator:
	@echo "building allocator"
	@$(INSTALL) $(CMDDIR)/allocator

access:
	@echo "building access"
	@$(INSTALL) $(CMDDIR)/access

mqproxy:
	@echo "building mqproxy"
	@$(INSTALL) $(CMDDIR)/mqproxy

scheduler:
	@echo "building scheduler"
	@$(INSTALL) $(CMDDIR)/scheduler

tinker:
	@echo "building tinker"
	@$(INSTALL) $(CMDDIR)/tinker

worker:
	@echo "building worker"
	@$(INSTALL) $(CMDDIR)/worker

cli:
	@echo "building cli"
	@$(CGOINSTALL) $(PROJECTDIR)/cli/cli

clean:
	@go clean -i ./...
	@rm -f $(BINDIR)/*
