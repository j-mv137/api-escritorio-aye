package main

import "database/sql"

type Storage interface {
	AddRevision(*Revision) error
	GetRevisionByState(state string) ([]Revision, error)
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgressDB() (*PostgresDB, error) {
	conStr := "postgresql://jacobo:jacobon@localhost:5432/aye_bd"
	postgresStore, err := sql.Open("postgres", conStr)

	if err != nil {
		return nil, err
	}

	if err = postgresStore.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		store: postgresStore,
	}, nil
}

func (s *PostgresDB) AddRevision(r *Revision) error {
	q := `INSERT INTO revisiones 
		(tipo, fecha, nombre, descripcion, domicilio, telefono) 
		values ($1, $2, $3, $4, $5, $6);`

	_, err := s.db.Exec(
		q,
		r.Tipo,
		r.Fecha,
		r.Nombre,
		r.Descripcion,
		r.Domicilio,
		r.Telefono,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresDB) GetRevisionByState(state string) ([]*Revision, error) {
	q := `SELECT * FROM revisiones WHERE estado=$1;`

	rows, err := s.db.Query(q, state)

	revisions := []*Revision{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		revision, err := scanRevision(rows)

		if err != nil {
			return nil, err
		}

		revisions = append(revisions, revision)
	}

	return revisions, nil
}

func scanRevision(rows *sql.Rows) (*Revision, error) {
	rev := &Revision{}

	err := rows.Scan(
		&rev.Id,
		&rev.Tipo,
		&rev.Fecha,
		&rev.Nombre,
		&rev.Domicilio,
		&rev.Telefono,
		&rev.Estado,
		&rev.Descripcion,
	)

	if err != nil {
		return nil, err
	}

	return rev, nil

}
