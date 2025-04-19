package service

import (
	"context"
	"errors"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model"
	"sarana-dafa-ai-service/model/web"

	"gorm.io/gorm"
)

/**
* Interface
**/
type BumameB2BAdminService interface {
	// Public
	FindAllPagination(ctx context.Context, reqPage web.RequestPaginationNumber, reqQuery web.QueryAdminB2B) ([]map[string]interface{}, helper.PaginationNumber)
	FindById(ctx context.Context, adminId string) (map[string]interface{}, error)
	Create(ctx context.Context, request web.RequestAdmin) (map[string]interface{}, error)
	Update(ctx context.Context, adminId string, request web.RequestAdmin) (map[string]interface{}, error)
	Delete(ctx context.Context, adminId string) (map[string]interface{}, error)
}

/**
* Object implementation creation
**/
type BumameB2BAdminServiceImpl struct {
	DB *gorm.DB
}

func NewBumameB2BAdminService(db *gorm.DB) BumameB2BAdminService {
	return &BumameB2BAdminServiceImpl{
		DB: db,
	}
}

func (service BumameB2BAdminServiceImpl) FindAllPagination(
	ctx context.Context, reqPage web.RequestPaginationNumber, reqQuery web.QueryAdminB2B) (
	[]map[string]interface{}, helper.PaginationNumber) {
	listData := []model.Admin{}
	tx := service.DB.Begin()
	defer tx.Rollback()

	pagination := helper.PaginationNumber{}
	pagination.Page = reqPage.Page
	pagination.PerPage = reqPage.Limit
	pagination.Sort = reqPage.Sort

	// get data where role contain b2b
	tx = tx.Where("is_deleted = 0")
	if reqQuery.Role != "" {
		tx = tx.Where("role = ?", reqQuery.Role)
	} else {
		tx = tx.Where("role LIKE ? OR role = ?", "%b2b%", "admin")
	}
	err := tx.Scopes(helper.Paginate(listData, &pagination, tx)).Find(&listData).Error

	helper.PanicIfError(err)

	// Convert data to response
	responseData := []map[string]interface{}{}
	for _, v := range listData {
		response := map[string]interface{}{
			"id":         v.Id,
			"name":       v.Name,
			"email":      v.Email,
			"phone":      v.Phone,
			"role":       v.Role,
			"position":   v.Position,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		}
		responseData = append(responseData, response)
	}
	return responseData, pagination
}

// Find company client by ID
func (service BumameB2BAdminServiceImpl) FindById(ctx context.Context, adminId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	admin := model.Admin{}
	err := tx.Where("id = ? AND is_deleted = 0", adminId).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":         admin.Id,
		"name":       admin.Name,
		"email":      admin.Email,
		"phone":      admin.Phone,
		"role":       admin.Role,
		"position":   admin.Position,
		"created_at": admin.CreatedAt,
		"updated_at": admin.UpdatedAt,
	}, nil
}

// Create new company client
func (service BumameB2BAdminServiceImpl) Create(ctx context.Context, request web.RequestAdmin) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	// check admin by email and username is already exist, if exist return error
	admin := model.Admin{}
	err := tx.Where("email = ? OR user_name = ?", request.Email, request.Email).Where("is_deleted = 0").First(&admin).Error
	helper.PanicIfError(err)
	if admin.Id != 0 {
		return nil, errors.New("admin email or username already exist")
	}

	admin.Name = request.Name
	admin.Email = request.Email
	admin.Password = request.Password
	admin.Role = request.Role
	admin.Position = request.Position
	admin.UserName = request.Email
	err = tx.Create(&admin).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":         admin.Id,
		"name":       admin.Name,
		"email":      admin.Email,
		"phone":      admin.Phone,
		"role":       admin.Role,
		"position":   admin.Position,
		"created_at": admin.CreatedAt,
	}, nil
}

// Update company client
func (service BumameB2BAdminServiceImpl) Update(ctx context.Context, adminId string, request web.RequestAdmin) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	admin := model.Admin{}
	err := tx.Where("id = ? AND is_deleted = 0", adminId).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		helper.PanicIfError(err)
	}

	// check admin by email and username is already exist, if exist return error
	existingAdmin := model.Admin{}
	service.DB.Where("id != ? AND is_deleted = 0", admin.Id).Where("email = ? OR user_name = ?", request.Email, request.Email).First(&existingAdmin)
	if existingAdmin.Id != 0 {
		return nil, errors.New("admin email or username already exist")
	}

	admin.Name = request.Name
	admin.Email = request.Email
	admin.UserName = request.Email
	admin.Password = request.Password
	admin.Role = request.Role
	admin.Position = request.Position

	err = tx.Save(&admin).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":         admin.Id,
		"name":       admin.Name,
		"email":      admin.Email,
		"phone":      admin.Phone,
		"role":       admin.Role,
		"position":   admin.Position,
		"created_at": admin.CreatedAt,
		"updated_at": admin.UpdatedAt,
	}, nil
}

// Delete company client
func (service BumameB2BAdminServiceImpl) Delete(ctx context.Context, adminId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	admin := model.Admin{}
	err := tx.Where("id = ? AND is_deleted = 0", adminId).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		helper.PanicIfError(err)
	}

	admin.IsDeleted = 1
	err = tx.Save(&admin).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"message": "admin deleted successfully",
	}, nil
}
