package models


type Users struct {
	Users []User `json:"users,omitempty"`
}

type User struct {
	UserName string `json:"username,omitempty"`
}


type RequestCreateUser struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ResponseCreateUser struct {
	UserName string `json:"username,omitempty"`
	UserID uint64 `json:"user_id,omitempty"`
}