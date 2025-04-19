package service

import (
	"context"
	"errors"
	"fmt"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model"
	"sarana-dafa-ai-service/model/web"
	"strconv"

	"gorm.io/gorm"
)

/**
* Interface
**/
type BumameB2BProductService interface {
	FindAllPagination(ctx context.Context, reqPage web.RequestPaginationNumber) ([]map[string]interface{}, helper.PaginationNumber)
	FindById(ctx context.Context, productId string) (map[string]interface{}, error)
	Create(ctx context.Context, request web.RequestB2BProduct) (map[string]interface{}, error)
	Update(ctx context.Context, productId string, request web.RequestB2BProduct) (map[string]interface{}, error)
	Delete(ctx context.Context, productId string) (map[string]interface{}, error)

	BulkUpdate(ctx context.Context, request []web.RequestB2BProduct) (map[string]interface{}, error)
	GenerateSlugs(ctx context.Context) (map[string]interface{}, error)
}

/**
* Object implementation creation
**/
type BumameB2BProductServiceImpl struct {
	DB *gorm.DB
}

func NewBumameB2BProductService(db *gorm.DB) BumameB2BProductService {
	return &BumameB2BProductServiceImpl{
		DB: db,
	}
}

func (service BumameB2BProductServiceImpl) FindAllPagination(
	ctx context.Context, reqPage web.RequestPaginationNumber) (
	[]map[string]interface{}, helper.PaginationNumber) {
	listData := []model.B2BProduct{}
	tx := service.DB.Begin()
	defer tx.Rollback()

	pagination := helper.PaginationNumber{}
	pagination.Page = reqPage.Page
	pagination.PerPage = reqPage.Limit
	pagination.Sort = reqPage.Sort

	tx = tx.Where("is_deleted = 0")
	err := tx.Scopes(helper.Paginate(listData, &pagination, tx)).Find(&listData).Error
	helper.PanicIfError(err)

	tx.Commit()

	responseData := []map[string]interface{}{}
	for _, v := range listData {
		response := map[string]interface{}{
			"id":         v.Id,
			"name":       v.Name,
			"price":      v.Price,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		}
		responseData = append(responseData, response)
	}
	return responseData, pagination
}

func (service BumameB2BProductServiceImpl) FindById(ctx context.Context, productId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	product := model.B2BProduct{}
	err := tx.Where("id = ? AND is_deleted = 0", productId).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":         product.Id,
		"name":       product.Name,
		"price":      product.Price,
		"created_at": product.CreatedAt,
		"updated_at": product.UpdatedAt,
	}, nil
}

func (service BumameB2BProductServiceImpl) Create(ctx context.Context, request web.RequestB2BProduct) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	product := model.B2BProduct{
		Name:  request.Name,
		Price: request.Price,
	}

	err := tx.Create(&product).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":         product.Id,
		"name":       product.Name,
		"price":      product.Price,
		"created_at": product.CreatedAt,
	}, nil
}

func (service BumameB2BProductServiceImpl) Update(ctx context.Context, productId string, request web.RequestB2BProduct) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	product := model.B2BProduct{}
	err := tx.Where("id = ? AND is_deleted = 0", productId).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		helper.PanicIfError(err)
	}

	product.Name = request.Name
	product.Price = request.Price
	err = tx.Save(&product).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":         product.Id,
		"name":       product.Name,
		"price":      product.Price,
		"created_at": product.CreatedAt,
		"updated_at": product.UpdatedAt,
	}, nil
}

func (service BumameB2BProductServiceImpl) BulkUpdate(ctx context.Context, request []web.RequestB2BProduct) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	for _, v := range request {
		if v.Id == 0 {
			// Create new product
			newProduct := model.B2BProduct{
				Name:  v.Name,
				Price: v.Price,
			}
			err := tx.Create(&newProduct).Error
			if err != nil {
				return nil, err
			}

		} else {
			// Update existing product
			product := model.B2BProduct{}
			err := tx.Where("id = ? AND is_deleted = 0", v.Id).First(&product).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("product with id " + strconv.FormatUint(v.Id, 10) + " not found")
				}
				return nil, err
			}

			product.Name = v.Name
			product.Price = v.Price
			err = tx.Save(&product).Error
			if err != nil {
				return nil, err
			}
		}
	}

	// Commit transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "product updated successfully",
	}, nil
}

func (service BumameB2BProductServiceImpl) Delete(ctx context.Context, productId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	product := model.B2BProduct{}
	err := tx.Where("id = ? AND is_deleted = 0", productId).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		helper.PanicIfError(err)
	}

	product.IsDeleted = 1
	err = tx.Save(&product).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"message": "product deleted successfully",
	}, nil
}

func (service BumameB2BProductServiceImpl) GenerateSlugs(ctx context.Context) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	// Get all products where slug is null
	var products []model.B2BProduct
	err := tx.Where("slug IS NULL AND is_deleted = 0").Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Generate and update slugs for each product
	for _, product := range products {
		// Generate base slug from product name
		baseSlug := helper.GenerateSlug(product.Name)
		slug := baseSlug
		counter := 1

		// Check if slug already exists
		for {
			var existingProduct model.B2BProduct
			err := tx.Where("slug = ? AND id != ?", slug, product.Id).First(&existingProduct).Error

			// If no record found, we can use this slug
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}

			// If other error occurred
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}

			// If slug exists, append counter and try again
			slug = fmt.Sprintf("%s-%d", baseSlug, counter)
			counter++
		}

		// Update the product with the new slug
		err = tx.Model(&product).Update("slug", slug).Error
		if err != nil {
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "slugs generated successfully",
		"count":   len(products),
	}, nil
}
