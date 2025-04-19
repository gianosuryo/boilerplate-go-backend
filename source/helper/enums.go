package helper

var ListExamination = []map[string]string{
	{
		"label": "Vital Signs",
		"value": "vital-signs",
	},
	{
		"label": "Injection",
		"value": "injection",
	},
	{
		"label": "Visus",
		"value": "visus",
	},
	{
		"label": "Physical Test",
		"value": "physical-test",
	},
	{
		"label": "Phlebotomy",
		"value": "phlebotomy",
	},
	{
		"label": "Specimen Mandiri",
		"value": "specimen-mandiri",
	},
	{
		"label": "Pap Smear",
		"value": "pap-smear",
	},
	{
		"label": "Rontgen",
		"value": "rontgen",
	},
	{
		"label": "EKG",
		"value": "ekg",
	},
	{
		"label": "Audiometri",
		"value": "audiometri",
	},
	{
		"label": "Spirometri",
		"value": "spirometri",
	},
	{
		"label": "Treadmill",
		"value": "treadmill",
	},
	{
		"label": "USG Abdomen",
		"value": "usg-abdomen",
	},
	{
		"label": "USG Mammae",
		"value": "usg-mammae",
	},
}

func GetDataExaminationChecklist(value string) map[string]string {
	for _, v := range ListExamination {
		if v["value"] == value {
			return v
		}
	}

	return nil
}

func IsAnyLabService(value string) bool {
	return value == "phlebotomy" || value == "specimen-mandiri" || value == "pap-smear"
}

func IsAnyElectromedicalService(value string) bool {
	return value == "rontgen" || value == "ekg" || value == "audiometri" || value == "spirometri" || value == "treadmill" || value == "usg-abdomen" || value == "usg-mammae"
}
