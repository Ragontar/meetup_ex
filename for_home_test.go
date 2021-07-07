package ex_meetup

import (
	"fmt"
	"testing"
)

type TestInput struct {
	name string
	a float64
	b float64
	c float64
	expResult float64
}

func TestTestedFunction(t *testing.T) {
	var (
		caseExpResultCheck = TestInput{
			name: "case: correct input",
			a: 3,
			b: 4,
			c: 5,
			expResult: 6}
		caseZeroEdge = TestInput{
			name: "case: zero edge",
			a: 0,
			b: 2,
			c: 3,
			expResult: -1}
		caseNotTriangle = TestInput{
			name: "case: not a triangle",
			a: 10,
			b: 1,
			c: 1,
			expResult: -1}
	)
	var cases = []TestInput{caseExpResultCheck, caseNotTriangle, caseZeroEdge}

	for _, cs := range cases{
		fmt.Printf("Testing %s\n", cs.name)
		result, _ := TestedFunction(cs.a, cs.b, cs.c)
		if result != cs.expResult {
			panic(fmt.Sprintf("test: %s\n", cs.name))
		}
	}
}
