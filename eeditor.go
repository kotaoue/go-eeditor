package eeditor

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

type Editor struct {
	path    string
	command string
}

func NewEditor(opts ...Option) *Editor {
	e := &Editor{
		path:    "./tmp",
		command: "vim",
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

type Option func(*Editor)

func Path(s string) Option {
	return func(e *Editor) { e.path = s }
}
func Command(s string) Option {
	return func(e *Editor) { e.command = s }
}

func (e *Editor) Open() ([]byte, error) {
	if err := e.safeCreate(); err != nil {
		return nil, err
	}
	defer os.Remove(e.path)

	if err := e.run(); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(e.path)
}

func (e *Editor) safeCreate() error {
	_, err := os.Stat(e.path)
	if !os.IsNotExist(err) {
		return errors.New("specified path is already exists")
	}

	_, err = os.Create(e.path)
	return err
}

func (e *Editor) run() error {
	c := exec.Command(e.command, e.path)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
