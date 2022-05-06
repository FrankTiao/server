// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/bangumi/server/internal/dal/dao"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newWebSession(db *gorm.DB) webSession {
	_webSession := webSession{}

	_webSession.webSessionDo.UseDB(db)
	_webSession.webSessionDo.UseModel(&dao.WebSession{})

	tableName := _webSession.webSessionDo.TableName()
	_webSession.ALL = field.NewField(tableName, "*")
	_webSession.Key = field.NewString(tableName, "key")
	_webSession.UserID = field.NewUint32(tableName, "user_id")
	_webSession.Value = field.NewField(tableName, "value")
	_webSession.CreatedAt = field.NewInt64(tableName, "created_at")
	_webSession.ExpiredAt = field.NewInt64(tableName, "expired_at")

	_webSession.fillFieldMap()

	return _webSession
}

type webSession struct {
	webSessionDo webSessionDo

	ALL       field.Field
	Key       field.String
	UserID    field.Uint32
	Value     field.Field
	CreatedAt field.Int64
	ExpiredAt field.Int64

	fieldMap map[string]field.Expr
}

func (w webSession) Table(newTableName string) *webSession {
	w.webSessionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webSession) As(alias string) *webSession {
	w.webSessionDo.DO = *(w.webSessionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webSession) updateTableName(table string) *webSession {
	w.ALL = field.NewField(table, "*")
	w.Key = field.NewString(table, "key")
	w.UserID = field.NewUint32(table, "user_id")
	w.Value = field.NewField(table, "value")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.ExpiredAt = field.NewInt64(table, "expired_at")

	w.fillFieldMap()

	return w
}

func (w *webSession) WithContext(ctx context.Context) *webSessionDo {
	return w.webSessionDo.WithContext(ctx)
}

func (w webSession) TableName() string { return w.webSessionDo.TableName() }

func (w webSession) Alias() string { return w.webSessionDo.Alias() }

func (w *webSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webSession) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["key"] = w.Key
	w.fieldMap["user_id"] = w.UserID
	w.fieldMap["value"] = w.Value
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["expired_at"] = w.ExpiredAt
}

func (w webSession) clone(db *gorm.DB) webSession {
	w.webSessionDo.ReplaceDB(db)
	return w
}

type webSessionDo struct{ gen.DO }

func (w webSessionDo) Debug() *webSessionDo {
	return w.withDO(w.DO.Debug())
}

func (w webSessionDo) WithContext(ctx context.Context) *webSessionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webSessionDo) Clauses(conds ...clause.Expression) *webSessionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webSessionDo) Returning(value interface{}, columns ...string) *webSessionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webSessionDo) Not(conds ...gen.Condition) *webSessionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webSessionDo) Or(conds ...gen.Condition) *webSessionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webSessionDo) Select(conds ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webSessionDo) Where(conds ...gen.Condition) *webSessionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webSessionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *webSessionDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w webSessionDo) Order(conds ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webSessionDo) Distinct(cols ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webSessionDo) Omit(cols ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webSessionDo) Join(table schema.Tabler, on ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webSessionDo) Group(cols ...field.Expr) *webSessionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webSessionDo) Having(conds ...gen.Condition) *webSessionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webSessionDo) Limit(limit int) *webSessionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webSessionDo) Offset(offset int) *webSessionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *webSessionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webSessionDo) Unscoped() *webSessionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webSessionDo) Create(values ...*dao.WebSession) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webSessionDo) CreateInBatches(values []*dao.WebSession, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webSessionDo) Save(values ...*dao.WebSession) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webSessionDo) First() (*dao.WebSession, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.WebSession), nil
	}
}

func (w webSessionDo) Take() (*dao.WebSession, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.WebSession), nil
	}
}

func (w webSessionDo) Last() (*dao.WebSession, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.WebSession), nil
	}
}

func (w webSessionDo) Find() ([]*dao.WebSession, error) {
	result, err := w.DO.Find()
	return result.([]*dao.WebSession), err
}

func (w webSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.WebSession, err error) {
	buf := make([]*dao.WebSession, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webSessionDo) FindInBatches(result *[]*dao.WebSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webSessionDo) Attrs(attrs ...field.AssignExpr) *webSessionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webSessionDo) Assign(attrs ...field.AssignExpr) *webSessionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webSessionDo) Joins(field field.RelationField) *webSessionDo {
	return w.withDO(w.DO.Joins(field))
}

func (w webSessionDo) Preload(field field.RelationField) *webSessionDo {
	return w.withDO(w.DO.Preload(field))
}

func (w webSessionDo) FirstOrInit() (*dao.WebSession, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.WebSession), nil
	}
}

func (w webSessionDo) FirstOrCreate() (*dao.WebSession, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.WebSession), nil
	}
}

func (w webSessionDo) FindByPage(offset int, limit int) (result []*dao.WebSession, count int64, err error) {
	if limit <= 0 {
		count, err = w.Count()
		return
	}

	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w webSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w *webSessionDo) withDO(do gen.Dao) *webSessionDo {
	w.DO = *do.(*gen.DO)
	return w
}
