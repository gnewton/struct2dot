package struct2dot

import (
	"fmt"
	"reflect"
	"strings"
)

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
	d.edges = make(map[string]bool)
	fmt.Println("digraph foo {")
	fmt.Println("\tnode [color=Red]")
	fmt.Println("\tedge [color=Blue]")
}

func (d *DotDriver) PrintFooter() {
	fmt.Println("}")
}

func (d *DotDriver) PrintType(t *reflect.Type) {
	printType(nil, t, d.edges)
}

func printType(parent *reflect.Type, t *reflect.Type, edges map[string]bool) {
	kind := (*t).Kind().String()
	if kind == "ptr" || kind == "slice" {
		actualType := (*t).Elem()
		printType(parent, &actualType, edges)
		return
	}

	if kind != "struct" {
		return
	}

	n := (*t).NumField()
	//fmt.Println(n)
	for i := 0; i < n; i++ {
		//fmt.Println(i)
		//newT := reflect.TypeOf((*t).Field(i))
		newT := (*t).Field(i).Type
		newKind := newT.Kind().String()
		if newKind == "ptr" || newKind == "slice" {
			newT = newT.Elem()
		}
		if newT.Name() != "int8" && newT.Name() != "string" && newT.Name() != "int16" {
			parentString := clean((*t).String())
			var child string
			child = clean(newT.String())
			edgeName := parentString + "_" + child
			if _, ok := edges[edgeName]; !ok {
				// if newKind == "slice" {
				// 	child = clean(newT.String() + " [label=\"[]" + clean(newT.String()) + "\"]")
				// }
				fmt.Println("\t", parentString, "->", child, ";")
				edges[edgeName] = true
			}
			printType(t, &newT, edges)
		}
	}
}

func clean(s string) string {
	s = strings.Replace(s, "*", "", 1)
	//s = strings.Replace(s, "gomesh2016.", "", 1)
	s = strings.Replace(s, "pubmedstruct.", "", 1)
	//s = strings.Replace(s, "[]", "", -1)

	return s
}
