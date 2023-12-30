package app

import (
	"ed"
	"ed/internal/app/postgres"
	"ed/internal/domain/classroom"
	"ed/internal/domain/lesson"
	"ed/internal/domain/logger"
	"ed/internal/domain/request"
	"ed/internal/domain/upload"
	"ed/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func InitEngine() (*gin.Engine, error) {
	database, err := postgres.OpenDb()
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}

	loggerRepo := logger.NewLoggerRepo(database)
	loggerService := logger.NewLoggerService(loggerRepo)

	userRepo := user.NewUserRepo(database)
	userService := user.NewUserService(userRepo)
	userHand := user.NewHandler(userService, loggerService)

	classroomRepo := classroom.NewClassroomRepo(database)
	classroomService := classroom.NewClassroomService(classroomRepo)
	classroomHand := classroom.NewHandler(classroomService)

	lessonRepo := lesson.NewLessonRepo(database)
	lessonService := lesson.NewLessonService(lessonRepo)
	lessonHand := lesson.NewHandler(lessonService)

	reqService := request.NewRequestService(loggerService)

	uploadRepo := upload.NewUploadRepo(database)
	uploadService := upload.NewUploadService(uploadRepo)
	uploadHand := upload.NewHandler(uploadService, reqService, loggerService)

	r := gin.Default()

	r = userHand.InitUserRoutes(r)
	r = classroomHand.InitClassroomRoutes(r)
	r = lessonHand.InitLessonRoutes(r)
	r = uploadHand.InitUploadRoutes(r)

	return r, nil
}
