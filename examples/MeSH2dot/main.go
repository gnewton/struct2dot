package main

import (
	"github.com/gnewton/gomesh2016"
	"github.com/gnewton/struct2dot"
	"reflect"
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
	t := reflect.TypeOf(dr)
	pt.PrintType(&t)

	sr := new(gomesh2016.SupplementalRecord)
	t = reflect.TypeOf(sr)
	pt.PrintType(&t)

	par := new(gomesh2016.PharmacologicalAction)
	t = reflect.TypeOf(par)
	pt.PrintType(&t)

	qr := new(gomesh2016.QualifierRecord)
	t = reflect.TypeOf(qr)
	pt.PrintType(&t)

	pt.PrintFooter()
}
