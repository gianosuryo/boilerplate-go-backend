package web

type RequestB2BPackage struct {
	Id                              uint64   `json:"id" example:"1" default:"0"`
	Name                            string   `json:"name" validate:"required" example:"Package Name"`
	IsUseCustomExaminationChecklist uint     `json:"is_use_custom_examination_checklist" example:"1"`
	CustomExaminationChecklist      []string `json:"custom_examination_checklist" example:"injection,phlebotomy"`
	ProductIds                      []uint64 `json:"product_ids" validate:"required" example:"1,2,3"`
}
