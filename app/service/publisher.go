package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type PublisherService struct {
	repo *repository.PublisherRepository
}

// Corrected function signature with matching parameter names
func newPublisherService(publisherRepo *repository.PublisherRepository) *PublisherService {
	return &PublisherService{repo: publisherRepo}
}

func (s *PublisherService) Create(params *dto.PublisherCreateReq) (*dto.PublisherCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.PublisherCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *PublisherService) GetByID(id uint) (dto.PublisherDetailResp, error) {
	var resp dto.PublisherDetailResp

	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrUserNotFound
	}

	resp.FromEntity(item)

	return resp, nil
}

func (s *PublisherService) Update(id uint, params *dto.PublisherCreateReq) (*dto.PublisherDetailResp, error) {
	
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, exception.ErrUserNotFound
	}

	
	item.Name = params.Name
	item.City = params.City

	// Simpan data yang diperbarui ke repository
	err = s.repo.Update(item)
	if err != nil {
		return nil, err
	}

	// Convert updated item ke DTO untuk response
	var resp dto.PublisherDetailResp
	resp.FromEntity(item)

	return &resp, nil
}

func (s *PublisherService) GetList() ([]dto.PublisherDetailResp, error) {
	items, err := s.repo.GetList()
	if err != nil {
		return nil, err
	}

	var resp []dto.PublisherDetailResp
	for _, item := range items {
		var publisherResp dto.PublisherDetailResp
		publisherResp.FromEntity(&item)
		resp = append(resp, publisherResp)
	}

	return resp, nil
}

func (s *PublisherService) Delete(id uint) error {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if item == nil {
		return exception.ErrUserNotFound
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
