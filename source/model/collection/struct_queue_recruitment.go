package collection

type CachedProgressRecruitmentGenerateType struct {
	GenerateID       string                                   `json:"GenerateID"`
	RecruitmentId    string                                   `json:"JobId"`
	TrainingId       uint64                                   `json:"TrainingId"`
	UserId           uint64                                   `json:"UserRealId"`
	UserUUID         string                                   `json:"UserId"`
	UserName         string                                   `json:"UserName"`
	UserEmail        string                                   `json:"UserEmail"`
	CompanyName      string                                   `json:"CompanyName"`
	Prompt           string                                   `json:"Prompt"`
	ProgressStatus   CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage  string                                   `json:"ProgressMessage"`
	EnableAutoAssess uint                                     `json:"EnableAutoAssess"`
	Result           ResultRecruitmentType                    `json:"Result"`
}

type CachedProgressRecruitmentAutoAssessType struct {
	IndexItem       int                                      `json:"IndexItem"`
	GenerateID      string                                   `json:"GenerateID"`
	ScreeningTestId uint64                                   `json:"ScreeningTestId"`
	ProgressStatus  CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage string                                   `json:"ProgressMessage"`
	CultureValue    string                                   `json:"CultureValue"`
	RubricScore     string                                   `json:"RubricScore"`
	EssayAnswer     string                                   `json:"EssayAnswer"`
	Question        string                                   `json:"Question"`
	VideoURL        string                                   `json:"VideoURL"`
	Result          interface{}                              `json:"Result"`
}

type RequestQueueRecruitmentGenerate struct {
	GenerateID    string `json:"GenerateID"`
	RecruitmentId string `json:"JobID"`
	UserId        string `json:"UserId"`
	UserName      string `json:"UserName"`
	UserEmail     string `json:"UserEmail"`
	CompanyName   string `json:"CompanyName"`
	Prompt        string `json:"Prompt"`
}
type ResponseQueueRecruitmentGenerate struct {
	GenerateID                      string                              `json:"GenerateID"`
	RecruitmentId                   string                              `json:"JobID"`
	JobDescription                  string                              `json:"JobDescription"`
	OutreachMessage                 string                              `json:"OutreachMessage"`
	GeneratedMultipleChoiceQuestion []RecruitmentMultipleChoiceQuestion `json:"GeneratedMultipleChoiceQuestion"`
	GeneratedCodeTestQuestion       string                              `json:"GeneratedCodeTestQuestion"`
	GeneratedEssayQuestion          []EssayQuestion                     `json:"GeneratedEssayQuestion"`
	GeneratedVideoInterview         []VideoInterviewQuestion            `json:"GeneratedVideoInterview"`
	Status                          string                              `json:"Status"`
	Message                         string                              `json:"State"`
}

type RequestQueueRecruitmentJobDescRegenerate struct {
	GenerateID        string `json:"GenerateID"`
	RecruitmentId     string `json:"JobID"`
	CompanyName       string `json:"CompanyName"`
	Prompt            string `json:"Prompt"`
	OldJobDescription string `json:"OldJobDescription"`
}

type ResponseQueueRecruitmentJobDescRegenerate struct {
	GenerateID     string `json:"GenerateID"`
	RecruitmentId  string `json:"JobID"`
	Status         string `json:"Status"`
	Message        string `json:"State"`
	JobDescription string `json:"JobDescription"`
}

type RequestQueueRecruitmentAssessmentRegenerate struct {
	RecruitmentId             string                              `json:"JobID"`
	GenerateID                string                              `json:"GenerateID"`
	Prompt                    string                              `json:"Prompt"`
	CompanyName               string                              `json:"CompanyName"`
	JobDescription            string                              `json:"JobDescription"`
	OldMultipleChoiceQuestion []RecruitmentMultipleChoiceQuestion `json:"OldMultipleChoiceQuestion,omitempty"`
	OldEssayQuestion          interface{}                         `json:"OldEssayQuestion,omitempty"`
	OldVideoInterview         interface{}                         `json:"OldVideoInterview,omitempty"`
}

type ResponseQueueRecruitmentAssessmentRegenerate struct {
	GenerateID                      string                              `json:"GenerateID"`
	RecruitmentId                   string                              `json:"JobID"`
	Status                          string                              `json:"Status"`
	Message                         string                              `json:"State"`
	GeneratedMultipleChoiceQuestion []RecruitmentMultipleChoiceQuestion `json:"GeneratedMultipleChoiceQuestion,omitempty"`
	GeneratedEssayQuestion          []EssayQuestion                     `json:"GeneratedEssayQuestion,omitempty"`
	GeneratedVideoInterview         []VideoInterviewQuestion            `json:"GeneratedVideoInterview,omitempty"`
}
type EssayQuestion struct {
	Question string `json:"Question"`
	Type     string `json:"Type"`
}
type VideoInterviewQuestion struct {
	Question string `json:"Question"`
}

type RecruitmentMultipleChoiceQuestion struct {
	Question        string                      `json:"Question"`
	MultipleChoices []RecruitmentMultipleChoice `json:"MultipleChoices"`
	Difficulty      int                         `json:"Difficulty"`
}

type RecruitmentMultipleChoice struct {
	Choice    string `json:"Choice"`
	IsCorrect bool   `json:"IsCorrect"`
}

type ResultRecruitmentType struct {
	JobDescription                  string                              `json:"JobDescription"`
	OutreachMessage                 string                              `json:"OutreachMessage"`
	GeneratedMultipleChoiceQuestion []RecruitmentMultipleChoiceQuestion `json:"GeneratedMultipleChoiceQuestion"`
	GeneratedCodeTestQuestion       string                              `json:"GeneratedCodeTestQuestion"`
	GeneratedEssayQuestion          []EssayQuestion                     `json:"GeneratedEssayQuestion"`
	GeneratedVideoInterview         []VideoInterviewQuestion            `json:"GeneratedVideoInterview"`
}
