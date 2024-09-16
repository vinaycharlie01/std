package db

import (
	"myapp/studentservice/model"

	"gorm.io/gorm"
)

type Queries struct {
	DB *gorm.DB
}

func NewDB(in *gorm.DB) *Queries {
	return &Queries{in}
}

type I interface {
	CreateStudeNt(in *model.Student) error
	Update(in *model.Student) error
	Remove(in *model.Student) error
	Read(id int) ([]model.Student, error)
}

func (q *Queries) CreateStudeNt(in *model.Student) error {
	return q.DB.Create(in).Error
}

func (q *Queries) Update(in *model.Student, id int) error {
	err := q.DB.Model(&model.Student{}).Where("id", id).Update("name", in.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func (q *Queries) Remove(in int) error {
	err := q.DB.Where("student_id = ?", in).Delete(&model.Course{}).Error
	if err != nil {
		return err
	}

	err = q.DB.Where("id = ?", in).Delete(&model.Student{}).Error
	if err != nil {
		return err
	}
	return nil

}

func (q *Queries) Read(id int) ([]model.Student, error) {
	var res []model.Student
	err := q.DB.Model(&model.Student{}).Preload("Student").Where("id=?", id).Scan(&res).Error
	// err := q.DB.Select("*").Joins("JOIN students ON students.id=courses.student_id").Where("students.id=?", id).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (q *Queries) ReadAll() ([]model.Student, error) {
	var res []model.Student
	err := q.DB.Model(&model.Student{}).Preload("Student").Scan(&res).Error
	// err := q.DB.Select("*").Joins("JOIN students ON students.id=courses.student_id").Where("students.id=?", id).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
