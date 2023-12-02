package repository

import (
	"ed/pkg/util"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type commonRepo struct {
	db *sqlx.DB
}

func newCommonRepo(db *sqlx.DB) *commonRepo {
	return &commonRepo{db: db}
}

func (db *commonRepo) CheckExistInDb(table, param, value string) (bool, error) {
	str := fmt.Sprintf("select exists(select 1 from %s where %s = '%s') as result",
		table, param, value)

	var result bool

	err := db.db.QueryRow(str).Scan(&result)
	if err != nil {
		return false, util.ErrDbTrace(err, str, util.Trace())
	}

	return result, nil
}

func (db *commonRepo) InitTables() error {
	str := `create table if not exists accounts(id serial primary key, login text unique, password text not null,
		id_person int not null, active boolean not null)`
	_, err := db.db.DB.Exec(str)
	if err != nil {
		return util.ErrDbTrace(err, str, util.Trace())
	}

	str = `create table if not exists persons(id serial primary key, name text default '', last_name text default '', 
	middle_name text default '', age int default 0, date timestamp default '10-19-2023 08:35:34.000', 
	sex boolean default false, role text not null)`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return util.ErrDbTrace(err, str, util.Trace())
	}

	str = `create table if not exists classrooms(id serial primary key, number text not null, max_persons int not null,
		description text default '', active boolean default false)`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return util.ErrDbTrace(err, str, util.Trace())
	}

	str = `create table if not exists lessons(id serial primary key, name text default '', 
	date timestamp default '10-19-2023 08:35:34.000', id_teacher int not null, description text default '')`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return util.ErrDbTrace(err, str, util.Trace())
	}

	return nil
}
