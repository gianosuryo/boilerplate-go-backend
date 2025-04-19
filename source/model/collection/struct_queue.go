package collection

import (
	"html/template"
	"time"
)

type GenerateType string
type QueueTrainingGenerateStatusType string
type CachedProgressTrainingGenerateStatusType string

const (
	GT_COMPLETE             GenerateType = "Generate Complete"
	GT_REGENERATE_THUMBNAIL GenerateType = "Regenerate Thumbnail"
	GT_REGENERATE_CHAPTER   GenerateType = "Regenerate Chapter"
)

const (
	Q_TG_QUEUE_STATUS      QueueTrainingGenerateStatusType = "Queue"
	Q_TG_PROGRESS_STATUS   QueueTrainingGenerateStatusType = "Progress"
	Q_TG_FINISH_STATUS     QueueTrainingGenerateStatusType = "Finish"
	Q_TG_ERROR_STATUS      QueueTrainingGenerateStatusType = "Error"
	Q_TG_FATALERROR_STATUS QueueTrainingGenerateStatusType = "FatalError"
)

const (
	CP_TG_QUEUE_STATUS      CachedProgressTrainingGenerateStatusType = "queue"
	CP_TG_PROGRESS_STATUS   CachedProgressTrainingGenerateStatusType = "progress"
	CP_TG_FINISH_STATUS     CachedProgressTrainingGenerateStatusType = "finish"
	CP_TG_ERROR_STATUS      CachedProgressTrainingGenerateStatusType = "error"
	CP_TG_FATALERROR_STATUS CachedProgressTrainingGenerateStatusType = "fatalError"

	CP_TG_DOCUMENT_QUEUE_STATUS       CachedProgressTrainingGenerateStatusType = "document_queue"
	CP_TG_DOCUMENT_PROGRESS_STATUS    CachedProgressTrainingGenerateStatusType = "document_progress"
	CP_TG_DOCUMENT_FINISH_STATUS      CachedProgressTrainingGenerateStatusType = "document_finish"
	CP_TG_DOCUMENT_ERROR_STATUS       CachedProgressTrainingGenerateStatusType = "document_error"
	CP_TG_DOCUMENT_FATALERROR_STATUS  CachedProgressTrainingGenerateStatusType = "document_fatalerror"
	CP_TG_DOCUMENT_UNPROCESSED_STATUS CachedProgressTrainingGenerateStatusType = "document_unprocessed"

	CP_TG_SKELETON_QUEUE_STATUS       CachedProgressTrainingGenerateStatusType = "skeleton_queue"
	CP_TG_SKELETON_PROGRESS_STATUS    CachedProgressTrainingGenerateStatusType = "skeleton_progress"
	CP_TG_SKELETON_FINISH_STATUS      CachedProgressTrainingGenerateStatusType = "skeleton_finish"
	CP_TG_SKELETON_ERROR_STATUS       CachedProgressTrainingGenerateStatusType = "skeleton_error"
	CP_TG_SKELETON_FATALERROR_STATUS  CachedProgressTrainingGenerateStatusType = "skeleton_fatalerror"
	CP_TG_SKELETON_UNPROCESSED_STATUS CachedProgressTrainingGenerateStatusType = "skeleton_unprocessed"

	CP_TG_THUMBNAIL_QUEUE_STATUS      CachedProgressTrainingGenerateStatusType = "thumbnail_queue"
	CP_TG_THUMBNAIL_PROGRESS_STATUS   CachedProgressTrainingGenerateStatusType = "thumbnail_progress"
	CP_TG_THUMBNAIL_FINISH_STATUS     CachedProgressTrainingGenerateStatusType = "thumbnail_finish"
	CP_TG_THUMBNAIL_ERROR_STATUS      CachedProgressTrainingGenerateStatusType = "thumbnail_error"
	CP_TG_THUMBNAIL_FATALERROR_STATUS CachedProgressTrainingGenerateStatusType = "thumbnail_fatalerror"

	CP_TG_CHAPTER_QUEUE_STATUS      CachedProgressTrainingGenerateStatusType = "chapter_queue"
	CP_TG_CHAPTER_PROGRESS_STATUS   CachedProgressTrainingGenerateStatusType = "chapter_progress"
	CP_TG_CHAPTER_FINISH_STATUS     CachedProgressTrainingGenerateStatusType = "chapter_finish"
	CP_TG_CHAPTER_ERROR_STATUS      CachedProgressTrainingGenerateStatusType = "chapter_error"
	CP_TG_CHAPTER_FATALERROR_STATUS CachedProgressTrainingGenerateStatusType = "chapter_fatalerror"

	CP_TG_ASSESSMENT_QUEUE_STATUS       CachedProgressTrainingGenerateStatusType = "assessment_queue"
	CP_TG_ASSESSMENT_PROGRESS_STATUS    CachedProgressTrainingGenerateStatusType = "assessment_progress"
	CP_TG_ASSESSMENT_FINISH_STATUS      CachedProgressTrainingGenerateStatusType = "assessment_finish"
	CP_TG_ASSESSMENT_ERROR_STATUS       CachedProgressTrainingGenerateStatusType = "assessment_error"
	CP_TG_ASSESSMENT_FATALERROR_STATUS  CachedProgressTrainingGenerateStatusType = "assessment_fatalerror"
	CP_TG_ASSESSMENT_UNPROCESSED_STATUS CachedProgressTrainingGenerateStatusType = "assessment_unprocessed"
)

type DocumentAvailability struct {
	DocumentUUID    string `json:"DocumentUUID"`
	IsAvailable     bool   `json:"IsAvailable"`
	HasKnowledge    bool   `json:"HasKnowledge"`
	FatalErrorCount uint   `json:"FatalErrorCount"`
}

type ChapterDataType struct {
	ChapterTitle       string   `json:"ChapterTitle"`
	ChapterOverview    []string `json:"ChapterOverview,omitempty"`
	ChapterContentType string   `json:"ChapterContentType"`
	ChapterContent     string   `json:"ChapterContent"`
	ChapterContentExam []string `json:"ChapterContentExam"`
}

type CourseDataType struct {
	CourseTitle       string            `json:"CourseTitle"`
	CourseImage       string            `json:"CourseImage"`
	CourseDescription string            `json:"CourseDescription"`
	ChapterData       []ChapterDataType `json:"ChapterData"`
}

type YoutubeEmbedDataType struct {
	EmbedURL string `json:"YoutubeURL"`
	Title    string `json:"YoutubeTitle"`
}

type ChapterExamDataType struct {
	Question        string                               `json:"Question"`
	MultipleChoices []ChapterExamMultipleChoicesDataType `json:"MultipleChoices"`
	Difficulty      int                                  `json:"Difficulty"`
}

type AssessmentAttachmentDataType struct {
	Question   string `json:"Question"`
	Difficulty int    `json:"Difficulty"`
}
type AssessmentEssayDataType struct {
	Question   string `json:"Question"`
	Difficulty int    `json:"Difficulty"`
}

type ChapterExamMultipleChoicesDataType struct {
	Choice    string `json:"Choice"`
	IsCorrect bool   `json:"IsCorrect"`
}

type ResponseGeneratePriceType struct {
	InputToken  uint    `json:"InputToken"`
	OutputToken uint    `json:"OutputToken"`
	InputChar   uint    `json:"InputChar"`
	OutputChar  uint    `json:"OutputChar"`
	InputPrice  float32 `json:"InputPrice"`
	OutputPrice float32 `json:"OutputPrice"`
	ImageUsage  uint    `json:"ImageUsage"`
	ImagePrice  float32 `json:"ImagePrice"`
	TotalPrice  float32 `json:"TotalPrice"`
}

type RequestQueueDocumentType struct {
	GenerateID     string   `json:"GenerateID"`
	DocumentID     string   `json:"DocumentID"`
	DocumentIDs    []string `json:"DocumentIDs,omitempty"`
	FileURL        string   `json:"FileURL"`
	Tone           string   `json:"Tone"`
	Language       string   `json:"Language"`
	AdditionalNote string   `json:"AdditionalNotes"`
	SubCompanyId   string   `json:"SubCompanyId"`
	SubCompanyName string   `json:"SubCompanyName"`
	UserId         string   `json:"UserId"`
	UserName       string   `json:"UserName"`
	UserEmail      string   `json:"UserEmail"`
	ApiKey         string   `json:"ApiKey"`
}

type ResponseQueueDocumentType struct {
	GenerateID       string                    `json:"GenerateID"`
	DocumentID       string                    `json:"DocumentID"`
	Status           string                    `json:"Status"`
	Message          string                    `json:"Message"`
	HasKnowledge     bool                      `json:"HasKnowledge"`
	TimeTakenSeconds float32                   `json:"TimeTakenSeconds"`
	ApiKey           string                    `json:"ApiKey"`
	Pricing          ResponseGeneratePriceType `json:"Pricing"`
}

type RequestQueueThumbnailType struct {
	GenerateID     string `json:"GenerateID"`
	CourseID       string `json:"CourseID"`
	SubCompanyId   string `json:"SubCompanyId"`
	SubCompanyName string `json:"SubCompanyName"`
	UserId         string `json:"UserId"`
	UserName       string `json:"UserName"`
	UserEmail      string `json:"UserEmail"`
	ApiKey         string `json:"ApiKey"`
}

type ResponseQueueThumbnailType struct {
	GenerateID       string                    `json:"GenerateID"`
	CourseID         string                    `json:"CourseID"`
	Status           string                    `json:"Status"`
	Message          string                    `json:"Message"`
	ThumbnailURL     string                    `json:"ThumbnailURL"`
	TimeTakenSeconds float32                   `json:"TimeTakenSeconds"`
	ApiKey           string                    `json:"ApiKey"`
	Pricing          ResponseGeneratePriceType `json:"Pricing"`
}

type RequestQueueSkeletonType struct {
	GenerateID      string   `json:"GenerateID"`
	CourseID        string   `json:"CourseID"`
	DocumentIDs     []string `json:"DocumentIDs"`
	SubCompanyId    string   `json:"SubCompanyId"`
	SubCompanyName  string   `json:"SubCompanyName"`
	UserId          string   `json:"UserId"`
	UserName        string   `json:"UserName"`
	UserEmail       string   `json:"UserEmail"`
	Tone            string   `json:"Tone"`
	Language        string   `json:"Language"`
	AdditionalNotes string   `json:"AdditionalNotes"`
	ApiKey          string   `json:"ApiKey"`
}

type ResponseQueueSkeletonType struct {
	GenerateID       string                    `json:"GenerateID"`
	CourseID         string                    `json:"CourseID"`
	CourseData       CourseDataType            `json:"CourseData"`
	Status           string                    `json:"Status"`
	Message          string                    `json:"Message"`
	Pricing          ResponseGeneratePriceType `json:"Pricing"`
	TimeTakenSeconds float32                   `json:"TimeTakenSeconds"`
	ApiKey           string                    `json:"ApiKey"`
}

type RequestQueueChapterType struct {
	GenerateID         string   `json:"GenerateID"`
	CourseID           string   `json:"CourseID"`
	ChapterID          string   `json:"ChapterID"`
	ChapterTitle       string   `json:"ChapterTitle"`
	ChapterOverview    []string `json:"ChapterOverview"`
	ChapterContentType string   `json:"ChapterContentType"`
	Tone               string   `json:"Tone"`
	Language           string   `json:"Language"`
	AdditionalNotes    string   `json:"AdditionalNotes"`
	SubCompanyId       string   `json:"SubCompanyId"`
	SubCompanyName     string   `json:"SubCompanyName"`
	UserId             string   `json:"UserId"`
	UserName           string   `json:"UserName"`
	UserEmail          string   `json:"UserEmail"`
	ApiKey             string   `json:"ApiKey"`
}

type ResponseQueueChapterType struct {
	GenerateID         string                    `json:"GenerateID"`
	CourseID           string                    `json:"CourseID"`
	ChapterID          string                    `json:"ChapterID"`
	ChapterContent     string                    `json:"ChapterContent"`
	ChapterContentExam []ChapterExamDataType     `json:"ChapterContentExam"`
	Status             string                    `json:"Status"`
	Message            string                    `json:"Message"`
	Pricing            ResponseGeneratePriceType `json:"Pricing"`
	TimeTakenSeconds   float32                   `json:"TimeTakenSeconds"`
	ApiKey             string                    `json:"ApiKey"`
}

type RequestQueueAssessmentType struct {
	GenerateID   string `json:"GenerateID"`
	CourseID     string `json:"CourseID"`
	ChapterID    string `json:"ChapterID"`
	PromptText   string `json:"PromptText"`
	Language     string `json:"Language"`
	Type         string `json:"Type"` // Essay, MultipleChoices, PPT, XLS
	ContentCount int    `json:"ContentCount"`
	UserId       string `json:"UserId"`
	UserName     string `json:"UserName"`
	UserEmail    string `json:"UserEmail"`
	ApiKey       string `json:"ApiKey"`
}

type ResponseQueueAssessmentType struct {
	GenerateID               string                         `json:"GenerateID"`
	CourseID                 string                         `json:"CourseID"`
	ChapterID                string                         `json:"ChapterID"`
	Status                   string                         `json:"Status"`
	Message                  string                         `json:"Message"`
	ContentTitle             string                         `json:"ContentTitle"`
	Type                     string                         `json:"Type"` // Essay, MultipleChoices, PPT, XLS, AttachmentVideo
	ContentAttachmentGeneral []AssessmentAttachmentDataType `json:"ContentAttachmentGeneral"`
	ContentAttachmentXLS     []AssessmentAttachmentDataType `json:"ContentAttachmentXLS"`
	ContentAttachmentPPT     []AssessmentAttachmentDataType `json:"ContentAttachmentPPT"`
	ContentAttachmentVideo   []AssessmentAttachmentDataType `json:"ContentAttachmentVideo"`
	ContentEssay             []AssessmentEssayDataType      `json:"ContentEssay"`
	ContentMCQ               []ChapterExamDataType          `json:"ContentMCQ"`
	ContentYoutubeEmbed      YoutubeEmbedDataType           `json:"ContentYoutubeEmbed"`
	Pricing                  ResponseGeneratePriceType      `json:"Pricing"`
	TimeTakenSeconds         float32                        `json:"TimeTakenSeconds"`
	ApiKey                   string                         `json:"ApiKey"`
}

type RequestQueueEmailFeedback struct {
	GenerateID      string  `json:"GenerateID"`
	EmailFeedbackID string  `json:"EmailFeedbackID"`
	Name            string  `json:"Name"`
	Email           string  `json:"Email"`
	Feedback        string  `json:"Feedback"`
	SatisfyRate     float32 `json:"SatisfyRate"`
	CompetenceRate  float32 `json:"CompetenceRate"`
	UserId          string  `json:"UserId"`
	UserName        string  `json:"UserName"`
	UserEmail       string  `json:"UserEmail"`
	ApiKey          string  `json:"ApiKey"`
}

type ResponseQueueEmailFeedback struct {
	GenerateID       string      `json:"GenerateID"`
	EmailFeedbackID  string      `json:"EmailFeedbackID"`
	Status           string      `json:"Status"`
	Message          string      `json:"Message"`
	Subject          string      `json:"Subject"`
	Body             string      `json:"Body"`
	Pricing          interface{} `json:"Pricing"`
	TimeTakenSeconds float32     `json:"TimeTakenSeconds"`
	ApiKey           string      `json:"ApiKey"`
}

type CachedProgressTrainingGenerateType struct {
	GenerateID      string                                   `json:"GenerateID"`
	ProgressStatus  CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage string                                   `json:"ProgressMessage"`
	TrainingData    TrainingDataQueueMessageType             `json:"TrainingData"`
	PricingObject   interface{}                              `json:"Pricing"`
}

type CachedProgressTrainingDetailGenerateType struct {
	GenerateID                            string                                   `json:"GenerateID"`
	CourseUUID                            string                                   `json:"CourseUUID"`
	ChapterUUID                           string                                   `json:"ChapterUUID"`
	TrainingId                            string                                   `json:"TrainingId"`
	TrainingDetailId                      string                                   `json:"TrainingDetailId"`
	ProgressStatus                        CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage                       string                                   `json:"ProgressMessage"`
	TrainingDetailTitle                   string                                   `json:"TrainingDetailTitle"`
	TrainingDetailContentType             string                                   `json:"TrainingDetailContentType"`
	TrainingDetailContent                 string                                   `json:"TrainingDetailContent"`
	TrainingDetailContentAttachmentBundle []AssessmentAttachmentDataType           `json:"TrainingDetailContentAttachmentBundle"`
	TrainingDetailContentEssay            []AssessmentEssayDataType                `json:"TrainingDetailContentEssay"`
	TrainingDetailContentExam             []ChapterExamDataType                    `json:"TrainingDetailContentExam"`
	TrainingDetailContentYoutubeEmbed     YoutubeEmbedDataType                     `json:"TrainingDetailContentYoutubeEmbed"`
	TrainingDetailContentYoutubeUrl       string                                   `json:"TrainingDetailContentYoutubeUrl"`
	PricingObject                         interface{}                              `json:"Pricing"`
	TimeTakenSeconds                      float32                                  `json:"TimeTakenSeconds"`
	VersionIndex                          int                                      `json:"VersionIndex"`
}

type CachedProgressTrainingDetailGenerateHistoryType struct {
	GenerateID                            string                           `json:"GenerateID"`
	CourseUUID                            string                           `json:"CourseUUID"`
	ChapterUUID                           string                           `json:"ChapterUUID"`
	TrainingDetailContentType             string                           `json:"TrainingDetailContentType"`
	TrainingDetailContent                 []string                         `json:"TrainingDetailContent"`
	TrainingDetailContentAttachmentBundle [][]AssessmentAttachmentDataType `json:"TrainingDetailContentAttachmentBundle"`
	TrainingDetailContentEssay            [][]AssessmentEssayDataType      `json:"TrainingDetailContentEssay"`
	TrainingDetailContentExam             [][]ChapterExamDataType          `json:"TrainingDetailContentExam"`
	TrainingDetailContentYoutubeEmbed     []YoutubeEmbedDataType           `json:"TrainingDetailContentYoutubeEmbed"`
	TrainingDetailContentYoutubeUrl       []string                         `json:"TrainingDetailContentYoutubeUrl"`
}

type CachedProgressEssaySubmissionType struct {
	GenerateID      string                                   `json:"GenerateID"`
	CourseUUID      string                                   `json:"CourseUUID"`
	ChapterUUID     string                                   `json:"ChapterUUID"`
	ResutlEssayID   string                                   `json:"ResutlEssayID"`
	RubricScore     string                                   `json:"RubricScore"`
	ProjectBrief    string                                   `json:"ProjectBrief"`
	Result          string                                   `json:"Result"`
	ProgressStatus  CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage string                                   `json:"ProgressMessage"`
}

type CachedProgressEmailFeedbackGenerateType struct {
	GenerateID        string                                   `json:"GenerateID"`
	ProgressStatus    CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage   string                                   `json:"ProgressMessage"`
	EmailFeedbackData EmailFeedbackDataQueueMessageType        `json:"EmailFeedbackData"`
	PricingObject     interface{}                              `json:"Pricing"`
	TimeTakenSeconds  float32                                  `json:"TimeTakenSeconds"`
	GeneratedAt       time.Time                                `json:"GeneratedAt"`
}

type CachedProgressGenerateLogType struct {
	GenerateID              string                                   `json:"GenerateID"`
	UserUUID                string                                   `json:"UserId"`
	UserName                string                                   `json:"UserName"`
	UserEmail               string                                   `json:"UserEmail"`
	Tone                    string                                   `json:"Tone"`
	Language                string                                   `json:"Language"`
	AdditionalNotes         string                                   `json:"AdditionalNotes"`
	DocumentUUIDs           []string                                 `json:"DocumentUUIDs"`
	ProgressStatus          CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage         string                                   `json:"ProgressMessage"`
	SkeletonFatalErrorCount uint                                     `json:"SkeletonFatalErrorCount"`
	StreamToken             string                                   `json:"StreamToken"`
}

type CachedProgressElasticGenerateLogType struct {
	ApiKeyUUID               string                                       `json:"ApiKeyUUID"`
	ApiKeyName               string                                       `json:"ApiKeyName"`
	GenerateID               string                                       `json:"GenerateID"`
	GenerateType             GenerateType                                 `json:"GenerateType"` // CompleteGenerate, RegenerateThumbnail, RegenerateChapter
	CourseUUID               string                                       `json:"CourseUUID"`
	UserUUID                 string                                       `json:"UserUUID"`
	UserName                 string                                       `json:"UserName"`
	UserEmail                string                                       `json:"UserEmail"`
	Tone                     string                                       `json:"Tone"`
	Language                 string                                       `json:"Language"`
	AdditionalNotes          string                                       `json:"AdditionalNotes"`
	CourseTitle              string                                       `json:"CourseTitle"`
	CourseImage              string                                       `json:"CourseImage"`
	CourseDescription        string                                       `json:"CourseDescription"`
	ChapterAllCount          uint                                         `json:"ChapterCount"`
	ChapterCompleteCount     uint                                         `json:"ChapterCompleteCount"`
	TotalPrice               float32                                      `json:"TotalPrice"`
	TotalPriceStr            string                                       `json:"TotalPriceStr"`
	TotalTimeTaken           float32                                      `json:"TotalTimeTaken"`
	DetailElasticGenerateLog []CachedProgressElasticGenerateLogDetailType `json:"DetailLog"`
}

type CachedProgressElasticGenerateLogDetailType struct {
	ChapterUUID   string   `json:"ChapterUUID"`
	DocumentUUID  string   `json:"DocumentUUID"`
	ChapterUUIDs  []string `json:"ChapterUUIDs"`
	Thumbnail     string   `json:"Thumbnail"`
	Type          string   `json:"Type"`
	TimeTaken     float32  `json:"TimeTaken"`
	Price         float32  `json:"Price"`
	PriceStr      string   `json:"PriceStr"`
	PricingObject string   `json:"PricingObject"`
}

// Log detail done generate - chapter uuid (if exist), type generate(document, skeleton, thumbnail, chapter), time taken, price, price object, character

// InputToken: int
// OutputToken: int
// InputChar: int
// OutputChar: int
// InputPrice: float
// OutputPrice: float
// ImageUsage: int
// ImagePrice: float
// TotalPrice: float

type EmailFeedbackDataQueueMessageType struct {
	EmailFeedbackUUID string        `json:"EmailFeedbackUUID"`
	EmailFeedbackId   string        `json:"EmailFeedbackId"`
	Subject           string        `json:"Subject"`
	Body              string        `json:"Body"`
	BodyParsed        template.HTML `json:"BodyParsed"`
	Name              string        `json:"Name"`
	Email             string        `json:"Email"`
	Feedback          string        `json:"Feedback"`
	CompetenceRate    float32       `json:"CompetenceRate"`
	SatisfyRate       float32       `json:"SatisfyRate"`
	BannerMoodUrl     string        `json:"BannerMoodUrl"`
}
type TrainingDataQueueMessageType struct {
	CourseUUID          string                               `json:"CourseUUID"`
	TrainingId          string                               `json:"TrainingId"`
	TrainingTitle       string                               `json:"TrainingTitle"`
	TrainingImage       string                               `json:"TrainingImage"`
	TrainingDescription string                               `json:"TrainingDescription"`
	TrainingDetailData  []TrainingDetailDataQueueMessageType `json:"TrainingDetailData"`
}

type TrainingDetailDataQueueMessageType struct {
	TrainingDetailTitle       string                                   `json:"TrainingDetailTitle"`
	TrainingDetailOverview    []string                                 `json:"TrainingDetailOverview,omitempty"`
	TrainingDetailContentType string                                   `json:"TrainingDetailContentType"`
	TrainingId                string                                   `json:"TrainingId"`
	TrainingDetailId          string                                   `json:"TrainingDetailId"`
	CourseUUID                string                                   `json:"CourseUUID"`
	ChapterUUID               string                                   `json:"ChapterUUID"`
	ProgressStatus            CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	Message                   string                                   `json:"Message"`
	TimeTakenSeconds          float32                                  `json:"TimeTakenSeconds"`

	TrainingDetailContent     string                                   `json:"TrainingDetailContent"`
	TrainingDetailContentExam []TrainingDetailExamDataQueueMessageType `json:"TrainingDetailContentExam"`
}

type TrainingDetailExamDataQueueMessageType struct {
	Question        string                                                  `json:"Question"`
	MultipleChoices []TrainingDetailExamMultipleChoicesDataQueueMessageType `json:"MultipleChoices"`
}

type TrainingDetailExamMultipleChoicesDataQueueMessageType struct {
	Choice    string `json:"Choice"`
	IsCorrect bool   `json:"IsCorrect"`
}
