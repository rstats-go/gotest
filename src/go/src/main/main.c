#include <R.h>
#include <Rinternals.h>
#include "_cgo_export.h"

SEXP godouble(SEXP x){
  return Rf_ScalarInteger( DoubleIt( INTEGER(x)[0] ) ) ;
}
