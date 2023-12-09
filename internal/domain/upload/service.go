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
	AppendChank(file *os.File, data *FileData) error
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

func (s *uploadService) updateFile(name string, size int64, data *FileData) error {
	hash := sha256.Sum256([]byte(strconv.Itoa(int(data.Id_account))))

	directory := fmt.Sprintf("media/%s/%s/%s/",
		hex.EncodeToString(hash[:]), time.Now().Format("02-01-2006"), data.File_type)

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	parts := strings.Split(name, ".")
	if len(parts) < 2 {
		return ed.ErrTrace(errors.New("bad file name"), ed.Trace())
	}

	fmt.Println("-------------------------")
	fmt.Println(directory+parts[0]+"."+data.File_type)

	err = os.Rename(name, directory+parts[0]+"."+data.File_type)
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	return s.r.CreateFile(FileHistory{
		Id_account: data.Id_account,
		Directory:  directory,
		File:       name + "." + data.File_type,
		Date:       time.Now().Format("02-01-2006 15:04:05.000"),
		Size:       size})
}

func (s *uploadService) AppendChank(file *os.File, data *FileData) error {
	_, err := file.Write([]byte(data.Data))
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if data != nil && data.Last {
		inf, err := file.Stat()
		if err != nil {
			return ed.ErrTrace(err, ed.Trace())
		}

		err = file.Close()
		if err != nil {
			return ed.ErrTrace(err, ed.Trace())
		}

		return s.updateFile(file.Name(), inf.Size(), data)
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
