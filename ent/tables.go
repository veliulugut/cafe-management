// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/tables"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Tables is the model entity for the Tables schema.
type Tables struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NumberOfGuests holds the value of the "number_of_guests" field.
	NumberOfGuests int `json:"number_of_guests,omitempty"`
	// TableNumber holds the value of the "table_number" field.
	TableNumber int `json:"table_number,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// TableType holds the value of the "table_type" field.
	TableType string `json:"table_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TablesQuery when eager-loading is set.
	Edges        TablesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TablesEdges holds the relations/edges for other nodes in the graph.
type TablesEdges struct {
	// Reservation holds the value of the reservation edge.
	Reservation []*Reservation `json:"reservation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReservationOrErr returns the Reservation value or an error if the edge
// was not loaded in eager-loading.
func (e TablesEdges) ReservationOrErr() ([]*Reservation, error) {
	if e.loadedTypes[0] {
		return e.Reservation, nil
	}
	return nil, &NotLoadedError{edge: "reservation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tables) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tables.FieldID, tables.FieldNumberOfGuests, tables.FieldTableNumber:
			values[i] = new(sql.NullInt64)
		case tables.FieldDescription, tables.FieldTableType:
			values[i] = new(sql.NullString)
		case tables.FieldCreatedAt, tables.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tables fields.
func (t *Tables) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tables.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tables.FieldNumberOfGuests:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number_of_guests", values[i])
			} else if value.Valid {
				t.NumberOfGuests = int(value.Int64)
			}
		case tables.FieldTableNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field table_number", values[i])
			} else if value.Valid {
				t.TableNumber = int(value.Int64)
			}
		case tables.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case tables.FieldTableType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_type", values[i])
			} else if value.Valid {
				t.TableType = value.String
			}
		case tables.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tables.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tables.
// This includes values selected through modifiers, order, etc.
func (t *Tables) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryReservation queries the "reservation" edge of the Tables entity.
func (t *Tables) QueryReservation() *ReservationQuery {
	return NewTablesClient(t.config).QueryReservation(t)
}

// Update returns a builder for updating this Tables.
// Note that you need to call Tables.Unwrap() before calling this method if this Tables
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tables) Update() *TablesUpdateOne {
	return NewTablesClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tables entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tables) Unwrap() *Tables {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tables is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tables) String() string {
	var builder strings.Builder
	builder.WriteString("Tables(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("number_of_guests=")
	builder.WriteString(fmt.Sprintf("%v", t.NumberOfGuests))
	builder.WriteString(", ")
	builder.WriteString("table_number=")
	builder.WriteString(fmt.Sprintf("%v", t.TableNumber))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("table_type=")
	builder.WriteString(t.TableType)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TablesSlice is a parsable slice of Tables.
type TablesSlice []*Tables
