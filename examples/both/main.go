package main

import (
	"github.com/gnewton/gomesh2016"
	"github.com/gnewton/pubmedstruct"
	"github.com/gnewton/struct2dot"
)

func main() {
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: false,
		NotTypes: []string{
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
	pt.PrintType(d)

	dr := new(gomesh2016.DescriptorRecord)
	pt.PrintType(dr)

	sr := new(gomesh2016.SupplementalRecord)
	pt.PrintType(sr)

	par := new(gomesh2016.PharmacologicalAction)
	pt.PrintType(par)

	qr := new(gomesh2016.QualifierRecord)
	pt.PrintType(qr)

	pt.PrintFooter()
}
