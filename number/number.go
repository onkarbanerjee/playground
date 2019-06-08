package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Number struct {
	N,
	Fact,
	Fib int
}

func main() {
	http.HandleFunc("/number", numberHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("Could not start server", err)
	}
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	if r != nil && r.URL != nil {
		values := r.URL.Query()
		if len(values["n"]) > 0 {
			i, err := strconv.Atoi(values["n"][0])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			num := Number{N: i}
			num.Fact = factorial(i)
			num.Fib = fibonacci(i)
			buf := bytes.Buffer{}
			err = json.NewEncoder(&buf).Encode(num)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(buf.Bytes())
		}
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
