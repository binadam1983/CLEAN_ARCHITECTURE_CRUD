package repo

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/the_clean_architecture_crud/capitals"
	"github.com/the_clean_architecture_crud/model"
)

type CapitalsRepoImpl struct {
	db *sql.DB
}

//we could instead place a switch statement here for the type of DB available/receive & initialize repoImpl accordingly
//but since we are working with only sql.DB, we punched it in directly

func CreateCapitalsRepo(DB *sql.DB) capitals.CapitalsRepo {
	return &CapitalsRepoImpl{DB}
}

func (repo *CapitalsRepoImpl) InsertCapital(capital *model.Capitals) (int, error) {

	insForm, err := repo.db.Prepare("INSERT INTO capitals (country, capital) VALUES (?,?)")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	insForm.Exec(capital.Country, capital.Capital)
	return http.StatusOK, nil
}

func (repo *CapitalsRepoImpl) ShowAll() (*[]model.Capitals, error) {

	selDB, err := repo.db.Query("SELECT * FROM capitals ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	row := model.Capitals{}
	result := []model.Capitals{}
	for selDB.Next() {
		_ = selDB.Scan(&row.Id, &row.Country, &row.Capital)

		result = append(result, row)
	}
	return &result, nil
}

func (repo *CapitalsRepoImpl) ShowById(id int) (*model.Capitals, error) {

	row, err := repo.db.Query("SELECT * FROM capitals WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	temp := &model.Capitals{}
	for row.Next() {
		_ = row.Scan(&temp.Id, &temp.Country, &temp.Capital)
	}
	return temp, nil
}

func (repo *CapitalsRepoImpl) UpdateCapital(capital *model.Capitals) (int, error) {

	sqlResult, err := repo.db.Prepare("UPDATE Capitals SET country=?, capital=? WHERE id=? ;")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	sqlResult.Exec(capital.Country, capital.Capital, capital.Id)

	return http.StatusOK, nil
}

func (repo *CapitalsRepoImpl) DeleteCapital(id int) (int, error) {

	sqlResult, err := repo.db.Prepare("DELETE FROM capitals WHERE id=?")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	sqlResult.Exec(id)
	log.Println(repo)
	return http.StatusOK, nil
}

func (repo *CapitalsRepoImpl) EditHTML(id int) (*model.Capitals, error) {

	row, err := repo.db.Query("SELECT * FROM capitals WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	temp := &model.Capitals{}
	for row.Next() {
		_ = row.Scan(&temp.Id, &temp.Country, &temp.Capital)
	}
	return temp, nil
}

func (repo *CapitalsRepoImpl) DeleteHTML(id int) (*model.Capitals, error) {

	row, err := repo.db.Query("SELECT * FROM capitals WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	temp := &model.Capitals{}
	for row.Next() {
		_ = row.Scan(&temp.Id, &temp.Country, &temp.Capital)
	}
	return temp, nil
}
