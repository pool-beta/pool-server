package models

type RequestPayDebit struct {
	UserAuth UserAuth `json:"user_auth,omitempty"`
	Name     string   `json:"name,omitempty"`
}
