package models

type Test struct {
	Test string `json:"test,omitempty"`
}

type TestUser struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
}

/* RequestSetupTestEnv */
type RequestSetupTestEnv struct {
	Secret string `json:"password,omitempty"`
}

type ResponseSetupTestEnv struct {
	Users []TestUser `json:"users,omitempty"`
}
