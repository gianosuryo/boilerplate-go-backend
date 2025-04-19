package web

type RequestMessaging struct {
	Body          string `json:"body,omitempty" form:"body"`
	Subject       string `json:"subject,omitempty" form:"subject"`
	RecepientList string `json:"recepient_list,omitempty" form:"recepient_list"`
}
