package main

import (
	//"github.com/gnewton/gomesh2016"
	"github.com/gnewton/pubmedstruct"
	"github.com/gnewton/struct2dot"
	"reflect"
)

func main() {

	ignoreTypes := []string{"pubmedstruct.PubmedData"}
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		IgnoreTypes:         ignoreTypes,
	}

	pt := struct2dot.DotDriver{Config: &config}
	pt.PrintHeader()

	d := new(pubmedstruct.PubmedArticle)
	t := reflect.TypeOf(*d)
	pt.PrintType(&t)

	pt.PrintFooter()

}
