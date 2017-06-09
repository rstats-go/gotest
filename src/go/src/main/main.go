package main

import "C"

import "twice"

//export DoubleIt
func DoubleIt(x int) int {
  return twice.Twice(x) ;
}

func main() {}
