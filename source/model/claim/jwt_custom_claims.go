package claim

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserId         uint64 `json:"user_id"`
	UserUUID       string `json:"user_uuid"`
	UserName       string `json:"user_name"`
	UserUserName   string `json:"user_user_name"`
	UserEmail      string `json:"user_email"`
	UserRole       string `json:"user_role"`
	SubCompanyId   string `json:"sub_company_id"`
	SubCompanyName string `json:"sub_company_name"`
	// Fixed Data
	CompanySchema  string `json:"company_schema"`
	CompanyDomain  string `json:"company_domain"`
	CompanyCode    string `json:"company_code"`
	ServiceUrl     string `json:"service_url"`
	WebUrl         string `json:"web_url"`
	MainServiceUrl string `json:"main_service_url"`
	MainWebUrl     string `json:"main_web_url"`

	jwt.RegisteredClaims
}

// "user_id": 1,
// "user_name": "Admin Metrodata",
// "user_user_name": "Admin Metrodata",
// "user_email": "rozi@mejakita.com",
// "user_role": "Superuser",
// "sub_company_name": "ADD,Added by non-superuser (tester),Coding Weekend,Metrodata,MSHR",
// "sub_company_id": "1,2,3,8,9",
// "user_division_id": null,
// "user_division_name": null,
// "user_level_id": null,
// "user_level_name": null,
// "company_schema": "mejakerja_nya_metrodata",
// "company_domain": "metrodata.co.id;soltius.co.id;mii.co.id;ms.mii.co.id;codingweekend.edu.id;",
// "company_code": "metrodata",
// "service_url": "https://mejakerja.azurewebsites.net/",
// "web_url": "https://mejakerja.azurewebsites.net/",
// "main_service_url": "https://service.mejakerja.net/",
// "main_web_url": "https://www.mejakerja.net/",
// "rand": "7753028160.52254300 1678943709"
