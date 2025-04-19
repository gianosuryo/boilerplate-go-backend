package web

type RequestAppointmentPatientExaminationChecklist struct {
	Status          string     `json:"status" form:"status" validate:"required" example:"pending" enums:"pending,completed,skipped"`
	VitalSigns      [][]string `json:"vital_signs" form:"vital_signs" example:[['Tensi (mmHg)', '103/70 mmhg'], ['Nadi (X/menit)', '60 x/m']]`
	Injection       [][]string `json:"injection" form:"injection" example:['Form Prescreening 1', 'Jawaban Form Screening 1'], ['Form Prescreening 2', 'Jawaban Form Screening 2'], ["E-Signature URL", "URL"], ["Refuse Reason", "Jawaban"]` // examination information
	Visus           [][]string `json:"visus" form:"visus" example:[['Visus (Kiri)', '20/20'], ['Visus (Kanan)', '20/20']]`
	PhysicalTest    [][]string `json:"physical_test" form:"physical_test" example:[['Kulit', 'Normal'], ['Kesadaran Umum', 'Normal']]`
	Phlebotomy      [][]string `json:"phlebotomy" form:"phlebotomy" example:[['Refuse Reason', 'Jawaban']]`             // examination information
	SpecimenMandiri [][]string `json:"specimen_mandiri" form:"specimen_mandiri" example:[['Refuse Reason', 'Jawaban']]` // examination information
	PapSmear        [][]string `json:"pap_smear" form:"pap_smear" example:[['Refuse Reason', 'Jawaban']]`               // examination information
	// Rontgen just checklist aja
	// EKG just checklist aja
	Audiometri [][]string `json:"audiometri" form:"audiometri" example:[['Telinga Kanan', 'Air Conduction (AC) 250Hz'], ['Telinga Kiri', 'Air Conduction (AC) 250Hz']]`
	// Spirometri just checklist aja
	// Treadmill just checklist aja
	// USG Abdomen just checklist aja
	// USG Mammae just checklist aja
	AdminId uint64 `json:"admin_id" form:"admin_id" validate:"required" example:"1"`
}

type RequestAppointmentPatientExaminationChecklistInputted struct {
	AdminId uint64 `json:"admin_id" form:"admin_id" validate:"required" example:"1"`
}

type RequestAppointmentPatientExaminationChecklistAnalyzed struct {
	AdminId uint64 `json:"admin_id" form:"admin_id" validate:"required" example:"1"`
}
