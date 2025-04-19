package web

type RequestAppointmentPatientAnalysis struct {
	VitalSigns      [][]string `json:"vital_signs" form:"vital_signs" example:[['Tensi (mmHg)', '103/70 mmhg'], ['Nadi (X/menit)', '60 x/m']]`
	Visus           [][]string `json:"visus" form:"visus" example:[['Visus (Kiri)', '20/20'], ['Visus (Kanan)', '20/20']]`
	PhysicalTest    [][]string `json:"physical_test" form:"physical_test" example:[['Kulit', 'Normal'], ['Kesadaran Umum', 	'Normal']]`
	Injection       [][]string `json:"injection" form:"injection" example:['Form Prescreening 1', 'Jawaban Form Screening 1'], ['Form Prescreening 2', 'Jawaban Form Screening 2'], ["E-Signature URL", "URL"], ["Refuse Reason", "Jawaban"]` // examination information
	Phlebotomy      [][]string `json:"phlebotomy" form:"phlebotomy" example:[['Refuse Reason', 'Jawaban']]`                                                                                                                                   // examination information
	SpecimenMandiri [][]string `json:"specimen_mandiri" form:"specimen_mandiri" example:[['Refuse Reason', 'Jawaban']]`                                                                                                                       // examination information
	PapSmear        [][]string `json:"pap_smear" form:"pap_smear" example:[['Refuse Reason', 'Jawaban']]`                                                                                                                                     // examination information
	Audiometri      [][]string `json:"audiometri" form:"audiometri" example:[['Telinga Kiri', 'Jawaban'], ['Telinga Kanan', 'Jawaban']]`
}

type RequestElectromedicalExamination struct {
	ElectromedicalExaminationJson string `json:"electromedical_examination_json" form:"electromedical_examination_json" example:"{rontgen: {status: 'normal', description: 'Normal'}, ekg: {status: 'normal', description: 'Normal'}}"`
	AdminId                       uint64 `json:"admin_id" form:"admin_id" example:"1"`
}

type RequestOverallExamination struct {
	Saran           string `json:"saran" form:"saran" example:"Sarankan untuk melakukan pemeriksaan lebih lanjut"`
	OverallAnalysis string `json:"overall_analysis" form:"overall_analysis" example:"Fit to work"`
	AdminId         uint64 `json:"admin_id" form:"admin_id" example:"1"`
}
