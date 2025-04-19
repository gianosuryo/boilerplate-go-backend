package claim

type TokenInfo struct {
	UserId    uint64 `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
}
