package course

func NewCourse(id int, name string) *Course {
	return &Course{CourseId: int64(id), CourseName: name}
}

