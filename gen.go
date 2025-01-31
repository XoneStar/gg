package gg

import (
	"fmt"
	"io"
	"os"
)

type Generator struct {
	g *Group
}

// New will create a new generator which hold the Group reference.
func New() *Generator {
	return &Generator{g: NewGroup()}
}

func (g *Generator) NewGroup() (ng *Group) {
	ng = NewGroup()
	g.g.append(ng)
	return ng
}

// Write will write the Group into the given writer.
func (g *Generator) Write(w io.Writer) {
	g.g.render(w)
}

// WriteFile will write the Group into the given path.
func (g *Generator) WriteFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %s", path, err)
	}
	g.g.render(file)
	return nil
}

// AppendFile will append the Group after the give path.
func (g *Generator) AppendFile(path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("create file %s: %s", path, err)
	}
	g.g.render(file)
	return nil
}
