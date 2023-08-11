package models

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	
	
)

func WithdrawIR(username string,bookId int) (error) {
	var userID int
	db, err := Connection()
	error := db.QueryRow("SELECT user_id FROM user WHERE name=?", username).Scan(&userID)
	if error != nil {
		return error
	}
   fmt.Println(bookId,userID)
	rows, err := db.Query(`DELETE FROM request WHERE book_id = ? and user_id=? and status != "owned";`,bookId,userID)
	if err != nil {
		return err
		
	}
	defer rows.Close()

	
	return nil
}