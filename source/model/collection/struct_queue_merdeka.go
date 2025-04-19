package collection

type ResponseQueueMerdekaFetchSignal struct {
	Type      string `json:"type"`
	CourseID  string `json:"course_id"`
	ChapterID string `json:"chapter_id"`
	Status    string `json:"status"`
}
