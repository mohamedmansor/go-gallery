package service

import (
	"blog/api/repository"
	"blog/models"
)

type ImageService struct {
	repository repository.ImageRepository
}

//NewImageService : returns the ImageService struct instance
func NewImageService(r repository.ImageRepository) ImageService {
	return ImageService{
		repository: r,
	}
}

//Save -> calls image repository save method
func (p ImageService) Save(image models.Image) error {
	return p.repository.Save(image)
}

//FindAll -> calls image repo find all method
func (p ImageService) FindAll(image models.Image, keyword string) (*[]models.Image, int64, error) {
	return p.repository.FindAll(image, keyword)
}

// Update -> calls image repo update method
func (p ImageService) Update(image models.Image) error {
	return p.repository.Update(image)
}

// Delete -> calls image repo delete method
func (p ImageService) Delete(id int64) error {
	var image models.Image
	image.ID = id
	return p.repository.Delete(image)
}

// Find -> calls image repo find method
func (p ImageService) Find(image models.Image) (models.Image, error) {
	return p.repository.Find(image)
}
