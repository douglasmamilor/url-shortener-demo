package dal

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"url-shortener/pkg/config"
)

// DAL data access layer
type DAL struct {
	DB *sqlx.DB

	// DAL Objects
	URLDAL IURLDAL
}

func (d *DAL) setupDALObjects(cfg *config.Config) error {
	connString := fmt.Sprintf("host=%s port=%d user=%v password=%v dbname=%v sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return errors.Wrapf(err, "%v - Postgres: Unable to open initial connection to postgres at: %v", ErrDBSetupFailed, cfg.PostgresHost)
	}

	d.DB = db
	d.URLDAL = NewURLDAL(db)

	return nil
}

// New creates, configures and returns a new DAL object
func New(cfg *config.Config) (*DAL, error) {
	dal := &DAL{}

	if err := dal.setupDALObjects(cfg); err != nil {
		return nil, err
	}

	return dal, nil
}
