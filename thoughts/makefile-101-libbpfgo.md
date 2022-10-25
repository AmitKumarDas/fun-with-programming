## Learn Makefile By Tinkering An Existing One

### [DESIGN] Parts of a Makefile
```yaml
- .ONESHELL
- SHELL = /bin/bash
- DIRs
- CLIs
- OS & Arch
```

### [IMP] Rules
```yaml
- $@ # target/file name itself
- $< # first input file
- $^ # all input files
- $? # all input files newer than target
- $$
- $* # part that matched in the rule definition's % bit
- $(@D) # to refer to just the dir portion of $@
- $(@F) # to refer to just the file portions of $@
- $(<D) # work the same way on the $< variable 
- $(<F) # work the same way on the $< variable 
```

### [TIL] Variable Set
```makefile
v1 = val # lazy
v2 := val # immediate
v3 ?= val # lazy set if absent
v4 += val # appending to existing value (or setting the value if variable does not exist)
```

### [TIL] The special var $(MAKE) means "the make currently in use"

### [TIL] false program
```makefile
# What happens if there's an error!?  Let's say you're building stuff, and
# one of the commands fails.  Make will abort and refuse to proceed if any
# of the commands exits with a non-zero error code.
#
# To demonstrate this, we'll use the `false` program, which just exits with
# a code of 1 and does nothing else.
badkitty:
    $(MAKE) kitty
    false # <-- this will fail
    echo "should not get here"
```

### [QUIRK] $* and $% together
```makefile
srcfiles := $(shell echo src/{00..99}.txt)

# How do we make a text file in the src dir?
# We define the filename using a "stem" with the % as a placeholder.
# What this means is "any file named src/*.txt", and it puts whatever
# matches the "%" bit into the $* variable.
src/%.txt:
    @# First things first, create the dir if it doesn't exist.
    @# Prepend with @ because srsly who cares about dir creation
    @[ -d src ] || mkdir src
    @# then, we just echo some data into the file
    @# The $* expands to the "stem" bit matched by %
    @# So, we get a bunch of files containing their names
    echo $* > $@

# Running `make source` will make ALL of the files in the src/ dir.  Before
# it can make any of them, it'll first make the src/ dir itself.  Then
# it'll copy the "stem" value (that is, the number in the filename matched
# by the %) into the file, like the rule says above.
#
# Try typing "make source" to make all this happen.
source: $(srcfiles)
```

### $< and $@ together
```makefile
# So, to make a dest file, let's copy a source file into its destination.
# Also, it has to create the destination folder first.
#
# The destination of any dest/*.txt file is the src/*.txt file with
# the matching stem.
dest/%.txt: src/%.txt
    @[ -d dest ] || mkdir dest
    cp $< $@
```

### [TIL] Dont Repeat Yourself
```makefile
# We don't want to type `make dest/#.txt` 100 times!
#
# We can use the built-in pattern substitution "patsubst" so we don't have
# to re-build the list.  This patsubst function uses the "stem" concept

destfiles := $(patsubst src/%.txt,dest/%.txt,$(srcfiles))
destination: $(destfiles)
```

### [TIL] Make uses := for assignment instead of =

### [TIL] Execute the shell program to generate a list of files
```makefile
srcfiles := $(shell echo src/{00..99}.txt)
```

### [QUIRK] $(var) vs $$(cmd args) vs $$var
```yaml
- We have to use a double-$ in the command line
- Since each line of a makefile is parsed first using the makefile syntax
- And THEN the result is passed to the shell
```

### [QUIRK] ifeq() vs if [[]] :: make vs bash

### [NICE] Variable Substitution
```makefile
OUTPUT = ./output
VMLINUXH = $(OUTPUT)/vmlinux.h

LIBBPF_SRC = $(abspath ./libbpf/src)
LIBBPF_OBJ = $(abspath ./$(OUTPUT)/libbpf.a)
LIBBPF_OBJDIR = $(abspath ./$(OUTPUT)/libbpf)
LIBBPF_DESTDIR = $(abspath ./$(OUTPUT))
```

### [NICE] Make Directories
```makefile
OUTPUT = ./output

$(OUTPUT):
    mkdir -p $(OUTPUT)

$(OUTPUT)/libbpf:
    mkdir -p $(OUTPUT)/libbpf
```

### [TIL] Implement function like body in Dynamic Make Target
### [NICE] Only echo & @if to Implement Any Logic
```makefile
.PHONY: vmlinuxh
vmlinuxh: $(VMLINUXH)

$(VMLINUXH): $(OUTPUT)
    @if [ ! -f $(BTFFILE) ]; then \
        echo "ERROR: kernel does not seem to support BTF"; \
        exit 1; \
    fi
    @if [ ! $(BPFTOOL) ]; then \
        echo "ERROR: could not find bpftool"; \
        exit 1; \
    fi;

    @echo "INFO: generating $(VMLINUXH) from $(BTFFILE)"; \
    if ! $(BPFTOOL) btf dump file $(BTFFILE) format c > $(VMLINUXH); then \
        echo "ERROR: could not create $(VMLINUXH)"; \
        rm -f "$(VMLINUXH)"; \
        exit 1; \
    fi;
```

### [TIL] .ONESHELL
```yaml
- Each line of the command list is run as a separate invocation of the shell
- We can run both of the commands in the *same* shell invocation, by escaping the \n character
- Make provides a phony target .ONESHELL which basically lets you write a full script inline 
- The full contents of the recipe are passed to a single shell to be executed
```

#### Following does not work since each command runs in a separate shell
```makefile
newfile:
    @printf "Filename: "
    @read FILE
    touch $$FILE
```

#### This works
```makefile
newfile:
    @printf "Filename: " && read FILE && touch $$FILE
```

#### Better But Invalid
```yaml
- It is invalid since second @ would be passed verbatim to your shell
```
```makefile
.ONESHELL:
newfile:
    @printf "Filename: "
    @read FILE
    touch $$FILE
```

#### Better & Works
```makefile
.ONESHELL:
newfile:
    @printf "Filename: "
    read FILE
    touch $$FILE
```

### [TIL] Silence Output
```yaml
- Declare the .SILENT target
- This is a special target which will silence all echoing by default
- It is generally a good idea to sill leave a verbosity option, e.g. make VERBOSE=1
```

```makefile
ifndef VERBOSE
.SILENT:
endif
```

### [TIL] Files & Functions
```yaml
- $(wildcard <pattern>): generate a list of files in the current directory that match the pattern
- $(addsuffix <suffix>, <list>): add a given suffix to each member of a list
- $(basename <list>): get the base name (without extensions) of each member of a list
```

```makefile
files = $(wildcard test-*)

all: $(files)

.PHONY: all $(files)
$(files): 
    @echo "Building $(basename $@)..."
```

## References
```yaml
- https://github.com/aquasecurity/libbpfgo/blob/main/Makefile
- https://www.gnu.org/software/make/manual/html_node/One-Shell.html
- https://til.zqureshi.in/makefile-oneshell/
- https://gist.github.com/isaacs/62a2d1825d04437c6f08
```
