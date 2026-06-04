CDKTF_CLI_VER := 0.21.0
ROOT_MODULE   := github.com/unimelb-platform/cdktf-provider-awscc-go

# Discover every direct sub-directory of awscc/ and derive their go.mod paths.
SUBMOD_DIRS   := $(wildcard awscc/*/)
SUBMOD_GOMODS := $(SUBMOD_DIRS:%/=%/go.mod)

.PHONY: all generate submodules

## Default target: generate bindings then stamp+tidy every sub-package go.mod.
all: generate submodules

## Run cdktf get to (re)generate the provider bindings under awscc/.
generate:
	npx cdktf-cli@$(CDKTF_CLI_VER) get

## Ensure every sub-package go.mod is up-to-date relative to go.mod.tmpl.
submodules: $(SUBMOD_GOMODS)

## jsii has no intra-repo imports so it only needs the template.
awscc/jsii/go.mod: go.mod.tmpl
	@{ printf 'module $(ROOT_MODULE)/awscc/jsii\n\n'; cat $<; } > $@
	cd awscc/jsii && go mod tidy

## All other sub-packages import awscc/jsii, so they depend on its go.mod
## existing first (via the replace directive) before go mod tidy runs.
awscc/%/go.mod: go.mod.tmpl awscc/jsii/go.mod
	@{ printf 'module $(ROOT_MODULE)/awscc/$*\n\n'; cat $<; } > $@
	@printf '\nreplace $(ROOT_MODULE)/awscc/jsii => ../jsii\n' >> $@
	cd awscc/$* && go mod tidy
