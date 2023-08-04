package repository

import (
	"fmt"
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Repository defines a interface for access the database.
type Repository interface {
	Model(value any) *gorm.DB
	Select(query any, args ...any) *gorm.DB
	Find(out any, where ...any) *gorm.DB
	Exec(sql string, values ...any) *gorm.DB
	First(out any, where ...any) *gorm.DB
	Raw(sql string, values ...any) *gorm.DB
	Create(value any) *gorm.DB
	Save(value any) *gorm.DB
	Updates(value any) *gorm.DB
	Delete(value any) *gorm.DB
	Where(query any, args ...any) *gorm.DB
	Preload(column string, conditions ...any) *gorm.DB
	Association(column string) *gorm.Association
	Close() error
	AutoMigrate(dst ...any) error
	GetDB() *gorm.DB
}

// repository defines a repository for access the database.
type repository struct {
	db *gorm.DB
}

// Model specify the models you would like to run db operations
func (r *repository) Model(value any) *gorm.DB {
	return r.db.Model(value)
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields
func (r *repository) Select(query any, args ...any) *gorm.DB {
	return r.db.Select(query, args)
}

// Find finds records that match given conditions
func (r *repository) Find(out any, where ...any) *gorm.DB {
	return r.db.Select(out, where...)
}

// Exec execs given SQL using by gorm.DB
func (r *repository) Exec(sql string, values ...any) *gorm.DB {
	return r.db.Exec(sql, values...)
}

// First returns the first record that match given conditions, order by primary key
func (r *repository) First(out any, where ...any) *gorm.DB {
	return r.db.First(out, where...)
}

// Raw returns the record that executed the given SQL using gorm.DB
func (r *repository) Raw(sql string, values ...any) *gorm.DB {
	return r.db.Raw(sql, values...)
}

// Create insert the value into database.
func (r *repository) Create(value any) *gorm.DB {
	return r.db.Create(value)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (r *repository) Save(value any) *gorm.DB {
	return r.db.Save(value)
}

// Updates update value in database
func (r *repository) Updates(value any) *gorm.DB {
	return r.db.Updates(value)
}

// Delete delete value match given conditions
func (r *repository) Delete(value any) *gorm.DB {
	return r.db.Delete(value)
}

// Where returns a new relation
func (r *repository) Where(query any, args ...any) *gorm.DB {
	return r.db.Where(query, args...)
}

// Preload preloads associations with given conditions
func (r *repository) Preload(column string, conditions ...any) *gorm.DB {
	return r.db.Preload(column, conditions...)
}

// Association relates relationship table
func (r *repository) Association(column string) *gorm.Association {
	return r.db.Association(column)
}

// Close closes current db connection. If database connection is not an io.Closer, returns an error
func (r *repository) Close() error {
	sqlDB, _ := r.db.DB()
	return sqlDB.Close()
}

// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func (r *repository) AutoMigrate(dst ...any) error {
	return r.db.AutoMigrate(dst...)
}

func (r *repository) GetDB() *gorm.DB {
	return r.db
}

// NewRepository is constructor for bookRepository.
func NewRepository(l logger.Logger, conf *config.Configuration) (Repository, error) {
	l.GetZapLogger().Infof("Try database connection")
	db, err := connectDatabase(l, conf)
	if err != nil {
		l.GetZapLogger().Errorf("Failure database connection")
		return nil, err
	}
	l.GetZapLogger().Infof("Success database connection, %s:%d", conf.Database.Host, conf.Database.Port)
	return &repository{db: db}, nil
}

const (
	// POSTGRES represents PostgresSQL
	POSTGRES = "postgres"
	// MYSQL represents MySQL
	MYSQL = "mysql"
	// SQLITE represents SQLite3
	SQLITE = "sqlite3"
)

func connectDatabase(l logger.Logger, conf *config.Configuration) (*gorm.DB, error) {
	gormConfig := &gorm.Config{Logger: l}

	var dsn string
	if conf.Database.Dialect == POSTGRES {
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Dbname,
			conf.Database.Username,
			conf.Database.Password,
		)
		return gorm.Open(postgres.Open(dsn), gormConfig)
	} else if conf.Database.Dialect == MYSQL {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=True&loc=Local&timeout=10000ms",
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Dbname,
			conf.Database.Charset,
			conf.Database.Collation,
		)
		return gorm.Open(mysql.Open(dsn), gormConfig)
	} else {
		return gorm.Open(sqlite.Open(conf.Database.Host), gormConfig)
	}
}
