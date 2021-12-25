package structs

type NewTokenReq struct {
	UserID string `json:"user_id"`
}

type NewTokenResp struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

type CheckTokenReq struct {
	Token string `json:"token"`
}

type CheckTokenResp struct {
	UserID string `json:"user_id"`
	OK     bool   `json:"ok"`

	Error string `json:"error"`
}

type RegisterReq struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RegisterResp struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Error  string `json:"error"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Error  string `json:"error"`
}
