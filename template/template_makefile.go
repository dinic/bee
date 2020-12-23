package template

var tplMakefile string = `echo "auto makefile ..."
cd $ProjectHolder_ROOT
cat << END > Makefile

default: all

.PHONY: all
all: ProjectHolder

.PHONY: pre
pre:
	@./build/run.sh pre

.PHONY: vet
vet: pre
	@./build/run.sh vet

.PHONY: unit_test
unit_test: pre
	@./build/run.sh unit_tests

.PHONY: ProjectHolder
ProjectHolder: pre
	@./build/run.sh gobuild ProjectHolder ProjectHolder/cmd/ProjectHolder

.PHONY: clean
clean:
	@./build/run.sh clean

.PHONY: install
install:
	@./build/run.sh install
END
`
