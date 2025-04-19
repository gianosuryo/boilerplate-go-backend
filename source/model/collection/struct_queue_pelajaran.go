package collection

type CachedProgressSkeletonPelajaranGenerateType struct {
	GenerateID             string                                   `json:"GenerateID"`
	CourseID               string                                   `json:"CourseID"`
	UserUUID               string                                   `json:"UserId"`
	UserName               string                                   `json:"UserName"`
	UserEmail              string                                   `json:"UserEmail"`
	Tone                   string                                   `json:"Tone"`
	Language               string                                   `json:"Language"`
	Prompt                 string                                   `json:"Prompt"`
	Documents              []string                                 `json:"Documents"`
	ProgressStatus         CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage        string                                   `json:"ProgressMessage"`
	TrainingData           TrainingDataPelajaranType                `json:"TrainingData"`
	InProgressChapterCount int                                      `json:"InProgressChapterCount"`
}
type RequestQueueSkeletonPelajaranType struct {
	GenerateID string   `json:"GenerateID"`
	CourseID   string   `json:"CourseID"`
	Documents  []string `json:"Documents"`
	UserId     string   `json:"UserId"`
	UserName   string   `json:"UserName"`
	UserEmail  string   `json:"UserEmail"`
	Tone       string   `json:"Tone"`
	Language   string   `json:"Language"`
	Prompt     string   `json:"Prompt"`
}
type ResponseQueueSkeletonPelajaranType struct {
	GenerateID   string                    `json:"GenerateID"`
	CourseID     string                    `json:"CourseID"`
	TrainingData TrainingDataPelajaranType `json:"Data"`
	Status       string                    `json:"Status"`
	Message      string                    `json:"Message"`
}

type RequestQueuePelajaranChapterType struct {
	GenerateID         string `json:"GenerateID"`
	CourseID           string `json:"CourseID"`
	ChapterID          string `json:"ChapterID"`
	Prompt             string `json:"Prompt"`
	ChapterContentType string `json:"ContentType"`
	Tone               string `json:"Tone"`
	Language           string `json:"Language"`
}

type ResponseQueuePelajaranChapterType struct {
	GenerateID                    string                    `json:"GenerateID"`
	CourseID                      string                    `json:"CourseID"`
	ChapterID                     string                    `json:"ChapterID"`
	Title                         string                    `json:"Title"`
	ContentMarkdown               string                    `json:"ContentMarkdown"`
	ContentMultipleChoiceQuestion []ChapterExamDataType     `json:"ContentMultipleChoiceQuestion"`
	ContentYoutube                YoutubeEmbedDataType      `json:"ContentYoutube"`
	ContentProjectBasedLearning   string                    `json:"ContentProjectBasedLearning"`
	ProjectBrief                  string                    `json:"ProjectBrief"`
	RubricScore                   string                    `json:"RubricScore"`
	Status                        string                    `json:"Status"`
	Message                       string                    `json:"State"`
	Pricing                       ResponseGeneratePriceType `json:"Pricing"`
	TimeTakenSeconds              float32                   `json:"TimeTakenSeconds"`
}

type RequestQueueRegenerateCourseDetail struct {
	GenerateID          string   `json:"GenerateID"`
	CourseID            string   `json:"CourseID"`
	ChapterID           string   `json:"ChapterID"`
	ContentType         string   `json:"ContentType"`
	Prompt              string   `json:"Prompt"`
	CourseDetailHistory []string `json:"ContentHistory"`
}
type ResponseQueueRegenerateCourseDetail struct {
	GenerateID                    string                `json:"GenerateID"`
	CourseID                      string                `json:"CourseID"`
	ChapterID                     string                `json:"ChapterID"`
	ContentMarkdown               string                `json:"Content"`
	ContentMultipleChoiceQuestion []ChapterExamDataType `json:"ContentMultipleChoiceQuestion"`
	ContentYoutube                YoutubeEmbedDataType  `json:"ContentYoutube"`
	ContentProjectBasedLearning   string                `json:"ContentProjectBasedLearning"`
	Status                        string                `json:"Status"`
	Message                       string                `json:"Message"`
}
type RequestQueueEssaySubmission struct {
	GenerateID   string `json:"GenerateID"`
	CourseID     string `json:"CourseID"`
	ChapterID    string `json:"ChapterID"`
	ProjectBrief string `json:"ProjectBrief"`
	RubricScore  string `json:"RubricScore"`
	UserResponse string `json:"UserResponse"`
}
type ResponseQueueEssaySubmission struct {
	GenerateID string `json:"GenerateID"`
	CourseID   string `json:"CourseID"`
	ChapterID  string `json:"ChapterID"`
	Content    string `json:"Content"`
	Status     string `json:"Status"`
	Message    string `json:"State"`
}

type TrainingDataPelajaranType struct {
	Title       string                 `json:"Title"`
	Description string                 `json:"Description"`
	Chapters    []ChapterPelajaranType `json:"Chapters"`
}

type ChapterPelajaranType struct {
	Title       string                    `json:"Title"`
	SubChapters []SubChapterPelajaranType `json:"SubChapters"`
}

type SubChapterPelajaranType struct {
	ChapterID   string `json:"ChapterID"`
	Prompt      string `json:"Prompt"`
	Title       string `json:"Title"`
	ContentType string `json:"ContentType"`
	Status      string `json:"Status"`
}
