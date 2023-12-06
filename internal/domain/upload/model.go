package upload

type FileData struct{
	Id_account uint `json:"id_account"`
	Data string `json:"data"`
	File_name string `json:"file_name"`
	File_type string `json:"file_type"`
}

type FileHistory struct{
	Id uint `db:"id"`
	Id_account uint `db:"id_account"`
	Directory string `db:"directory"`
	File string `db:"file"`
	Date string `db:"date"`
	Size float32 `db:"size"`
}