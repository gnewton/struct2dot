package main

import (
	//"github.com/gnewton/gomesh2016"
	"github.com/gnewton/pubmedstruct"
	"github.com/gnewton/struct2dot"
)

func main() {

	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		NotTypes: []string{
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
	}

	pt := struct2dot.DotDriver{Config: &config}
	pt.PrintHeader()

	d := new(pubmedstruct.PubmedArticle)
	pt.PrintType(d)

	pt.PrintFooter()

}
