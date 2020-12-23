package template

import (
	"github.com/pkg/errors"
	"strings"
)

const (
	placeHolder   = "$$$$"
	projectHolder = "ProjectHolder"
)

const (
	TplConfig     = "tpl_config"
	TplConfigure  = "tpl_configure"
	TplGitignore  = "tpl_gitignore"
	TplMakefile   = "tpl_makefile"
	TplOption     = "tpl_option"
	TplRunsh      = "tpl_runsh"
	TplMain       = "tpl_main"
	TplVersion    = "tpl_version"
	TplGolanglint = "tpl_golanglint"
)

var (
	templates map[string]string = map[string]string{
		TplConfig:     tplConfig,
		TplConfigure:  tplConfigure,
		TplGitignore:  tplGitignore,
		TplMakefile:   tplMakefile,
		TplOption:     tplOption,
		TplRunsh:      tplRunsh,
		TplMain:       tplMain,
		TplVersion:    tplVersion,
		TplGolanglint: tplGolanglint,
	}
)

func Gen(project, tpl string) (string, error) {
	v, ok := templates[tpl]
	if !ok {
		return "", errors.Errorf("project: \"%s\" not found template: \"%s\"",
			project, tpl)
	}

	v = strings.ReplaceAll(v, placeHolder, "`")
	v = strings.ReplaceAll(v, projectHolder, project)

	return v, nil
}
