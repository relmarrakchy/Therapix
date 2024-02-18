package models

type User struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type SignUpData struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Response struct {
	Response string `json:"Response"`
	Data     string `json:"Data"`
}

type LoginData struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type Verify struct {
	Id       string `json:"_id"`
	Rev      string `json:"Rev"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
