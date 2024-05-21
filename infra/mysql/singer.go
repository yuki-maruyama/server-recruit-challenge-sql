package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type singerRepository struct {
	db *sql.DB
}

var _ repository.SingerRepository = (*singerRepository)(nil)

func NewSingerRepository(db *sql.DB) *singerRepository {

	return &singerRepository{
		db,
	}
}

func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	singers := []*model.Singer{}
	rows, err := r.db.QueryContext(ctx, `
	SELECT
		id,
		name
	FROM
		singers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id   model.SingerID
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		singer := model.Singer{
			ID:   id,
			Name: name,
		}
		singers = append(singers, &singer)
	}
	return singers, nil
}

func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	var singer model.Singer
	if err := r.db.QueryRowContext(ctx, `
	SELECT
		id,
		name
	FROM
		singers
	WHERE
		id = ?
	`, id).Scan(&singer.ID, &singer.Name); err != nil {
		return nil, err
	}
	return &singer, nil
}

func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	return fmt.Errorf("need to implement")
}

func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	return fmt.Errorf("need to implement")
}
