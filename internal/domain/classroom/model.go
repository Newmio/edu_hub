package classroom

type Classroom struct {
	Id          uint   `json:"id" db:"id"`
	Number      int    `json:"number" db:"number"`
	Max_person  uint   `json:"max_person" db:"max_person"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}
