package main

import "database/sql"

type Storage interface {
	AddOrder(*Order) error
	GetOrderByState(state string) ([]*Order, error)
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
		db: postgresStore,
	}, nil
}

func (s *PostgresDB) AddOrder(r *Order) error {
	q := `INSERT INTO ordenes 
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

func (s *PostgresDB) GetOrderByState(state string) ([]*Order, error) {
	q := `SELECT * FROM revisiones WHERE estado=$1;`

	rows, err := s.db.Query(q, state)

	orders := []*Order{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order, err := scanOrder(rows)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func scanOrder(rows *sql.Rows) (*Order, error) {
	rev := &Order{}

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
