package models

type RequestCreatePool struct {
	UserAuth UserAuth `json:"user_auth,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type ResponseCreatePool struct {
	UserName string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"pool_type,omitempty"`
	ID       uint64 `json:"pool_id,omitempty"`
}
