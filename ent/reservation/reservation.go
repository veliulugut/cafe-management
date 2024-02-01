// Code generated by ent, DO NOT EDIT.

package reservation

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the reservation type in the database.
	Label = "reservation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTableID holds the string denoting the table_id field in the database.
	FieldTableID = "table_id"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// Table holds the table name of the reservation in the database.
	Table = "reservations"
)

// Columns holds all SQL columns for reservation fields.
var Columns = []string{
	FieldID,
	FieldFullName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTableID,
	FieldPhoneNumber,
	FieldStartTime,
	FieldEndTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime time.Time
	// DefaultEndTime holds the default value on creation for the "end_time" field.
	DefaultEndTime time.Time
)

// OrderOption defines the ordering options for the Reservation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTableID orders the results by the table_id field.
func ByTableID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTableID, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}
