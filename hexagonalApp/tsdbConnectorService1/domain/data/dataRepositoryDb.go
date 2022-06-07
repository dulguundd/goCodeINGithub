package data

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/dulguundd/logError-lib/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RepositoryDb struct {
	pool *sqlx.DB
}

func (d RepositoryDb) GetLastDataOfTodayByHour() ([]Data, *errs.AppError) {
	var datas []Data
	var err error

	queryCommand := "select bucket, last as temp, device_id from temp_metrics_h_last where bucket > now() - interval '1 day' ORDER BY bucket ASC;"

	//queryCommand := "select time, last as temp, device_id from temp_metrics_h_last where time > now() - interval '1 day' ORDER BY time ASC;"
	err = d.pool.Select(&datas, queryCommand)

	if err != nil {
		logger.Error("Error while querying data table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return datas, nil
}
func (d RepositoryDb) GetLastDataOfTodayByHourOfDevice(id int) ([]Data, *errs.AppError) {
	var datas []Data
	var err error

	queryCommand := "select bucket, last as temp, device_id from temp_metrics_h_last WHERE bucket > now() - interval '1 day' and device_id = $1 ORDER BY bucket ASC;"
	err = d.pool.Select(&datas, queryCommand, id)

	if err != nil {
		logger.Error("Error while querying data table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return datas, nil
}

func (d RepositoryDb) PostData(newData Data) (*Data, *errs.AppError) {

	queryCommand := " INSERT INTO data(time, temperature, device_id) VALUES($1,$2,$3);"
	_, err := d.pool.Exec(queryCommand, newData.Bucket.Format("2006-01-02 15:04:05"), newData.Temp, newData.Device_id)
	if err != nil {
		logger.Error("Error while inserting data: " + err.Error())
		return nil, errs.NewUnexpectedError("Cannot insert data")
	}
	return &newData, nil
}

func NewRepositoryDb(dbClient *sqlx.DB) RepositoryDb {
	return RepositoryDb{dbClient}
}
