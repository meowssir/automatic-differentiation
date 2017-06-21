package main

import "strings"

// https://en.wikipedia.org/wiki/Automatic_differentiation
// Computational differentiation algorithm for computing implicit derivatives

func main() {

    // a string of a polynomial expression we want to compute the implicit derivative of.
    // expression: 2*x*y^2

    // interface to handle variant types.
    a := []interface{}{2, "*", "x", "*", "y", "^", 2}
    // 0 -> constant: define its properties and it should be / have an operator like property or simply use an operator on each coefficient.
    // 1 -> product operator
    // 2 -> x term: variable not being y, define it's properties.
    // 3 -> product operator
    // 4 -> exponent operator: let j be an idx who's element is a variable, if j+1 is ^ -
    // - j+2 is the exponent on the j term. handle this..
    // - j+2 becomes the coefficient on the j term, j-1 - and j+1 is decremented n-1. handle powers of 1.
    b := []interface{}{}
    // define each operators properties.
    ...

    // define the chain rule for a given term.
    // let j be an idx who's element is y which is has a composite function at j+n.
    // based on prefined rules of the composite function, compute its derivative.
    // the expression becomes j*j+n assuming we've placed its derivative in j+n.
}
