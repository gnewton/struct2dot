package struct2dot

import (
	"fmt"
	"reflect"
	"strings"
)

var numberTypes = []string{
	"uint8",
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

type isAttributeFunc func(string) bool

type Config struct {
	ShowStrings         bool
	ShowNumbers         bool
	ShowAttrubutes      bool
	RemovePackagePrefix bool

	ManualLinks map[string][]string

	NotTypes []string

	NotAttributes   []string
	IsAttributeFunc isAttributeFunc
}

type DotDriver struct {
	Config            *Config
	edges             map[string]bool
	nodeLabelsPrinted map[string]bool
	notTypes          map[string]bool
}

func (d *DotDriver) PrintHeader() {
	d.init()
	fmt.Println("digraph foo {")
	fmt.Println("\tnode [color=Red]")
	fmt.Println("\tedge [color=Blue]")
	fmt.Println("\toverlap=false;")
	fmt.Println("\tsplines=true;")
	fmt.Println("\t# End Header")
	fmt.Println("")

}

func (d *DotDriver) init() {
	d.edges = make(map[string]bool)
	d.notTypes = make(map[string]bool)
	d.nodeLabelsPrinted = make(map[string]bool)
	if d.Config == nil {
		d.Config = &(Config{
			ShowStrings:         false,
			ShowNumbers:         false,
			RemovePackagePrefix: true,
			NotTypes:            nil,
		})
	}
	if d.Config.NotTypes != nil && len(d.Config.NotTypes) > 0 {
		for i, _ := range d.Config.NotTypes {
			d.notTypes[d.Config.NotTypes[i]] = true
			d.notTypes["*"+d.Config.NotTypes[i]] = true
		}
	}

	if !d.Config.ShowStrings {
		d.notTypes["string"] = true
		d.notTypes["[]string"] = true
	}

	if !d.Config.ShowNumbers {
		for i, _ := range numberTypes {
			d.notTypes[numberTypes[i]] = true
		}
	}
}

func (d *DotDriver) PrintFooter() {
	if d.Config.ManualLinks != nil && len(d.Config.ManualLinks) > 0 {
		fmt.Println("\n\n\t# MANUAL LINKS")
		fmt.Println("\tedge [color=Green]")
		for src, dests := range d.Config.ManualLinks {
			for _, dest := range dests {
				printNodeEdgeNode(src, dest, d.nodeLabelsPrinted)
			}

		}
	}
	fmt.Println("}")
}

func (d *DotDriver) PrintType(i interface{}) {
	t := reflect.TypeOf(i)
	printType(nil, &t, d)
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
		newKind := newT.Kind().String()

		if _, ok := d.notTypes[removePointerLeadingAsterisk(newT.String())]; ok {
			continue
		}

		if newKind == "ptr" || newKind == "slice" {
			newT = newT.Elem()
		}

		parentString := clean((*t).String(), d.Config.RemovePackagePrefix)
		childString := clean(newT.String(), d.Config.RemovePackagePrefix)
		edgeName := parentString + "_" + childString
		if _, ok := d.edges[edgeName]; !ok {
			printNodeEdgeNode(parentString, childString, d.nodeLabelsPrinted)
			d.edges[edgeName] = true
		}
		printType(t, &newT, d)
	}
}

func removePointerLeadingAsterisk(s string) string {
	return strings.Replace(s, "*", "", 1)

}

func printNodeEdgeNode(parent string, child string, nodeLabelsPrinted map[string]bool) {
	fmt.Println("")
	printLabel(parent, nodeLabelsPrinted)
	printLabel(child, nodeLabelsPrinted)

	fmt.Println("\t\"" + parent + "\"  ->  \"" + child + "\";")
}

func printLabel(label string, nodeLabelsPrinted map[string]bool) {
	if _, ok := nodeLabelsPrinted[label]; !ok {
		fmt.Println("\t\"" + label + "\"[label=\"" + label + "\"];")
		nodeLabelsPrinted[label] = true
	}
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
