package db

import (
	"fmt"
	"strings"
)

type Todo struct {
	Id   int64  `db:"id, primarykey, autoincrement" json:"id"`
	Text string `db:"text" json:"text"`
	Done bool   `db:"done" json:"done"`
}

func (db *DB) TodoList() ([]Todo, error) {
	var todos []Todo
	_, err := db.dbmap.Select(&todos, "SELECT * FROM todo ORDER BY Id")
	if err != nil {
		return nil, fmt.Errorf("select error: %v", err)
	}
	return todos, err
}

func (db *DB) TodoCreate(text string) error {
	new := Todo{
		Text: text,
		Done: false,
	}
	err := db.dbmap.Insert(&new)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) TodoSwitch(id int) error {
	var update Todo
	err := db.dbmap.SelectOne(&update, "SELECT * FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}
	update.Done = !update.Done
	_, err = db.dbmap.Update(&update)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) TodoDelete(ids []int) error {
	q := "1 = 0"
	if 0 < len(ids) {
		q = "id IN (?" + strings.Repeat(", ?", len(ids)-1) + ")"
	}
	var args []interface{}
	for _, v := range ids {
		args = append(args, v)
	}
	_, err := db.dbmap.Exec("DELETE FROM todo WHERE "+q, args...)
	if err != nil {
		return err
	}
	return nil

}
