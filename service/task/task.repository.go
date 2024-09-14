package task

import (
	"github.com/sandeshtamanq/jwt/database"
	"github.com/sandeshtamanq/jwt/entity"
)

type Repository struct{}

func TaskRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateTask(t *entity.Task) error {
	createdTask := database.DB.Create(&t)

	err := createdTask.Error

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTask(userId uint) (*[]entity.Task, error) {
	var tasks []entity.Task

	err := database.DB.Preload("User").Where("user_id = ?", userId).Find(&tasks).Error

	return &tasks, err

}
