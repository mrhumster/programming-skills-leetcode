/*
Package t860

At a lemonade stand, each lemonade costs $5.
Customers are standing in a queue to buy from you and order one at a time
(in the order specified by bills). Each customer will only buy one lemonade
and pay with either a $5, $10, or $20 bill. You must provide the correct
change to each customer so that the net transaction is that the customer pays $5.

Note that you do not have any change in hand at first.

Given an integer array bills where bills[i] is the bill the ith customer pays,
return true if you can provide every customer with the correct change, or false otherwise.
*/
package t860

import (
	"errors"
	"fmt"
)

func needMoneyBack(v int) int {
	return v - 5
}


func takeMoneyFromTheTill(till *map[int]int, v int) error {
	if v == 15 && (*till)[10] >= 1 && (*till)[5] >= 1 {
		(*till)[10] -= 1
		(*till)[5] -= 1
		return nil
	}
	if v == 15 && (*till)[10] == 0 && (*till)[5] >= 3 {
		(*till)[5] -= 3
		return nil
	}
	if v == 10 && (*till)[10] >= 1 {
		(*till)[10] -= 1
		return nil
	}
	if v == 10 && (*till)[5] >= 2 {
		(*till)[5] -= 2
		return nil
	}
	if v == 5 && (*till)[5] >= 1 {
		(*till)[5] -= 1
		return nil
	}
	return errors.New("no back money")
}

func lemonadeChange(bills []int) bool {
	fmt.Printf("bills %#v\n", bills)
	till := map[int]int{5: 0, 10: 0, 20: 0}
	for _, v := range bills {
		fmt.Printf("%#v\n", till)
		amound := needMoneyBack(v)
		if amound == 0 {
			// If does't need money back
			till[v] += 1
			continue
		}
		if amound != 0 {
			// If you need to return the money and till contains the required amound
			err := takeMoneyFromTheTill(&till, amound)
			if err != nil {
				fmt.Println(err)
				return false
			}
			till[v] += 1
		}
	}
	return true
}
