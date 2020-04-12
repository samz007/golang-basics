package src

import (
	"testing"
)


// run test|debug test
func TestPalindromeChecker(t *testing.T){
	if palindromeChecker("yoy") != true {
		t.Fail()
	}
	if palindromeChecker("yoyo") != false {
		t.Fail()
	}
}