package mysql

import (
	"SpendSmartAPI/internal/domain"
	"context"
	"database/sql"
)

// This type implements user_repository interface
type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{
		db: db,
	}
}

func (r *UserMySQLRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO user (name, email, password, phone) VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Phone)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}

func (r *UserMySQLRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	query := `
		SELECT id, name, email, password, phone FROM user 
	`

	// Mantém cursor aberto no banco para usar com Next, então é necessário fechar rows
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var row domain.User

		// Go pega cada coluna da query e coloca nas variáveis
		err := rows.Scan(&row.ID, &row.Name, &row.Email, &row.Password, &row.Phone)

		if err != nil {
			return nil, err
		}

		users = append(users, row)
	}

	return users, nil
}

func (r *UserMySQLRepository) FindById(ctx context.Context, id int) (*domain.User, error) {
	query := `
		SELECT id, name, email, password, phone FROM user
		WHERE id = ?
	`

	var u domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Phone)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}
