package repository

import (
	"context"
	"database/sql"
	"user-service/service"

	_ "github.com/lib/pq"
)

type DB struct {
	conn  *sql.DB
}

func NewDB() *DB {

	conn, err := sql.Open("postgres", "user=elbek2001 password=1234 dbname=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	return &DB{
		conn:  conn,
	}
}

func (d *DB) CreateUser(ctx context.Context, name string) (int, error) {

	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	id := 0
	err := d.conn.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (d *DB) UpdateUser(ctx context.Context, user service.User) error {
	query := `UPDATE users SET name = $2 WHERE id = $1`

	_, err := d.conn.Exec(query, user.ID, user.Name)
	if err != nil {
		return err
	}
   
	return nil
}

func (d *DB) DeleteUser(ctx context.Context, user service.User) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := d.conn.Exec(query, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) ListUsers(ctx context.Context) ([]service.User, error) {
	var (
		id int
		name string
		query = `SELECT * FROM users`
	)
	users := make([]service.User, 0)
	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		if err := rows.Scan(&id, &name); err != nil{
			return nil, err
		}

		users = append(users, service.User{
			ID: id,
			Name: name,
		})
	}

	return users, nil
}

func (d *DB) GetUserBYID(ctx context.Context, id int) (service.User, error) {
	return service.User{}, nil
}