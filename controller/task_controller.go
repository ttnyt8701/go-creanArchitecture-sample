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
	// リクエストパラメーラからidを取得し整数に変換
    vars := mux.Vars(r)
    idStr := vars["id"]

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid id", http.StatusBadRequest)
        return
    }

	// 取得
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
	task, err := t.u.GetAllTasks()
	if err != nil{
		fmt.Fprintf(w, "Invalid id: %s", err)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request){
	// ボディから情報取得してセット
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	// 作成
	id, err := t.u.CreateTask(task.Title)
	if err != nil{
		message := fmt.Sprintf("投稿失敗:%s",err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("投稿成功 {id:%v , Title:%s}",id,task.Title)
	json.NewEncoder(w).Encode(message)
	
}

func (t *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request){
	// リクエストパラメーラからidを取得し整数に変換
	vars := mux.Vars(r)
    idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid id", http.StatusBadRequest)
        return
    }

	// ボディから情報取得してセット
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := t.u.UpdateTask(id,task.Title);err != nil{
		http.Error(w, "Invalid request update", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "更新成功")
}

func(t *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request){
		// リクエストパラメーラからidを取得し整数に変換
		vars := mux.Vars(r)
		idStr := vars["id"]
	
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		if err := t.u.DeleteTask(id);err != nil{
			http.Error(w, "Invalid request delete", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "削除成功")
}