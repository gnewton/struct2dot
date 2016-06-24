package struct2dot

import (
	"github.com/gnewton/pubmedstruct"
	//	"reflect"
	"testing"
)

func TestDefaultRun(test *testing.T) {
	config := new(Config)

	pt := DotDriver{Config: config}

	pt.PrintHeader()
	d := new(pubmedstruct.PubmedArticle)
	pt.PrintType(d)

	pt.PrintFooter()
}
