package repository

import (
	pageDto "app/src/shared/dto"
)

type IBaseRepository[T comparable] interface {
	Create(model *T) (*T, error)
	FindBy(options *T) (*[]T, error)
	FindOneBy(options *T) (*T, error)
	Delete(*T) error
	Update(model *T) (*T, error)
	Paging(dto *pageDto.PageOptionsDto) (*[]T, int, error)
}
