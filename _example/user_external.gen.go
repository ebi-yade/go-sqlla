// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla - DO NOT EDIT.
package example

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"database/sql"
	"time"

	"github.com/mackee/go-sqlla/v2"
)

type userExternalSQL struct {
	where sqlla.Where
}

func NewUserExternalSQL() userExternalSQL {
	q := userExternalSQL{}
	return q
}

var userExternalAllColumns = []string{
	"`id`", "`user_id`", "`created_at`", "`updated_at`",
}

type userExternalSelectSQL struct {
	userExternalSQL
	Columns     []string
	order       string
	limit       *uint64
	offset      *uint64
	tableAlias  string
	joinClauses []string

	additionalWhereClause     string
	additionalWhereClauseArgs []interface{}

	groupByColumns []string

	isForUpdate bool
}

func (q userExternalSQL) Select() userExternalSelectSQL {
	return userExternalSelectSQL{
		q,
		userExternalAllColumns,
		"",
		nil,
		nil,
		"",
		nil,
		"",
		nil,
		nil,
		false,
	}
}

func (q userExternalSelectSQL) Or(qs ...userExternalSelectSQL) userExternalSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q userExternalSelectSQL) Limit(l uint64) userExternalSelectSQL {
	q.limit = &l
	return q
}

func (q userExternalSelectSQL) Offset(o uint64) userExternalSelectSQL {
	q.offset = &o
	return q
}

func (q userExternalSelectSQL) ForUpdate() userExternalSelectSQL {
	q.isForUpdate = true
	return q
}

func (q userExternalSelectSQL) TableAlias(alias string) userExternalSelectSQL {
	q.tableAlias = "`" + alias + "`"
	return q
}

func (q userExternalSelectSQL) SetColumns(columns ...string) userExternalSelectSQL {
	q.Columns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsRune(column, '(') {
			q.Columns = append(q.Columns, column)
		} else {
			q.Columns = append(q.Columns, "`"+column+"`")
		}
	}
	return q
}

func (q userExternalSelectSQL) JoinClause(clause string) userExternalSelectSQL {
	q.joinClauses = append(q.joinClauses, clause)
	return q
}

func (q userExternalSelectSQL) AdditionalWhereClause(clause string, args ...interface{}) userExternalSelectSQL {
	q.additionalWhereClause = clause
	q.additionalWhereClauseArgs = args
	return q
}

func (q userExternalSelectSQL) appendColumnPrefix(column string) string {
	if q.tableAlias == "" || strings.ContainsRune(column, '(') {
		return column
	}
	return q.tableAlias + "." + column
}

func (q userExternalSelectSQL) GroupBy(columns ...string) userExternalSelectSQL {
	q.groupByColumns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsRune(column, '(') {
			q.groupByColumns = append(q.groupByColumns, column)
		} else {
			q.groupByColumns = append(q.groupByColumns, "`"+column+"`")
		}
	}
	return q
}

func (q userExternalSelectSQL) ID(v uint64, exprs ...sqlla.Operator) userExternalSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) IDIn(vs ...uint64) userExternalSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) userExternalSelectSQL {
	v := uint64(pk)
	return q.ID(v, exprs...)
}

func (q userExternalSelectSQL) OrderByID(order sqlla.Order) userExternalSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userExternalSelectSQL) UserID(v uint64, exprs ...sqlla.Operator) userExternalSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: q.appendColumnPrefix("`user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) UserIDIn(vs ...uint64) userExternalSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) OrderByUserID(order sqlla.Order) userExternalSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`user_id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userExternalSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userExternalSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) CreatedAtIn(vs ...time.Time) userExternalSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) OrderByCreatedAt(order sqlla.Order) userExternalSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`created_at`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userExternalSelectSQL) UpdatedAt(v time.Time, exprs ...sqlla.Operator) userExternalSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) UpdatedAtIn(vs ...time.Time) userExternalSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userExternalSelectSQL) OrderByUpdatedAt(order sqlla.Order) userExternalSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`updated_at`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userExternalSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	tableName := "user_external"
	if q.tableAlias != "" {
		tableName = tableName + " AS " + q.tableAlias
		pcs := make([]string, 0, len(q.Columns))
		for _, column := range q.Columns {
			pcs = append(pcs, q.appendColumnPrefix(column))
		}
		columns = strings.Join(pcs, ", ")
	}
	query := "SELECT " + columns + " FROM " + tableName
	if len(q.joinClauses) > 0 {
		jc := strings.Join(q.joinClauses, " ")
		query += " " + jc
	}
	if wheres != "" {
		query += " WHERE" + wheres
	}
	if q.additionalWhereClause != "" {
		query += " " + q.additionalWhereClause
		if len(q.additionalWhereClauseArgs) > 0 {
			vs = append(vs, q.additionalWhereClauseArgs...)
		}
	}
	if len(q.groupByColumns) > 0 {
		query += " GROUP BY "
		gbcs := make([]string, 0, len(q.groupByColumns))
		for _, column := range q.groupByColumns {
			gbcs = append(gbcs, q.appendColumnPrefix(column))
		}
		query += strings.Join(gbcs, ", ")
	}
	query += q.order
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}
	if q.offset != nil {
		query += " OFFSET " + strconv.FormatUint(*q.offset, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (s UserExternal) Select() userExternalSelectSQL {
	return NewUserExternalSQL().Select().ID(s.Id)
}
func (q userExternalSelectSQL) Single(db sqlla.DB) (UserExternal, error) {
	q.Columns = userExternalAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return UserExternal{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q userExternalSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (UserExternal, error) {
	q.Columns = userExternalAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return UserExternal{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
	return q.Scan(row)
}

func (q userExternalSelectSQL) All(db sqlla.DB) ([]UserExternal, error) {
	rs := make([]UserExternal, 0, 10)
	q.Columns = userExternalAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q userExternalSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]UserExternal, error) {
	rs := make([]UserExternal, 0, 10)
	q.Columns = userExternalAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q userExternalSelectSQL) Scan(s sqlla.Scanner) (UserExternal, error) {
	var row UserExternal
	err := s.Scan(
		&row.Id,
		&row.UserId,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	return row, err
}

type userExternalUpdateSQL struct {
	userExternalSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q userExternalSQL) Update() userExternalUpdateSQL {
	return userExternalUpdateSQL{
		userExternalSQL: q,
		setMap:          sqlla.SetMap{},
	}
}

func (q userExternalUpdateSQL) SetID(v uint64) userExternalUpdateSQL {
	q.setMap["`id`"] = v
	return q
}

func (q userExternalUpdateSQL) WhereID(v uint64, exprs ...sqlla.Operator) userExternalUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) WhereIDIn(vs ...uint64) userExternalUpdateSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) SetUserID(v uint64) userExternalUpdateSQL {
	q.setMap["`user_id`"] = v
	return q
}

func (q userExternalUpdateSQL) WhereUserID(v uint64, exprs ...sqlla.Operator) userExternalUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: "`user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) WhereUserIDIn(vs ...uint64) userExternalUpdateSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) SetCreatedAt(v time.Time) userExternalUpdateSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q userExternalUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) userExternalUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) WhereCreatedAtIn(vs ...time.Time) userExternalUpdateSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) SetUpdatedAt(v time.Time) userExternalUpdateSQL {
	q.setMap["`updated_at`"] = v
	return q
}

func (q userExternalUpdateSQL) WhereUpdatedAt(v time.Time, exprs ...sqlla.Operator) userExternalUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) WhereUpdatedAtIn(vs ...time.Time) userExternalUpdateSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = UserExternal{}
	if t, ok := s.(userExternalDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	setColumns, svs, err := q.setMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, wvs, err := q.where.ToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE user_external SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s UserExternal) Update() userExternalUpdateSQL {
	return NewUserExternalSQL().Update().WhereID(s.Id)
}

func (q userExternalUpdateSQL) Exec(db sqlla.DB) ([]UserExternal, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userExternalSQL

	return qq.Select().All(db)
}

func (q userExternalUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]UserExternal, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userExternalSQL

	return qq.Select().AllContext(ctx, db)
}

type userExternalDefaultUpdateHooker interface {
	DefaultUpdateHook(userExternalUpdateSQL) (userExternalUpdateSQL, error)
}

type userExternalInsertSQL struct {
	userExternalSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q userExternalSQL) Insert() userExternalInsertSQL {
	return userExternalInsertSQL{
		userExternalSQL: q,
		setMap:          sqlla.SetMap{},
	}
}

func (q userExternalInsertSQL) ValueID(v uint64) userExternalInsertSQL {
	q.setMap["`id`"] = v
	return q
}

func (q userExternalInsertSQL) ValueUserID(v uint64) userExternalInsertSQL {
	q.setMap["`user_id`"] = v
	return q
}

func (q userExternalInsertSQL) ValueCreatedAt(v time.Time) userExternalInsertSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q userExternalInsertSQL) ValueUpdatedAt(v time.Time) userExternalInsertSQL {
	q.setMap["`updated_at`"] = v
	return q
}

func (q userExternalInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.userExternalInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q userExternalInsertSQL) userExternalInsertSQLToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = UserExternal{}
	if t, ok := s.(userExternalDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO user_external " + qs

	return query, vs, nil
}

func (q userExternalInsertSQL) OnDuplicateKeyUpdate() userExternalInsertOnDuplicateKeyUpdateSQL {
	return userExternalInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q userExternalInsertSQL) Exec(db sqlla.DB) (UserExternal, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return UserExternal{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return UserExternal{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserExternal{}, err
	}
	return NewUserExternalSQL().Select().PkColumn(id).Single(db)
}

func (q userExternalInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (UserExternal, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return UserExternal{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return UserExternal{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserExternal{}, err
	}
	return NewUserExternalSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q userExternalInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userExternalDefaultInsertHooker interface {
	DefaultInsertHook(userExternalInsertSQL) (userExternalInsertSQL, error)
}

type userExternalInsertSQLToSqler interface {
	userExternalInsertSQLToSql() (string, []interface{}, error)
}

type userExternalInsertOnDuplicateKeyUpdateSQL struct {
	insertSQL               userExternalInsertSQLToSqler
	onDuplicateKeyUpdateMap sqlla.SetMap
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateID(v uint64) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) SameOnUpdateID() userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = sqlla.SetMapRawValue("VALUES(`id`)")
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateUserID(v uint64) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`user_id`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateUserID(v sqlla.SetMapRawValue) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`user_id`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) SameOnUpdateUserID() userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`user_id`"] = sqlla.SetMapRawValue("VALUES(`user_id`)")
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateCreatedAt(v time.Time) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) SameOnUpdateCreatedAt() userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = sqlla.SetMapRawValue("VALUES(`created_at`)")
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateUpdatedAt(v time.Time) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateUpdatedAt(v sqlla.SetMapRawValue) userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = v
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) SameOnUpdateUpdatedAt() userExternalInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = sqlla.SetMapRawValue("VALUES(`updated_at`)")
	return q
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = UserExternal{}
	if t, ok := s.(userExternalDefaultInsertOnDuplicateKeyUpdateHooker); ok {
		q, err = t.DefaultInsertOnDuplicateKeyUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}

	query, vs, err := q.insertSQL.userExternalInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	os, ovs, err := q.onDuplicateKeyUpdateMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	query += " ON DUPLICATE KEY UPDATE" + os
	vs = append(vs, ovs...)

	return query + ";", vs, nil
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) (UserExternal, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return UserExternal{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return UserExternal{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserExternal{}, err
	}
	return NewUserExternalSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q userExternalInsertOnDuplicateKeyUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userExternalDefaultInsertOnDuplicateKeyUpdateHooker interface {
	DefaultInsertOnDuplicateKeyUpdateHook(userExternalInsertOnDuplicateKeyUpdateSQL) (userExternalInsertOnDuplicateKeyUpdateSQL, error)
}

type userExternalBulkInsertSQL struct {
	insertSQLs []userExternalInsertSQL
}

func (q userExternalSQL) BulkInsert() *userExternalBulkInsertSQL {
	return &userExternalBulkInsertSQL{
		insertSQLs: []userExternalInsertSQL{},
	}
}

func (q *userExternalBulkInsertSQL) Append(iqs ...userExternalInsertSQL) {
	q.insertSQLs = append(q.insertSQLs, iqs...)
}

func (q *userExternalBulkInsertSQL) userExternalInsertSQLToSql() (string, []interface{}, error) {
	if len(q.insertSQLs) == 0 {
		return "", []interface{}{}, fmt.Errorf("sqlla: This userExternalBulkInsertSQL's InsertSQL was empty")
	}
	iqs := make([]userExternalInsertSQL, len(q.insertSQLs))
	copy(iqs, q.insertSQLs)

	var s interface{} = UserExternal{}
	if t, ok := s.(userExternalDefaultInsertHooker); ok {
		for i, iq := range iqs {
			var err error
			iq, err = t.DefaultInsertHook(iq)
			if err != nil {
				return "", []interface{}{}, err
			}
			iqs[i] = iq
		}
	}

	sms := make(sqlla.SetMaps, 0, len(q.insertSQLs))
	for _, iq := range q.insertSQLs {
		sms = append(sms, iq.setMap)
	}

	query, vs, err := sms.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	return "INSERT INTO `user_external` " + query, vs, nil
}

func (q *userExternalBulkInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.userExternalInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q *userExternalBulkInsertSQL) OnDuplicateKeyUpdate() userExternalInsertOnDuplicateKeyUpdateSQL {
	return userExternalInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q *userExternalBulkInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userExternalDeleteSQL struct {
	userExternalSQL
}

func (q userExternalSQL) Delete() userExternalDeleteSQL {
	return userExternalDeleteSQL{
		q,
	}
}

func (q userExternalDeleteSQL) ID(v uint64, exprs ...sqlla.Operator) userExternalDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) IDIn(vs ...uint64) userExternalDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) UserID(v uint64, exprs ...sqlla.Operator) userExternalDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: v, Op: op, Column: "`user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) UserIDIn(vs ...uint64) userExternalDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userExternalDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) CreatedAtIn(vs ...time.Time) userExternalDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) UpdatedAt(v time.Time, exprs ...sqlla.Operator) userExternalDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) UpdatedAtIn(vs ...time.Time) userExternalDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userExternalDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM user_external"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func (q userExternalDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (q userExternalDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s UserExternal) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserExternalSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s UserExternal) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserExternalSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
