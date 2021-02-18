package main

import (
	"encoding/json"
	"testing"
	"github.com/intel-go/fastjson"
)

type Human struct{

	// defining struct variables
	Name string
	Address string
	Age int
}

func Benchmark_Json(b *testing.B) {
	var human1 Human

	// data in JSON format which
	// is to be decoded
	Data := []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`)

	// decoding human1 struct
	// from json format
	err := fastjson.Unmarshal(Data, &human1)

	if nil != err {
		panic(err)
	}
}

func Benchmark_Fastjson(b *testing.B) {
	var human1 Human

	// data in JSON format which
	// is to be decoded
	Data := []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`)

	// decoding human1 struct
	// from json format
	err := json.Unmarshal(Data, &human1)

	if nil != err {
		panic(err)
	}
}