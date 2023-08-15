package main

import (
	"encoding/json"
	"fmt"
)

// struct
type User struct {
	Name string 
	Email string
	Password string `json:"-"` // adding json:"-" , will not show the field to anyone .
	Age int 
	Role []string `json:"devRoles,omitempty"` //  `json:"devRoles, omitempty"` -> No Space with space it will give error
	// omitempty ignores if the field entered by user is nil or empty 
}

func main() {
	EncodeJson()
	DecodeData()
}

func EncodeJson(){
	newUsers:=[]User{
		{
			"Unkown",
			"unknow@mail.com",
			"84129enfcdf",
			33,
			[]string{"backendDev","devOPS"},
		},
		{
			"Unkown2",
			"unkno43w@mail.com",
			"84129e34nfcdf",
			21,
			[]string{"frontendDev","cloudDev"},
		},
		{
			"Unkow4n",
			"unkn334ow@mail.com",
			"84ss129enfcdf",
			27,
			nil, 
		},
	}

	// package this as JSON 
	// res,err:=json.Marshal(newUsers) // use marshalIndent for fomatted data 
	res,err:=json.MarshalIndent(newUsers,"","\t") 
	ErrHandler(err)
	fmt.Printf("%s\n ", res)
}

func DecodeData(){
	jsonData := []byte(
		`
		{
			"Name": "Unkown",
			"Email": "unknow@mail.com",
			"Age": 33,
			"devRoles": ["backendDev","devOPS"]
		}
			
		`,
	)

	var users User
	valid:=json.Valid(jsonData)
	switch valid{
		case true:
			fmt.Println("Json is valid")
			json.Unmarshal(jsonData,&users)
			fmt.Printf("%#v\n",users)
		case false:
			fmt.Println("Invalid json")
		default:
			fmt.Println("Error")
	}

	// other cases when one wants to add data to key value 
	var keyData map[string]interface{}
	json.Unmarshal(jsonData,&keyData)
	fmt.Printf("%#v\n",keyData)
	fmt.Println("")
	// loop through it 
	for x,y := range keyData{
		fmt.Printf("Key is %v and value is %v and the type is %T\n ", x,y,y)
	}
}

func ErrHandler(err error){
	if err != nil {
		panic(err)
	}
}