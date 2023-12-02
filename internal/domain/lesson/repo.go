package lesson

import "github.com/jmoiron/sqlx"

type ILessonRepo interface{

}

type lessonRepo struct{
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB)*lessonRepo{
	return &lessonRepo{db: db}
}