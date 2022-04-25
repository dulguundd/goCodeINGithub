package device

import (
	"database/sql"
	"github.com/dulguundd/logError-lib/errs"
	"github.com/dulguundd/logError-lib/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RepositoryDb struct {
	pool *sqlx.DB
}

func (d RepositoryDb) GetAllDevice() ([]Device, *errs.AppError) {
	var devices []Device
	var err error

	queryCommand := "SELECT * FROM device ORDER BY device_id DESC"
	err = d.pool.Select(&devices, queryCommand)

	if err != nil {
		logger.Error("Error while querying device table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return devices, nil
}

func (d RepositoryDb) GetDevice(id int) (*Device, *errs.AppError) {
	var r Device
	var err error

	queryCommand := "SELECT * FROM device where device_id = $1"
	err = d.pool.Get(&r, queryCommand, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Device not found")
		} else {
			logger.Error("Error while querying device table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}

	return &r, nil
}

func NewRepositoryDb(dbClient *sqlx.DB) RepositoryDb {
	return RepositoryDb{dbClient}
}
