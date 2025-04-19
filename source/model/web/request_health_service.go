package web

// Sementara price boleh kosong
type RequestBumameHealthService struct {
	Name                     string   `json:"name" form:"name" validate:"required"`
	Description              string   `json:"description" form:"description"`
	Price                    uint64   `json:"price" form:"price"`
	HourDuration             float32  `json:"hour_duration" form:"hour_duration"`
	Equipments               string   `json:"equipments" form:"equipments"`
	Consumables              string   `json:"consumables" form:"consumables"`
	HealthServiceCategoryIds []uint64 `json:"health_service_category_ids" form:"health_service_category_ids"`
	SkillIds                 []uint64 `json:"skill_ids" form:"skill_ids"`
	PrescreeningQuestionIds  []uint64 `json:"prescreening_question_ids" form:"prescreening_question_ids"`
	DetailPemeriksaan        string   `json:"detail_pemeriksaan" form:"detail_pemeriksaan"`
	ManfaatPemeriksaan       string   `json:"manfaat_pemeriksaan" form:"manfaat_pemeriksaan"`
	TentangPemeriksaan       string   `json:"tentang_pemeriksaan" form:"tentang_pemeriksaan"`
	IsNeedFasting            uint     `json:"is_need_fasting" form:"is_need_fasting"`
}

type RequestHealthServiceAvailableTime struct {
	Date string `json:"date" form:"date"`
}
