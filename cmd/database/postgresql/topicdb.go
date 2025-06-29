package postgresql

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"back-end/core/logger"
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type TopicDB struct {
	table               string
	connect             *sqlx.DB
	IgnoreInsertColumns []string
	builder             sq.StatementBuilderType
}

func NewTopicDB() (repo.TopicRepo, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:524020@localhost:5432/nckh?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &TopicDB{
		table:               "topic",
		connect:             db,
		IgnoreInsertColumns: []string{"id"},
		builder:             sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

func (topic *TopicDB) Close() {
	_ = topic.connect.Close()
}

func (u *TopicDB) GetTopic(ctx context.Context, condition *repo.GetCondition) ([]*model.Topic, error) {
	ctxLogger := logger.NewContextLog(ctx)

	db := u.builder.Select("*").From(u.table)
	if condition != nil {
		db = db.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", condition.Key)})
	}
	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return nil, err
	}
	var result []*model.Topic
	err = u.connect.SelectContext(ctx, &result, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed while select, error: %s", err.Error())
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}
	return result, nil
}

func (u *TopicDB) CreateTopic(ctx context.Context, create *model.Topic) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := u.builder.Insert(u.table).
		Columns(GetListColumn(create, u.IgnoreInsertColumns, []string{})...).
		Values(GetListValues(create, u.IgnoreInsertColumns, []string{})...)
	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build insert query, error: %s", err.Error())
		return err
	}
	_, err = u.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed while insert topic, error: %s", err.Error())
		return err
	}
	return nil
}

func (u *TopicDB) UpdateTopic(ctx context.Context, update *model.Topic) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := u.builder.Update(u.table).
		Set("name", update.TopicName).
		Set("group_student_id", update.GroupStudentID).
		Set("lecturer_id", update.LecturerID).
		Set("start_day", update.StartDay).
		Set("end_day", update.EndDay).
		Set("allowance", update.Allowance).
		Set("status", update.TopicStatus).
		Where(sq.Eq{"id": update.TopicID})

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build update query, error: %s", err.Error())
		return err
	}
	_, err = u.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed while update topic, error: %s", err.Error())
		return err
	}
	return nil
}
