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
func CompileTemplates(dir string, suffix string) (*template.Template, error) {
	tpl := template.New("")

	tpl.Funcs(template.FuncMap{
		"dec": func(i int) int {
			i--
			return i
		},
		"htmlSafe": func(html string) template.HTML {
			/* #nosec G203 -- command is meant to display html as is */
			return template.HTML(html)
		},
		"inc": func(i int) int {
			i++
			return i
		},
	})

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
