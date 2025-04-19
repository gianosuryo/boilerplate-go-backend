package web

type ResponseToken struct {
	Token      string               `json:"token"`
	User       ResponseUser         `json:"user"`
}


type ResponseUser struct {
	// Id field
	Id uint64 `json:"id"`

	// Main field
	Name           string `json:"name"`
	Role           string `json:"role"`
	Email          string `json:"email"`
	UserName       string `json:"user_name"`
}
