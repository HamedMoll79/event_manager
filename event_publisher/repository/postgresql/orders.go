package postgresql

import (
	"context"
	"database/sql"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
)

type ordersRepository struct {
	conn *sql.DB
}

func NewOrdersRepository(getter postgresql.DataContextGetter) *ordersRepository {
	return &ordersRepository{
		conn: getter.GetDataContext(),
	}
}

func (d *ordersRepository) Save(ctx context.Context, order entity.Order) (entity.Order, error) {
	const op = "postgresql.ordersRepository.Save"

	id := ""
	err := d.conn.QueryRowContext(ctx, `INSERT INTO orders(store_id, meta_data, amount, user_id, type, is_published) VALUES ($1 , $2, $3, $4, $5,  $6) RETURNING id`,
		order.StoreID, order.MetaData, order.Amount, order.UserID, order.Type.String(), order.IsPublished).Scan(&id)
	if err != nil {
		return entity.Order{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	order.ID = id

	return order, nil
}
