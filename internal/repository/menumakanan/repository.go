package menumakanan

import (
	"log"

	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type MenuMakananRepository interface {
	CreateMenuMakanan(menu models.MenuMakanan) error
	GetAllMenuMakanan() ([]models.MenuMakanan, error)
	DeleteMenuMakanan(id uint) error
	GetMenuMakananByID(id uint) (models.MenuMakanan, error)
	UpdateMenuMakanan(menuMakanan models.MenuMakanan) (models.MenuMakanan, error)
}

type menuMakananRepository struct {
	db *gorm.DB
}

func NewMenuMakananRepo(db *gorm.DB) MenuMakananRepository {
	return &menuMakananRepository{db: db}
}

func (r *menuMakananRepository) UpdateMenuMakanan(menuMakanan models.MenuMakanan) (models.MenuMakanan, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Save(&menuMakanan).Error; err != nil {
		tx.Rollback()
		return models.MenuMakanan{}, err
	}

	log.Println("hello world")
	tx.Commit()
	return menuMakanan, nil
}

func (r *menuMakananRepository) DeleteMenuMakanan(id uint) error {
	tx := r.db.Begin()

	if err := tx.Debug().Delete(&models.MenuMakanan{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *menuMakananRepository) CreateMenuMakanan(menu models.MenuMakanan) error {
	return r.db.Create(&menu).Error
}

func (r *menuMakananRepository) GetAllMenuMakanan() ([]models.MenuMakanan, error) {
	tx := r.db.Begin()

	var menus []models.MenuMakanan
	if err := tx.Debug().Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuMakananRepository) GetMenuMakananByID(id uint) (models.MenuMakanan, error) {
	tx := r.db.Begin()
	var menuMakanan models.MenuMakanan
	if err := tx.Debug().Take(&menuMakanan, "id = ?", id).Error; err != nil {
		return menuMakanan, err
	}

	return menuMakanan, nil
}
