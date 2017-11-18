package api

type IndexTables struct {
	Tables []TableIndex
}

// GET /api/ => { "Tables": [ ... ] }
// GET /api/mibs/ => [ { "Tables": [ ... ] } ]
// GET /api/mibs/:mib => { "Tables": [ ... ] }
// GET /api/hosts/:host/ => { "Tables": [ ... ] }
// GET /api/tables => { "Tables": [ ... ] }
type TableIndex struct {
	ID string

	IndexKeys  []string
	ObjectKeys []string
}

// GET /api/tables/ => [ { ... }, ... ]
// GET /api/tables/:table => { ... }
//
// GET /api/hosts/:host/tables/ => [ { ... }, ... ]
// GET /api/hosts/:host/tables/:table => { ... }
type Table struct {
	TableIndex

	Entries []TableEntry
	Error   *Error `json:",omitempty"`
}

type TableIndexMap map[string]interface{}
type TableObjectsMap map[string]interface{}

type TableEntry struct {
	HostID  string `json:",omitempty"` // XXX: always?
	Index   TableIndexMap
	Objects TableObjectsMap
}

// GET /api/tables/:table
// GET /api/hosts/:host/tables/:table
type TableQuery struct {
	Hosts []string `schema:"host"`
}

// GET /api/tables/
// GET /api/hosts/:host/tables/
type TablesQuery struct {
	Hosts  []string `schema:"host"`
	Tables []string `schema:"table"`
}
