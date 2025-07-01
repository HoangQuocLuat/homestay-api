package postgresql

// import (
// 	"back-end/cmd/database/model"
// 	"back-end/cmd/database/repo"
// 	"context"
// 	"database/sql"
// 	"fmt"

// 	sq "github.com/Masterminds/squirrel"
// 	"github.com/jmoiron/sqlx"
// )

// type ServiceDB struct {
// 	table               string
// 	connect             *sqlx.DB
// 	IgnoreInsertColumns []string
// 	builder             sq.StatementBuilderType
// }

// func NewServiceDB() (repo.ServiceRepo, error) {
// 	db, err := sqlx.Open("postgres", "postgres://postgres:524020@localhost:5432/nckh?sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &ServiceDB{
// 		table:               "services",
// 		connect:             db,
// 		IgnoreInsertColumns: []string{"id"},
// 		builder:             sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
// 	}, nil
// }

// func (s *ServiceDB) Close() {
// 	_ = s.connect.Close()
// }

// func (s *ServiceDB) GetServices(ctx context.Context, condition *repo.GetCondition) ([]*model.Service, error) {

// 	db := s.builder.Select("*").From(s.table)
// 	if condition != nil && condition.Key != "" {
// 		db = db.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", condition.Key)})
// 	}

// 	query, args, err := db.ToSql()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result []*model.Service
// 	err = s.connect.SelectContext(ctx, &result, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(result) == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	return result, nil
// }

// func (s *ServiceDB) CreateService(ctx context.Context, service *model.Service) error {

// 	db := s.builder.Insert(s.table).
// 		Columns(GetListColumn(service, s.IgnoreInsertColumns, []string{})...).
// 		Values(GetListValues(service, s.IgnoreInsertColumns, []string{})...)

// 	query, args, err := db.ToSql()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.connect.ExecContext(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *ServiceDB) UpdateService(ctx context.Context, service *model.Service) error {

// 	db := s.builder.Update(s.table).
// 		Set("name", service.Name).
// 		Set("updated_at", service.UpdatedAt).
// 		Where(sq.Eq{"id": service.Id})

// 	query, args, err := db.ToSql()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.connect.ExecContext(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
