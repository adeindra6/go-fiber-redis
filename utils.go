package main

import "encoding/json"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  `json:"address"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func toJson(val []byte) User {
	user := User{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}

	return user
}
