package logger

import "ed"

func (db *loggerRepo) MigrateLogger()error{
	str := `create table if not exists logs(id serial primary key, url text default '',
	body_req text default '', headers_req text default '', status int default 0,
	body_resp text default '', headers_resp text default '', method text default '', 
	date_start timestamp default '19-10-2023 08:08:25.000', date_stop timestamp default '19-10-2023 08:08:25.000',
	milliseconds int default 0, ip text default '')`

	_, err := db.db.Exec(str)
	if err != nil{
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	str = `create table if not exists errors(id serial primary key, error text default '',
	date timestamp default '19-10-2023 08:08:25.000')`

	if err != nil{
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}