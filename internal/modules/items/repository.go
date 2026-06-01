package items

import (
	"database/sql"
	"errors"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindAll() ([]Item, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, created_at, updated_at
		FROM items
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) FindByID(id int) (*Item, error) {
	var item Item

	err := r.DB.QueryRow(`
		SELECT id, name, description, created_at, updated_at
		FROM items
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("item não encontrado")
	}

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *Repository) Create(item Item) (*Item, error) {
	var createdItem Item

	err := r.DB.QueryRow(`
		INSERT INTO items (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at, updated_at
	`, item.Name, item.Description).Scan(
		&createdItem.ID,
		&createdItem.Name,
		&createdItem.Description,
		&createdItem.CreatedAt,
		&createdItem.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdItem, nil
}

func (r *Repository) Update(id int, item Item) (*Item, error) {
	var updatedItem Item

	err := r.DB.QueryRow(`
		UPDATE items
		SET name = $1,
			description = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`, item.Name, item.Description, id).Scan(
		&updatedItem.ID,
		&updatedItem.Name,
		&updatedItem.Description,
		&updatedItem.CreatedAt,
		&updatedItem.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("item não encontrado")
	}

	if err != nil {
		return nil, err
	}

	return &updatedItem, nil
}

func (r *Repository) Delete(id int) error {
	result, err := r.DB.Exec(`
		DELETE FROM items
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("item não encontrado")
	}

	return nil
}
