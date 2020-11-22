package models

/* CreatePool */
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

/* GetPool */
type RequestGetPool struct {
	UserAuth UserAuth `json:"user_auth,omitempty"`
	ID       uint64   `json:"pool_id,omitempty"`
}

type ResponseGetPool struct {
	Name string `json:"name,omitempty"`
	Type string `json:"pool_type,omitempty"`
	ID   uint64 `json:"pool_id,omitempty"`
}

/* GetAllPools */
type RequestGetAllPools struct {
	UserAuth UserAuth `json:"user_auth,omitempty"`
}

type ResponseGetAllPools struct {
}
