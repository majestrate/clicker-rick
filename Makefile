
GIT ?= $(shell which git)

all: update

update:
 $(GIT) pull
