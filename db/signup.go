package db

import (
	"fmt"
	"github.com/brendsanchez/go-project/models"
	"github.com/brendsanchez/go-project/tools"
)

func SignUp(up models.SignUp) error {
	err := Connect()
	if err != nil {
		fmt.Print("> error to connect")
		return err
	}
	defer Db.Close()
	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + up.UserEmail + "','" + up.UserUUID + "','" + tools.DateMySql() + "')"
	fmt.Printf("> query : %s \n", query)

	_, err = Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
