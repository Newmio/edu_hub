package upload

import (
	"bufio"
	"crypto/sha256"
	"ed"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

type IUploadService interface {
	CreateFile(file FileHistory) error
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

func (s *uploadService) CreateFile(file FileHistory) error {
	file.Date = time.Now().Format("02-01-2006 15:04:05.000")
	return s.r.CreateFile(file)
}


func (s *uploadService) UploadFile(file FileData)error{
	hash := sha256.Sum256([]byte(strconv.Itoa(int(file.Id_account))))

	directory := fmt.Sprintf("%s/%s/%s/%s", 
	hex.EncodeToString(hash[:]), time.Now().Format("02-01-2006"), file.File_type, file.File_name + "." + file.File_type)

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	out, err := os.Create(directory)
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	_, err = writer.WriteString(file.Data)
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	err = writer.Flush()
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	return nil
}