package main

import "encoding/json"

var (
	homeData  map[string]interface{}
	itemsData = make(map[string]interface{})
)

func init() {
	loadHomeData()
	loadItemsData()
}

func loadHomeData() {
	j := `{"data":{
    "item":{
      "id": "home", "kind": "folder", "name":"Home",
			"path": [],
			"is_folder": "true",
			"is_report": "",
      "children": [
        {
					"id": "fin", "kind": "folder", "name":"Financial Reports",
					"is_folder": "true",
					"is_report": ""
				},
        {
					"id": "rh", "kind": "folder", "name":"Human Resources",
					"is_folder": "true",
					"is_report": ""
				},
				{
					"id": "readme", "kind": "report", "name":"README",
					"is_folder": "",
					"is_report": "true",
					"content": "Content"
				}
      ],
			"readme": "<h1>Content</h1><p>This is the main folder.</p>"
    }
  }}`

	mustLoad(j, &homeData)
}

func loadItemsData() {
	j := `{"data":{
		"item":{
			"id": "fin", "kind": "folder", "name":"Financial Reports",
			"path": [{"id":"","name":"home"}],
			"is_folder": "true",
			"is_report": "",
			"children": [
				{
					"id": "con", "kind": "folder", "name":"Planning",
					"is_folder": "",
					"is_report": "true"
				},
				{
					"id": "ro", "kind": "report", "name":"Operational",
					"is_folder": "",
					"is_report": "true"
				}
			]
		}
	}}`

	var i1 map[string]interface{}
	mustLoad(j, &i1)
	itemsData["fin"] = i1

	j = `{"data":{
		"item":{
			"id": "rh", "kind": "folder", "name":"Human resources",
			"path": [{"id":"","name":"home"}],
			"is_folder": "true",
			"is_report": "",
			"children": [
				{
					"id": "hdc", "kind": "report", "name":"Headcount",
					"is_folder": "",
					"is_report": "true"
				},
				{
					"id": "adm", "kind": "report", "name":"Admissions",
					"is_folder": "",
					"is_report": "true"
				}
			]
		}
	}}`

	var i2 map[string]interface{}
	mustLoad(j, &i2)
	itemsData["rh"] = i2
}

func mustLoad(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}
