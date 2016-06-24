package main

import (
	"github.com/gnewton/gomesh2016"
	"github.com/gnewton/struct2dot"
)

func main() {

	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		NotTypes: []string{
			"gomesh2016.Day",
			"gomesh2016.Month",
			"gomesh2016.Year",
		},
		ManualLinks: map[string][]string{"DescriptorReferredTo": []string{"DescriptorRecord"}},
	}

	pt := struct2dot.DotDriver{Config: &config}
	pt.PrintHeader()

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
