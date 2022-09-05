## testscript
https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript

### Notes:
- testscript is a better way to `test` files & commands
  - No need to write bash scripts
  - Yet achieve the agility of scripting
  - If needed use golang to write complex logic
- A testscript file is one with extension .txt or .txtar
- A testscript can be run via `go test`
- Each testscript file defines a subtest
- Hence, a bunch of testscripts can be run as subtests within a go test function
- To run a specific testscript e.g. `foo.txt` run `go test -run=TestXYZ/^foo$`
  - Where `TestXYZ` is the name of the test function in `*_test.go` file 
  - This enables one to run a specific subtest within `TestXYZ`
  - Alternatively run `go test -run=TestXYZ/foo`

### Best Practices:
- Keep the scripts modular i.e. Single Responsibility Pattern is your friend
