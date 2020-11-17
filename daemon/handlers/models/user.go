package models

type UserName struct {
	UserName string `json:"username,omitempty"`
}

type UserAuth struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ResponseGetUsers struct {
	UserNames []UserName `json:"users,omitempty"`
}

type RequestCreateUser struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ResponseCreateUser struct {
	UserName string `json:"username,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
}
