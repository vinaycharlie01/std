package model

// Student Details:
// ID (unique identifier, integer)
// Name (string)
// Phone (string)
// Email (unique, string)

type Student struct {
	ID     int      `json:"id" gorm:"primaryKey,autoincrement"`
	Name   string   `json:"name"`
	Phone  string   `json:"phone"`
	Email  string   `json:"email" gorm:"unique"`
	Course []Course `json:"courses"`
}

/*
Course Details:
ID (unique identifier, integer)
Name (string)
Student_ID (foreign key referencing the Student ID)

*/

type Course struct {
	ID         int    `json:"id" gorm:"primaryKey,autoincrement"`
	Name       string `json:"name"`
	Student_ID int    `json:"student_id"`
}

func NEWStudnetTable() (*Student, *Course) {
	return &Student{}, &Course{}
}
