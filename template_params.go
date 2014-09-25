package main

import (
	"fmt"
	"strings"
)

type Import struct {
	Alias, Path string
}

type Params struct {
	Package string

	Imports  []Import
	Types    []string
	RawTypes []string
}

func Parse(params string) *Params {
	result := &Params{}

	result.RawTypes = strings.Split(params, PARAMETER_DIVIDER)
	result.Types = make([]string, len(result.RawTypes))

	for i, v := range result.RawTypes {
		splits := strings.Split(v, ".")
		typeName := splits[len(splits)-1]

		pack := strings.Join(splits[:len(splits)-1], ".")

		if pack == "" {
			result.Types[i] = typeName
		} else {
			depName := fmt.Sprintf("dep%d", i)

			trimmedTypeName := strings.TrimLeft(typeName, "*")

			result.Imports = append(result.Imports, Import{Alias: depName, Path: pack})
			result.Types[i] = fmt.Sprintf("%s%s.%s", strings.Repeat("*", len(typeName)-len(trimmedTypeName)), depName, trimmedTypeName)
		}
	}

	return result
}
