package lesson

type Lesson struct {
	Id          int   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Date        string `json:"date" db:"date"`
	Id_teacher  int   `json:"id_teacher" db:"id_teacher"`
	Description string `json:"description" db:"description"`
}