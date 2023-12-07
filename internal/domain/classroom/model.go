package classroom

type Classroom struct {
	Id          int   `json:"id" db:"id"`
	Number      int    `json:"number" db:"number"`
	Max_person  int   `json:"max_person" db:"max_person"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}
