package database

import (
	"go-was-example/api-db-server/datatype"
)

func (rdb *RDBHandler) GetUser() (bool, []datatype.User, error) {
	var users []datatype.User
	resultDB := rdb.db.Table(User{}.TableName()).Find(&users)
	if resultDB.Error != nil {
		return false, users, resultDB.Error
	}
	exist := resultDB.RowsAffected > 0
	return exist, users, nil
}
