package lesson

type Lesson struct {
	Id          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Date        string `json:"date" db:"date"`
	Id_teacher  uint   `json:"id_teacher" db:"id_teacher"`
	Description string `json:"description" db:"description"`
}