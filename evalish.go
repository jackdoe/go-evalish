package evalish

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"plugin"
)

// Mostly taken from https://lemire.me/blog/2021/10/14/calling-a-dynamically-compiled-function-from-go/
// Compile the code, it is not going to recompile if the .so exists
func Compile(code string, root string, gobinary string) (*plugin.Plugin, error) {
	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(code)))
	root = path.Join(root, sum)
	os.MkdirAll(root, 0700)

	soPath := path.Join(root, "code.so")
	codePath := path.Join(root, "code.go")

	if _, err := os.Stat(soPath); errors.Is(err, os.ErrNotExist) {
		err := ioutil.WriteFile(codePath, []byte(code), 0700)
		if err != nil {
			return nil, err
		}

		output, err := exec.Command(gobinary, "build", "-buildmode=plugin", "-o", soPath, codePath).CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("err: %v, output:\n%s", err, output)
		}
	}
	plug, err := plugin.Open(soPath)
	if err != nil {
		return nil, err
	}

	return plug, nil
}

func CompileP(code string, root string, gobinary string) *plugin.Plugin {
	p, err := Compile(code, root, gobinary)
	if err != nil {
		panic(err)
	}
	return p
}

func LookupP(p *plugin.Plugin, sym string) plugin.Symbol {
	f, err := p.Lookup(sym)
	if err != nil {
		panic(err)
	}
	return f
}
