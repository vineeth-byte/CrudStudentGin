package dto

type Student struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Email       string `json:"email,omitempty" bson:"email,omitempty"`
	Mobile      string `json:"mobile,omitempty" bson:"mobile,omitempty"`
	TeacherName string `json:"teacherName,omitempty" bson:"teacherName,omitempty"`
	Department  string `json:"department,omitempty" bson:"department,omitempty"`
	Marks       Marks  `json:"marks,omitempty" bson:"marks,omitempty"`
}

type Marks struct {
	Maths   int `json:"maths,omitempty" bson:"maths,omitempty"`
	Science int `json:"science,omitempty" bson:"science,omitempty"`
	Social  int `json:"social,omitempty" bson:"social,omitempty"`
	English int `json:"english,omitempty" bson:"english,omitempty"`
}
