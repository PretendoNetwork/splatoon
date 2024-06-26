# TODO - Assumes a UNIX-like OS

RED    := $(shell tput setaf 1)
BLUE   := $(shell tput setaf 4)
CYAN   := $(shell tput setaf 14)
ORANGE := $(shell tput setaf 202)
YELLOW := $(shell tput setaf 214)
RESET  := $(shell tput sgr0)

ifeq ($(shell which go),)
# TODO - Read contents from .git folder instead?
$(error "$(RED)go command not found. Install go to continue $(BLUE)https://go.dev/doc/install$(RESET)")
endif

ifneq ($(wildcard .git),)
# * .git folder exists, build server build string from repo info
ifeq ($(shell which git),)
# TODO - Read contents from .git folder instead?
$(error "$(RED)git command not found. Install git to continue $(ORANGE)https://git-scm.com/downloads$(RESET)")
endif
$(info "$(CYAN)Building server build string from repository info$(RESET)")
# * Build server build string from repo info
BRANCH        := $(shell git rev-parse --abbrev-ref HEAD)
REMOTE_ORIGIN := $(shell git config --get remote.origin.url)

# * Handle multiple origin URL formats
HTTPS_PREFIX_CHECK := $(shell echo $(REMOTE_ORIGIN) | head -c 8)
HTTP_PREFIX_CHECK  := $(shell echo $(REMOTE_ORIGIN) | head -c 7)
GIT@_PREFIX_CHECK  := $(shell echo $(REMOTE_ORIGIN) | head -c 4)

ifeq ($(HTTPS_PREFIX_CHECK), https://)
REMOTE_PATH := $(shell echo $(REMOTE_ORIGIN) | cut -d/ -f4-)
else ifeq ($(HTTP_PREFIX_CHECK), http://)
REMOTE_PATH := $(shell echo $(REMOTE_ORIGIN) | cut -d/ -f4-)
else ifeq ($(GIT@_PREFIX_CHECK), git@)
REMOTE_PATH := $(shell echo $(REMOTE_ORIGIN) | cut -d: -f2-)
else
REMOTE_PATH := $(shell echo $(REMOTE_ORIGIN) | cut -d/ -f2-)
endif

HASH         := $(shell git rev-parse --short HEAD)
SERVER_BUILD := $(BRANCH):$(REMOTE_PATH)@$(HASH)

else
# * .git folder not present, assume downloaded from zip file and just use folder name
$(info "$(CYAN)git repository not found. Building server build string from folder name$(RESET)")
SERVER_BUILD := $(notdir $(CURDIR))
endif

# * Final build string
DATE_TIME    := $(shell date --iso=seconds)
BUILD_STRING := $(SERVER_BUILD), $(DATE_TIME)

all:
ifeq ($(wildcard .env),)
	$(warning "$(YELLOW).env file not found, environment variables may not be populated correctly$(RESET)")
endif
	go get -u
	go mod tidy
	go build -ldflags "-X 'main.serverBuildString=$(BUILD_STRING)'" -o ./build/$(notdir $(CURDIR))