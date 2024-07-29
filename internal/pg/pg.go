package pg

import (
	"dobrino/config"
	"fmt"
	"time"

	"dobrino/internal/pg/dbmodels"
	"dobrino/internal/pg/dbquery"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PG struct {
	*dbquery.Query
}

var DB = Init(config.POSTGRES_CONN_STRING)

func Init(connString string) PG {

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprint("Failed to connect to database at dsn: ", connString))
	}

	return PG{
		dbquery.Use(db),
	}
}

func (pg PG) CreateAdmin(chatId string) error {
	values := &dbmodels.Admin{
		ID:          uuid.New().String(),
		DateCreated: time.Now(),
		ChatID:      chatId,
	}

	err := pg.Admin.Create(values)
	if err != nil {
		return err
	}

	return nil
}

func (pg PG) GetAdmin(chatId string) (*dbmodels.Admin, error) {
	a := dbquery.Admin
	admin, err := pg.Admin.Where(a.ChatID.Eq(chatId)).First()
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (pg PG) CreateUser(chatId string) error {
	values := &dbmodels.User{
		ID:          uuid.New().String(),
		DateCreated: time.Now(),
		ChatID:      chatId,
	}
	err := pg.User.Create(values)
	if err != nil {
		return err
	}

	return nil
}

func (pg PG) GetUser(chatId string) (*dbmodels.User, error) {
	u := pg.User
	user, err := pg.User.Where(u.ChatID.Eq(chatId)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (pg PG) GetButtons() ([]*dbmodels.Buttons, error) {
	buttons, err := pg.Buttons.Find()
	if err != nil {
		return nil, err
	}

	return buttons, err
}
