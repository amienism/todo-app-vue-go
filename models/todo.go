package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID			uint	`gorm:"primaryKey"`
	Activity	string	
	Status		bool	
}

// Get Todos
func GetTodos(db *gorm.DB, Todo *[]Todo) (err error) {
	err = db.Find(Todo).Error
	if err != nil {
		return err
	}
	return nil
}

//Get Todos by id
func GetTodo(db *gorm.DB, Todo *Todo, id int) (err error) {
	err = db.Where("id = ?", id).First(Todo).Error
	if err != nil {
		return err
	}
	return nil
}

func PostTodo(db *gorm.DB, Todo *Todo) (err error) {
	err = db.Create(Todo).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(db *gorm.DB, Todo *Todo, id int) (err error) {
	db.Save(Todo)
	return nil
}

func DeleteTodo(db *gorm.DB, Todo *Todo, id int) (err error) {
	db.Where("id = ?", id).Delete(Todo)
	return nil
}
