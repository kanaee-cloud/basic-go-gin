package repository

import (
	"base-gin/app/domain/dao"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

func newPublisherRepo(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{db: db}
}

func (r *PublisherRepository) Create(newItem *dao.Publisher) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PublisherRepository) GetList() ([]dao.Publisher, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var items []dao.Publisher
	tx := r.db.WithContext(ctx).Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return items, nil
}

func (r *PublisherRepository) Delete(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Delete(&dao.Publisher{}, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return exception.ErrUserNotFound
		}
		return tx.Error
	}

	return nil
}

func (r *PublisherRepository) Update(publisher *dao.Publisher) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Save(publisher)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PublisherRepository) GetByID(id uint) (*dao.Publisher, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Publisher
	tx:= r.db.WithContext(ctx).First(&item, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			return nil, exception.ErrUserNotFound
		}

		return nil, tx.Error
	}

	return &item, nil
}



