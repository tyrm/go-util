package pkgerutil

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCompileTemplates(t *testing.T) {
	templates, err := CompileTemplates("", ".gohtml")
	if err != nil {
		t.Errorf("compiling templates: %s", err.Error())
		return
	}

	tables := []struct {
		// inputs
		template string
		input    interface{}

		// outputs
		output string
		err    error
	}{
		{"static", nil, "Static Text.", nil},
		{"dynamic", "Dynamic Text.", "Dynamic Text.", nil},
		{"dynamic", "<a href=\"/link\">link</a>", "&lt;a href=&#34;/link&#34;&gt;link&lt;/a&gt;", nil},
		{"htmlsafe", "<a href=\"/link\">link</a>", "<a href=\"/link\">link</a>", nil},
		{"dec", 5, "4", nil},
		{"inc", 1, "2", nil},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Executing template %s", i, table.template)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var resp bytes.Buffer
			err := templates.ExecuteTemplate(&resp, table.template, table.input)
			if err != table.err {
				t.Errorf("[%d] template execution error, got: '%v', want: '%v'", i, err, table.err)
				return
			}

			respStr := resp.String()
			if respStr != table.output {
				t.Errorf("[%d] got invalid output for %s, got: '%v', want: '%v'", i, table.template, respStr, table.output)
			}
		})
	}
}
