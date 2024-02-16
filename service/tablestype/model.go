package tablestype

type CreateTableTypeModel struct {
	TablesID int    `json:"tables_id"`
	Name     string `json:"name"`
}

type UpdateTableTypeModel struct {
	Name     string `json:"name"`
	TablesID int    `json:"tables_id"`
}
