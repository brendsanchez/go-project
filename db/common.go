package db

import (
	"database/sql"
	"fmt"
	"github.com/brendsanchez/go-project/models"
	"github.com/brendsanchez/go-project/secret"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Db *sql.DB
var err error
var SecretModel models.SecretRDSJson

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}

func Connect() error {
	Db, err = sql.Open("mysql", connStr(SecretModel))

	if err != nil {
		fmt.Print("> fail connexion")
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Print("> fail ping")
		return err
	}

	fmt.Print("> connexion successful")
	return nil
}

func connStr(claves models.SecretRDSJson) string {
	dbUsername := claves.Username
	dbAuthToken := claves.Password
	dbEndpoint := claves.Host
	dbName := os.Getenv("DbName")

	strConnexion := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUsername, dbAuthToken, dbEndpoint, dbName)

	fmt.Printf("> connexion %s", strConnexion)
	return strConnexion
}
