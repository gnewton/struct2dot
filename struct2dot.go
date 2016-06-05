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
	ManualLinks         map[string][]string
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
	fmt.Println("\toverlap=false;")
	fmt.Println("\tsplines=true;")
}

func (d *DotDriver) init() {
	d.edges = make(map[string]bool)
	d.ignoreTypes = make(map[string]bool)
	if d.Config == nil {
		d.Config = &(Config{
			ShowStrings:         false,
			ShowNumbers:         false,
			RemovePackagePrefix: true,
			IgnoreTypes:         nil,
		})
	}
	if d.Config.IgnoreTypes != nil && len(d.Config.IgnoreTypes) > 0 {
		for i, _ := range d.Config.IgnoreTypes {
			d.ignoreTypes[d.Config.IgnoreTypes[i]] = true
			d.ignoreTypes["*"+d.Config.IgnoreTypes[i]] = true
		}
	}

	if !d.Config.ShowStrings {
		d.ignoreTypes["string"] = true
		d.ignoreTypes["[]string"] = true
	}

	if !d.Config.ShowNumbers {
		for i, _ := range numberTypes {
			d.ignoreTypes[numberTypes[i]] = true
		}
	}
}

func (d *DotDriver) PrintFooter() {
	if d.Config.ManualLinks != nil && len(d.Config.ManualLinks) > 0 {
		fmt.Println("\n\n### MANUAL LINKS")
		fmt.Println("\tedge [color=Red]")
		for src, dests := range d.Config.ManualLinks {
			for _, dest := range dests {
				fmt.Println(src, "->", dest)
			}

		}
	}
	fmt.Println("}")
}

func (d *DotDriver) PrintType(t *reflect.Type) {
	printType(nil, t, d)
}

func printType(parent *reflect.Type, t *reflect.Type, d *DotDriver) {

	kind := (*t).Kind().String()
	if parent != nil {
		//fmt.Println("-----------  ", (*parent).String(), (*t).String(), "  =", kind)
	}
	if kind == "ptr" || kind == "slice" {
		fmt.Println("   ###DEEP")
		actualType := (*t).Elem()
		printType(parent, &actualType, d)
		return
	}

	fmt.Println("   ###KIND=", kind)
	if kind != "struct" {
		return
	}

	n := (*t).NumField()
	if n == 0 {
		fmt.Println("   ###NOT CHILDREN")
	}
	//fmt.Println("   ### 12")

	for i := 0; i < n; i++ {
		fmt.Println("#++++++++++++++++++++++", (*t).Field(i).Name)
		newT := (*t).Field(i).Type
		newKind := newT.Kind().String()

		fmt.Println("### CHILD=", newT.String(), ":", newT.Name())
		fmt.Println("###", d.ignoreTypes)
		if _, ok := d.ignoreTypes[removePointerLeadingAsterisk(newT.String())]; ok {
			fmt.Println("#################### RETURNING")
			continue
		}

		if newKind == "ptr" || newKind == "slice" {
			newT = newT.Elem()
		}

		parentString := clean((*t).String(), d.Config.RemovePackagePrefix)
		var child string
		child = clean(newT.String(), d.Config.RemovePackagePrefix)
		edgeName := parentString + "_" + child
		if _, ok := d.edges[edgeName]; !ok {
			fmt.Println("\t\"" + parentString + "\"[label=\"" + parentString + "\"];")
			fmt.Println("\t\"" + parentString + "\"->\"" + child + "\";")
			d.edges[edgeName] = true
		}
		printType(t, &newT, d)
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
