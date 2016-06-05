# struct2dot

##Convert Go structs to GraphViz dot format.

**Note that this was built to deal with the use case of wanting to visualize complex XML, so this is not a generic visualizer for Go data structures.**

The worksflow to convert XMl to [Graphviz](html://www.graphviz.org)'s [`dot`](https://en.wikipedia.org/wiki/DOT_%28graph_description_language%29) format, for subsequent processing by [`neato`](http://linux.die.net/man/1/neato):

1. XML converted to Go structs using [`chidley`](https://github.com/gnewton/chidley)
2. Write simple Go program using this `struct2dot` library and structs you want from `chidley` generated Go code
3. Compile and run #2, producing `dot` format output to stdio
4. Run `neato` on `dot` file, like: `neato -Tsvg out.dot > out.svg`

##
Example code to print structs generated by `chidley` from NLM's [Medical Subject Headings (MeSH®) 2016 XML](https://www.nlm.nih.gov/mesh/download_mesh.html) at [gomesh2016](https://github.com/gnewton/gomesh2016):
```
	pt := new(struct2dot)
	pt.PrintHeader()

	dr := new(gomesh2016.DescriptorRecord)
	t := reflect.TypeOf(*dr)
	pt.PrintType(t)

	sr := new(gomesh2016.SupplementalRecord)
	t = reflect.TypeOf(*sr)
	pt.PrintType(t)

	par := new(gomesh2016.PharmacologicalAction)
	t = reflect.TypeOf(*par)
	pt.PrintType(t)

	qr := new(gomesh2016.QualifierRecord)
	t = reflect.TypeOf(*qr)
	pt.PrintType(t)

	pt.PrintFooter()
}
```

Output:

<img src="https://gnewton.github.io/repos/struct2dot/meshDefault.svg">

##Configuration

The `struct2dot` type can take a `struct2dot.Config`, that can be used to alter the output and is generally used to make the diagram more readable. Defaults:

```
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		IgnoreTypes:         nil,
		ManualLinks:         nil,
	}
```

Links in the diagram to Go fundamental types tends to make the diagram very busy.
The following are used to control if they are included.
* `Config.ShowStrings` is a flag that indicates whether Go string types (including `[]string`) should be included
* `Config.ShowNumbers`: is a flag that indicates whether Go number types (including `[]type`)should be included

The package name prefix can also make the diagram larger and busier.
* `Config.RemovePackagePrefix` can control this

```
	config := struct2dot.Config{
		ShowStrings:         true,
		ShowNumbers:         true,
		RemovePackagePrefix: false,
		IgnoreTypes:         nil,
		ManualLinks:         nil,
	}
```

Here is an example with the above `config`, with strings, numbers turned on, and package names not removed:
<img src="https://gnewton.github.io/repos/struct2dot/meshWithStringsAndNumbersAndPackage.svg">


### Removing types in XML
In addition to fundamental types, some of the types generated by `chidley` from the XML might also cause the `dot` diagram to be too busy.
It is possible to indicate which types should be ignored. Here, the 'gomesh2016.Year', 'gomesh2016.Month', 'gomesh2016.Day' are used by many types and make the diagram more busy:

```
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		IgnoreTypes:         []string{"gomesh2016.Year", "gomesh2016.Month", "gomesh2016.Day"},
	}
```

Output:
<img src="https://gnewton.github.io/repos/struct2dot/meshDefaultWithIgnoreTypes.svg">

### Adding nodes and edges
It is possible to programmatically add additional nodes and links.
Here a edge is added from `DescriptorReferredTo` to `DescriptorRecord` (this is from the example directory [https://github.com/gnewton/struct2dot/tree/master/examples/MeSH2dot](https://github.com/gnewton/struct2dot/tree/master/examples/MeSH2dot):

```
	config := struct2dot.Config{
		ShowStrings:         false,
		ShowNumbers:         false,
		RemovePackagePrefix: true,
		IgnoreTypes:         []string{"gomesh2016.Year", "gomesh2016.Month", "gomesh2016.Day"},
		ManualLinks:         map[string][]string{"DescriptorReferredTo": []string{"DescriptorRecord"}},
	}
```

**NB: Bug: The text used needs to match the internal string, which is effected by whether the package name is removed.
So if the above had set `RemovePackagePrefix: true`, then this would require `ManualLinks: map[string][]string{"gomesh2016.DescriptorReferredTo": []string{"gomesh2016.DescriptorRecord"}}`.

This is more a function of my `dot` knowledge than a Go issue. Soon to be fixed**

Output (note added link line colour is red):
<img src="https://gnewton.github.io/repos/struct2dot/meshDefaultWithIgnoreTypesManualLink.svg">

# More complex example: Pubmed article

<img src="https://gnewton.github.io/repos/struct2dot/pubmedDefaultIgnoreTypes.svg">

