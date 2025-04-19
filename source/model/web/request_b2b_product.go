package web

type RequestB2BProduct struct {
	Id                   uint64   `json:"id" example:"1" default:"0"`
	Name                 string   `json:"name" validate:"required" example:"Product Name"`
	Price                uint64   `json:"price" validate:"required" example:"10000000"`
	ExaminationChecklist []string `json:"examination_checklist" example:"injection,phlebotomy"`
}

type RequestB2BProductBulkUpdateWithUploadCSV struct {
	File string `json:"file" validate:"required" example:"file.csv"`
}

type RequestB2BProductGetExaminationChecklist struct {
	ProductIds []uint64 `json:"product_ids" validate:"required" example:"1,2,3"`
}
