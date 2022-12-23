package main

import "fmt"

type apiPayload interface {
	map[string]interface{} | []map[string]interface{}
}

type Request[P apiPayload] struct {
	Endpoint string
	Method   string
	Data     P
}

func main() {
	d := make(map[string]interface{})
	d["key_1"] = "value_1"
	d["key_2"] = 1

	wrapper := Request[map[string]interface{}]{Endpoint: "abc.com", Method: "GET", Data: d}
	fmt.Printf("%s %T\n", wrapper, wrapper)

	result := []map[string]interface{}{}
	mp1 := map[string]interface{}{
		"one": 1,
		"two": 2,
	}

	mp2 := map[string]interface{}{
		"three": 3,
		"four":  4,
	}
	result = append(result, mp1, mp2)

	wrapper1 := Request[[]map[string]interface{}]{Endpoint: "abc.com", Method: "GET", Data: result}
	fmt.Printf("%s %T\n", wrapper1, wrapper1)
}
