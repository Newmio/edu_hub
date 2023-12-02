package classroom

type IClassroomService interface{

}

type classroomService struct{
	r IClassroomRepo
}

func NewClassroomService(r IClassroomRepo)*classroomService{
	return &classroomService{r: r}
}