package lesson

import "ed"

func (db *lessonRepo) MigrateLesson() error {
	str := `create table if not exists lessons(id serial primary key, name text default '', 
	date timestamp default '10-19-2023 08:35:34.000', id_teacher int not null, description text default '')`
	_, err := db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}
