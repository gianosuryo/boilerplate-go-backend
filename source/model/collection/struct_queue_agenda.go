package collection

type CachedProgressAgendaGenerateType struct {
	GenerateID        string                                   `json:"GenerateID"`
	AgendaID          string                                   `json:"AgendaID"`
	UserUUID          string                                   `json:"UserId"`
	UserName          string                                   `json:"UserName"`
	UserEmail         string                                   `json:"UserEmail"`
	Tone              string                                   `json:"Tone"`
	Language          string                                   `json:"Language"`
	Documents         []string                                 `json:"Documents"`
	ProgressStatus    CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage   string                                   `json:"ProgressMessage"`
	Title             string                                   `json:"Title"`
	AgendaActionItems []AgendaActionItems                      `json:"AgendaActionItems"`
}

type RequestQueueAgendaGenerate struct {
	GenerateID  string              `json:"GenerateID"`
	AgendaID    string              `json:"AgendaID"`
	UserId      string              `json:"UserId"`
	UserName    string              `json:"UserName"`
	UserEmail   string              `json:"UserEmail"`
	Tone        string              `json:"Tone"`
	Language    string              `json:"Language"`
	Title       string              `json:"Title"`
	ActionItems []AgendaActionItems `json:"ActionItems"`
}
type ResponseQueueAgendaGenerate struct {
	GenerateID string `json:"GenerateID"`
	AgendaID   string `json:"AgendaID"`
	Subject    string `json:"Subject"`
	Body       string `json:"Body"`
	Status     string `json:"Status"`
	Message    string `json:"Message"`
}

type AgendaActionItems struct {
	EmailPIC          string `json:"EmailPIC"`
	DeadlineAt        string `json:"DeadlineAt"`
	ActionDescription string `json:"ActionDescription"`
	IsTimeSensitive   bool   `json:"IsTimeSensitive"`
}
