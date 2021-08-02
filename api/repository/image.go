package repository

import (
	"blog/infrastructure"
	"blog/models"
)

//ImageRepository -> ImageRepository
type ImageRepository struct {
	db infrastructure.Database
}

// NewImageRepository : fetching database
func NewImageRepository(db infrastructure.Database) ImageRepository {
	return ImageRepository{
		db: db,
	}
}

//Save -> Method for saving image to database
func (p ImageRepository) Save(image models.Image) error {
	return p.db.DB.Create(&image).Error
}

//FindAll -> Method for fetching all images from database
func (p ImageRepository) FindAll(image models.Image, keyword string) (*[]models.Image, int64, error) {
	var images []models.Image
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("Created_at desc").Model(&models.Image{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("image.title LIKE ? ", queryKeyword))
	}

	err := queryBuilder.
		Where(image).
		Find(&images).
		Count(&totalRows).Error

	return &images, totalRows, err
}

//Update -> Method for updating Image
func (p ImageRepository) Update(image models.Image) error {
	return p.db.DB.Save(&image).Error
}

//Find -> Method for fetching image by id
func (p ImageRepository) Find(image models.Image) (models.Image, error) {
	var images models.Image
	err := p.db.DB.
		Debug().
		Model(&models.Image{}).
		Where(&image).
		Take(&images).Error
	return images, err
}

//Delete Deletes Image
func (p ImageRepository) Delete(image models.Image) error {
	return p.db.DB.Delete(&image).Error
}
