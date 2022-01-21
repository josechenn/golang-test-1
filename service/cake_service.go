package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const cakes = 20
const apple = 25

func CountTotalBox(w http.ResponseWriter, r *http.Request) {
	var array_cakes []int
	var array_apple []int
	var result string

	for x := 2; x < 10; x++ {
		if cakes%x == 0 {
			array_cakes = append(array_cakes, x)
		}
		if apple%x == 0 {
			array_apple = append(array_apple, x)
		}
	}

	for _, a1 := range array_cakes {
		for _, a2 := range array_apple {
			if a1 == a2 {
				result = fmt.Sprintf("Ainun dapat membuat sebanyak %s box dan setiap box nya ada %s kue dan ada %s apel", strconv.Itoa(a1), strconv.Itoa(cakes/a1), strconv.Itoa(apple/a2))
			} else {
				result = "kue dan apel tidak bisa di bagi rata"
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
