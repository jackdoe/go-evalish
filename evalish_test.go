package evalish

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestEval(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	c, err := Compile(
		`
package main

func Sum(a int,b int) int {
    return a + b
}

`, dir, "go")
	if err != nil {
		panic(err)
	}

	f := LookupP(c, "Sum").(func(int, int) int)
	if f(5, 6) != 11 {
		panic("WTF")
	}
}

func TestExample(t *testing.T) {
	dir, err := ioutil.TempDir("", "testing")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	code := `
package main

func Sum(a int,b int) int {
    return a + b
}
`
	f := LookupP(
		CompileP(code, dir, "go"),
		"Sum",
	).(func(int, int) int)

	if f(5, 6) != 11 {
		panic("WTF")
	}

	f2 := LookupP(
		CompileP(code, dir, "go"),
		"Sum",
	).(func(int, int) int)

	if f2(5, 6) != 11 {
		panic("WTF")
	}

}
