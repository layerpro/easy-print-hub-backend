package profile

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) GetProfile(id int) (*Profile, error) {
	row := r.db.QueryRow(
		`SELECT id, name, email FROM users WHERE id = $1`,
		id,
	)

	var profile Profile
	err := row.Scan(&profile.ID, &profile.Name, &profile.Email)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
