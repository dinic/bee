package template

var tplConfigure string = `#! /bin/sh

cd $$$$dirname $0$$$$

mkdir -p auto

ProjectHolder_ROOT=$$$$pwd$$$$
TMPGOPATH=$$$$echo ${ProjectHolder_ROOT%%/src/github.com*}$$$$

. build/options

ProjectHolder_PREFIX=${ProjectHolder_PREFIX:-$ProjectHolder_ROOT}
ProjectHolder_BIN=${ProjectHolder_BIN:-$ProjectHolder_PREFIX/bin}
ProjectHolder_CONF=${ProjectHolder_CONF:-$ProjectHolder_PREFIX/conf}
. build/config
. build/makefile

export GOPATH=$ProjectHolder_BUILD_DIR
`
