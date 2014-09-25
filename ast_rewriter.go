package main

import (
	"fmt"
	"go/ast"
	"regexp"
)

var (
	letters = []string{
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}
	replaceRegexp = regexp.MustCompile(`\b[A-Z]\b`)
)

type ASTModifier struct {
	mapping    map[string]int
	parameters *Params
}

func NewASTModifier(parameters *Params) (*ASTModifier, error) {
	result := &ASTModifier{
		mapping:    make(map[string]int),
		parameters: parameters,
	}

	for i := 0; i < len(parameters.Types); i++ {
		if i >= len(letters) {
			return nil, fmt.Errorf("Too many parameters given, maximum of %d accepted", len(letters))
		}

		result.mapping[letters[i]] = i
	}

	return result, nil
}

func (a *ASTModifier) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.Ident:
		if i, ok := a.mapping[n.Name]; ok {
			n.Name = a.parameters.Types[i]
		}
	case *ast.ImportSpec:
		n.Path.Value = replaceRegexp.ReplaceAllStringFunc(n.Path.Value, func(s string) string {
			if i, ok := a.mapping[s]; ok {
				return a.parameters.RawTypes[i]
			} else {
				return s
			}
		})
	}

	return a
}

// Note: modifies the ast in-place
func rewriteAst(tree *ast.File, parameters *Params) error {
	astModifier, err := NewASTModifier(parameters)
	if err != nil {
		return err
	}
	ast.Walk(astModifier, tree)
	return nil
}
