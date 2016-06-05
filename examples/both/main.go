package main

import (
	"github.com/gnewton/gomesh2016"
	"github.com/gnewton/pubmedstruct"
	"github.com/gnewton/struct2dot"
	"reflect"
)

func main() {
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: false,
		IgnoreTypes: []string{
			"gomesh2016.Day",
			"gomesh2016.Month",
			"gomesh2016.Year",
			"pubmedstruct.Affiliation",
			"pubmedstruct.CollectiveName",
			"pubmedstruct.Day",
			"pubmedstruct.ForeName",
			"pubmedstruct.Hour",
			"pubmedstruct.Identifier",
			"pubmedstruct.Identifier",
			"pubmedstruct.Initials",
			"pubmedstruct.LastName",
			"pubmedstruct.Minute",
			"pubmedstruct.Month",
			"pubmedstruct.Season",
			"pubmedstruct.Suffix",
			"pubmedstruct.Year",
		},
		ManualLinks: map[string][]string{
			"pubmedstruct.DescriptorName": []string{"gomesh2016.DescriptorRecord"},
			"pubmedstruct.QualifierName":  []string{"gomesh2016.QualifierRecord"},
		},
	}
	pt := struct2dot.DotDriver{Config: &config}
	pt.PrintHeader()

	d := new(pubmedstruct.PubmedArticle)
	t := reflect.TypeOf(*d)
	pt.PrintType(&t)

	dr := new(gomesh2016.DescriptorRecord)
	t = reflect.TypeOf(*dr)
	pt.PrintType(&t)

	sr := new(gomesh2016.SupplementalRecord)
	t = reflect.TypeOf(*sr)
	pt.PrintType(&t)

	par := new(gomesh2016.PharmacologicalAction)
	t = reflect.TypeOf(*par)
	pt.PrintType(&t)

	qr := new(gomesh2016.QualifierRecord)
	t = reflect.TypeOf(*qr)
	pt.PrintType(&t)

	pt.PrintFooter()
}
