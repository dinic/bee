package main

import (
	"bee/project"

	"github.com/pkg/errors"
)

type CreateCmd struct {
	dir     string
	project string
}

func newCreateCmd(args []string) (Cmd, error) {
	if len(args) == 0 {
		return nil, errors.Errorf("useage: bee create project [dir]")
	}

	cmd := new(CreateCmd)
	cmd.project = args[0]

	if len(args) >= 2 {
		cmd.dir = args[1]
	}

	return cmd, nil
}

func init() {
	registCreator("create", newCreateCmd)
}

func (c *CreateCmd) Run() error {
	p, err := project.CreateProject(c.dir, c.project)
	if err != nil {
		return err
	}

	if err = p.Gen(); err != nil {
		return err
	}

	return nil
}
