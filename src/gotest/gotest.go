package main

// #cgo CFLAGS: -I/Library/Frameworks/R.framework/Resources/include -DNDEBUG   -I/usr/local/include
// #cgo LDFLAGS: -dynamiclib -Wl,-headerpad_max_install_names -undefined dynamic_lookup -single_module -multiply_defined suppress -L/Library/Frameworks/R.framework/Resources/lib -L/usr/local/lib
import "C"

//export DoubleIt
func DoubleIt(x int) int {
  return 2 * x ;
}

func main() {}
