echo "auto makefile ..."
cd $bee_ROOT
cat << END > Makefile

default: all

.PHONY: all
all: bee

.PHONY: pre
pre:
	@./build/run.sh pre

.PHONY: vet
vet: pre
	@./build/run.sh vet

.PHONY: unit_test
unit_test: pre
	@./build/run.sh unit_tests

.PHONY: bee
bee: pre
	@./build/run.sh gobuild bee bee/cmd/bee

.PHONY: clean
clean:
	@./build/run.sh clean

.PHONY: install
install:
	@./build/run.sh install
END
