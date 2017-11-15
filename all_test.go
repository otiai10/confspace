package confspace

import (
	"fmt"
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	space, _ := New("myproject")
	if err := space.DropAll(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(code)
}

func TestSpace(t *testing.T) {
	space, err := New("myproject")
	Expect(t, err).ToBe(nil)
	Expect(t, space).TypeOf("*confspace.Space")
}

func TestSpace_Open(t *testing.T) {
	space, err := New("myproject")
	Expect(t, err).ToBe(nil)
	f, err := space.Open("hoge/fuga")
	Expect(t, err).ToBe(nil)
	defer f.Close()

	f2, err := space.Open("hoge/fuga")
	Expect(t, err).ToBe(nil)
	defer f2.Close()
}
