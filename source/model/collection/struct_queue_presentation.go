package collection

type CachedProgressPresentationGenerateType struct {
	GenerateID      string                                   `json:"GenerateID"`
	PresentationID  string                                   `json:"PresentationID"`
	UserUUID        string                                   `json:"UserId"`
	UserName        string                                   `json:"UserName"`
	UserEmail       string                                   `json:"UserEmail"`
	Tone            string                                   `json:"Tone"`
	Language        string                                   `json:"Language"`
	Prompt          string                                   `json:"Prompt"`
	Documents       []string                                 `json:"Documents"`
	ProgressStatus  CachedProgressTrainingGenerateStatusType `json:"ProgressStatus"`
	ProgressMessage string                                   `json:"ProgressMessage"`
	Title           string                                   `json:"Title"`
	SlideContent    []PresentationSlideContent               `json:"SlideContent"`
}

type RequestQueuePresentationGenerate struct {
	GenerateID     string   `json:"GenerateID"`
	PresentationID string   `json:"PresentationID"`
	Documents      []string `json:"Documents"`
	UserId         string   `json:"UserId"`
	UserName       string   `json:"UserName"`
	UserEmail      string   `json:"UserEmail"`
	Tone           string   `json:"Tone"`
	Language       string   `json:"Language"`
	Prompt         string   `json:"Prompt"`
}
type ResponseQueuePresentationGenerate struct {
	GenerateID     string                     `json:"GenerateID"`
	PresentationID string                     `json:"PresentationID"`
	Title          string                     `json:"Title"`
	SlideContent   []PresentationSlideContent `json:"Slides"`
	Status         string                     `json:"Status"`
	Message        string                     `json:"State"`
}

type PresentationSlideContent struct {
	SlideType              string                                  `json:"SlideType"`
	Content                interface{}                             `json:"Content"`
	ContentTitle           PresentationSlideContentTitle           `json:"ContentTitle"`
	ContentBody            PresentationSlideContentBody            `json:"ContentBody"`
	ContentConclusion      PresentationSlideContentConclusion      `json:"ContentConclusion"`
	ContentTableOfContents PresentationSlideContentTableOfContents `json:"ContentTableOfContents"`
}

type PresentationSlideContentTitle struct {
	Title string `json:"Title"`
	Body  string `json:"Subtitle"`
}
type PresentationSlideContentBody struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
}
type PresentationSlideContentConclusion struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
}
type PresentationSlideContentTableOfContents struct {
	Title   string   `json:"Title"`
	Content []string `json:"Content"`
}
