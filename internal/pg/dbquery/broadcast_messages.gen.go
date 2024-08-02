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

func newBroadcastMessage(db *gorm.DB, opts ...gen.DOOption) broadcastMessage {
	_broadcastMessage := broadcastMessage{}

	_broadcastMessage.broadcastMessageDo.UseDB(db, opts...)
	_broadcastMessage.broadcastMessageDo.UseModel(&dbmodels.BroadcastMessage{})

	tableName := _broadcastMessage.broadcastMessageDo.TableName()
	_broadcastMessage.ALL = field.NewAsterisk(tableName)
	_broadcastMessage.ID = field.NewInt32(tableName, "id")
	_broadcastMessage.IsSent = field.NewBool(tableName, "is_sent")
	_broadcastMessage.Message = field.NewString(tableName, "message")
	_broadcastMessage.Image = field.NewString(tableName, "image")

	_broadcastMessage.fillFieldMap()

	return _broadcastMessage
}

type broadcastMessage struct {
	broadcastMessageDo

	ALL     field.Asterisk
	ID      field.Int32
	IsSent  field.Bool
	Message field.String
	Image   field.String

	fieldMap map[string]field.Expr
}

func (b broadcastMessage) Table(newTableName string) *broadcastMessage {
	b.broadcastMessageDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b broadcastMessage) As(alias string) *broadcastMessage {
	b.broadcastMessageDo.DO = *(b.broadcastMessageDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *broadcastMessage) updateTableName(table string) *broadcastMessage {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt32(table, "id")
	b.IsSent = field.NewBool(table, "is_sent")
	b.Message = field.NewString(table, "message")
	b.Image = field.NewString(table, "image")

	b.fillFieldMap()

	return b
}

func (b *broadcastMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *broadcastMessage) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 4)
	b.fieldMap["id"] = b.ID
	b.fieldMap["is_sent"] = b.IsSent
	b.fieldMap["message"] = b.Message
	b.fieldMap["image"] = b.Image
}

func (b broadcastMessage) clone(db *gorm.DB) broadcastMessage {
	b.broadcastMessageDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b broadcastMessage) replaceDB(db *gorm.DB) broadcastMessage {
	b.broadcastMessageDo.ReplaceDB(db)
	return b
}

type broadcastMessageDo struct{ gen.DO }

type IBroadcastMessageDo interface {
	gen.SubQuery
	Debug() IBroadcastMessageDo
	WithContext(ctx context.Context) IBroadcastMessageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBroadcastMessageDo
	WriteDB() IBroadcastMessageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBroadcastMessageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBroadcastMessageDo
	Not(conds ...gen.Condition) IBroadcastMessageDo
	Or(conds ...gen.Condition) IBroadcastMessageDo
	Select(conds ...field.Expr) IBroadcastMessageDo
	Where(conds ...gen.Condition) IBroadcastMessageDo
	Order(conds ...field.Expr) IBroadcastMessageDo
	Distinct(cols ...field.Expr) IBroadcastMessageDo
	Omit(cols ...field.Expr) IBroadcastMessageDo
	Join(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo
	Group(cols ...field.Expr) IBroadcastMessageDo
	Having(conds ...gen.Condition) IBroadcastMessageDo
	Limit(limit int) IBroadcastMessageDo
	Offset(offset int) IBroadcastMessageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBroadcastMessageDo
	Unscoped() IBroadcastMessageDo
	Create(values ...*dbmodels.BroadcastMessage) error
	CreateInBatches(values []*dbmodels.BroadcastMessage, batchSize int) error
	Save(values ...*dbmodels.BroadcastMessage) error
	First() (*dbmodels.BroadcastMessage, error)
	Take() (*dbmodels.BroadcastMessage, error)
	Last() (*dbmodels.BroadcastMessage, error)
	Find() ([]*dbmodels.BroadcastMessage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodels.BroadcastMessage, err error)
	FindInBatches(result *[]*dbmodels.BroadcastMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dbmodels.BroadcastMessage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBroadcastMessageDo
	Assign(attrs ...field.AssignExpr) IBroadcastMessageDo
	Joins(fields ...field.RelationField) IBroadcastMessageDo
	Preload(fields ...field.RelationField) IBroadcastMessageDo
	FirstOrInit() (*dbmodels.BroadcastMessage, error)
	FirstOrCreate() (*dbmodels.BroadcastMessage, error)
	FindByPage(offset int, limit int) (result []*dbmodels.BroadcastMessage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBroadcastMessageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b broadcastMessageDo) Debug() IBroadcastMessageDo {
	return b.withDO(b.DO.Debug())
}

func (b broadcastMessageDo) WithContext(ctx context.Context) IBroadcastMessageDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b broadcastMessageDo) ReadDB() IBroadcastMessageDo {
	return b.Clauses(dbresolver.Read)
}

func (b broadcastMessageDo) WriteDB() IBroadcastMessageDo {
	return b.Clauses(dbresolver.Write)
}

func (b broadcastMessageDo) Session(config *gorm.Session) IBroadcastMessageDo {
	return b.withDO(b.DO.Session(config))
}

func (b broadcastMessageDo) Clauses(conds ...clause.Expression) IBroadcastMessageDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b broadcastMessageDo) Returning(value interface{}, columns ...string) IBroadcastMessageDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b broadcastMessageDo) Not(conds ...gen.Condition) IBroadcastMessageDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b broadcastMessageDo) Or(conds ...gen.Condition) IBroadcastMessageDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b broadcastMessageDo) Select(conds ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b broadcastMessageDo) Where(conds ...gen.Condition) IBroadcastMessageDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b broadcastMessageDo) Order(conds ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b broadcastMessageDo) Distinct(cols ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b broadcastMessageDo) Omit(cols ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b broadcastMessageDo) Join(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b broadcastMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b broadcastMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b broadcastMessageDo) Group(cols ...field.Expr) IBroadcastMessageDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b broadcastMessageDo) Having(conds ...gen.Condition) IBroadcastMessageDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b broadcastMessageDo) Limit(limit int) IBroadcastMessageDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b broadcastMessageDo) Offset(offset int) IBroadcastMessageDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b broadcastMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBroadcastMessageDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b broadcastMessageDo) Unscoped() IBroadcastMessageDo {
	return b.withDO(b.DO.Unscoped())
}

func (b broadcastMessageDo) Create(values ...*dbmodels.BroadcastMessage) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b broadcastMessageDo) CreateInBatches(values []*dbmodels.BroadcastMessage, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b broadcastMessageDo) Save(values ...*dbmodels.BroadcastMessage) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b broadcastMessageDo) First() (*dbmodels.BroadcastMessage, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.BroadcastMessage), nil
	}
}

func (b broadcastMessageDo) Take() (*dbmodels.BroadcastMessage, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.BroadcastMessage), nil
	}
}

func (b broadcastMessageDo) Last() (*dbmodels.BroadcastMessage, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.BroadcastMessage), nil
	}
}

func (b broadcastMessageDo) Find() ([]*dbmodels.BroadcastMessage, error) {
	result, err := b.DO.Find()
	return result.([]*dbmodels.BroadcastMessage), err
}

func (b broadcastMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodels.BroadcastMessage, err error) {
	buf := make([]*dbmodels.BroadcastMessage, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b broadcastMessageDo) FindInBatches(result *[]*dbmodels.BroadcastMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b broadcastMessageDo) Attrs(attrs ...field.AssignExpr) IBroadcastMessageDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b broadcastMessageDo) Assign(attrs ...field.AssignExpr) IBroadcastMessageDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b broadcastMessageDo) Joins(fields ...field.RelationField) IBroadcastMessageDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b broadcastMessageDo) Preload(fields ...field.RelationField) IBroadcastMessageDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b broadcastMessageDo) FirstOrInit() (*dbmodels.BroadcastMessage, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.BroadcastMessage), nil
	}
}

func (b broadcastMessageDo) FirstOrCreate() (*dbmodels.BroadcastMessage, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodels.BroadcastMessage), nil
	}
}

func (b broadcastMessageDo) FindByPage(offset int, limit int) (result []*dbmodels.BroadcastMessage, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b broadcastMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b broadcastMessageDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b broadcastMessageDo) Delete(models ...*dbmodels.BroadcastMessage) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *broadcastMessageDo) withDO(do gen.Dao) *broadcastMessageDo {
	b.DO = *do.(*gen.DO)
	return b
}
