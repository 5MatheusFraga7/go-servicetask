package repository

import (
	"database/sql"
	"fmt"
	"task-manager/internal/db"
	"task-manager/internal/db/adapters"
	"task-manager/internal/models"
)

type TaskRepository struct {
}

func (tr *TaskRepository) CreateTask(t models.Task) error {

	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	defer db.CloseConnectionToDatabase(postgresAdapter)

	_, err = postgresAdapter.Exec(getCreateTaskSql(),
		t.Name,
		t.Description,
		t.UserName,
		t.Deadline,
	)

	if err != nil {
		return fmt.Errorf("falha ao inserir tarefa no banco de dados: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return nil
}
func (tr TaskRepository) ReadTask(taskId int) []models.Task {
	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return nil
	}

	rows, err := postgresAdapter.Query(getReadTaskSql(), taskId)

	if err != nil {
		fmt.Println("Erro ao executar consulta no banco:", err)
		return nil
	}

	defer rows.Close()
	defer db.CloseConnectionToDatabase(postgresAdapter)

	return mapRowsToTasks(rows)
}

func (tr TaskRepository) UpdateTask(t models.Task) error {

	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return err
	}

	_, err = postgresAdapter.Exec(getUpdateTaskSql(),
		t.Id,
		t.Name,
		t.Description,
		t.UserName,
		t.Deadline,
		t.Done,
		t.Duration,
		t.FinishedAT,
	)

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return err
	}

	return nil
}
func (tr TaskRepository) DeleteTask(taskId int) error {
	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return err
	}

	_, err = postgresAdapter.Exec(getDeleteTaskSql(), taskId)

	return nil
}

func getCreateTaskSql() string {
	query := `
			INSERT INTO tasks (name, description, username, deadline, done, created_at)
			VALUES ($1, $2, $3, $4, false, now())
	`
	return query
}

func getReadTaskSql() string {
	query := `SELECT * FROM tasks WHERE id = $1`
	return query
}

func getUpdateTaskSql() string {

	query := `
			UPDATE tasks 
			SET name = $2, description = $3, username = $4, deadline = $5, done = $6, duration = $7, finished_at = $8
			WHERE id = $1
	`
	return query
}

func getDeleteTaskSql() string {
	query := `DELETE FROM tasks WHERE id = $1`
	return query
}

func mapRowsToTasks(rows *sql.Rows) []models.Task {
	var tasks []models.Task
	var durationNil sql.NullString
	var dateNil sql.NullTime

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &durationNil, &task.UserName, &task.Deadline, &task.Done, &dateNil, &task.CreatedAt)

		if err != nil {
			fmt.Println("Erro ao escanear linha:", err)
		} else {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
