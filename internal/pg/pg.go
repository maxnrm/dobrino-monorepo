package pg

import (
	"dobrino/config"
	"errors"
	"fmt"
	"time"

	"dobrino/internal/pg/dbmodels"
	"dobrino/internal/pg/dbquery"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PG struct {
	*dbquery.Query
}

var DB = Init(config.POSTGRES_CONN_STRING)

func Init(connString string) *PG {

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprint("Failed to connect to database at dsn: ", connString))
	}

	return &PG{
		dbquery.Use(db),
	}
}

func (pg *PG) CreateAdmin(chatId string) error {
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

func (pg *PG) GetAdmin(chatId string) (*dbmodels.Admin, error) {
	a := dbquery.Admin
	admin, err := pg.Admin.Where(a.ChatID.Eq(chatId)).First()
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (pg *PG) CreateUser(chatId string) error {
	user := &dbmodels.User{
		ID:          uuid.New().String(),
		DateCreated: time.Now(),
		ChatID:      chatId,
	}

	err := pg.User.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PG) IncrementUserInteractions(chatId string) error {
	u := pg.User

	affected, err := pg.User.Where(u.ChatID.Eq(chatId)).Update(u.Interactions, u.Interactions.Add(1))
	if err != nil {
		return err
	}

	if affected.RowsAffected == 0 {
		return errors.New("interactions update: no rows affected")
	}

	return nil
}

func (pg *PG) GetUser(chatId string) (*dbmodels.User, error) {
	u := pg.User
	user, err := pg.User.Where(u.ChatID.Eq(chatId)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (pg *PG) GetUsers() ([]*dbmodels.User, error) {
	users, err := pg.User.Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (pg *PG) GetButtons() ([]*dbmodels.Button, error) {
	b := pg.Button
	f := pg.File
	buttons, err := pg.Button.Order(b.Sort.Asc()).Find()
	if err != nil {
		return nil, err
	}

	for _, button := range buttons {
		if button.Image != nil {
			file, err := pg.File.Where(f.ID.Eq(*button.Image)).Take()
			if err != nil {
				button.Image = nil
				continue
			}
			// dirty hack to avoid writing excessive logic with another field
			button.Image = file.FilenameDisk
		}
	}

	return buttons, nil
}

func (pg *PG) GetBroadcastMessageForSend() (*dbmodels.BroadcastMessage, error) {
	bm := pg.BroadcastMessage
	f := pg.File
	msg, err := pg.BroadcastMessage.Where(bm.IsSent.Is(true)).Where(bm.SendStatus.IsNull()).First()
	if err != nil {
		return nil, err
	}

	if msg.Image != nil {
		file, err := pg.File.Where(f.ID.Eq(*msg.Image)).Take()
		if err != nil {
			msg.Image = nil
			return msg, nil
		}
		msg.Image = file.FilenameDisk
	}

	return msg, nil
}

func (pg *PG) SetBroadcastMessageStatus(id int32, status bool) error {
	bm := pg.BroadcastMessage
	_, err := pg.BroadcastMessage.Where(bm.ID.Eq(id)).Update(bm.SendStatus, status)
	if err != nil {
		return err
	}

	return nil
}
