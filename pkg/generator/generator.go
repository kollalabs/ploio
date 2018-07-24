package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/ploio/ploio/pkg/application"
)

// Run starts the config generator for k8s manifests
func Run(templateDir string, outputDir string, data *application.Version) error {
	tmpls, err := LoadTemplates(templateDir)
	if err != nil {
		return err
	}
	err = ExecuteTemplates(tmpls, outputDir, data)
	if err != nil {
		return err
	}
	return nil
}

// Template object holds a root Template.Template object and all the associated file names for generation/execution
type Template struct {
	Root      *template.Template
	Filenames []string
}

// LoadTemplates is a helper function that takes a directory and loads all the files with the extension .tmpl
// into a template object. It returns a Generator.Templates object which has the template.TemplateRoot object and all the filenames
// including paths stored in the Generator.Templates.Filenames slice. The Generator.Templates object can then
// be passed to Generator.Execute() method along with your data and output directory and will create executed
// versions of your template files in the same directory structure as they were loaded.
func LoadTemplates(rootDir string) (*Template, error) {

	newTemplate := &Template{}

	var filenames []string
	cleanRoot := filepath.Clean(rootDir)
	fmt.Println("Loading templates from " + cleanRoot)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if e1 != nil {
			return e1
		}
		if !info.IsDir() && strings.HasSuffix(path, ".tmpl") {

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]

			t := root.New(name).Funcs(sprig.TxtFuncMap())
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
			filenames = append(filenames, name)
		}

		return nil
	})
	if err != nil {
		return newTemplate, err
	}
	newTemplate.Root = root
	newTemplate.Filenames = filenames
	return newTemplate, err
}

// ExecuteTemplates is a helper function that will execute a Generator.Template object by generating
// all the necessary files in the same exact directory structure as they were loaded from the
// template directory. It will also replace any directory names that use template variables with
// the values from `data`.
func ExecuteTemplates(t *Template, outputDir string, data interface{}) error {

	for _, tname := range t.Filenames {
		fullPath := strings.TrimSuffix(filepath.Join(outputDir, tname), ".tmpl")
		// Parse directory path for any template variables
		buf := new(bytes.Buffer)
		tmpl, err := template.New("filename").Parse(fullPath)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(buf, data)
		if err != nil {
			panic(err)
		}
		fullPath = buf.String()
		err = os.MkdirAll(filepath.Dir(fullPath), 0777)
		if err != nil {
			return err
		}
		fileWriter, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		err = t.Root.ExecuteTemplate(fileWriter, tname, data)
		if err != nil {
			return err
		}
	}
	return nil
}
