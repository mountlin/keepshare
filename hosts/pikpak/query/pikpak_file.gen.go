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

	"github.com/KeepShareOrg/keepshare/hosts/pikpak/model"
)

func newFile(db *gorm.DB, opts ...gen.DOOption) file {
	_file := file{}

	_file.fileDo.UseDB(db, opts...)
	_file.fileDo.UseModel(&model.File{})

	tableName := _file.fileDo.TableName()
	_file.ALL = field.NewAsterisk(tableName)
	_file.AutoID = field.NewInt64(tableName, "auto_id")
	_file.MasterUserID = field.NewString(tableName, "master_user_id")
	_file.WorkerUserID = field.NewString(tableName, "worker_user_id")
	_file.FileID = field.NewString(tableName, "file_id")
	_file.TaskID = field.NewString(tableName, "task_id")
	_file.Status = field.NewString(tableName, "status")
	_file.IsDir = field.NewBool(tableName, "is_dir")
	_file.Size = field.NewInt64(tableName, "size")
	_file.Name = field.NewString(tableName, "name")
	_file.CreatedAt = field.NewTime(tableName, "created_at")
	_file.UpdatedAt = field.NewTime(tableName, "updated_at")
	_file.OriginalLinkHash = field.NewString(tableName, "original_link_hash")

	_file.fillFieldMap()

	return _file
}

type file struct {
	fileDo

	ALL              field.Asterisk
	AutoID           field.Int64
	MasterUserID     field.String
	WorkerUserID     field.String
	FileID           field.String
	TaskID           field.String
	Status           field.String
	IsDir            field.Bool
	Size             field.Int64
	Name             field.String
	CreatedAt        field.Time
	UpdatedAt        field.Time
	OriginalLinkHash field.String

	fieldMap map[string]field.Expr
}

func (f file) Table(newTableName string) *file {
	f.fileDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f file) As(alias string) *file {
	f.fileDo.DO = *(f.fileDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *file) updateTableName(table string) *file {
	f.ALL = field.NewAsterisk(table)
	f.AutoID = field.NewInt64(table, "auto_id")
	f.MasterUserID = field.NewString(table, "master_user_id")
	f.WorkerUserID = field.NewString(table, "worker_user_id")
	f.FileID = field.NewString(table, "file_id")
	f.TaskID = field.NewString(table, "task_id")
	f.Status = field.NewString(table, "status")
	f.IsDir = field.NewBool(table, "is_dir")
	f.Size = field.NewInt64(table, "size")
	f.Name = field.NewString(table, "name")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.OriginalLinkHash = field.NewString(table, "original_link_hash")

	f.fillFieldMap()

	return f
}

func (f *file) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *file) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 12)
	f.fieldMap["auto_id"] = f.AutoID
	f.fieldMap["master_user_id"] = f.MasterUserID
	f.fieldMap["worker_user_id"] = f.WorkerUserID
	f.fieldMap["file_id"] = f.FileID
	f.fieldMap["task_id"] = f.TaskID
	f.fieldMap["status"] = f.Status
	f.fieldMap["is_dir"] = f.IsDir
	f.fieldMap["size"] = f.Size
	f.fieldMap["name"] = f.Name
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["original_link_hash"] = f.OriginalLinkHash
}

func (f file) clone(db *gorm.DB) file {
	f.fileDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f file) replaceDB(db *gorm.DB) file {
	f.fileDo.ReplaceDB(db)
	return f
}

type fileDo struct{ gen.DO }

type IFileDo interface {
	gen.SubQuery
	Debug() IFileDo
	WithContext(ctx context.Context) IFileDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFileDo
	WriteDB() IFileDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFileDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFileDo
	Not(conds ...gen.Condition) IFileDo
	Or(conds ...gen.Condition) IFileDo
	Select(conds ...field.Expr) IFileDo
	Where(conds ...gen.Condition) IFileDo
	Order(conds ...field.Expr) IFileDo
	Distinct(cols ...field.Expr) IFileDo
	Omit(cols ...field.Expr) IFileDo
	Join(table schema.Tabler, on ...field.Expr) IFileDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFileDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFileDo
	Group(cols ...field.Expr) IFileDo
	Having(conds ...gen.Condition) IFileDo
	Limit(limit int) IFileDo
	Offset(offset int) IFileDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFileDo
	Unscoped() IFileDo
	Create(values ...*model.File) error
	CreateInBatches(values []*model.File, batchSize int) error
	Save(values ...*model.File) error
	First() (*model.File, error)
	Take() (*model.File, error)
	Last() (*model.File, error)
	Find() ([]*model.File, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.File, err error)
	FindInBatches(result *[]*model.File, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.File) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFileDo
	Assign(attrs ...field.AssignExpr) IFileDo
	Joins(fields ...field.RelationField) IFileDo
	Preload(fields ...field.RelationField) IFileDo
	FirstOrInit() (*model.File, error)
	FirstOrCreate() (*model.File, error)
	FindByPage(offset int, limit int) (result []*model.File, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFileDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fileDo) Debug() IFileDo {
	return f.withDO(f.DO.Debug())
}

func (f fileDo) WithContext(ctx context.Context) IFileDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileDo) ReadDB() IFileDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileDo) WriteDB() IFileDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileDo) Session(config *gorm.Session) IFileDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileDo) Clauses(conds ...clause.Expression) IFileDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileDo) Returning(value interface{}, columns ...string) IFileDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileDo) Not(conds ...gen.Condition) IFileDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileDo) Or(conds ...gen.Condition) IFileDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileDo) Select(conds ...field.Expr) IFileDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileDo) Where(conds ...gen.Condition) IFileDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFileDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f fileDo) Order(conds ...field.Expr) IFileDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileDo) Distinct(cols ...field.Expr) IFileDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileDo) Omit(cols ...field.Expr) IFileDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileDo) Join(table schema.Tabler, on ...field.Expr) IFileDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFileDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileDo) RightJoin(table schema.Tabler, on ...field.Expr) IFileDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileDo) Group(cols ...field.Expr) IFileDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileDo) Having(conds ...gen.Condition) IFileDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileDo) Limit(limit int) IFileDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileDo) Offset(offset int) IFileDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFileDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileDo) Unscoped() IFileDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileDo) Create(values ...*model.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileDo) CreateInBatches(values []*model.File, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileDo) Save(values ...*model.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileDo) First() (*model.File, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Take() (*model.File, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Last() (*model.File, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) Find() ([]*model.File, error) {
	result, err := f.DO.Find()
	return result.([]*model.File), err
}

func (f fileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.File, err error) {
	buf := make([]*model.File, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileDo) FindInBatches(result *[]*model.File, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileDo) Attrs(attrs ...field.AssignExpr) IFileDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileDo) Assign(attrs ...field.AssignExpr) IFileDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileDo) Joins(fields ...field.RelationField) IFileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileDo) Preload(fields ...field.RelationField) IFileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileDo) FirstOrInit() (*model.File, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) FirstOrCreate() (*model.File, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.File), nil
	}
}

func (f fileDo) FindByPage(offset int, limit int) (result []*model.File, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileDo) Delete(models ...*model.File) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileDo) withDO(do gen.Dao) *fileDo {
	f.DO = *do.(*gen.DO)
	return f
}
