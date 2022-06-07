package data

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

func Test_should(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	client := sqlx.NewDb(db, "sqlmock")
	DataRepo := NewRepositoryDb(client)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newData := Data{
		Bucket:    time.Now(),
		Temp:      10,
		Device_id: 1,
	}
	// now we execute our method
	if _, AppError := DataRepo.PostData(newData); AppError == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

//func Test_should_return(t *testing.T) {
//	mockDB, mocksql, _ := sqlmock.New()
//	defer mockDB.Close()
//
//	//create a db client for account repository
//	client := sqlx.NewDb(mockDB, "sqlmock")
//	DataRepo := RepositoryDb{client}
//
//	// fake transaction instance
//	newData := Data{time.Now(), 10, 100}
//
//	// setting expectations and mock objects
//	mocksql.ExpectBegin()
//
//	mocksql.ExpectExec(`INSERT INTO data`).
//		WithArgs(newData.Time, newData.Temp, newData.Device_id).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	mocksql.ExpectRollback()
//
//	// now we execute our method
//	if _, appError := DataRepo.PostData(newData); appError == nil {
//		t.Errorf("was expecting an error, but there was none")
//	}
//}

func Test_should_return_GetLastDataOfTodayByHour(t *testing.T) {
	mockDB, mocksql, _ := sqlmock.New()
	defer mockDB.Close()

	//create a db client for account repository
	client := sqlx.NewDb(mockDB, "sqlmock")
	DataRepo := RepositoryDb{client}

	rows := sqlmock.NewRows([]string{"time", "temp", "device_id"}).
		AddRow(time.Now(), 20, 1)
	mocksql.ExpectPrepare("select time, last as temp, device_id from temp_metrics_h_last where time > now() - interval '1 day' ORDER BY time ASC;").ExpectQuery().WillReturnRows(rows)

	_, appError := DataRepo.GetLastDataOfTodayByHour()

	if appError == nil {
		t.Errorf("Test failed while quering data")
	}
}

func Test_should_return_GetLastDataOfTodayByHourOfDevice(t *testing.T) {
	mockDB, mocksql, _ := sqlmock.New()
	defer mockDB.Close()

	//create a db client for account repository
	client := sqlx.NewDb(mockDB, "sqlmock")
	DataRepo := RepositoryDb{client}

	deviceid := 1

	rows := sqlmock.NewRows([]string{"time", "temp", "device_id"}).
		AddRow(time.Now(), 20, 1)
	mocksql.ExpectPrepare("select time, last as temp, device_id from temp_metrics_h_last where time > now() - interval '1 day' and device_id = $1 ORDER BY time ASC;").ExpectQuery().WithArgs(deviceid).WillReturnRows(rows)

	_, appError := DataRepo.GetLastDataOfTodayByHourOfDevice(deviceid)

	if appError == nil {
		t.Errorf("Test failed while quering data")
	}
}
