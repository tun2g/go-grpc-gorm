package dto

import (
	constants "app/src/shared/constants"
)

type PageOptionsDto struct {
	Q          *string          `form:"q" binding:"omitempty"`
	Order      *constants.Order `form:"order" binding:"omitempty,oneof=ASC DESC"`
	OrderField *string          `form:"order_field" binding:"omitempty"`
	Offset     *int             `form:"offset" binding:"omitempty,min=0"`
	Limit      *int             `form:"limit" binding:"omitempty,min=1,max=50"`
}

func NewPageOptionsDto(pageOptionsDto *PageOptionsDto) *PageOptionsDto {
	limit := 50
	if pageOptionsDto.Limit != nil {
		limit = *pageOptionsDto.Limit
	}

	offset := 0
	if pageOptionsDto.Offset != nil {
		offset = *pageOptionsDto.Offset
	}
	return &PageOptionsDto{
		Limit:      &limit,
		Offset:     &offset,
		Q:          pageOptionsDto.Q,
		Order:      pageOptionsDto.Order,
		OrderField: pageOptionsDto.OrderField,
	}
}

type PageMetaDtoParameters struct {
	PageOptionsDto PageOptionsDto
	Total          int
}

type PageMetaDto struct {
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	Offset  int  `json:"offset"`
	HasNext bool `json:"hasNext"`
	HasPrev bool `json:"hasPrev"`
}

func NewPageMetaDto(pageOptionsDto *PageOptionsDto, total int) *PageMetaDto {
	if pageOptionsDto.Limit == nil {
		*pageOptionsDto.Limit = 50
	}
	if pageOptionsDto.Offset == nil {
		*pageOptionsDto.Offset = 0
	}

	return &PageMetaDto{
		Total:   total,
		Limit:   *pageOptionsDto.Limit,
		Offset:  *pageOptionsDto.Offset,
		HasNext: *pageOptionsDto.Offset+*pageOptionsDto.Limit < total,
		HasPrev: *pageOptionsDto.Offset > *pageOptionsDto.Limit,
	}
}

type PageDto[T comparable] struct {
	Data []T          `json:"data"`
	Meta *PageMetaDto `json:"meta"`
}

func NewPageDto[T comparable](data []T, meta *PageMetaDto) *PageDto[T] {
	return &PageDto[T]{
		Data: data,
		Meta: meta,
	}
}
