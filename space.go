package confspace

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/otiai10/ternary"
)

// Space ...
type Space struct {
	Name string
	Dir  string
}

// New ...
func New(name string) (*Space, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	dotted := ternary.If(strings.HasPrefix(name, ".")).String(name, "."+name)
	return &Space{
		Name: name,
		Dir:  filepath.Join(u.HomeDir, dotted),
	}, nil
}

// Open creates or opens specified file.
func (s *Space) Open(name string) (*os.File, error) {
	if !strings.HasPrefix(name, s.Dir) {
		name = filepath.Join(s.Dir, name)
	}
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return os.Create(name)
		}
		return nil, err
	}
	return os.Open(name)
}

// DropAll ...
func (s *Space) DropAll() error {
	return os.RemoveAll(s.Dir)
}
