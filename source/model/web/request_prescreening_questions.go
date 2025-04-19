package web

type RequestPrescreeningQuestions struct {
	QuestionJson string `json:"question_json" form:"question_json" validate:"required"`
}

type RequestPrescreeningQuestionJson struct {
	AnswerType string   `json:"answer_type" form:"answer_type"`
	Question   string   `json:"question" form:"question"`
	Options    []string `json:"options" form:"options"`
}

type RequestGeneratePrescreeningQuestionsByHealthServiceIds struct {
	HealthServiceIds []uint64 `json:"health_service_ids" form:"health_service_ids" validate:"required"`
}
