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

	t.UserID = uint(userId)

	err := taskRepository.CreateTask(&t)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "task created successfully"})
}

func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, 200, map[string]string{"message": "task created successfully"})
}
