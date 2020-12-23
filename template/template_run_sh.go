package template

var tplRunsh string = `#!/bin/bash

set -e
ROOT="$( cd "$( dirname $0 )/.." && pwd )"
cd $ROOT
. auto/buildinfo

export GO111MODULE=on
export GOPROXY=https://goproxy.cn

cd $ROOT

function pre() {
	if [ ! -f "go.mod" ]; then
		go mod init ProjectHolder
	fi
}


function vet() {
	cd $ROOT
	GOPACKGES=$$$$go list ./... $$$$
	go vet $GOPACKGES 2>&1
}

function unit_tests() {
	cd $ROOT
	test -f coverage.out && rm -f coverage.out >/dev/null 2>&1
	for package in $$$$go list ./... $$$$
	do
		go test -coverprofile=profile.out -covermode=atomic $package
		if [ -f profile.out ]; then
			cat profile.out >> coverage.out
			rm -f profile.out
		fi
	done
}

function gobuild() {
	echo "go build -o $ROOT/bin/$1 $2"
	cd $ROOT
	go build -o $ROOT/bin/$1 $2
}

function install() {
	cd $ROOT
	# TODO
}

function clean() {
	rm -f Makefile
	rm -rf auto
	rm -f coverage.out profile.out
}

function main () {
	if [[ $# < 1 ]]; then
		targets="pre"
	else
		targets=($@)
	fi

	if [ $1 == "gobuild" ]; then
		$1 $2 $3
	else
		for target in ${targets[@]}; do
			$target
		done
	fi
}

main "$@"
`
