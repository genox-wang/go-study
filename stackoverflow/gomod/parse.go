package gomod

import (
	"golang.org/x/mod/modfile"
	"io/ioutil"
)

func parseGoMod() string {
	buf, err := ioutil.ReadFile("../../go.mod")
	if err != nil {
		panic(err)
	}
	f, err := modfile.Parse("go.mod", buf, nil)
	if err != nil {
		panic(err)
	}
	return f.Go.Version
}
