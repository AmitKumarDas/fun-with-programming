## Golang scripting to manage releases
- This codebase showcases use of Golang to script a project's release
- Note: _Release management is done by Carvel_
- Motivation:
  - Goodness of bash coupled with correctness of Golang
  - Granular error handling
  - 100% unit tested
  - Modular code base
  - Reduced bash & make spaghetti

## Usage
- `make`

### References
- https://github.com/gohugoio/hugo/blob/master/magefile.go
- https://github.com/magefile/mage/blob/master/sh/cmd.go
- https://godoc.org/gopkg.in/pipe.v2
- https://github.com/bitfield/script
