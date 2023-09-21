// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/KeepShareOrg/keepshare/server/model"
)

func newSharedLink(db *gorm.DB, opts ...gen.DOOption) sharedLink {
	_sharedLink := sharedLink{}

	_sharedLink.sharedLinkDo.UseDB(db, opts...)
	_sharedLink.sharedLinkDo.UseModel(&model.SharedLink{})

	tableName := _sharedLink.sharedLinkDo.TableName()
	_sharedLink.ALL = field.NewAsterisk(tableName)
	_sharedLink.AutoID = field.NewInt64(tableName, "auto_id")
	_sharedLink.UserID = field.NewString(tableName, "user_id")
	_sharedLink.State = field.NewString(tableName, "state")
	_sharedLink.Host = field.NewString(tableName, "host")
	_sharedLink.CreatedBy = field.NewString(tableName, "created_by")
	_sharedLink.CreatedAt = field.NewTime(tableName, "created_at")
	_sharedLink.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sharedLink.Size = field.NewInt64(tableName, "size")
	_sharedLink.Visitor = field.NewInt32(tableName, "visitor")
	_sharedLink.Stored = field.NewInt32(tableName, "stored")
	_sharedLink.LastVisitedAt = field.NewTime(tableName, "last_visited_at")
	_sharedLink.LastStoredAt = field.NewTime(tableName, "last_stored_at")
	_sharedLink.Revenue = field.NewInt64(tableName, "revenue")
	_sharedLink.Title = field.NewString(tableName, "title")
	_sharedLink.OriginalLinkHash = field.NewString(tableName, "original_link_hash")
	_sharedLink.HostSharedLinkHash = field.NewString(tableName, "host_shared_link_hash")
	_sharedLink.OriginalLink = field.NewString(tableName, "original_link")
	_sharedLink.HostSharedLink = field.NewString(tableName, "host_shared_link")

	_sharedLink.fillFieldMap()

	return _sharedLink
}

type sharedLink struct {
	sharedLinkDo

	ALL                field.Asterisk
	AutoID             field.Int64
	UserID             field.String
	State              field.String
	Host               field.String
	CreatedBy          field.String
	CreatedAt          field.Time
	UpdatedAt          field.Time
	Size               field.Int64
	Visitor            field.Int32
	Stored             field.Int32
	LastVisitedAt      field.Time
	LastStoredAt       field.Time
	Revenue            field.Int64
	Title              field.String
	OriginalLinkHash   field.String
	HostSharedLinkHash field.String
	OriginalLink       field.String
	HostSharedLink     field.String

	fieldMap map[string]field.Expr
}

func (s sharedLink) Table(newTableName string) *sharedLink {
	s.sharedLinkDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sharedLink) As(alias string) *sharedLink {
	s.sharedLinkDo.DO = *(s.sharedLinkDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sharedLink) updateTableName(table string) *sharedLink {
	s.ALL = field.NewAsterisk(table)
	s.AutoID = field.NewInt64(table, "auto_id")
	s.UserID = field.NewString(table, "user_id")
	s.State = field.NewString(table, "state")
	s.Host = field.NewString(table, "host")
	s.CreatedBy = field.NewString(table, "created_by")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.Size = field.NewInt64(table, "size")
	s.Visitor = field.NewInt32(table, "visitor")
	s.Stored = field.NewInt32(table, "stored")
	s.LastVisitedAt = field.NewTime(table, "last_visited_at")
	s.LastStoredAt = field.NewTime(table, "last_stored_at")
	s.Revenue = field.NewInt64(table, "revenue")
	s.Title = field.NewString(table, "title")
	s.OriginalLinkHash = field.NewString(table, "original_link_hash")
	s.HostSharedLinkHash = field.NewString(table, "host_shared_link_hash")
	s.OriginalLink = field.NewString(table, "original_link")
	s.HostSharedLink = field.NewString(table, "host_shared_link")

	s.fillFieldMap()

	return s
}

func (s *sharedLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sharedLink) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 18)
	s.fieldMap["auto_id"] = s.AutoID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["state"] = s.State
	s.fieldMap["host"] = s.Host
	s.fieldMap["created_by"] = s.CreatedBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["size"] = s.Size
	s.fieldMap["visitor"] = s.Visitor
	s.fieldMap["stored"] = s.Stored
	s.fieldMap["last_visited_at"] = s.LastVisitedAt
	s.fieldMap["last_stored_at"] = s.LastStoredAt
	s.fieldMap["revenue"] = s.Revenue
	s.fieldMap["title"] = s.Title
	s.fieldMap["original_link_hash"] = s.OriginalLinkHash
	s.fieldMap["host_shared_link_hash"] = s.HostSharedLinkHash
	s.fieldMap["original_link"] = s.OriginalLink
	s.fieldMap["host_shared_link"] = s.HostSharedLink
}

func (s sharedLink) clone(db *gorm.DB) sharedLink {
	s.sharedLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sharedLink) replaceDB(db *gorm.DB) sharedLink {
	s.sharedLinkDo.ReplaceDB(db)
	return s
}

type sharedLinkDo struct{ gen.DO }

type ISharedLinkDo interface {
	gen.SubQuery
	Debug() ISharedLinkDo
	WithContext(ctx context.Context) ISharedLinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISharedLinkDo
	WriteDB() ISharedLinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISharedLinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISharedLinkDo
	Not(conds ...gen.Condition) ISharedLinkDo
	Or(conds ...gen.Condition) ISharedLinkDo
	Select(conds ...field.Expr) ISharedLinkDo
	Where(conds ...gen.Condition) ISharedLinkDo
	Order(conds ...field.Expr) ISharedLinkDo
	Distinct(cols ...field.Expr) ISharedLinkDo
	Omit(cols ...field.Expr) ISharedLinkDo
	Join(table schema.Tabler, on ...field.Expr) ISharedLinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISharedLinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISharedLinkDo
	Group(cols ...field.Expr) ISharedLinkDo
	Having(conds ...gen.Condition) ISharedLinkDo
	Limit(limit int) ISharedLinkDo
	Offset(offset int) ISharedLinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISharedLinkDo
	Unscoped() ISharedLinkDo
	Create(values ...*model.SharedLink) error
	CreateInBatches(values []*model.SharedLink, batchSize int) error
	Save(values ...*model.SharedLink) error
	First() (*model.SharedLink, error)
	Take() (*model.SharedLink, error)
	Last() (*model.SharedLink, error)
	Find() ([]*model.SharedLink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SharedLink, err error)
	FindInBatches(result *[]*model.SharedLink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SharedLink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISharedLinkDo
	Assign(attrs ...field.AssignExpr) ISharedLinkDo
	Joins(fields ...field.RelationField) ISharedLinkDo
	Preload(fields ...field.RelationField) ISharedLinkDo
	FirstOrInit() (*model.SharedLink, error)
	FirstOrCreate() (*model.SharedLink, error)
	FindByPage(offset int, limit int) (result []*model.SharedLink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISharedLinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sharedLinkDo) Debug() ISharedLinkDo {
	return s.withDO(s.DO.Debug())
}

func (s sharedLinkDo) WithContext(ctx context.Context) ISharedLinkDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sharedLinkDo) ReadDB() ISharedLinkDo {
	return s.Clauses(dbresolver.Read)
}

func (s sharedLinkDo) WriteDB() ISharedLinkDo {
	return s.Clauses(dbresolver.Write)
}

func (s sharedLinkDo) Session(config *gorm.Session) ISharedLinkDo {
	return s.withDO(s.DO.Session(config))
}

func (s sharedLinkDo) Clauses(conds ...clause.Expression) ISharedLinkDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sharedLinkDo) Returning(value interface{}, columns ...string) ISharedLinkDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sharedLinkDo) Not(conds ...gen.Condition) ISharedLinkDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sharedLinkDo) Or(conds ...gen.Condition) ISharedLinkDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sharedLinkDo) Select(conds ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sharedLinkDo) Where(conds ...gen.Condition) ISharedLinkDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sharedLinkDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISharedLinkDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sharedLinkDo) Order(conds ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sharedLinkDo) Distinct(cols ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sharedLinkDo) Omit(cols ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sharedLinkDo) Join(table schema.Tabler, on ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sharedLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sharedLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sharedLinkDo) Group(cols ...field.Expr) ISharedLinkDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sharedLinkDo) Having(conds ...gen.Condition) ISharedLinkDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sharedLinkDo) Limit(limit int) ISharedLinkDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sharedLinkDo) Offset(offset int) ISharedLinkDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sharedLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISharedLinkDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sharedLinkDo) Unscoped() ISharedLinkDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sharedLinkDo) Create(values ...*model.SharedLink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sharedLinkDo) CreateInBatches(values []*model.SharedLink, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sharedLinkDo) Save(values ...*model.SharedLink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sharedLinkDo) First() (*model.SharedLink, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SharedLink), nil
	}
}

func (s sharedLinkDo) Take() (*model.SharedLink, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SharedLink), nil
	}
}

func (s sharedLinkDo) Last() (*model.SharedLink, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SharedLink), nil
	}
}

func (s sharedLinkDo) Find() ([]*model.SharedLink, error) {
	result, err := s.DO.Find()
	return result.([]*model.SharedLink), err
}

func (s sharedLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SharedLink, err error) {
	buf := make([]*model.SharedLink, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sharedLinkDo) FindInBatches(result *[]*model.SharedLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sharedLinkDo) Attrs(attrs ...field.AssignExpr) ISharedLinkDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sharedLinkDo) Assign(attrs ...field.AssignExpr) ISharedLinkDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sharedLinkDo) Joins(fields ...field.RelationField) ISharedLinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sharedLinkDo) Preload(fields ...field.RelationField) ISharedLinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sharedLinkDo) FirstOrInit() (*model.SharedLink, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SharedLink), nil
	}
}

func (s sharedLinkDo) FirstOrCreate() (*model.SharedLink, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SharedLink), nil
	}
}

func (s sharedLinkDo) FindByPage(offset int, limit int) (result []*model.SharedLink, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sharedLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sharedLinkDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sharedLinkDo) Delete(models ...*model.SharedLink) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sharedLinkDo) withDO(do gen.Dao) *sharedLinkDo {
	s.DO = *do.(*gen.DO)
	return s
}