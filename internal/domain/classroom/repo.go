package classroom

import "github.com/jmoiron/sqlx"

type IClassroomRepo interface{

}

type classroomRepo struct{
	db *sqlx.DB
}

func NewClassroomRepo(db *sqlx.DB)*classroomRepo{
	return &classroomRepo{db: db}
}