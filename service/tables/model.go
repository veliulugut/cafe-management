package tables

type TableModel struct {
	NumberOfGuests int    `json:"number_of_guests"`
	TableNumber    int    `json:"table_number"`
	Description    string `json:"description"`
	TableType      string `json:"table_type"`
}

type CreateTableModel struct {
	NumberOfGuests int    `json:"number_of_guests"`
	TableNumber    int    `json:"table_number"`
	Description    string `json:"description"`
	TableType      string `json:"table_type"`
}

type UpdateTableModel struct {
	NumberOfGuests int    `json:"number_of_guests"`
	TableNumber    int    `json:"table_number"`
	Description    string `json:"description"`
	TableType      string `json:"table_type"`
}
