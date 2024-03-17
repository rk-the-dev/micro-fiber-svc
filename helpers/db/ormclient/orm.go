package ormclient

import (
	"gorm.io/gorm"
)

type ORMHelper struct{}

func NewORMHelper() *ORMHelper {
	return &ORMHelper{}
}

func (*ORMHelper) GetDB(dbType string) (*gorm.DB, error) {
	var db *gorm.DB
	//var err error

	// switch dbType {
	// case app.DB_MY_SQL:
	// Create DNS
	// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// case "postgres":
	// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// default:
	// 	return nil, fmt.Errorf("unsupported database type: %s", dbType)
	// }

	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
