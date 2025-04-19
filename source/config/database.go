package config

import (
	"fmt"
	"os"
	"sarana-dafa-ai-service/model"
	"sarana-dafa-ai-service/storage/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv(env.DB_HOST),
		os.Getenv(env.DB_USER),
		os.Getenv(env.DB_PASSWORD),
		os.Getenv(env.DB_NAME),
		os.Getenv(env.DB_PORT),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate database model to database
	// Please use CLI command to migrate database model to database
	// Usage: go run server.go -migrate
	// MigrateTable(db)

	return db
}
func MigrateTable(db *gorm.DB) {
	enableSeed := false
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_manpower")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_patient_examination_checklist")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_patient_product_selection")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_product_examination_checklist")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_held_date")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_detail_product")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_patient_analysis")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment_patient")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_admin_company_client")

	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_appointment")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_company_client")
	// db.Exec("DROP TABLE IF EXISTS b2b_bumame_product")

	if err := db.AutoMigrate(&model.B2BProduct{}); err == nil {
		if enableSeed {
			var count int64
			db.Model(&model.B2BProduct{}).Count(&count)
			if count == 0 {
				for _, product := range DataSeedB2BProduct {
					db.Create(&product)
				}
			}
		}
	}

	var adminIds []uint64
	if err := db.AutoMigrate(&model.Admin{}); err == nil {
		if enableSeed {
			for _, admin := range DataSeedAdmin {
				db.Where("email = ?", admin.Email).First(&admin)
				if admin.Id == 0 {
					db.Create(&admin)
					if admin.Role == "b2b-sales" {
						adminIds = append(adminIds, admin.Id)
					}
				} else {
					adminIds = append(adminIds, admin.Id)
				}
			}
		}
	}
	db.Exec(`DO $$ 
DECLARE 
    _table text;
    _sequence text;
BEGIN
    FOR _table IN 
        SELECT table_name 
        FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_type = 'BASE TABLE'
        AND table_name != 'fiber_storage'
    LOOP
        _sequence := (
            SELECT pg_get_serial_sequence('public.' || _table, 'id')
        );
        
        IF _sequence IS NOT NULL THEN
            EXECUTE format(
                'SELECT setval(%L, COALESCE((SELECT MAX(id) FROM %I), 1))', 
                _sequence, _table
            );
        END IF;
    END LOOP;
END $$;`)
}
