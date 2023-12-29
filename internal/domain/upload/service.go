package upload

import (
	"crypto/sha256"
	"ed"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type IUploadService interface {
	CreateFile() (*os.File, error)
	AppendChank(file *os.File, chank []byte) error
	UpdateFile(file *os.File, data *FileSettings) error
	ReadLastChanks(data *FileSettings) ([][]byte, error)
}

type uploadService struct {
	r iUploadRepo
}

func NewUploadService(r iUploadRepo) *uploadService {
	err := r.MigrateFiles()
	if err != nil {
		return nil
	}

	return &uploadService{r: r}
}

func checkExistsFile(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDirectory(data *FileSettings) string {
	hash := sha256.Sum256([]byte(strconv.Itoa(int(data.Id_account))))

	return fmt.Sprintf("media/%s/%s/%s/",
		hex.EncodeToString(hash[:]), time.Now().Format("02-01-2006"), data.File_type)
}

func (s *uploadService) RealAllFile(data *FileSettings) ([][]byte, error) {
	//directory := createDirectory(data)

	//if checkExistsFile(directory) {
		// var chanks [][]byte

		// content, err := os.ReadFile(directory)
		// if err != nil{
		// 	return nil, ed.ErrTrace(err, ed.Trace())
		// }

		// for i := int64(0); i <= 4; i++ {
		// 	offset := (len(content) + CHANK_SIZE - 1) / CHANK_SIZE

		// 	if offset < 0{
		// 		offset = 0
		// 	}

		// 	_, err = file.Seek(offset, 0)
		// 	if err != nil {
		// 		return nil, ed.ErrTrace(err, ed.Trace())
		// 	}

		// 	data := make([]byte, CHANK_SIZE)

		// 	_, err = file.Read(data)
		// 	if err != nil {
		// 		return nil, ed.ErrTrace(err, ed.Trace())
		// 	}

		// 	chanks[i] = append([]byte(nil), data...)
		// }
	//}

	return nil, nil
}

// Возвращает последние 5 чанков по 8Кб
func (s *uploadService) ReadLastChanks(data *FileSettings) ([][]byte, error) {
	directory := createDirectory(data)

	if checkExistsFile(directory) {

		file, err := os.Open(directory)
		if err != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}

		inf, err := file.Stat()
		if err != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}

		chanks := make([][]byte, 5)

		for i := int64(0); i <= 4; i++ {
			offset := inf.Size() - (CHANK_SIZE*i + 1)

			if offset < 0{
				offset = 0
			}

			_, err = file.Seek(offset, 0)
			if err != nil {
				return nil, ed.ErrTrace(err, ed.Trace())
			}

			data := make([]byte, CHANK_SIZE)

			_, err = file.Read(data)
			if err != nil {
				return nil, ed.ErrTrace(err, ed.Trace())
			}

			chanks[i] = append([]byte(nil), data...)
		}

		return chanks, nil
	}

	return nil, nil
}

// Переносит файл в нужную папку, создает запись про файл в базе
func (s *uploadService) UpdateFile(file *os.File, data *FileSettings) error {
	name := file.Name()

	inf, err := file.Stat()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	err = file.Close()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	directory := createDirectory(data)

	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	parts := strings.Split(name, ".")
	if len(parts) < 2 {
		return ed.ErrTrace(errors.New("bad file name"), ed.Trace())
	}

	err = os.Rename(name, directory+parts[0]+"."+data.File_type)
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	return s.r.CreateFile(FileHistory{
		Id_account: data.Id_account,
		Directory:  directory,
		File:       name + "." + data.File_type,
		Date:       time.Now().Format("02-01-2006 15:04:05.000"),
		Size:       inf.Size()})
}

// Добавляет чанк в файл
func (s *uploadService) AppendChank(file *os.File, chank []byte) error {
	_, err := file.Write(chank)
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	return nil
}

func (s *uploadService) CreateFile() (*os.File, error) {
	for {
		name := fmt.Sprint(time.Now().UnixNano())

		if checkExistsFile(name) {
			continue
		}

		out, err := os.Create(name + ".txt")
		if err != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}

		return out, nil
	}
}
