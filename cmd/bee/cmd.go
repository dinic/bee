package main

import (
	"github.com/pkg/errors"
)

type Cmd interface {
	Run() error
}

type Creator func([]string) (Cmd, error)

var (
	cmdCreators map[string]Creator
)

func registCreator(cmd string, creator Creator) {
	if len(cmdCreators) == 0 {
		cmdCreators = make(map[string]Creator)
	}

	cmdCreators[cmd] = creator
}

func createCmd(args []string) (Cmd, error) {
	if len(cmdCreators) == 0 {
		return nil, errors.Errorf("fatal no cmd regited")
	}

	if len(args) == 0 {
		return nil, errors.Errorf("error cmd line, useage like \"bee run|check [...]\"")
	}

	creator, ok := cmdCreators[args[0]]
	if !ok {
		return nil, errors.Errorf("cmd: \"%s\" not surppor", args[0])
	}
	return creator(args[1:])
}
