package model

type Classroom struct {
	Id          uint   `json:"id" db:"id"`
	Number      int    `json:"number" db:"number"`
	Max_person  uint   `json:"max_person" db:"max_person"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}

type Lesson struct {
	Id          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Date        string `json:"date" db:"date"`
	Id_teacher  uint   `json:"id_teacher" db:"id_teacher"`
	Description string `json:"description" db:"description"`
}