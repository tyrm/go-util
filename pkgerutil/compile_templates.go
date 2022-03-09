package pkgerutil

import (
	"github.com/markbates/pkger"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// CompileTemplates takes a pkger dir and turns it into templates.
// Note: dir should be provided by pkger.Include
func CompileTemplates(dir string, suffix string, funcs *template.FuncMap) (*template.Template, error) {
	tpl := template.New("")

	if funcs != nil {
		tpl.Funcs(*funcs)
	}

	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || !strings.HasSuffix(path, suffix) {
			return nil
		}
		f, err := pkger.Open(path)
		if err != nil {
			return err
		}

		// Now read it.
		sl, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(sl))
		if err != nil {
			return err
		}

		return nil
	})

	return tpl, err
}
