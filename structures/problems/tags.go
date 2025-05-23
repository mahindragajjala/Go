package problems

import (
	"encoding/json"
	"fmt"
)

/*
Use struct tags to define json field names for a struct APIUser.
Marshal and unmarshal JSON into/from this struct.
*/
type APIUser struct {
	UserName string `json:"username"`
	Data     string `json:"data"`
}

func Data() {
	input := APIUser{UserName: "OPENAPI", Data: "010100101000100"}
	jsondata := Marshal(input)
	fmt.Println(string(jsondata))
	Unmarshal(jsondata, &input)

}
func Marshal(data APIUser) []byte {
	jsondata, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error in marshalling...")
	}
	return jsondata
}
func Unmarshal(marshaldata []byte, da *APIUser) {
	err := json.Unmarshal(marshaldata, da) // <- no & here
	if err != nil {
		fmt.Println("error in the unmarshal:", err)
		return
	}
	fmt.Printf("Unmarshaled Data: %+v\n", da)
}
