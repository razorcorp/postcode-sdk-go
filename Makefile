RELEASE_DATE=$(shell date +"%Y-%m-%d")
OS=$(shell uname)
NL?=\n
SED_OPTS?=-i
BUILD_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
RELEASE_BRANCH=develop
TAG_PREFIX=v

VERSION=

export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION}"

.PHONY: clean

checks:
	$(info Checking required parameters)

ifeq ("${VERSION}", "")
	$(error Undefined: VERSION!)
endif

ifeq ("${OS}", "Darwin")
	$(warning There maybe `CHANGELOG.md.bak` created after release process)
	$(eval NL := $$\\\n)
	$(eval SED_OPTS := "-i .bak")
endif

release_start: checks
	$(info Starting releasing for new version v${VERSION})

ifneq ("${BUILD_BRANCH}", "${RELEASE_BRANCH}")
	$(error Use develop branch to performed a release. Current branch is ${BUILD_BRANCH})
endif

	$(info Creating release branch)
	@git checkout -b release/${VERSION}

update_changelog:
	$(info Updating CHANGELOG.md)
	@sed ${SED_OPTS} -e "/^## \[Unreleased\]/p; s/## \[Unreleased\]/${NL}## \[${VERSION}\] - ${RELEASE_DATE}/" \
		-e "s/^\[Unreleased\]: \(.*\)\/\(.*\)\.\{3\}\(.*\)$$/[Unreleased]: \1\/${VERSION}...HEAD${NL}[${VERSION}]: \1\/\2...${VERSION}/g" \
		-e "s/^\[Unreleased\]: \(.*\)\/$$/[Unreleased]: \1\/${VERSION}...HEAD${NL}[${VERSION}]: \1\/tree\/${VERSION}/g" CHANGELOG.md
	$(info CHANGELOG.md updated!)

release_finish:
	$(info Finishing release process)
	@git add CHANGELOG.md
	@git commit -m "Update CHANGELOG and version"

	@git checkout master
	@git merge --no-ff --no-commit release/${VERSION}
	@git commit -m "Merge release/${VERSION}"
	@git branch -D release/${VERSION}
	@git tag ${TAG_PREFIX}${VERSION}
	@git checkout develop
	@git merge --no-commit master
	$(info Release finished. ${TAG_PREFIX}${VERSION} is now available)

release: release_start update_changelog release_finish clean

test:
	$(error Testing not implemented yet)

clean:
ifneq ("$(wildcard CHANGELOG.md.bak)","")
	$(info Removing CHANGELOG.md.bak file)
	@rm CHANGELOG.md.bak
endif
