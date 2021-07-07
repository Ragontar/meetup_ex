package ex_meetup

import (
	"errors"
	"math"
)

func TestedFunction(a, b, c float64) (float64, error) {
	//Negative length check
	if (a <= 0) || (b <= 0) || (c <= 0) {
		return -1, errors.New("wrong edges' length")
	}
	//Being-a-triangle condition. Also covers a zero-length case
	if !((a + b > c) && (a + c > b) && (b + c > a)) {
		return -1, errors.New("not a triangle")
	}

	p := (a + b + c)/2
	answer := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	return answer, nil
}
