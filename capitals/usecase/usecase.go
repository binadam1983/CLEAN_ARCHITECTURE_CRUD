package usecase

import (
	"github.com/the_clean_architecture_crud/capitals"
	"github.com/the_clean_architecture_crud/model"
)

type UsecaseImpl struct {
	capitalsRepo capitals.CapitalsRepo
}

func CreateCapitalsUsecase(capitalsRepo capitals.CapitalsRepo) capitals.CapitalsUsecase {
	return &UsecaseImpl{capitalsRepo}
}

func (u *UsecaseImpl) InsertCapital(capital *model.Capitals) (int, error) {
	return u.capitalsRepo.InsertCapital(capital)
}
func (u *UsecaseImpl) ShowAll() (*[]model.Capitals, error) {
	return u.capitalsRepo.ShowAll()
}
func (u *UsecaseImpl) ShowById(id int) (*model.Capitals, error) {
	return u.capitalsRepo.ShowById(id)
}
func (u *UsecaseImpl) UpdateCapital(capital *model.Capitals) (int, error) {
	return u.capitalsRepo.UpdateCapital(capital)
}
func (u *UsecaseImpl) DeleteCapital(id int) (int, error) {
	return u.capitalsRepo.DeleteCapital(id)
}
func (u *UsecaseImpl) EditHTML(id int) (*model.Capitals, error) {
	return u.capitalsRepo.EditHTML(id)
}
func (u *UsecaseImpl) DeleteHTML(id int) (*model.Capitals, error) {
	return u.capitalsRepo.DeleteHTML(id)
}
