package task

import (
	"context"

	"github.com/lgu-elo/task/internal/task/model"
	"github.com/pkg/errors"
)

type Repository interface {
	GetById(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
	Create(task *model.Task) error
}

func (db *storage) GetById(id int) (*model.Task, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var task model.Task

	err := db.QueryRow(
		context.Background(),
		"SELECT * FROM tasks WHERE id=$1",
		id,
	).Scan(&task.ID, &task.Description, &task.Project_id, &task.User_id, &task.Status)

	db.log.Println(task.Description)
	if err != nil {
		return nil, errors.Wrap(err, "task not found")
	}

	return &task, err
}

func (db *storage) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, errors.Wrap(err, "can't get all tasks")
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Project_id, &task.User_id, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *storage) Update(task *model.Task) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"UPDATE tasks set description=$1, project_id=$2, user_id=$3, status=$4 WHERE id=$5",
		task.Description, task.Project_id, task.User_id, task.Status, task.ID,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't update task")
	}

	if err := rows.Err(); err != nil {
		return errors.Wrap(err, "can't update task")
	}

	return nil
}

func (db *storage) Create(task *model.Task) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"INSERT INTO tasks(description, project_id, user_id, status) VALUES($1,$2,$3,$4) RETURNING id",
		task.Description,
		task.Project_id,
		task.User_id,
		task.Status,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't create task")
	}

	if err := rows.Err(); err != nil {
		return errors.Wrap(err, "can't create task")
	}

	return nil
}

func (db *storage) Delete(id int) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	err := db.QueryRow(
		context.Background(),
		"DELETE FROM tasks WHERE id=$1",
		id,
	).Scan()

	if err != nil {
		return errors.Wrap(err, "can't delete task")
	}

	return nil
}
