/*
Package baseballgame
You are keeping the scores for a baseball game with strange rules.
At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i]
is the ith operation you must apply to the record and is one of the
following:

An integer x.

	Record a new score of x.
	'+'.
	Record a new score that is the sum of the previous two scores.
	'D'.
	Record a new score that is the double of the previous score.
	'C'.
	Invalidate the previous score, removing it from the record.

Return the sum of all the scores on the record after applying all the operations.

The test cases are generated such that the answer and all intermediate calculations
fit in a 32-bit integer and that all operations are valid.
*/
package baseballgame

import (
	"log"
	"strconv"
)

func doubleLast(s *[]int) {
	if len(*s) == 0 {
		log.Println("Operation error")
		return
	}
	*s = append(*s, (*s)[len(*s)-1]*2)
}

func cancelLast(s *[]int) {
	if len(*s) == 0 {
		log.Println("Operation error")
	}
	*s = (*s)[:len(*s)-1]
}

func sumLastTwo(s *[]int) {
	if len(*s) < 2 {
		log.Println("Operation error")
	}
	*s = append(*s, (*s)[len(*s)-1]+(*s)[len(*s)-2])
}

func addScore(s *[]int, v int) {
	*s = append(*s, v)
}

func process(s *[]int, o string) {
	log.Printf("Score before operation: %#v", *s)
	log.Printf("Operation: %s", o)
	switch o {
	case "D":
		doubleLast(s)
	case "C":
		cancelLast(s)
	case "+":
		sumLastTwo(s)
	default:
		num, err := strconv.Atoi(o)
		if err != nil {
			log.Println("operation error")
			return
		}
		addScore(s, num)
	}
	log.Printf("Score after operation: %#v", *s)
}

func sum(s *[]int) int {
	_s := 0
	for _, v := range *s {
		_s += v
	}
	return _s
}

func calPoints(operations []string) int {
	log.Println("\nStart new game!")
	score := []int{}
	for _, operation := range operations {
		process(&score, operation)
	}
	return sum(&score)
}
