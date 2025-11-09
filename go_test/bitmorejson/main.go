package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// encoding of json. the above data can converted onto valid JSON

func main() {
	fmt.Println("welcome to json handling")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "oncloudev.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "oncloudev.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "oncloudev.com", "hit123", nil},
	}
	// package this data as a json data
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "oncloudev.com",
		"tags": ["web-dev","js"]
	}
	`)
	// verify the data its actually in a correct json format or not
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")

		//decode json
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT IN UPPERCASE")
	}

	// some cases whare you just want to add data to key value
	// scenario when we dont want the structure

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("the key is  %v and value is %v and type of data: %T\n", k, v, v)
	}
}
