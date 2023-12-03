package lesson

type ILessonService interface{

}

type lessonService struct{
	r iLessonRepo
}

func NewLessonService(r iLessonRepo)*lessonService{
	err := r.MigrateLesson()
	if err != nil{
		return nil
	}
	
	return &lessonService{r: r}
}