package service

import (
	"base-gin/app/repository"
	"base-gin/app/domain/dto"
)

type PublisherService struct {
	repo *repository.PublisherRepository
}

// Corrected function signature with matching parameter names
func newPublisherService(publisherRepo *repository.PublisherRepository) *PublisherService {
	return &PublisherService{repo: publisherRepo}
}

func (s *PublisherService) Create(params *dto.PublisherCreateReq)(*dto.PublisherCreateResp, error){
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.PublisherCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

