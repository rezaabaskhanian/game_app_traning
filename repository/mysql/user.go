package mysql

import (
	"database/sql"
	"fmt"
	"game_app-traning/entity"
)

func (d *MySqlDB) IsPhoneNumberUniqe(phonenumber string) (bool, error) {

	user := entity.User{}
	var created_At []uint8

	row := d.db.QueryRow(`select * from usersTraning where phone_number = ?`, phonenumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &created_At)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("cant scan query result : %w", err)
	}
	return false, nil
}

func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into usersTraning(name , phone_number) values(?,?)`, u.Name, u.PhoneNumber)

	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()

	u.ID = uint(id)

	return u, nil
}
