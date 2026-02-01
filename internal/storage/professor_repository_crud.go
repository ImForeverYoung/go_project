package storage

import (
	"HW_5/internal/model"
	"context"
	"fmt"
	"os"
)

// new professor
func (s *Storage) CreateProfessor(ctx context.Context, p model.ProfessorRequest) (int, error) {
	query := `INSERT INTO professors (name, department, title, email) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := s.conn.QueryRow(ctx, query, p.Name, p.Department, p.Title, p.Email).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateProfessor failed: %v\n", err)
		return 0, err
	}
	return id, nil
}

// by id
func (s *Storage) GetProfessor(ctx context.Context, id int) (model.Professor, error) {
	query := `SELECT id, name, department, title, email FROM professors WHERE id = $1`
	var p model.Professor
	err := s.conn.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.Department, &p.Title, &p.Email)
	if err != nil {
		return model.Professor{}, err
	}
	return p, nil
}

// get all
func (s *Storage) ListProfessors(ctx context.Context) ([]model.Professor, error) {
	query := `SELECT id, name, department, title, email FROM professors ORDER BY id`
	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var professors []model.Professor
	for rows.Next() {
		var p model.Professor
		if err := rows.Scan(&p.ID, &p.Name, &p.Department, &p.Title, &p.Email); err != nil {
			return nil, err
		}
		professors = append(professors, p)
	}
	return professors, nil
}

// updates row
func (s *Storage) UpdateProfessor(ctx context.Context, id int, p model.ProfessorRequest) error {
	query := `UPDATE professors SET name=$1, department=$2, title=$3, email=$4 WHERE id=$5`
	cmd, err := s.conn.Exec(ctx, query, p.Name, p.Department, p.Title, p.Email, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("professor with id %d not found", id)
	}
	return nil
}

// delete by id
func (s *Storage) DeleteProfessor(ctx context.Context, id int) error {
	query := `DELETE FROM professors WHERE id=$1`
	cmd, err := s.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("professor with id %d not found", id)
	}
	return nil
}
