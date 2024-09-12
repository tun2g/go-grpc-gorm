package repository

import (
	"errors"

	pageDto "app/src/shared/dto"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository[T comparable] struct {
	DB *gorm.DB
}

func (repo *BaseRepository[T]) Create(model *T) (*T, error) {
	if err := repo.DB.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *BaseRepository[T]) FindBy(options *T) (*[]T, error) {
	var models []T
	query := repo.DB
	query = query.Where(options)

	if err := query.Find(&models).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &models, nil
}

func (repo *BaseRepository[T]) FindOneBy(options *T) (*T, error) {
	var model T
	query := repo.DB

	if options != nil {
		query = query.Where(options)
	}

	err := query.First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &model, nil
}

func (repo *BaseRepository[T]) Delete(model *T) error {
	if err := repo.DB.Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepository[T]) Update(model *T) (*T, error) {
	if err := repo.DB.Save(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *BaseRepository[T]) Paging(dto *pageDto.PageOptionsDto) (*[]T, int, error) {
	var entities []T
	query := repo.DB.
		Offset(*dto.Offset).
		Limit(*dto.Limit)

	if dto.Order != nil {
		orderField := "createdAt"
		if dto.OrderField != nil {
			dto.OrderField = &orderField
		}
		query.Order(fmt.Sprintf("%s %s", *dto.OrderField, *dto.Order))
	}

	query.Find(&entities)

	err := query.Error

	if err != nil {
		return nil, 0, err
	}

	var count int64
	var model T
	err = repo.DB.Model(&model).Count(&count).Error

	if err != nil {
		return nil, 0, err
	}
	return &entities, int(count), nil
}
