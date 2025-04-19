package web

type RequestChatSessionAsk struct {
	SessionCode string `json:"session_code" form:"session_code" default:"-"`
	Message     string `json:"message" form:"message" validate:"required"`
}
