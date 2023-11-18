package main

import (
	"database/sql"
	"fmt"
	"go-sample/controller"
	"go-sample/repository"
	"go-sample/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// ミドルウェア
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		log.Printf("Started %s %s",r.Method,r.URL.Path)

		next.ServeHTTP(w,r)

		log.Printf("completed %s in %v", r.URL.Path, time.Since(start))
	})
}

// DB
func initDB()(*sql.DB, error){
	db, err := sql.Open("sqlite3","./test.db")
	return db, err
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}


func main(){
	// DB
	db, err := initDB()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT)")
	if err != nil{
		log.Fatal(err)
	}

	// ルーター
	r := mux.NewRouter()

	// ルーターにミドルウェアを登録
	r.Use(loggingMiddleware)

	// エンドポイント登録
	r.HandleFunc("/",mainHandler)

	// 注入
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	// ハンドラー
	r.HandleFunc("/tasks", taskController.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskController.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}/update", taskController.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}/delete", taskController.DeleteTask).Methods("DELETE")

	// サーバー起動 - r: リクエストの処理 - log.Fatal: エラーが発生したときにログ出力してプログラム終了
	log.Fatal(http.ListenAndServe(":8199", r))
}