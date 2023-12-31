package user

import (
	"ed"
)

func (db *userRepo) MigrateUser() error {
	str := `create table if not exists accounts(id serial primary key, login text unique, password text not null,
		id_person int not null, role text default 'user', active boolean not null)`
	_, err := db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	str = `create table if not exists persons(id serial primary key, name text default '', last_name text default '', 
	middle_name text default '', age int default 0, date timestamp default '10-19-2023 08:35:34.000', 
	sex boolean default false, role text not null)`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}
