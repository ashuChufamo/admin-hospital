package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/webProj/Admin"
	"github.com/webProj/entity"
)

// ManageDoctorsRepository implements Admin.ManageDoctorsRepository interface
type ManageDoctorsRepository struct {
	conn *gorm.DB
}

// NewManageDoctorsRepository returns new object of ManageDoctorsRepository
func NewManageDoctorsRepository(db *gorm.DB) Admin.ManageDoctorsRepository {
	return &ManageDoctorsRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mdRepo *ManageDoctorsRepository) Doctors() ([]entity.Doctor, []error) {
	docs := []entity.Doctor{}
	errs := mdRepo.conn.Find(&docs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return docs, errs
}

// Doctor retrieves a doctor from the database by its id
func (mdRepo *ManageDoctorsRepository) Doctor(id uint) (*entity.Doctor, []error) {
	docs := entity.Doctor{}
	errs := mdRepo.conn.First(&docs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &docs, errs
}

// UpdateDoctor updats a given doctor in the database
func (mdRepo *ManageDoctorsRepository) UpdateDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	doc := user
	errs := mdRepo.conn.Save(doc).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return doc, errs
}

// DeleteDoctor deletes a given doctor from the database
func (mdRepo *ManageDoctorsRepository) DeleteDoctor(id uint) (*entity.Doctor, []error) {
	doc, errs := mdRepo.Doctor(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mdRepo.conn.Delete(doc, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return doc, errs
}

// StoreDoctor stores a given doctor in the database
func (mdRepo *ManageDoctorsRepository) StoreDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	doc := user
	errs := mdRepo.conn.Create(doc).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return doc, errs
}
