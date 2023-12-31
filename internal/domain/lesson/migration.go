package lesson

import "ed"

func (db *lessonRepo) MigrateLesson() error {
	str := `create table if not exists lessons(id serial primary key, name text default '', 
	date timestamp default '10-19-2023 08:35:34.000', id_teacher int not null, 
	id_subject int not null, description text default '')`
	_, err := db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	str = `create table if not exists subjects(id serial primary key, name text default '', min_hours int default 3, 
	max_hours int default 5, examination boolean default false, description text default '')`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	str = `create table if not exists lesson_visits(id serial primary key, id_account int not null, id_subject int not null, 
		id_lesson int not null, hours int default 0, eval int default 0, behavior int default 0)`
	_, err = db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}
