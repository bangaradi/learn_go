package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` //gives the key as 'coursename' in json data
	Price    int      // sometimes it's not necessary to give any json specifics
	Platform string   `json:"platform"`       // gives key as 'plaform'
	Password string   `json:"-"`              // this '-' reflects that it's not shown in the response, because we don't want our passowrds to be visible
	Tags     []string `json:"tags,omitempty"` // gives key name as 'tags' if value is not 'nil' if the value is 'nil' then the field is not returned
}

// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{
			"reactjs",
			200,
			"youtube",
			"aditya123",
			[]string{"web", "frontend", "blabla"},
		},
		{
			"golang",
			300,
			"youtube",
			"bangar123",
			[]string{"web", "backend", "blabla"},
		},
		{
			"api",
			230,
			"youtube",
			"adi123",
			nil,
		},
	}

	// package this data as json data
	//first attempt directly use the Marhshal.. this gives us kind of unreadable data
	//	// finalJson, err := json.Marshal(courses)
	// using MarshalIndent with parameters (data, prefix, indent string)
	//	// trying changing the parameters and running to get a better hold of it :tick:
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// ----------------
	// THE OUTPUT BEFORE THE JSON TAGGING
	// ----------------
	// [
	//        {
	//               "Name": "reactjs",
	//               "Price": 200,
	//               "Platform": "youtube",
	//               "Password": "aditya123",
	//               "Tags": [
	//                      "web",
	//                      "frontend",
	//                      "blabla"
	//               ]
	//        },
	//        {
	//               "Name": "golang",
	//               "Price": 300,
	//               "Platform": "youtube",
	//               "Password": "bangar123",
	//               "Tags": [
	//                      "web",
	//                      "backend",
	//                      "blabla"
	//               ]
	//        },
	//        {
	//               "Name": "api",
	//               "Price": 230,
	//               "Platform": "youtube",
	//               "Password": "adi123",
	//               "Tags": null
	//        }
	// ]

	// ---------------
	// THE OUTPUT AFTER THE JSON TAGGING
	// ----------------
	// [
	//
	//	{
	//	        "coursename": "reactjs",
	//	        "Price": 200,
	//	        "platform": "youtube",
	//	        "tags": [
	//	                "web",
	//	                "frontend",
	//	                "blabla"
	//	        ]
	//	},
	//	{
	//	        "coursename": "golang",
	//	        "Price": 300,
	//	        "platform": "youtube",
	//	        "tags": [
	//	                "web",
	//	                "backend",
	//	                "blabla"
	//	        ]
	//	},
	//	{
	//	        "coursename": "api",
	//	        "Price": 230,
	//	        "platform": "youtube"
	//	}
	//
	// ]
}

func DecodeJson() {
	jsonData := []byte(`
		{
		"coursename": "reactjs",
		"Price": 200,
		"platform": "youtube",
		"Password": "aditya123",
		"tags": [
				"web",
				"frontend",
				"blabla"
				]
		}
	`)
	var obtainedData course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("the data is valid")
		json.Unmarshal(jsonData, &obtainedData)
		fmt.Printf("%#v \n", obtainedData)
	} else {
		fmt.Println("json wasn't valid")
	}

	// sometimes we just want the data in the form of key-value pair
	var myData map[string]interface{} // for the key we know it's string, but for value we aren't sure whether it'll surely be a string or not
	// so to deal with that issue, the above given syntax is used :tick:
	json.Unmarshal(jsonData, &myData)
	fmt.Println("the data in the form of a map is -->")
	fmt.Println(myData)
}
