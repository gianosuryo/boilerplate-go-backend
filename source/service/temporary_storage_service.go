package service

import (
	"context"
	"fmt"
	"sarana-dafa-ai-service/helper"
	"time"

	"github.com/gofiber/storage"
)

/**
* Interface
**/
type TemporaryStorageService interface {
	// Public
	Create(ctx context.Context, storageName string, data string) (map[string]interface{}, error)
	FindOneById(ctx context.Context, storageName string, id string) string
	DeleteById(ctx context.Context, storageName string, id string) error
	getStorageId(storageName string, id string) string
}

/**
* Object implementation creation
**/
type TemporaryStorageServiceImpl struct {
	Storage storage.Storage
}

// Constructor function for TemporaryStorageServiceImpl
func NewTemporaryStorageService(storage storage.Storage) TemporaryStorageService {
	return &TemporaryStorageServiceImpl{
		Storage: storage,
	}
}

func (t *TemporaryStorageServiceImpl) Create(ctx context.Context, storageName string, data string) (map[string]interface{}, error) {
	uuid := helper.RandomString(5)

	err := t.Storage.Set(t.getStorageId(storageName, uuid), []byte(data), time.Hour*24*5)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":   uuid,
		"data": data,
	}, nil
}

func (t *TemporaryStorageServiceImpl) FindOneById(ctx context.Context, storageName string, id string) string {
	data, err := t.Storage.Get(t.getStorageId(storageName, id))
	if err != nil {
		return ""
	}
	return string(data)
}

func (t *TemporaryStorageServiceImpl) DeleteById(ctx context.Context, storageName string, id string) error {
	fullId := t.getStorageId(storageName, id)
	err := t.Storage.Delete(fullId)
	fmt.Println("Deleting cache data", fullId, err)
	if err != nil {
		helper.PanicIfError(err)
	}
	return nil
}
func (t *TemporaryStorageServiceImpl) getStorageId(storageName string, id string) string {
	return storageName + "---" + id
}
