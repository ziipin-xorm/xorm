package core

import (
	"fmt"
	"strings"
	"time"
)

type DbType string

type Uri struct {
	DbType  DbType
	Proto   string
	Host    string
	Port    string
	DbName  string
	User    string
	Passwd  string
	Charset string
	Laddr   string
	Raddr   string
	Timeout time.Duration
}

// a dialect is a driver's wrapper
type Dialect interface {
	Init(*Uri, string, string) error
	URI() *Uri
	DBType() DbType
	SqlType(t *Column) string
	SupportInsertMany() bool
	QuoteStr() string
	AndStr() string
	EqStr() string
	RollBackStr() string
	DropTableSql(tableName string) string
	AutoIncrStr() string
	SupportEngine() bool
	SupportCharset() bool
	IndexOnTable() bool
	ShowCreateNull() bool

	IndexCheckSql(tableName, idxName string) (string, []interface{})
	TableCheckSql(tableName string) (string, []interface{})
	ColumnCheckSql(tableName, colName string) (string, []interface{})

	GetColumns(tableName string) ([]string, map[string]*Column, error)
	GetTables() ([]*Table, error)
	GetIndexes(tableName string) (map[string]*Index, error)

	CreateTableSql(table *Table, tableName, storeEngine, charset string) string
	Filters() []Filter

	DriverName() string
	DataSourceName() string
}

type Base struct {
	dialect        Dialect
	driverName     string
	dataSourceName string
	*Uri
}

func (b *Base) Init(dialect Dialect, uri *Uri, drivername, dataSourceName string) error {
	b.dialect = dialect
	b.driverName, b.dataSourceName = drivername, dataSourceName
	b.Uri = uri
	return nil
}

func (b *Base) URI() *Uri {
	return b.Uri
}

func (b *Base) DBType() DbType {
	return b.Uri.DbType
}

func (b *Base) DriverName() string {
	return b.driverName
}

func (b *Base) ShowCreateNull() bool {
	return true
}

func (b *Base) DataSourceName() string {
	return b.dataSourceName
}

func (b *Base) Quote(c string) string {
	return b.dialect.QuoteStr() + c + b.dialect.QuoteStr()
}

func (b *Base) AndStr() string {
	return "AND"
}

func (b *Base) EqStr() string {
	return "="
}

func (db *Base) RollBackStr() string {
	return "ROLL BACK"
}

func (db *Base) DropTableSql(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`", tableName)
}

func (b *Base) CreateTableSql(table *Table, tableName, storeEngine, charset string) string {
	var sql string
	sql = "CREATE TABLE IF NOT EXISTS "
	if tableName == "" {
		tableName = table.Name
	}

	sql += b.Quote(tableName) + " ("

	pkList := table.PrimaryKeys

	for _, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		if col.IsPrimaryKey && len(pkList) == 1 {
			sql += col.String(b.dialect)
		} else {
			sql += col.StringNoPk(b.dialect)
		}
		sql = strings.TrimSpace(sql)
		sql += ", "
	}

	if len(pkList) > 1 {
		sql += "PRIMARY KEY ( "
		sql += b.Quote(strings.Join(pkList, b.Quote(",")))
		sql += " ), "
	}

	sql = sql[:len(sql)-2] + ")"
	if b.dialect.SupportEngine() && storeEngine != "" {
		sql += " ENGINE=" + storeEngine
	}
	if b.dialect.SupportCharset() {
		if charset == "" {
			charset = b.dialect.URI().Charset
		}
		sql += " DEFAULT CHARSET " + charset
	}
	sql += ";"
	return sql
}

var (
	dialects = map[DbType]Dialect{}
)

func RegisterDialect(dbName DbType, dialect Dialect) {
	if dialect == nil {
		panic("core: Register dialect is nil")
	}
	dialects[dbName] = dialect // !nashtsai! allow override dialect
}

func QueryDialect(dbName DbType) Dialect {
	return dialects[dbName]
}