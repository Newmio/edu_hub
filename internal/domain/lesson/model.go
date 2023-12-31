package lesson

type Subject struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Min_hours int    `json:"min_hours" db:"min_hours"`
	Max_hours int    `json:"max_hours" db:"max_hours"`
	Description string `json:"description" db:"description"`
}

type Lesson struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Date        string `json:"date" db:"date"`
	Id_teacher  int    `json:"id_teacher" db:"id_teacher"`
	Id_subject  int    `json:"id_subject" db:"id_subject"`
	Description string `json:"description" db:"description"`
}

type LessonVisit struct{
	Id int `json:"id" db:"id"`
	Id_acc int `json:"id_account" db:"id_account"`
	Id_subject int `json:"id_subject" db:"id_subject"`
	Id_lesson int `json:"id_lesson" db:"id_lesson"`
	Hours int `json:"hours" db:"hours"`
	Eval int `json:"evaluation" db:"eval"`
	Behavior int `json:"behavior" db:"behavior"`
}