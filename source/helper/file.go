package helper

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/storage/env"
	"sarana-dafa-ai-service/storage/filekind"
	"sarana-dafa-ai-service/storage/message"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ImageFileExtenstion() []string {
	return []string{"image/jpeg", "image/png"}
}
func ValidateImageUpload(c *fiber.Ctx, fieldName string, maxWidth int, maxHeight int) (web.ErrorResponse, *multipart.FileHeader) {
	// Check uploaded file
	imageFile, err := c.FormFile(fieldName)
	if err != nil {
		return web.ErrorResponse{}, nil
	}

	// Check file content
	contentType := imageFile.Header.Get("Content-Type")
	if !slices.Contains(ImageFileExtenstion(), contentType) {
		return CreateErrorResponse(fieldName, message.InvalidValue(contentType)), nil
	}

	/**
	* Image dimention validation
	**/
	// Check image dimention
	openFile, err := imageFile.Open()
	if err != nil {
		return CreateErrorResponse(fieldName, err.Error()), nil
	}

	m, _, err := image.Decode(openFile)
	if err != nil {
		return CreateErrorResponse(fieldName, err.Error()), nil
	}

	// Get dimention
	bounds := m.Bounds()
	if bounds.Dx() > maxWidth || bounds.Dy() > maxHeight {
		return CreateErrorResponse(fieldName, message.MaximumImageDimention(maxHeight, maxWidth)), nil
	}

	return web.ErrorResponse{}, imageFile
}
func UploadFile(c *fiber.Ctx, folderName string, fieldName string, prefixFileName string) (string, error) {
	// Check uploaded file
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	// Doing upload file
	newFileName := fmt.Sprintf("%s_%s%s", prefixFileName, time.Now().Format("20060102150405"), filepath.Ext(file.Filename))
	destination := fmt.Sprintf("%s/%s/%s", os.Getenv(env.UPLOAD_DIRECTORY), folderName, newFileName)
	fmt.Println(destination)
	if err := c.SaveFile(file, destination); err != nil {
		PanicIfError(err)
		return "", err
	}
	return newFileName, err
}
func UploadFileGenerateToBucket(c *fiber.Ctx, folderName string, fieldName string, isRequired bool) ([]string, error) {
	// Check uploaded file
	form, err := c.MultipartForm()
	if err != nil {
		return []string{}, err
	}
	files := form.File[fieldName]

	if isRequired && len(files) == 0 {
		return []string{}, errors.New("pdf_file;Pdf file is required")
	}

	newFileNameArr := []string{}
	for iter, file := range files {
		if isRequired && (file.Filename == "undefined" || filepath.Ext(file.Filename) != ".pdf") {
			return []string{}, errors.New("pdf_file;Pdf file is not valid")
		}

		filename := strings.Replace(file.Filename, filepath.Ext(file.Filename), "", -1)
		// Doing upload file
		newFileName := fmt.Sprintf("%s/%s_%s%s", folderName, filename, strconv.Itoa(iter), filepath.Ext(file.Filename))
		blobFile, err := file.Open()
		if err != nil {
			PanicIfError(err)
		}

		if err := uploader.UploadFileToBucket(blobFile, newFileName); err != nil {
			PanicIfError(err)
			return []string{}, err
		}

		newFileNameArr = append(newFileNameArr, newFileName)
	}

	return newFileNameArr, err
}
func DeleteFileGenerateFromBucket(fileNames []string) {
	for _, filename := range fileNames {
		if err := uploader.DeleteFileFromBucket(filename); err != nil {
			PanicIfError(err)
		}
	}
}
func DeleteFile(fileName string, fieldName string) error {
	if fileName != "" {
		deleteFileDestionation := fmt.Sprintf("%s/%s/%s", os.Getenv(env.UPLOAD_DIRECTORY), fieldName, fileName)
		if _, err := os.Stat(deleteFileDestionation); err != nil {
			return err
		}
		if err := os.Remove(deleteFileDestionation); err != nil {
			return err
		}
	}
	return nil

}
func GetFullFileUrl(fileName string, kindType string, kindCat string) (fullFileUrl string) {
	return getFullFileURL(fileName, kindType, kindCat, false)
}
func GetFullFileUrlLegacy(fileName string, kindType string, kindCat string) (fullFileUrl string) {
	return getFullFileURL(fileName, kindType, kindCat, true)
}
func getFullFileURL(fileName string, kindType string, kindCat string, isLegacy bool) (fullFileUrl string) {

	// Divide by kind
	switch kindType {
	case filekind.TYPE_IMAGE:
		fullFileUrl = imageDir(kindCat, isLegacy) + fileName
	}

	// Add http based on setting
	if isLegacy {
		fullFileUrl = os.Getenv(env.BASE_URL_FILE_LEGACY) + fullFileUrl
	} else {
		fullFileUrl = os.Getenv(env.BASE_URL_FILE) + fullFileUrl
	}

	return fullFileUrl
}
func imageDir(kindCat string, isLegacy bool) (dir string) {

	// If not legacy then add files dir
	if !isLegacy {
		dir = ""
	}

	// Divide by category
	switch kindCat {
	case filekind.CAT_LOGO_TEMPLATE_STYLE:
		dir = dir + "images/logo_template_style/"
	case filekind.CAT_TRAINING:
		dir = dir + "images/training/"
	case filekind.CAT_PROFILE:
		dir = dir + "images/profile/"

	default:
		dir = dir + kindCat + "/"
	}

	return dir
}
func UploadFileToBucketResumeFile(c *fiber.Ctx, folderName string, fieldName string, isRequired bool) ([]string, error) {
	// Check uploaded file
	form, err := c.MultipartForm()
	if err != nil {
		return []string{}, err
	}
	files := form.File[fieldName]

	if isRequired && len(files) == 0 {
		return []string{}, errors.New(fmt.Sprintf("%s;File is required", fieldName))
	}

	newFileNameArr := []string{}
	for iter, file := range files {
		if isRequired && (file.Filename == "undefined") {
			return []string{}, errors.New(fmt.Sprintf("%s;File is required, %s", fieldName, file.Filename))
		}

		filename := strings.Replace(file.Filename, filepath.Ext(file.Filename), "", -1)
		filename = strings.Replace(filename, " ", "-", -1)
		// Doing upload file
		newFileName := fmt.Sprintf("%s/%s_%s%s", folderName, filename, strconv.Itoa(iter), filepath.Ext(file.Filename))
		blobFile, err := file.Open()
		if err != nil {
			PanicIfError(err)
		}

		if err := uploaderResumeFile.UploadFileToBucketResumeFile(blobFile, newFileName); err != nil {
			PanicIfError(err)
			return []string{}, err
		}

		newFileNameArr = append(newFileNameArr, newFileName)
	}

	return newFileNameArr, err
}

func UploadFileToBucketVideoInterviewFile(c *fiber.Ctx, folderName string, fieldName string, isRequired bool) ([]string, error) {
	// Check uploaded file
	form, err := c.MultipartForm()
	if err != nil {
		return []string{}, err
	}

	files := form.File[fieldName]

	if isRequired && len(files) == 0 {
		return []string{}, errors.New(fmt.Sprintf("%s;File is required", fieldName))
	}

	newFileNameArr := []string{}
	for iter, file := range files {
		if isRequired && (file.Filename == "undefined") {
			return []string{}, errors.New(fmt.Sprintf("%s;File is required, %s", fieldName, file.Filename))
		}

		filename := strings.Replace(file.Filename, filepath.Ext(file.Filename), "", -1)
		filename = strings.Replace(filename, " ", "-", -1)
		// Doing upload file
		newFileName := fmt.Sprintf("%s/%s_%s%s", folderName, filename, strconv.Itoa(iter), filepath.Ext(file.Filename))
		blobFile, err := file.Open()
		if err != nil {
			PanicIfError(err)
		}

		if err := uploaderResumeFile.UploadFileToBucketResumeFile(blobFile, newFileName); err != nil {
			PanicIfError(err)
			return []string{}, err
		}

		newFileNameArr = append(newFileNameArr, newFileName)
	}

	return newFileNameArr, err
}

func ReadCSVFile(filename *multipart.FileHeader) ([]byte, error) {
	f, err := filename.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

func ProcessCSV(reader *csv.Reader) ([][]string, error) {
	records := [][]string{}
	iter := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			return nil, err
		}

		// skip header
		if iter == 0 {
			iter++
			continue
		}
		records = append(records, record)
	}
	return records, nil
}

func GetDataFromCSV(c *fiber.Ctx, fieldRequestName string) ([][]string, error) {
	file, err := c.FormFile(fieldRequestName)
	if err != nil {
		return nil, err
	}

	data, err := ReadCSVFile(file)
	if err != nil {
		return nil, err
	}

	reader, err := ParseCSV(data)
	if err != nil {
		return nil, err
	}

	records, err := ProcessCSV(reader)
	if err != nil {
		return nil, err
	}

	return records, nil
}
