package main

import (
	"encoding/json"
	"fmt"
)

const body = `{
	"msg": {
		"id": 123
	},
	"payment": {
		"id": 123,
		"site_id": "mla",
		"status": "approved"
	},
	"movement": [
		{
			"id": 456,
			"site_id": "mla",
			"type": "a"
		},
		{
			"id": 789,
			"site_id": "mla",
			"type": "b"
		}
	],
	"list": ["a", "b"]
}`

type Payload map[string]interface{}

func main() {
	payload := toPayload(body)

	//fmt.Println("Original...")
	//fmt.Println(body)
	//
	//fmt.Println("Normalized...")
	//fmt.Println(toString(normalize(payload)))

	result := normalize(payload)
	fmt.Println(toString(result))

}

func normalize(payload Payload) []Payload {
	payloads := []Payload{payload}

	for {
		normalized := make([]Payload, 0)

		for _, p := range payloads {
			key, array := getFirstList(p)

			splitted := spplit(p, key, array)

			normalized = append(normalized, splitted...)
		}

		if len(payloads) == len(normalized) {
			break
		}

		payloads = normalized
	}



	return payloads
}

func getFirstList(payload Payload) (string, []interface{})  {
	for key, value := range payload {
		if list, ok := value.([]interface{}); ok {
			return key, list
		}
	}

	return "", nil
}

func spplit(payload Payload, key string, list []interface{}) []Payload  {
	if key == "" {
		return []Payload{payload}
	}

	payloads := make([]Payload, len(list))

	s := toString(payload)
	for i, value := range list {
		payloads[i] = toPayload(s)
		payloads[i][key] = value
	}

	return payloads
}

func toPayload(in string) Payload {
	out := Payload{}

	json.Unmarshal([]byte(in), &out)

	return out
}

func toString(in interface{}) string {
	out, _ := json.Marshal(in)

	return string(out)
}
