#include <R.h>
#include <Rinternals.h>
#include "../gotest.h"

SEXP godouble(SEXP x){
  return Rf_ScalarInteger( DoubleIt( INTEGER(x)[0] ) ) ;
}
