package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type B2BProduct struct {
	// Id field
	Id uint64 `json:"id" gorm:"autoIncrement"`

	Slug  string `json:"slug" gorm:"unique; default:null"`
	Name  string `json:"name" gorm:"not null"`
	Price uint64 `json:"price" gorm:"not null"`

	// ProductExaminationChecklist []B2BProductExaminationChecklist `json:"product_examination_checklist" gorm:"foreignKey:ProductId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`

	IsDeleted uint `json:"is_deleted" gorm:"default:0"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (B2BProduct) TableName() string {
	return "b2b_bumame_product"
}

func (b2bProduct *B2BProduct) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate base slug from product name
	baseSlug := strings.ToLower(strings.ReplaceAll(b2bProduct.Name, " ", "-"))
	slug := baseSlug
	counter := 1

	// Check if slug already exists
	for {
		var existingProduct B2BProduct
		err := tx.Where("slug = ?", slug).First(&existingProduct).Error

		// If no record found, we can use this slug
		if errors.Is(err, gorm.ErrRecordNotFound) {
			break
		}

		// If other error occurred
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// If slug exists, append counter and try again
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	b2bProduct.Slug = slug
	return nil
}

func (b2bProduct *B2BProduct) BeforeUpdate(tx *gorm.DB) (err error) {
	baseSlug := strings.ToLower(strings.ReplaceAll(b2bProduct.Name, " ", "-"))
	slug := baseSlug
	counter := 1

	// Check if slug already exists
	for {
		var existingProduct B2BProduct
		err := tx.Where("slug = ? AND id != ?", slug, b2bProduct.Id).First(&existingProduct).Error

		// If no record found, we can use this slug
		if errors.Is(err, gorm.ErrRecordNotFound) {
			break
		}

		// If other error occurred
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// If slug exists, append counter and try again
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	b2bProduct.Slug = slug
	return nil
}
