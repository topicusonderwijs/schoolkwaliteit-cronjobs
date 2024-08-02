// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package mysql

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Account struct {
	ID            int32
	Guid          string
	Type          sql.NullString
	ContactID     sql.NullInt32
	Code          sql.NullString
	Name          sql.NullString
	ContactName   sql.NullString
	Subdomain     sql.NullString
	AuthorityName sql.NullString
	AuthorityCode sql.NullString
	Address1      sql.NullString
	Address2      sql.NullString
	Zipcode       sql.NullString
	City          sql.NullString
	Country       sql.NullString
	Email         sql.NullString
	Phone         sql.NullString
	Website       sql.NullString
	Logo          sql.NullString
	ConceptID     sql.NullInt32
	Settings      json.RawMessage
	Active        sql.NullBool
	CustomerID    sql.NullString
	ExpiredOn     sql.NullTime
	ExpireReason  sql.NullString
	DeletedAt     sql.NullTime
	UsersCount    int32
	PeopleCount   int32
	GroupsCount   int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Parnassys     bool
	Anonymized    bool
}
