# hello\_dygraph\_rice

Simple illustration of how to use Go templates with the Dygraphs
JavaScript library, using "go.rice" to embed the Web assets.

Note that if `dygraph-cdn.html` is used as the template instead of
`dygraph.html` in the Go source file, then `http.Handle()` call can be
commented out since the program doesn't need to serve up the
JavaScript and CSS assets.

To add the `assets` and `templates` directories to the binary using
the `rice` tool, make sure you do the following:

1. Execute: `go get github.com/GeertJohan/go.rice`
2. Execute: `go get github.com/GeertJohan/go.rice/rice`
3. Build the application from within the application code directory:
   `go build .` 
4. Assuming your `$GOPATH/bin` is in your `$PATH`, you can then, from
   within the code directory, execute: `rice append --exec
   hello_dygraph_rice` .
   
At this point, `hello_dygraph_rice` now includes the `templates` and
the `assets` directories embedded at Zip archives appended to the
binary.

As noted by the [go.rice](https://github.com/GeertJohan/go.rice)
GitHub project, if you run the binary from within the application
directory, `go.rice` will just access the files from the directories
on the file system.  Once you have appended the Zip archive(s) to the
Go binary, it will use the appended Zip archives instead.
