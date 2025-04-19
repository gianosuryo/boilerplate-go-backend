package controller

import (
	"fmt"
	"net/http"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/**
* Interface design
**/
type BumameB2BProductController interface {
	// Public
	FindAll(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error

	BulkUpdate(*fiber.Ctx) error
	GenerateSlugs(*fiber.Ctx) error
}

/**
* Object implementation creation
* And inject depedencies
**/
type BumameB2BProductControllerImpl struct {
	BumameB2BProductService service.BumameB2BProductService
	Validate                *validator.Validate
}

func NewBumameB2BProductController(buas service.BumameB2BProductService, valid *validator.Validate) BumameB2BProductController {
	return &BumameB2BProductControllerImpl{
		BumameB2BProductService: buas,
		Validate:                valid,
	}
}

// @Summary Get all b2b products
// @Description Get all b2b products
// @Tags b2b-product
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Param queryList query web.QueryAdminB2B false "Query List"
// @Accept json
// @Produce json
// @Success 200 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product [get]
func (cont *BumameB2BProductControllerImpl) FindAll(c *fiber.Ctx) error {
	reqPagination := new(web.RequestPaginationNumber)
	if err := c.QueryParser(reqPagination); err != nil {
		return err
	}

	data, pagination := cont.BumameB2BProductService.FindAllPagination(c.Context(), *reqPagination)
	return helper.BuildSuccessResponsePagination(c, http.StatusOK, data, pagination)
}

// @Summary Get b2b product by ID
// @Description Get b2b product by ID
// @Tags b2b-product
// @Param b2b_product_id path string true "B2B Product ID"
// @Accept json
// @Produce json
// @Success 200 {object} web.ResponseStandard
// @Failure 404 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product/{b2b_product_id} [get]
func (cont *BumameB2BProductControllerImpl) FindById(c *fiber.Ctx) error {
	b2bProductId := c.Params("b2b_product_id")

	data, err := cont.BumameB2BProductService.FindById(c.Context(), b2bProductId)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusNotFound, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary Create b2b product
// @Description Create b2b product
// @Tags b2b-product
// @Accept json
// @Produce json
// @Param request body web.RequestB2BProduct true "Request"
// @Success 201 {object} web.ResponseStandard
// @Failure 400 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product [post]
func (cont *BumameB2BProductControllerImpl) Create(c *fiber.Ctx) error {
	request := new(web.RequestB2BProduct)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	data, err := cont.BumameB2BProductService.Create(c.Context(), *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusCreated, data)
}

// @Summary Update b2b product
// @Description Update b2b product
// @Tags b2b-product
// @Accept json
// @Produce json
// @Param b2b_product_id path string true "B2B Product ID"
// @Param request body web.RequestB2BProduct true "Request"
// @Success 200 {object} web.ResponseStandard
// @Failure 400 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product/{b2b_product_id} [put]
func (cont *BumameB2BProductControllerImpl) Update(c *fiber.Ctx) error {
	b2bProductId := c.Params("b2b_product_id")

	request := new(web.RequestB2BProduct)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	data, err := cont.BumameB2BProductService.Update(c.Context(), b2bProductId, *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary Delete b2b product
// @Description Delete b2b product
// @Tags b2b-product
// @Accept json
// @Produce json
// @Param b2b_product_id path string true "B2B Product ID"
// @Success 200 {object} web.ResponseStandard
// @Failure 400 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product/{b2b_product_id} [delete]
func (cont *BumameB2BProductControllerImpl) Delete(c *fiber.Ctx) error {
	b2bProductId := c.Params("b2b_product_id")

	data, err := cont.BumameB2BProductService.Delete(c.Context(), b2bProductId)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary Bulk update b2b products
// @Description Bulk update b2b products
// @Tags b2b-product
// @Accept multipart/form-data
// @Produce json
// @Param csv_file formData file true "CSV File Structure:\n Header: id,name,price,vital-signs,visus,physical-test,injection,phlebotomy,specimen-mandiri,pap-smear,rontgen,ekg,audiometri,spirometri,treadmill,usg-abdomen,usg-mammae;\n Example Value: 1,MCU Basic,10000,1,1,1,1,1,1,1,1,1,1,1,1,1,1 \n Note: if you want to create new product, you can leave the id as empty, else if you want to update the product, you can fill the id with the product id"
// @Success 200 {object} web.ResponseStandard
// @Failure 400 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product/bulk-update [put]
func (cont *BumameB2BProductControllerImpl) BulkUpdate(c *fiber.Ctx) error {
	// Get file from request
	records, err := helper.GetDataFromCSV(c, "csv_file")
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusBadRequest, []web.ErrorResponse{{
			Title:    "Bad Request",
			Messages: []string{err.Error()},
		}})
	}

	fmt.Println(records)
	request := []web.RequestB2BProduct{}

	// header examination checklist : vital-signs,visus,physical-test,injection,phlebotomy,specimen-mandiri,pap-smear,rontgen,ekg,audiometri,spirometri,treadmill,usg-abdomen,usg-mammae
	headerExaminationChecklist := []string{"vital-signs", "visus", "physical-test", "injection", "phlebotomy", "specimen-mandiri", "pap-smear", "rontgen", "ekg", "audiometri", "spirometri", "treadmill", "usg-abdomen", "usg-mammae"}

	for _, record := range records {
		examinationChecklist := []string{}
		for itemIter, itemRecord := range record {
			if itemIter >= 3 && itemRecord == "1" {
				fmt.Println(itemIter-3, itemRecord)
				examinationChecklist = append(examinationChecklist, headerExaminationChecklist[itemIter-3])
			}
		}

		idUint, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			return helper.BuildErrorResponse(c, http.StatusBadRequest, []web.ErrorResponse{{
				Title:    "Bad Request",
				Messages: []string{err.Error() + ". Must be a number"},
			}})
		}

		priceUint, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			return helper.BuildErrorResponse(c, http.StatusBadRequest, []web.ErrorResponse{{
				Title:    "Bad Request",
				Messages: []string{err.Error() + ". Must be a number"},
			}})
		}

		request = append(request, web.RequestB2BProduct{
			Id:                   idUint,
			Name:                 record[1],
			Price:                priceUint,
			ExaminationChecklist: examinationChecklist,
		})
	}

	fmt.Println(request)

	responseData, err := cont.BumameB2BProductService.BulkUpdate(c.Context(), request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, responseData)
}

// @Summary Generate slugs for products
// @Description Generate slugs for products where slug is null
// @Tags b2b-product
// @Accept json
// @Produce json
// @Success 200 {object} web.ResponseStandard
// @Failure 500 {object} web.ResponseStandard
// @Router /b2b-product/generate-slugs [post]
func (cont *BumameB2BProductControllerImpl) GenerateSlugs(c *fiber.Ctx) error {
	data, err := cont.BumameB2BProductService.GenerateSlugs(c.Context())
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}
