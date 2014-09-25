package main

import (
	"reflect"
	"testing"
)

var (
	inputs = []struct {
		Params, RepoName string
		Expected         *Params
	}{
		{
			"string",
			&Params{
				Types:    []string{"string"},
				RawTypes: []string{"string"},
			},
		},
		{
			"*string",
			&Params{
				Types:    []string{"*string"},
				RawTypes: []string{"*string"},
			},
		},
		{
			"string_int",
			&Params{
				Types:    []string{"string", "int"},
				RawTypes: []string{"string", "int"},
			},
		},
		{
			"string_io.Reader",
			&Params{
				Types:    []string{"string", "dep1.Reader"},
				RawTypes: []string{"string", "io.Reader"},
				Imports:  []Import{{Alias: "dep1", Path: "io"}},
			},
		},
		{
			"string_io.Reader_net/http.Client",
			&Params{
				Types:    []string{"string", "dep1.Reader", "dep2.Client"},
				RawTypes: []string{"string", "io.Reader", "net/http.Client"},
				Imports:  []Import{{Alias: "dep1", Path: "io"}, {Alias: "dep2", Path: "net/http"}},
			},
		},
		{
			"string_gonerics.io/d/set/string/wow.git.Set",
			&Params{
				Types:    []string{"string", "dep1.Set"},
				RawTypes: []string{"string", "gonerics.io/d/set/string/wow.git.Set"},
				Imports:  []Import{{Alias: "dep1", Path: "gonerics.io/d/set/string/wow.git"}},
			},
		},
		{
			"string_gonerics.io/d/set/string/wow.git.**Set",
			&Params{
				Types:    []string{"string", "**dep1.Set"},
				RawTypes: []string{"string", "gonerics.io/d/set/string/wow.git.**Set"},
				Imports:  []Import{{Alias: "dep1", Path: "gonerics.io/d/set/string/wow.git"}},
			},
		},
	}
)

func TestParseParams(t *testing.T) {
	for _, test := range inputs {
		result := Parse(test.Param)

		if !reflect.DeepEqual(test.Expected, result) {
			t.Error("Failed", test.Params, test.Expected, result)
		}
	}
}
