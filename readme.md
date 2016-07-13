Merger
=====
Merger is a tiny little application that lets you run multiple apps at the same time in parallel and send the output to the same terminal.

You can run as many apps as you want by passing them as argument to the `merger` command line tool.

Examples:
=====
To run `nodemon` and `gulp` at same time:
`merger nodemon gulp`

To run commands with multiple arguments, add them in quotes:
`merger "echo hi" "killall node"`

To stop a merger, run `CTRL+C`, and it will also kill all the processes that were started using merger. This tool should help in cases where you need to run multiple build processes but you dont want to open multiple terminal windows to see the output. 

Compiling:
=====
Run `go build` and `go install` on the project directory to build and install to your path. The path's required by go need to be setup properly. You can also get a pre-built binary from the releases page.