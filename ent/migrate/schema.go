// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PricesColumns holds the columns for the "prices" table.
	PricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "price_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PricesTable holds the schema information for the "prices" table.
	PricesTable = &schema.Table{
		Name:       "prices",
		Columns:    PricesColumns,
		PrimaryKey: []*schema.Column{PricesColumns[0]},
	}
	// ReservationsColumns holds the columns for the "reservations" table.
	ReservationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "full_name", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "table_id", Type: field.TypeInt},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "tables_reservation", Type: field.TypeInt, Nullable: true},
	}
	// ReservationsTable holds the schema information for the "reservations" table.
	ReservationsTable = &schema.Table{
		Name:       "reservations",
		Columns:    ReservationsColumns,
		PrimaryKey: []*schema.Column{ReservationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reservations_tables_reservation",
				Columns:    []*schema.Column{ReservationsColumns[8]},
				RefColumns: []*schema.Column{TablesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TablesColumns holds the columns for the "tables" table.
	TablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number_of_guests", Type: field.TypeInt},
		{Name: "table_number", Type: field.TypeInt},
		{Name: "description", Type: field.TypeString},
		{Name: "table_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TablesTable holds the schema information for the "tables" table.
	TablesTable = &schema.Table{
		Name:       "tables",
		Columns:    TablesColumns,
		PrimaryKey: []*schema.Column{TablesColumns[0]},
	}
	// TablesTypesColumns holds the columns for the "tables_types" table.
	TablesTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tables_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
	}
	// TablesTypesTable holds the schema information for the "tables_types" table.
	TablesTypesTable = &schema.Table{
		Name:       "tables_types",
		Columns:    TablesTypesColumns,
		PrimaryKey: []*schema.Column{TablesTypesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PricesTable,
		ReservationsTable,
		TablesTable,
		TablesTypesTable,
		UsersTable,
	}
)

func init() {
	ReservationsTable.ForeignKeys[0].RefTable = TablesTable
}
