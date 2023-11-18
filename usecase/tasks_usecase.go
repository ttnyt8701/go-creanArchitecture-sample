/*
ユースケース
ビジネスロジックを実装する層
データの処理、変換、制御...
コントローラーとリポジトリーの橋渡し
コントローラーからの入力を受け取り、レポジトリー層を使い、データの保存や取得を行いコントローラーに返す。
*/

package usecase

import (
	"fmt"
	"go-sample/model"
	"go-sample/repository"
)

type TaskUsecase interface {
	CreateTask(title string) error
	GetTasks(id int)(*model.Task, error)
	GetAllTasks()([]model.Task, error)
	UpdateTask(id int,title string) error
	DeleteTask(id int) error
}

/*
usecaseはレポジトリーに依存している。
レポジトリーを使ってDB操作をしてコントローラーに渡す。
*/
type taskUsecase struct{
	r repository.TaskRepository

}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r: r}
}

func (u *taskUsecase) CreateTask(title string) error{
	task := model.Task{Title: title}
	err := task.Validate()
	if err != nil{
		return err
	}
	
	id, err := u.r.Create(&task)
	fmt.Println(id)
	return err
}

func (u *taskUsecase) GetTasks(id int) (*model.Task, error){
	task, err := u.r.Read(id)
	if err != nil{
		return nil, err
	}

	return task, nil
}

func (u *taskUsecase) GetAllTasks() ([]model.Task, error){
	task, err := u.r.Reads()
	if err != nil{
		return nil, err
	}

	return task, nil
}

func (u *taskUsecase) UpdateTask(id int,title string) error{
	task := model.Task{ID: id,Title: title}
	err := u.r.Update(&task)
	if err != nil{
		return err
	}
	return nil
}

func (u *taskUsecase) DeleteTask(id int) error{
	err := u.r.Delete(id)
	if err != nil{
		return err
	}

	return nil
}