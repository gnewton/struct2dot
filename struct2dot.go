package struct2dot

import (
	"fmt"
	"reflect"
	"strings"
)

var numberTypes = []string{"uint8",
	"uint16",
	"uint32",
	"uint64",
	"int8",
	"int16",
	"int32",
	"int64",
	"float32",
	"float64",
	"complex64",
	"complex128",
	"byte",
	"rune",
}

type StructDriver interface {
	PrintHeader()
	PrintType(t *reflect.Type)
	PrintFooter()
}

type Config struct {
	ShowStrings         bool
	ShowNumbers         bool
	RemovePackagePrefix bool
	IgnoreTypes         []string
}

type DotDriver struct {
	Config      *Config
	edges       map[string]bool
	ignoreTypes map[string]bool
}

func (d *DotDriver) PrintHeader() {
	d.init()

	fmt.Println("digraph foo {")
	fmt.Println("\tnode [color=Red]")
	fmt.Println("\tedge [color=Blue]")
}

func (d *DotDriver) init() {
	d.edges = make(map[string]bool)
	d.ignoreTypes = make(map[string]bool)
	if len(d.Config.IgnoreTypes) > 0 {
		for i, _ := range d.Config.IgnoreTypes {
			d.ignoreTypes[d.Config.IgnoreTypes[i]] = true
		}
	}
	/*
		if !d.Config.ShowStrings {
			d.ignoreTypes["string"] = true
		}

			if !d.Config.ShowNumbers {
				for i, _ := range numberTypes {
					d.ignoreTypes[numberTypes[i]] = true
				}
			}
	*/
}

func (d *DotDriver) PrintFooter() {
	fmt.Println("}")
}

func (d *DotDriver) PrintType(t *reflect.Type) {
	printType(nil, t, d)
}

func printType(parent *reflect.Type, t *reflect.Type, d *DotDriver) {

	kind := (*t).Kind().String()
	if kind == "ptr" || kind == "slice" {
		actualType := (*t).Elem()
		printType(parent, &actualType, d)
		return
	}

	if kind != "struct" {
		return
	}

	n := (*t).NumField()
	for i := 0; i < n; i++ {
		newT := (*t).Field(i).Type

		if _, ok := d.ignoreTypes[removePointerLeadingAsterisk(newT.String())]; ok {
			return
		}

		newKind := newT.Kind().String()
		if newKind == "ptr" || newKind == "slice" {
			newT = newT.Elem()
		}

		if newT.Name() != "int8" && newT.Name() != "string" && newT.Name() != "int16" {
			parentString := clean((*t).String(), d.Config.RemovePackagePrefix)
			var child string
			child = clean(newT.String(), d.Config.RemovePackagePrefix)
			edgeName := parentString + "_" + child
			if _, ok := d.edges[edgeName]; !ok {
				fmt.Println("\t\"" + parentString + "\"->\"" + child + "\";")
				d.edges[edgeName] = true
			}
			printType(t, &newT, d)
		}
	}
}

func removePointerLeadingAsterisk(s string) string {
	return strings.Replace(s, "*", "", 1)

}

func clean(s string, removePackagePrefix bool) string {
	s = strings.Replace(s, "*", "", 1)

	if removePackagePrefix {
		if strings.Contains(s, ".") {
			index := strings.Index(s, ".")
			if index > 0 {
				s = s[index+1:]
			}
		}
	}
	return s
}
