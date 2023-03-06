package models

import (
	"context"
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *DBModel) TodoGetAll() ([]*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, title, description, is_done from todos"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		var ctg Todo
		err := rows.Scan(
			&ctg.ID,
			&ctg.Title,
			&ctg.Description,
			&ctg.IsDone,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &ctg)
	}
	return todos, nil
}

// getOne
func (m *DBModel) GetTodo(id int) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, is_done from todos where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var ctg Todo
	err := row.Scan(
		&ctg.ID,
		&ctg.Title,
		&ctg.Description,
		&ctg.IsDone,
	)
	if err != nil {
		return nil, err
	}
	return &ctg, nil
}

// Create
func (m *DBModel) TodoCreate(todo Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into todos (title, description, is_done) values ($1, $2, $3)`

	_, err := m.DB.ExecContext(ctx, query,
		todo.Title,
		todo.Description,
		todo.IsDone,
	)

	if err != nil {
		return err
	}

	return nil
}

// UPDATE
func (m *DBModel) TodoUpdate(todo Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update todos set title = $1, description = $2, is_done = $4 where id = $3`

	_, err := m.DB.ExecContext(ctx, query,
		todo.Title,
		todo.Description,
		todo.ID,
		todo.IsDone,
	)

	if err != nil {
		return nil
	}

	return nil
}

// DELETE
func (m *DBModel) TodoDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// tx, _ := m.DB.Begin()
	// defer func() {
	// 	// panicが起きたらロールバック
	// 	if recover() != nil {
	// 		tx.Rollback()
	// 	}
	// }()

	query := `delete from todos where id = $1`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
