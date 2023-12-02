package lesson

type ILessonService interface{

}

type lessonService struct{
	r ILessonRepo
}

func NewLessonService(r ILessonRepo)*lessonService{
	return &lessonService{r: r}
}