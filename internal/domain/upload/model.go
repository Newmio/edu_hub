package upload

type FileSettings struct {
	Status string `json:"status"`
	Id_account int    `json:"id_account"`
	File_type  string `json:"file_type"`
	File_name string `json:"file_name"`
}

type FileHistory struct {
	Id         int    `db:"id"`
	Id_account int    `db:"id_account"`
	Directory  string `db:"directory"`
	File       string `db:"file"`
	Date       string `db:"date"`
	Size       int64  `db:"byte_size"`
}

const CHANK_SIZE = 8 * 1024