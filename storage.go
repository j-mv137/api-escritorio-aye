package main

import "database/sql"

type Storage interface {
	AddRevision(*Revision) error
	GetRevisionByID(int) error
	GetMultRevision([]int) error
}

type PostgresDB struct {
	store *sql.DB
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

	_, err := s.store.Exec(
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
