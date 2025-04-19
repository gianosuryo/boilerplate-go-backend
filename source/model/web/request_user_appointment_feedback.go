package web

type RequestUserAppointmentFeedback struct {
	Rating           int    `validate:"required,min=1,max=5" json:"rating"`
	UserFeedbackJson string `validate:"required" json:"user_feedback_json"`
}
