package collection

type CachedProgressJobRecStatusType string
type JobRecPromptType string

const (
	CP_JR_QUEUE_STATUS      CachedProgressJobRecStatusType = "queue"
	CP_JR_PROGRESS_STATUS   CachedProgressJobRecStatusType = "progress"
	CP_JR_FINISH_STATUS     CachedProgressJobRecStatusType = "finish"
	CP_JR_ERROR_STATUS      CachedProgressJobRecStatusType = "error"
	CP_JR_FATALERROR_STATUS CachedProgressJobRecStatusType = "fatalError"
)

const (
	JR_TYPE_SKILLS_GAP     JobRecPromptType = "skills_gap"
	JR_TYPE_RESUME         JobRecPromptType = "resume"
	JR_TYPE_INTERVIEW_PREP JobRecPromptType = "interview_prep"
)

type SelectedCoursesType struct {
	// Main field
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	ImageFileUrl string `json:"image_file_url"`
	Description  string `json:"description"`
}

type JobRecAvailableCourses struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
}
type JobRecSkillsGapResult struct {
	Description string                `json:"Description"`
	CourseList  []string              `json:"CourseList"`
	CourseData  []SelectedCoursesType `json:"CourseData"`
}
type CachedProgressJobRec struct {
	GenerateID          string                         `json:"GenerateID"`
	JobRecID            string                         `json:"JobRecID"`
	ProgressStatus      CachedProgressJobRecStatusType `json:"ProgressStatus"`
	ProgressMessage     string                         `json:"ProgressMessage"`
	SkillsGapResult     JobRecSkillsGapResult          `json:"SkillsGapResult"`
	ResumeResult        interface{}                    `json:"ResumeResult"`
	InterviewPrepResult string                         `json:"InterviewPrepResult"`
	UserExperience      string                         `json:"UserExperience"`
	JobDescription      string                         `json:"JobDescription"`
	AvailableCourses    []JobRecAvailableCourses       `json:"AvailableCourses"`
	SelectedCourses     []SelectedCoursesType          `json:"SelectedCourses"`
	PromptType          JobRecPromptType               `json:"PromptType"`
	UserId              string                         `json:"UserId"`
	UserName            string                         `json:"UserName"`
	UserEmail           string                         `json:"UserEmail"`
	CreatedAt           string                         `json:"CreatedAt"`
	TimeTakenSeconds    float32                        `json:"TimeTakenSeconds"`
	Pricing             interface{}                    `json:"Pricing"`
}

type RequestQueueJobRec struct {
	GenerateID       string                   `json:"GenerateID"`
	JobRecID         string                   `json:"JobRecID"`
	UserId           string                   `json:"UserId"`
	UserName         string                   `json:"UserName"`
	UserEmail        string                   `json:"UserEmail"`
	UserExperience   string                   `json:"UserExperience"`
	JobDescription   string                   `json:"JobDescription"`
	AvailableCourses []JobRecAvailableCourses `json:"AvailableCourses"`
	PromptType       JobRecPromptType         `json:"PromptType"`
}

type ResponseQueueJobRec struct {
	GenerateID          string                `json:"GenerateID"`
	Status              string                `json:"Status"`
	Message             string                `json:"Message"`
	SkillsGapResult     JobRecSkillsGapResult `json:"SkillsGapResult"`
	InterviewPrepResult string                `json:"InterviewPrepResult"`
	Pricing             interface{}           `json:"Pricing"`
	TimeTakenSeconds    float32               `json:"TimeTakenSeconds"`
	PromptType          JobRecPromptType      `json:"PromptType"`
}
