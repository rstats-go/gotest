#' Doubles an integer using go
#'
#' @param x an integer
#' @useDynLib gotest
#' @export
godouble <- function(x) {
  .Call("godouble", x, PACKAGE = "gotest")
}
