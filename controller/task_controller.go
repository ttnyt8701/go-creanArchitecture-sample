package controller

import (
	"encoding/json"
	"fmt"
	"go-sample/model"
	"go-sample/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskController struct {
	u usecase.TaskUsecase
}

func NewTaskController(u usecase.TaskUsecase) *TaskController {
	return &TaskController{u: u}
}

func (t *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"] // パスパラメータからIDを取得

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid id", http.StatusBadRequest)
        return
    }

    task, err := t.u.GetTasks(id)
    if err != nil {
        http.Error(w, "Error fetching task", http.StatusInternalServerError)
        return 
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}

func (t *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request){
	// DBから全てのタスクを取得
	task, err := t.u.GetAllTasks()
	if err != nil{
		fmt.Fprintf(w, "Invalid id: %s", err)
		return 
	}else{

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request){
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	err := t.u.CreateTask(task.Title)
	if err != nil{
		message := fmt.Sprintf("投稿失敗:%s",err)
		http.Error(w, message, http.StatusBadRequest)
	}else{
		fmt.Fprintf(w, "投稿成功:")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task.Title)
	}
}

// func (t *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request){
// 	// パスパラメータからIDを取得
// 	vars := mux.Vars(r)
//     idStr := vars["id"]

// 	id, err := strconv.Atoi(idStr)
//     if err != nil {
//         http.Error(w, "Invalid id", http.StatusBadRequest)
//         return
//     }

// 	// 更新情報取得

// 	err := t.u.UpdateTask(id,"title")

// }