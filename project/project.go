package project

import (
	"os"
	"path/filepath"

	"bee/template"

	"github.com/pkg/errors"
)

type Project struct {
	dir        string
	name       string
	dirProject string
	dirCmd     string
	dirBin     string
	dirLogs    string
	dirPkg     string
	dirConf    string
	dirBuild   string
}

func CreateProject(dir, name string) (*Project, error) {
	p := new(Project)

	p.dir = dir
	p.name = name

	if err := p.init(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Project) init() error {
	if p.dir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		p.dir = dir
	}

	if !filepath.IsAbs(p.dir) {
		abs, err := filepath.Abs(p.dir)
		if err != nil {
			return err
		}

		p.dir = abs
	}

	p.dirProject = filepath.Join(p.dir, p.name)
	p.dirBin = filepath.Join(p.dirProject, "bin")
	p.dirLogs = filepath.Join(p.dirProject, "logs")
	p.dirPkg = filepath.Join(p.dirProject, "pkg")
	p.dirConf = filepath.Join(p.dirProject, "conf")
	p.dirCmd = filepath.Join(p.dirProject, "cmd", p.name)
	p.dirBuild = filepath.Join(p.dirProject, "build")

	return nil
}

func (p *Project) Gen() error {
	dirs := []string{p.dirProject, p.dirBin, p.dirCmd, p.dirPkg, p.dirConf, p.dirLogs, p.dirBuild}

	files := []struct {
		dir  string
		file string
		tpl  string
		perm os.FileMode
	}{
		{p.dirCmd, p.name + ".go", template.TplMain, 0644},
		{p.dirBuild, "config", template.TplConfig, 0644},
		{p.dirBuild, "options", template.TplOption, 0644},
		{p.dirBuild, "makefile", template.TplMakefile, 0644},
		{p.dirBuild, "run.sh", template.TplRunsh, 0755},
		{p.dirProject, "configure", template.TplConfigure, 0744},
		{p.dirProject, ".gitignore", template.TplGitignore, 0644},
		{p.dirProject, "VERSION", template.TplVersion, 0644},
		{p.dirProject, ".golangci.yml", template.TplGolanglint, 0644},
	}

	for _, dir := range dirs {
		if err := p.genProjectDir(dir); err != nil {
			return err
		}
	}

	for _, file := range files {
		if err := p.genProjectFile(file.dir, file.file, file.tpl, file.perm); err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) genProjectDir(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		if os.IsExist(err) {
			return nil
		}

		return errors.Errorf("mkdir \"%s\" faild, err: %s", dir, err)
	}

	return nil
}

func (p *Project) genProjectFile(dir, file, tpl string, perm os.FileMode) error {
	if err := p.genProjectDir(dir); err != nil {
		return err
	}

	path := filepath.Join(dir, file)
	content, err := template.Gen(p.name, tpl)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR|os.O_EXCL, perm)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		return errors.Errorf("write file: \"%s\", faild, err: %s", path, err)
	}

	return nil
}
