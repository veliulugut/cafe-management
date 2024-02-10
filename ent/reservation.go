// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/reservation"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Reservation is the model entity for the Reservation schema.
type Reservation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TableID holds the value of the "table_id" field.
	TableID int `json:"table_id,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime holds the value of the "end_time" field.
	EndTime            time.Time `json:"end_time,omitempty"`
	tables_reservation *int
	selectValues       sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reservation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID, reservation.FieldTableID:
			values[i] = new(sql.NullInt64)
		case reservation.FieldFullName, reservation.FieldPhoneNumber, reservation.FieldStatus:
			values[i] = new(sql.NullString)
		case reservation.FieldCreatedAt, reservation.FieldUpdatedAt, reservation.FieldStartTime, reservation.FieldEndTime:
			values[i] = new(sql.NullTime)
		case reservation.ForeignKeys[0]: // tables_reservation
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reservation fields.
func (r *Reservation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case reservation.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				r.FullName = value.String
			}
		case reservation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reservation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reservation.FieldTableID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field table_id", values[i])
			} else if value.Valid {
				r.TableID = int(value.Int64)
			}
		case reservation.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				r.PhoneNumber = value.String
			}
		case reservation.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = value.String
			}
		case reservation.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				r.StartTime = value.Time
			}
		case reservation.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				r.EndTime = value.Time
			}
		case reservation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tables_reservation", value)
			} else if value.Valid {
				r.tables_reservation = new(int)
				*r.tables_reservation = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reservation.
// This includes values selected through modifiers, order, etc.
func (r *Reservation) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Reservation.
// Note that you need to call Reservation.Unwrap() before calling this method if this Reservation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reservation) Update() *ReservationUpdateOne {
	return NewReservationClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reservation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reservation) Unwrap() *Reservation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reservation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reservation) String() string {
	var builder strings.Builder
	builder.WriteString("Reservation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("full_name=")
	builder.WriteString(r.FullName)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("table_id=")
	builder.WriteString(fmt.Sprintf("%v", r.TableID))
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(r.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(r.Status)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(r.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(r.EndTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reservations is a parsable slice of Reservation.
type Reservations []*Reservation
