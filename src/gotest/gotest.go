package main

import "C"

//export DoubleIt
func DoubleIt(x int) int {
  return 2 * x ;
}

func main() {}
