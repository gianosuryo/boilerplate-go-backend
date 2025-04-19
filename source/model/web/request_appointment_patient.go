package web

// apa yang kurang di identitas diri
type RequestAppointmentPatient struct {
	Name      string `json:"name" validate:"required" example:"Patient Name"`
	NIK       string `json:"nik" validate:"required" example:"1234567890123456"`
	Barcode   string `json:"barcode" validate:"required" example:"1234567890123456"`
	BirthDate string `json:"birth_date" validate:"required" example:"2021-01-01"`
	Gender    string `json:"gender" validate:"required" example:"Laki-laki"`
	Group     string `json:"group" example:"34_HLI_OS"`

	Phone string `json:"phone" validate:"required" example:"081234567890"`
	Email string `json:"email" validate:"required" example:"patient@example.com"`

	// B2BProductId []uint64 `json:"b2b_product_id" validate:"required" example:"1,2"`
	B2BPackageSlug []string `json:"b2b_package_slug" validate:"required" example:"mcu-basic-proo,mcu-basic-plus-plus"`
}

type RequestAppointmentPatientBulkInput struct {
	Name      string `json:"name" validate:"required" example:"Patient Name"`
	NIK       string `json:"nik" validate:"required" example:"1234567890123456"`
	BirthDate string `json:"birth_date" validate:"required" example:"2021-01-01"`
	Gender    string `json:"gender" validate:"required" example:"Laki-laki"`
	Phone     string `json:"phone" validate:"required" example:"081234567890"`
	Email     string `json:"email" validate:"required" example:"patient@example.com"`
	Barcode   string `json:"barcode" validate:"required" example:"1234567890123456"`

	B2BPackageSlug []string `json:"b2b_package_slug" validate:"required" example:"mcu-basic-proo,mcu-basic-plus-plus"`
}

type RequestAppointmentPatientBulkInputByProductSlug struct {
	Name      string `json:"name" validate:"required" example:"Patient Name"`
	NIK       string `json:"nik" validate:"required" example:"1234567890123456"`
	BirthDate string `json:"birth_date" validate:"required" example:"2021-01-01"`
	Gender    string `json:"gender" validate:"required" example:"Laki-laki"`
	Phone     string `json:"phone" validate:"required" example:"081234567890"`
	Email     string `json:"email" validate:"required" example:"patient@example.com"`
	Barcode   string `json:"barcode" validate:"required" example:"1234567890123456"`

	B2BProductSlug []string `json:"b2b_product_slug" validate:"required" example:"mcu-basic-proo,mcu-basic-plus-plus"`
}

type RequestAppointmentPatientAnalysisAnamnesa struct {
	RiwayatPenyakitSendiri  [][]string `json:"riwayat_penyakit_sendiri" validate:"required" example:[["a. Riwayat Hepatitis","Tidak Ada"]]`
	RiwayatPenyakitKeluarga [][]string `json:"riwayat_penyakit_keluarga" validate:"required" example:[["a. Riwayat Hepatitis","Tidak Ada"]]`
	Kebiasaan               [][]string `json:"kebiasaan" validate:"required" example:[['a. Riwayat Hepatitis','Tidak Ada']]`
}

type RequestAppointmentPatientCheckOut struct {
	PhotoProofUrl string `json:"photo_proof_url" example:"https://example.com/photo_proof.jpg"`
}

type RequestAppointmentPatientCheckIn struct {
	NotesUrine         string `json:"notes_urine" example:"Aman"`
	NotesThorax        string `json:"notes_thorax" example:"Aman"`
	StatusPuasa        string `json:"status_puasa" example:"Puasa"`
	UrineRefuseReason  string `json:"urine_refuse_reason" example:"Karena sudah kencing"`
	ThoraxRefuseReason string `json:"thorax_refuse_reason" example:"Karena sudah berjabat tangan"`
}

type RequestAppointmentPatientUpdateTubeNumber struct {
	TubeNumber string `json:"tube_number" example:"1234567890"`
}
