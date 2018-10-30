package repos

import (
	"database/sql"
	"log"

	"github.com/allyraza/choto/models"
)

// URLRepo :
type URLRepo struct {
	Database *sql.DB
}

// FindByID : finds a url by id
func (ur *URLRepo) FindByID(id int) *models.URL {
	url := &models.URL{}

	err := ur.Database.QueryRow("SELECT id, short, url, created_at, updated_at FROM urls WHERE id = ?", id).
		Scan(&url.ID, &url.Short, &url.Long, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		log.Println(err)
	}

	return url
}

// FindByKey : finds a url by key
func (ur *URLRepo) FindByKey(key string) (*models.URL, error) {
	url := &models.URL{}

	err := ur.Database.QueryRow("SELECT id, short_url, url, created_at, updated_at FROM urls WHERE short_url = ?", key).
		Scan(&url.ID, &url.Short, &url.Long, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return url, nil
}
