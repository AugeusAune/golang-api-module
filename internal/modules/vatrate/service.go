package vatrate

import (
	"context"
	"golang-api-module/internal/queue"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	ctx         context.Context
	db          *gorm.DB
	queueClient *queue.Client
	log         *logrus.Logger
}

func NewService(ctx context.Context, db *gorm.DB, queueClient *queue.Client, log *logrus.Logger) *Service {
	return &Service{
		ctx:         ctx,
		db:          db,
		queueClient: queueClient,
		log:         log,
	}
}

func (s *Service) Create(req CreateVatRateRequest) (*VatRate, error) {
	vatRate := &VatRate{
		Rate:     req.Rate,
		Month:    req.Month,
		Year:     req.Year,
		IsActive: req.IsActive,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {

		if vatRate.IsActive {
			if err := tx.Unscoped().Model(&VatRate{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(vatRate).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return vatRate, nil
}

func (s *Service) Delete(id string) error {
	if err := s.db.Where("id = ?", id).Delete(&VatRate{}).Error; err != nil {
		return err
	}

	return nil
}
