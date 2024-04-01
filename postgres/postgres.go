package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Connect(connectionString string) (*Db, error) {
	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	return &Db{db}, err
}

func BuildConnectionString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbName)
}

func (db Db) GetAllPatients() ([]Patient, error) {
	rows, err := db.Query("SELECT * FROM patients")
	if err != nil {
		fmt.Println("GetAllPatients error: ", err)
		return []Patient{}, err
	}

	var patient Patient
	patients := []Patient{}
	for rows.Next() {
		err = rows.Scan(
			&patient.ID,
			&patient.Name,
			&patient.OrderId,
		)
		if err != nil {
			fmt.Println("Sacn rows error: ", err)
			return patients, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

type Patient struct {
	ID      int
	Name    string
	OrderId int
}

type Order struct {
	ID      int
	Message string
}
