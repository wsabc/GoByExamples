package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// built-in json support
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.23)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("ghost")
	fmt.Println(string(strB))

	slcB, _ := json.Marshal([]string{"apple", "banana", "candy"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(mapB))

	// note: leading lowercase is not exposed
	ceo := &NPerson{name: "Chen", titles: []string{"CEO", "CFO"}}
	structB, _ := json.Marshal(ceo)
	fmt.Println(string(structB))

	ceo2 := &NPerson2{Name: "Chen", Titles: []string{"CEO", "CFO"}}
	structB2, _ := json.Marshal(ceo2)
	fmt.Println(string(structB2))

	// decode, note this literal bytes
	bb := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	var dat map[string]interface{}
	// note: need an address parameter
	if err := json.Unmarshal(bb, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"].(float64))

	strs := dat["strs"].([]interface{})
	fmt.Println(strs[1].(string))

	np := `{"n":"Kiwi", "t":["ceo", "cfo"]}`
	res := NPerson2{}
	// cast: []byte()
	_ = json.Unmarshal([]byte(np), &res)
	fmt.Println(res)
	fmt.Println(res.Titles[0])

	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"good": 1, "bad": 0}
	_ = encoder.Encode(d)

}

type NPerson struct {
	name   string
	titles []string
}

// named json keys
type NPerson2 struct {
	Name   string   `json:"n"`
	Titles []string `json:"t"`
}
