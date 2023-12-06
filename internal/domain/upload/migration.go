package upload

import "ed"

func (db *uploadRepo) MigrateFiles()error{
	str := `create table if not exists files(id serial primary key, id_account int default 0, directory text default ''),
	file text default '', size float default 0, date timestamp default '19-10-2023 8:35:00.000'`

	_, err := db.db.Exec(str)
	if err != nil{
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}