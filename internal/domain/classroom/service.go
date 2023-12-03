package classroom

type IClassroomService interface{

}

type classroomService struct{
	r iClassroomRepo
}

func NewClassroomService(r iClassroomRepo)*classroomService{
	err := r.MigrateClassroom()
	if err != nil{
		return nil
	}
	
	return &classroomService{r: r}
}