package user

type Account struct {
	Id        int   `json:"id" db:"id"`
	Login     string `json:"login" db:"login"`
	Pass      string `json:"password" db:"password"`
	Id_person int   `json:"id_person" db:"id_person"`
	Active    bool   `json:"active" db:"active"`
}

type Person struct {
	Id          int   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Last_name   string `json:"last_name" db:"last_name"`
	Middle_name string `json:"middle_name" db:"middle_name"`
	Age         int   `json:"age" db:"age"`
	Date        string `json:"date" db:"date"`
	Sex         bool   `json:"sex" db:"sex"`
	Role        string `json:"role" db:"role"`
}