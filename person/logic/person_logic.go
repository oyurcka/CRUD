package logic

import (
	"context"
	"errors"
	"time"

	"github.com/oyurcka/CRUD/model/app"
	"github.com/oyurcka/CRUD/model/person"

	"github.com/sirupsen/logrus"
)

type personLogic struct {
	personRepo     person.Repository
	contextTimeout time.Duration
}

func NewPersonLogic(person person.Repository, timeout time.Duration) person.Logic {
	return &personLogic{
		personRepo:     person,
		contextTimeout: timeout,
	}
}

func (p *personLogic) Get(c context.Context) ([]*app.Person, error) {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	listPerson, err := person.personRepo.Get(ctx)
	if err != nil {
		return nil, "", err
	}

	return listPerson, nil
}

func (p *personLogic) GetByID(c context.Context, id int64) (*app.Person, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	res, err := person.personRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *personLogic) Store(c context.Context, per *app.Person) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	existedPerson, _ := p.GetByID(ctx, per.ID)
	if existedPerson != nil {
		logrus.Error("Person already exist")
		return errors.New("Person already exist")
	}

	err := p.personRepo.Store(ctx, per)
	if err != nil {
		return err
	}

	return nil
}

func (p *personLogic) Update(c context.Context, per app.Person) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.personRepo.Update(ctx, per)
}

func (p *personLogic) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	existedPerson, err := p.personRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if existedPerson == nil {
		logrus.Error("Person does not exist")
		return errors.New("Person does not exist")
	}

	return p.personRepo.Delete(ctx, id)
}
