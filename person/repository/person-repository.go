package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/oyurcka/CRUD/model/app"
	"github.com/oyurcka/CRUD/model/person"

	"github.com/sirupsen/logrus"
)

type postgresqllPersonRepository struct {
	Conn *sql.DB
}

func NewPostgresqlPersonRepository(Conn *sql.DB) person.Repository {
	return &postgresqllPersonRepository{Conn}
}

func (p *postgresqllPersonRepository) get(ctx context.Context, query string, args ...interface{}) ([]*app.Person, error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*app.Person, 0)
	for rows.Next() {
		t := new(app.Person)
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Phone,
			&t.FirstName,
			&t.LastName,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *postgresqllPersonRepository) Get(ctx context.Context) ([]*app.Person, error) {
	query := `SELECT id, email, phone, firstname, lastname
  				FROM person ORDER BY id `

	res, err := p.get(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return res, err
}

func (p *postgresqllPersonRepository) GetByID(ctx context.Context, id int64) (res *app.Person, err error) {
	query := `SELECT id, email, phone, firstname, lastname
				FROM person WHERE ID = ?`

	list, err := p.get(ctx, query, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		logrus.Error(errors.New("Person not found"))
		return nil, errors.New("Person not found")
	}

	return
}

func (p *postgresqllPersonRepository) Store(ctx context.Context, per *app.Person) error {
	query := `INSERT person SET email=?, phone=?, firstname=?, lastname=?`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return err
	}

	res, err := stmt.ExecContext(ctx, per.Email, per.Phone, per.FirstName, per.LastName)
	if err != nil {
		logrus.Error(err)
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return err
	}

	per.ID = lastID
	return nil
}

func (p *postgresqllPersonRepository) Update(ctx context.Context, per *app.Person) error {
	query := `UPDATE person SET email=?, phone=?, firstname=?, lastname=? WHERE id = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	res, err := stmt.ExecContext(ctx, per.Email, per.Phone, per.FirstName, per.LastName, per.ID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowsAffected != 1 {
		logrus.Error(errors.New("More than one row affected"))
		return err
	}

	return nil
}

func (p *postgresqllPersonRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM person WHERE id = ?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}

	if rowsAffected != 1 {
		logrus.Error(errors.New("More than one row affected"))
		return err
	}

	return nil
}
