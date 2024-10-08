// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbquery

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"dobrino/internal/pg/dbmodels"
)

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&dbmodels.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewString(tableName, "id")
	_admin.DateCreated = field.NewTime(tableName, "date_created")
	_admin.ChatID = field.NewString(tableName, "chat_id")

	_admin.fillFieldMap()

	return _admin
}

type admin struct {
	adminDo

	ALL         field.Asterisk
	ID          field.String
	DateCreated field.Time
	ChatID      field.String

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.DateCreated = field.NewTime(table, "date_created")
	a.ChatID = field.NewString(table, "chat_id")

	a.fillFieldMap()

	return a
}

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["date_created"] = a.DateCreated
	a.fieldMap["chat_id"] = a.ChatID
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

type IAdminDo interface {
	gen.SubQuery
	Debug() IAdminDo
	WithContext(ctx context.Context) IAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminDo
	WriteDB() IAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminDo
	Not(conds ...gen.Condition) IAdminDo
	Or(conds ...gen.Condition) IAdminDo
	Select(conds ...field.Expr) IAdminDo
	Where(conds ...gen.Condition) IAdminDo
	Order(conds ...field.Expr) IAdminDo
	Distinct(cols ...field.Expr) IAdminDo
	Omit(cols ...field.Expr) IAdminDo
	Join(table schema.Tabler, on ...field.Expr) IAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	Group(cols ...field.Expr) IAdminDo
	Having(conds ...gen.Condition) IAdminDo
	Limit(limit int) IAdminDo
	Offset(offset int) IAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo
	Unscoped() IAdminDo
	Create(values ...*dbmodels.Admin) error
	CreateInBatches(values []*dbmodels.Admin, batchSize int) error
	Save(values ...*dbmodels.Admin) error
	First() (*dbmodels.Admin, error)
	Take() (*dbmodels.Admin, error)
	Last() (*dbmodels.Admin, error)
	Find() ([]*dbmodels.Admin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodels.Admin, err error)
	FindInBatches(result *[]*dbmodels.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dbmodels.Admin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminDo
	Assign(attrs ...field.AssignExpr) IAdminDo
	Joins(fields ...field.RelationField) IAdminDo
	Preload(fields ...field.RelationField) IAdminDo
	FirstOrInit() (*dbmodels.Admin, error)
	FirstOrCreate() (*dbmodels.Admin, error)
	FindByPage(offset int, limit int) (result []*dbmodels.Admin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminDo) Debug() IAdminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) IAdminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() IAdminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() IAdminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) IAdminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) IAdminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) IAdminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Order(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) IAdminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) IAdminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() IAdminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*dbmodels.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*dbmodels.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*dbmodels.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*dbmodels.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.Admin), nil
	}
}

func (a adminDo) Take() (*dbmodels.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.Admin), nil
	}
}

func (a adminDo) Last() (*dbmodels.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.Admin), nil
	}
}

func (a adminDo) Find() ([]*dbmodels.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*dbmodels.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodels.Admin, err error) {
	buf := make([]*dbmodels.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*dbmodels.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*dbmodels.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*dbmodels.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*dbmodels.Admin, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*dbmodels.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}
