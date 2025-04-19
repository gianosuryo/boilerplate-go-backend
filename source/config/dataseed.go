package config

import (
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model"
	"time"
)

var DataSeedPrescreeningTestJson = `{"riwayat_penyakit_sendiri":[["a. Riwayat Hepatitis","Tidak Ada"],["b. Riwayat Pengobatan TBC","Tidak Ada"],["c. Hipertensi / Darah Tinggi","Tidak Ada"],["d. Kencing Manis / Diabetes","Tidak Ada"],["e. Batuk Menahun","Tidak Ada"],["f. Riwayat Operasi","Tidak Ada"],["g. Riwayat Rawat Inap","Tidak Ada"],["h. Alergi Obat / Makanan / lainnya","Tidak Ada"],["i. Lain-lain","-"]],"riwayat_penyakit_keluarga":[["a. Riwayat Penyakit Jantung","Tidak Ada"],["b. Riwayat Hipertensi / Darah Tinggi","Ada Ayah"],["c. Riwayat Kencing Manis / Diabetes","Ada Ayah"],["d. Riwayat Stroke","Tidak Ada"],["e. Riwayat Penyakit Paru / Asma / TBC","Tidak Ada"],["f. Riwayat Kanker / Tumor","Tidak Ada"],["g. Riwayat Penyakit Gangguan Jiwa","Tidak Ada"],["h. Riwayat Penyakit Ginjal","Tidak Ada"],["i. Riwayat Penyakit saluran Pencernaan","Tidak Ada"],["j. Riwayat Penyakit Lainnya","-"]],"kebiasaan":[["a. Minum Alkohol","Tidak"],["b. Merokok","Ya, 6 batang / hari"],["c. Olahraga","Ya / Yes"],["d. Jenis Olahraga","Jogging, Sepak bola / Football, Bulutangkis / Badminton"],["e. Olahraga Berapa kali per minggu","Kurang dari 1 kali / minggu"]]}`

var DataSeedPhysicalExaminationJson = `[["Kulit","Normal"],["Kesadaran Umum","Normal"],["Kesadaran Mental","Normal"],["Mata",""],["Buta Warna","Normal"],["Kelainan Mata","Normal"],["THT",""],["Telinga","Normal"],["Hidung","Normal"],["Tenggorokan","Normal"],["Sinus","Normal"],["Tonsil","Normal"],["Gigi & Mulut",""],["Lidah","Normal"],["Gusi","Normal"],["Gigi","Normal"],["Kepala & Leher",""],["Kelenjar Limfe","Normal"],["Kelenjar Tiroid/Gondok","Normal"],["Dada",""],["Jantung","Normal"],["Paru","Normal"],["Tulang Belakang","Normal"],["Abdomen",""],["Perabaan","Normal"],["Hati","Normal"],["Ginjal","Normal"],["Extrimitas","Normal"],["Neurologis","Normal"],["Musuculoskeletal","Normal"]]`

var DataSeedVitalSignExaminationJson = `[["Tensi (mmHg)","103/70 mmhg"],["Nadi (X/menit)","60 x/m"],["RR (X/menit)","19 x/m"],["Suhu","36,3 c"],["SpO2 (%)","99%"],["Berat Badan (kg)","58 kg"],["Tinggi Badan (cm)","178 cm"],["Lingkar Perut (cm)","67 cm"],["Visus (Ka)","20/20"],["Visus (Ki)","20/20"],["BMI","18,3, Underweight"]]`

var DataSeedLabExaminationJson = `{"header":{"no_rm":"00120250100506","no_barcode":"2502200026","nama":"Tn. Ginanjar Eko","dokter_penunjuk":"-","jenis_kelamin":"Laki - laki","lokasi_pengambilan":"PT. OCS Global Service","tgl_lahir":"22-06-2002 / 22 Tahun 7 Bulan","tanggal_jam_pengambilan":"20-02-2025 13:00:06","telepon":"089901768299","tanggal_hasil_keluar":"21-02-2025 13:35:01","alamat":"Jakarta Selatan"},"sections":[{"name":"HEMATOLOGI","subsections":[{"name":"Hematologi Lengkap","tests":[{"name":"Hemoglobin (HGB)","hasil":"15.6","satuan":"g/dL","nilai_rujukan":"13.2 - 17.3","keterangan":""},{"name":"Hematokrit (HCT)","hasil":"45.2","satuan":"%","nilai_rujukan":"40 - 52","keterangan":""},{"name":"Eritrosit (RBC)","hasil":"5.4","satuan":"10^6/uL","nilai_rujukan":"4.4 - 5.9","keterangan":""},{"name":"Index Eritrosit :","hasil":"","satuan":"","nilai_rujukan":"","keterangan":""},{"name":"- MCV","hasil":"83.9*","satuan":"fl","nilai_rujukan":"84 - 96","keterangan":""},{"name":"- MCH","hasil":"29","satuan":"pg","nilai_rujukan":"27 - 30","keterangan":""},{"name":"- MCHC","hasil":"34.5","satuan":"g/dl","nilai_rujukan":"30 - 35","keterangan":""},{"name":"RDW","hasil":"12.2","satuan":"%","nilai_rujukan":"11.5 - 14.1","keterangan":""},{"name":"Trombosit","hasil":"294,000","satuan":"uL","nilai_rujukan":"150,000 - 440,000","keterangan":""},{"name":"Leukosit (WBC)","hasil":"5,670","satuan":"uL","nilai_rujukan":"5,000 - 10,000","keterangan":""},{"name":"HITUNG JENIS :","hasil":"","satuan":"","nilai_rujukan":"","keterangan":""},{"name":"- Basofil","hasil":"0","satuan":"%","nilai_rujukan":"0 - 1","keterangan":""},{"name":"- Eosinofil","hasil":"2","satuan":"%","nilai_rujukan":"2 - 4","keterangan":""},{"name":"- Neutrofil","hasil":"54","satuan":"%","nilai_rujukan":"50 - 70","keterangan":""},{"name":"- Limfosit","hasil":"38","satuan":"%","nilai_rujukan":"25 - 40","keterangan":""},{"name":"- Monosit","hasil":"6","satuan":"%","nilai_rujukan":"2 - 8","keterangan":""},{"name":"Laju Endap Darah","hasil":"3","satuan":"mm/jam","nilai_rujukan":"0 - 10","keterangan":""}]}]},{"name":"KIMIA","subsections":[{"name":"Fungsi Hati","tests":[{"name":"SGOT / AST","hasil":"16","satuan":"U/L","nilai_rujukan":"0 - 35","keterangan":""},{"name":"SGPT / ALT","hasil":"14","satuan":"U/L","nilai_rujukan":"0 - 45","keterangan":""}]},{"name":"DIABETES","tests":[{"name":"Glukosa Puasa","hasil":"127*","satuan":"mg/dL","nilai_rujukan":"70 - 110","keterangan":""}]},{"name":"PROFIL LIPID","tests":[{"name":"Kolesterol Total","hasil":"172*","satuan":"mg/dL","nilai_rujukan":"<200","keterangan":"Normal : < 200\nBatas Tinggi : 200 - 239\nTinggi : >= 240"},{"name":"Trigliserida","hasil":"91","satuan":"mg/dL","nilai_rujukan":"","keterangan":"Normal : < 150\nTinggi : 200 - 499\nBatas Tinggi : 150 - 199\nSangat Tinggi : >= 500"},{"name":"Kolesterol HDL","hasil":"49","satuan":"mg/dL","nilai_rujukan":"","keterangan":"Rendah : < 40\nTinggi : >= 60"},{"name":"Kolesterol LDL Direk","hasil":"121*","satuan":"mg/dL","nilai_rujukan":"<100","keterangan":"Normal : < 100\nBatas Tinggi : 100 - 129\nTinggi : >= 130"}]}]},{"name":"IMUNOLOGI","subsections":[{"name":"HEPATITIS","tests":[{"name":"HBsAg Kualitatif","hasil":"Non Reaktif","satuan":"COI","nilai_rujukan":"Non Reaktif","keterangan":"Non Reaktif : < 0.90\nBorderline : 0.90 - 0.99\nReaktif : >= 1.00"}]}]},{"name":"URINALISA","subsections":[{"name":"Urine Lengkap","tests":[]},{"name":"MAKROSKOPIS","tests":[{"name":"- Warna","hasil":"Kuning","satuan":"","nilai_rujukan":"Kuning","keterangan":""},{"name":"- Kejernihan","hasil":"Jernih","satuan":"","nilai_rujukan":"Jernih","keterangan":""}]},{"name":"KIMIA","tests":[{"name":"- Berat Jenis","hasil":"1.030","satuan":"","nilai_rujukan":"1.003 - 1.035","keterangan":""},{"name":"- Leukosit (esterase)","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Eritrosit (haem)","hasil":"Negatif","satuan":"uL","nilai_rujukan":"Negatif","keterangan":""},{"name":"- pH","hasil":"5.0","satuan":"","nilai_rujukan":"4.5 - 8.0","keterangan":""},{"name":"- Nitrit","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Protein (Albumin)","hasil":"Negatif","satuan":"mg/dL","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Glukosa","hasil":"Negatif","satuan":"mg/dL","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Keton","hasil":"Negatif","satuan":"mg/dL","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Urobilinogen","hasil":"Normal","satuan":"mg/dL","nilai_rujukan":"Normal","keterangan":""},{"name":"- Bilirubin","hasil":"Negatif","satuan":"mg/dL","nilai_rujukan":"Negatif","keterangan":""}]},{"name":"MIKROSKOPIS","tests":[{"name":"- Eritrosit","hasil":"0 - 1","satuan":"LPB","nilai_rujukan":"0 - 2","keterangan":""},{"name":"- Leukosit","hasil":"0 - 1","satuan":"LPB","nilai_rujukan":"0 - 5","keterangan":""},{"name":"- Sel Epitel","hasil":"0 - 1","satuan":"LPK","nilai_rujukan":"5 - 10","keterangan":""},{"name":"- Silinder (cast)","hasil":"Negatif","satuan":"LPK","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Kristal","hasil":"Negatif","satuan":"/LPB","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Bakteri","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"- Lain-lain","hasil":"-","satuan":"","nilai_rujukan":"Negatif","keterangan":""}]}]},{"name":"NARKOBA","subsections":[{"name":"","tests":[{"name":"Amphetamin","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"Morphine","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"THC","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"Cocaine","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""},{"name":"Benzodiapin","hasil":"Negatif","satuan":"","nilai_rujukan":"Negatif","keterangan":""}]}]}]}`

var DataSeedElectromedicalExaminationJson = `{"rontgen":{"url":"https://github.com/daffaputra09/assets/blob/main/TN.%20ALFIN%20MUJIB_1.2.156.112677.1000.301.20250220114549.27.pdf?raw=true","title":"HASIL PEMERIKSAAN RADIOLOGI","subtitle":"THORAX FOTO","hasil":["Trakea di tengah.","Mediastinum tidak melebar.","Cor tidak membesar.","Pulmo:","- Hili dalam batas normal.","- Corakan bronkovaskuler normal.","- Tidak tampak perberaakan / perselubungan.","Skeletal dan soft tissue yang tervisualisasi dalam batas normal."],"kesimpulan":["KESAN :","- Tidak tampak bronkopneumonia / pneumonia/ TB","- Tidak tampak kardiomegali."],"dokter":{"name":"dr. Pratama Adityabintoro, Sp.Rad","title":"Dokter Pemeriksa"}}}`

var DataSeedExaminationConclusionJson = `[["Tanda Vital","Underweight (18,3)"],["Pemeriksaan Fisik","Dalam Batas Normal"],["Hasil Darah","Penurunan MCV (83.9), peningkatan glukosa puasa (127), peningkatan LDL (121)"],["Urin","Dalam Batas Normal"],["TDM/NAPZA","Negatif"],["Rontgen Thorax","Dalam Batas Normal"]]`

var DataSeedExaminationAdvice = "Olahraga ringan rutin 3x dalam seminggu dalam durasi 30-45 menit, menjaga asupan makan, cukup asupan air putih, makan bergizi dan nutrisi seimbang, beristirahat yang cukup, mengelola stres dengan baik, konsultasi dengan dokter spesialis gizi terkait berat badan kurang, konsultasi dengan dokter spesialis penyakit dalam terkait hasil laboratorium. Lakukan MCU 1 tahun mendatang."

var defaultPassword, _ = helper.PasswordHash("asdf1234")

// create admin for role b2b-sales and b2b-ops
var DataSeedAdmin = []model.Admin{
	{
		Name:        "B2B Sales",
		Email:       "b2b-sales@admin.com",
		UserName:    "b2b-sales@admin.com",
		Password:    defaultPassword,
		Role:        "b2b-sales",
		Position:    "manager",
		Status:      1,
		IsConfirmed: 1,
		IsDeleted:   0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		Name:        "B2B Ops",
		Email:       "b2b-ops@admin.com",
		UserName:    "b2b-ops@admin.com",
		Password:    defaultPassword,
		Role:        "b2b-ops",
		Position:    "manager",
		Status:      1,
		IsConfirmed: 1,
		IsDeleted:   0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
var DataSeedB2BProduct = []model.B2BProduct{
	{
		Id:    1,
		Name:  "MCU Basic",
		Price: 100000,
	},
	{
		Id:    2,
		Name:  "MCU Plus",
		Price: 200000,
	},
	{
		Id:    3,
		Name:  "MCU Premium",
		Price: 300000,
	},
}
