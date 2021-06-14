package store

// import (
// 	"database/sql"

// 	"github.com/Erwin011895/shorty-challenge/internal/component"
// )

// type store struct {
// 	db *component.DB
// }

// // Store general store
// type Store interface {
// 	Master() *sql.DB
// 	Slave() *sql.DB
// 	Begin() (tx *sql.Tx, err error)
// 	Commit(tx *sql.Tx) error
// 	Rollback(tx *sql.Tx) error
// }

// // NewStore init general store
// func NewStore(db *component.DB) Store {
// 	return &store{
// 		db: db,
// 	}
// }

// // Master return master
// func (s *store) Master() *sql.DB {
// 	return s.db.Master
// }

// // Slave return slave
// func (s *store) Slave() *sql.DB {
// 	return s.db.Slave
// }

// // Begin begin transaction
// func (s *store) Begin() (tx *sql.Tx, err error) {
// 	return s.db.Master.Begin()
// }

// // Commit commit transaction
// func (s *store) Commit(tx *sql.Tx) error {
// 	return tx.Commit()
// }

// // Rollback rollback transaction
// func (s *store) Rollback(tx *sql.Tx) error {
// 	return tx.Rollback()
// }
