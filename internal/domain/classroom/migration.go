package classroom

import "ed"

func (db *classroomRepo) MigrateClassroom() error {
	str := `create table if not exists classrooms(id serial primary key, number text not null, max_persons int default 0,
		description text default '', active boolean default false)`
	_, err := db.db.DB.Exec(str)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}
