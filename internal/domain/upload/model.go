package upload

type FileData struct {
	Id_account int   `json:"id_account"`
	Max_chank int `json:"max_chank"`
	Data       string `json:"data"`
	File_name  string `json:"file_name"`
	File_type  string `json:"file_type"`
	Last       bool   `json:"last"`
}

type FileHistory struct {
	Id         int    `db:"id"`
	Id_account int    `db:"id_account"`
	Directory  string  `db:"directory"`
	File       string  `db:"file"`
	Date       string  `db:"date"`
	Size       float32 `db:"size"`
}
