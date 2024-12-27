package auth

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email, password from users WHERE email = $1 AND deleted_at IS null`,
		email,
	)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
