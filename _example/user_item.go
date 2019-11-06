package example

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

//go:generate go run ../cmd/sqlla/main.go

//+table: user_item
type UserItem struct {
	Id           uint64         `db:"id,primarykey,autoincrement"`
	UserId       uint64         `db:"user_id"`
	ItemId       string         `db:"item_id"`
	IsUsed       bool           `db:"is_used"`
	HasExtension sql.NullBool   `db:"has_extension"`
	UsedAt       mysql.NullTime `db:"used_at"`
}
