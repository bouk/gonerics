package main

import (
	"bytes"
	"flag"
	"github.com/google/go-github/github"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	headerTemplateSource = `package {{.Package}}
{{range .Imports}}import {{.Alias}} "{{.Path}}"
{{end}}
`
	PARAMETER_DIVIDER = "_"
)

var (
	parameters, templateType, name, templatesDir string

	githubClient   = github.NewClient(&http.Client{})
	headerTemplate = template.New("HeaderTemplate")

	gist     *github.Gist
	gistFile github.GistFile
)

func init() {
	flag.StringVar(&parameters, "parameters", "", "The parameters int_string_io.Reader")
	flag.StringVar(&templateType, "template-type", "", "The template type d/f/g")
	flag.StringVar(&name, "name", "", "The name of the template, map, graph, set")
	flag.StringVar(&templatesDir, "templates-dir", "", "The directory that contains the templates")

	headerTemplate.Parse(headerTemplateSource)
}

func main() {
	log.SetOutput(os.Stderr)
	flag.Parse()

	if parameters == "" || templateType == "" || name == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if templatesDir == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal("Failed to get directory", err)
		}

		templatesDir = filepath.Join(dir, "templates")
	}

	buffer := new(bytes.Buffer)

	templateParameters := Parse(parameters)

	fset := token.NewFileSet()

	switch templateType {
	case "f", "d":
		templateParameters.Package = name
	case "g":
		gist, _, err := githubClient.Gists.Get(name)

		if err != nil {
			log.Fatal("Failed to get gist ", err)
		}

		if len(gist.Files) != 1 {
			log.Fatal("Need exactly one file")
		}
		var key github.GistFilename
		for key, gistFile = range gist.Files {
			break
		}

		if !strings.HasSuffix(string(key), ".go") {
			log.Fatal("The filename needs to end with .go")
		}

		templateParameters.Package = string(key[:len(key)-3])
	default:
		log.Fatal("Unknown template type ", templateType)
	}

	headerTemplate.Execute(buffer, templateParameters)

	switch templateType {
	case "f", "d":
		filename := filepath.Join(templatesDir, templateType, name+".go")
		log.Printf("Using %s as template", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal("Failed to open template file", err)
		}
		io.Copy(buffer, file)
		file.Close()
	case "g":
		buffer.WriteString(*gistFile.Content)
	}

	tree, err := parser.ParseFile(fset, "file.go", buffer, 0)
	if err != nil {
		log.Fatal("Error while parsing template", err)
	}

	if err = rewriteAst(tree, templateParameters); err != nil {
		log.Fatal("Error while rewriting AST", err)
	}
	printer.Fprint(os.Stdout, fset, tree)
}
