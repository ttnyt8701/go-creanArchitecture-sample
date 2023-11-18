package usecase

import (
	"go-sample/model"
	"go-sample/repository"
)

type TaskUsecase interface {
	CreateTask(title string) (int, error)
	GetTasks(id int)(*model.Task, error)
	GetAllTasks()([]model.Task, error)
	UpdateTask(id int,title string) error
	DeleteTask(id int) error
}

type taskUsecase struct{
	r repository.TaskRepository

}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r: r}
}

func (u *taskUsecase) CreateTask(title string) (int, error){
	task := model.Task{Title: title}
	err := task.Validate()
	if err != nil{
		return 0, err
	}
	
	id, err := u.r.Create(&task)
	if err != nil{
		return 0, err
	}

	return id, err
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