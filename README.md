# struct2dot

Convert Go structs to GraphViz dot format.

*Note that this was built to deal with the use case of wanting to visualize complex XML, so this is not a generic visualizer for Go data structures.*

The worksflow to convert XMl to [Graphviz])(html://www.graphviz.org)'s [dot](https://en.wikipedia.org/wiki/DOT_%28graph_description_language%29) format, for subsequent processing by [neato](http://linux.die.net/man/1/neato):

1. XML converted to Go structs using [chidley](https://github.com/gnewton/chidley)
2. Write simple Go program using struct2dot and structs you want from chidley generated Go code
3. Compile and run #2, producing dot format output to stdio
4. Run neato on dot file, like: `neato -Tsvg out.dot > out.svg`


<img src="https://gnewton.github.io/repos/struct2dot/meshNoStringsNumbers.svg">