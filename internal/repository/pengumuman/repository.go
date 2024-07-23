package pengumuman

import (
	"log"

	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type PengumumanRepository interface {
	CreatePengumuman(pengumuman *models.Pengumuman) error
	GetAllPengumuman() ([]models.Pengumuman, error)
	UpdatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error)
	GetPengumumanByID(pengumumanID uint) (models.Pengumuman, error)
	DeletePengumumanByID(pengumumanID uint) error
}

type pengumumanRepository struct {
	db *gorm.DB
}

func NewPengumumanRepository(db *gorm.DB) PengumumanRepository {
	return &pengumumanRepository{db}
}

func (r *pengumumanRepository) CreatePengumuman(pengumuman *models.Pengumuman) error {
	return r.db.Create(pengumuman).Error
}

func (r *pengumumanRepository) GetAllPengumuman() ([]models.Pengumuman, error) {
	// start := time.Now()

	tx := r.db.Begin()
	var pengumuman []models.Pengumuman
	if err := tx.Debug().Find(&pengumuman).Error; err != nil {
		return nil, err
	}
	// elapsed := time.Since(start)
	// log.Printf("get all pengumuman took %s", elapsed)
	return pengumuman, nil
}

func (r *pengumumanRepository) UpdatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Save(&pengumuman).Error; err != nil {
		tx.Rollback()
		return models.Pengumuman{}, err
	}
	log.Println("error 1 repo")
	tx.Commit()
	return pengumuman, nil
}

func (r *pengumumanRepository) GetPengumumanByID(pengumumanID uint) (models.Pengumuman, error) {
	tx := r.db.Begin()
	var pengumuman models.Pengumuman
	if err := tx.Debug().Take(&pengumuman, "id = ?", pengumumanID).Error; err != nil {
		return pengumuman, err
	}
	return pengumuman, nil
}

func (r *pengumumanRepository) DeletePengumumanByID(pengumumanID uint) error {
	tx := r.db.Begin()

	if err := tx.Debug().Delete(&models.Pengumuman{}, pengumumanID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
