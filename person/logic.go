package person

import (
	"context"

	"github.com/oyurcka/CRUD/model/app"
)

type Logic interface {
	Get(ctx context.Context) ([]*app.Person, error)
	GetByID(ctx context.Context, id int64) (*app.Person, error)
	Store(context.Context, *app.Person) error
	Update(ctx context.Context, per *app.Person) error
	Delete(ctx context.Context, id int64) error
}
