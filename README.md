
This is an attempt to use `go` code in `src/` directory of an R package. 

Instead of building the shared library as usual, I'm using `go build` with `-buildmode=c-shared` in the `Makevars` : 

```
.PHONY: go

go:
	GOPATH=$(shell pwd)/.. go build -o $(SHLIB) -buildmode=c-shared gotest

clean:
	rm gotest.h
	rm gotest.so
```

`gotest` here refers to the name of a go package found in the `GOPATH`. 

For this to work, I have to use a `main` package with a dummy `main` function. 

```
package main

// #cgo CFLAGS: -I/Library/Frameworks/R.framework/Resources/include -DNDEBUG   -I/usr/local/include
// #cgo LDFLAGS: -dynamiclib -Wl,-headerpad_max_install_names -undefined dynamic_lookup -single_module -multiply_defined suppress -L/Library/Frameworks/R.framework/Resources/lib -L/usr/local/lib
import "C"

//export DoubleIt
func DoubleIt(x int) int {
  return 2 * x ;
}

func main() {}

```

This defines the go function `DoubleIt` that is then exposed to C with cgo, so that we can use it in the `test.c` file: 

```
#include <R.h>
#include <Rinternals.h>
#include "../gotest.h"

SEXP godouble(SEXP x){
  return Rf_ScalarInteger( DoubleIt( INTEGER(x)[0] ) ) ;
}
```

I have to cheat with the cgo directives. There's probably a better way ... 

```
// #cgo CFLAGS: -I/Library/Frameworks/R.framework/Resources/include -DNDEBUG   -I/usr/local/include
// #cgo LDFLAGS: -dynamiclib -Wl,-headerpad_max_install_names -undefined dynamic_lookup -single_module -multiply_defined suppress -L/Library/Frameworks/R.framework/Resources/lib -L/usr/local/lib
```
