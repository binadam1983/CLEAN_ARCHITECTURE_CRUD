package capitals

import "github.com/the_clean_architecture_crud/model"

type CapitalsRepo interface {
	InsertCapital(capital *model.Capitals) (int, error)
	ShowAll() (*[]model.Capitals, error)
	ShowById(id int) (*model.Capitals, error)
	EditHTML(id int) (*model.Capitals, error)
	UpdateCapital(capital *model.Capitals) (int, error)
	DeleteCapital(id int) (int, error)
	DeleteHTML(id int) (*model.Capitals, error)
}
