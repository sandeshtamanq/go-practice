package task

import (
	"fmt"
	"net/http"

	"github.com/sandeshtamanq/jwt/entity"
	"github.com/sandeshtamanq/jwt/service/auth"
	"github.com/sandeshtamanq/jwt/utils"
)

type Service struct{}

func TaskService() *Service {
	return &Service{}
}

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var t entity.Task
	taskRepository := TaskRepository()

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("empty body"))
	}

	utils.ParseJSON(r, &t)

	userId := auth.GetCurrentUserId(r.Context())

	if userId < 1 {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"message": "something went wrong"})
		return
	}

	t.UserID = userId

	err := taskRepository.CreateTask(&t)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "task created successfully"})
}

func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	taskRepository := TaskRepository()
	userId := auth.GetCurrentUserId(r.Context())
	tasks, err := taskRepository.GetTask(userId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, 200, tasks)
}
