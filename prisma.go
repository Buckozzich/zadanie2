package prisma

import (
	"fmt"
	"time"
	"GORM"
	"application/json"
)

type config struct {
	provider := "postgresql"
	URL := env("DATABASE_URL")
}

type user struct{
	Id int (&default(autoincrement()), *id)
email string &unique()
name string
password string
account Account[]string
history History[]string

}

type modelAccount struct{
	Id int (&default(autoincrement()), *id)
	title string
	var Decimal string
	type String
	author Users *relation(fields[authorID]int, referenses[id])
	history History[]string
}

type History struct{
	Id int (&default(autoincrement()), *id)  `gorm:"primary_key", json:"id"`
	var Decimal string
	operation    string
	CreatedAt time.Time  `json:"created_at"`
     UpdateAt time.Time `json:"update_at"`
     DeleteAt Time.Time  `json: "delete_at"`
}
