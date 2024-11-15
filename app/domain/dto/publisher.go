package dto

import "base-gin/app/domain/dao"

type PublisherCreateReq struct {
	Name string `json:"name" binding:"required,max=48,min=6"`
	City string `json:"city" binding:"required,max=32,min=2"`
}

func (o PublisherCreateReq) ToEntity() dao.Publisher {
	return dao.Publisher{
		Name: o.Name,
		City: o.City,
	}
}

type PublisherUpdateReq struct {
	Name string `json:"name" binding:"omitempty,max=48,min=6"`
	City string `json:"city" binding:"omitempty,max=32,min=2"`
}

func (o PublisherUpdateReq) ToEntity(item *dao.Publisher) {
	if o.Name != "" {
		item.Name = o.Name
	}
	if o.City != "" {
		item.City = o.City
	}
}

type PublisherDetailResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (o *PublisherDetailResp) FromEntity(item *dao.Publisher) {
	o.ID = int(item.ID)
	o.Name = item.Name
	o.City = item.City
}

type PublisherCreateResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (o *PublisherCreateResp) FromEntity(item *dao.Publisher) {
	o.ID = int(item.ID)
	o.Name = item.Name
	o.City = item.City
}
