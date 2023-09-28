package dal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"database/sql"
	"fmt"

	"url-shortener/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

//counterfeiter:generate . IURLDAL
type IURLDAL interface {
	AddURL(url *model.URL) error
	FindByShortCode(shortCode string) (*model.URL, error)
}

// URLDAL ...
type URLDAL struct {
	DB *sqlx.DB
}

// NewURLDAL creates an instance of a URL DAL
func NewURLDAL(db *sqlx.DB) *URLDAL {
	return &URLDAL{
		DB: db,
	}
}

func (u URLDAL) AddURL(url *model.URL) error {
	_, err := u.DB.NamedExec(
		`INSERT INTO "url"(
			id,
      short_code,
			original_url,
			created_at
		) VALUES (
			:id,
      :short_code,
			:original_url,
			:created_at
		)
		ON CONFLICT (short_code)
		DO
		UPDATE SET
		short_code = EXCLUDED.short_code,
		original_url = EXCLUDED.original_url,
		created_at = EXCLUDED.created_at
		`,
		map[string]interface{}{
			"id":           url.ID,
			"short_code":   url.ShortCode,
			"original_url": url.OriginalURL,
			"created_at":   url.CreatedAt,
		})

	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok {
			if postgresErr.Code == "23505" {
				return errors.New(ErrDuplicateKey)
			}
		}

		return errors.Wrapf(err, fmt.Sprintf("Could not perform db insert for url: %v", url))
	}

	return nil
}

func (u URLDAL) FindByShortCode(shortCode string) (*model.URL, error) {
	urlRecord := &model.URL{}

	if err := u.DB.Get(urlRecord, `SELECT * FROM url WHERE short_code = $1`, shortCode); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ErrNoResult)
		}
		return nil, err
	}

	return urlRecord, nil
}
