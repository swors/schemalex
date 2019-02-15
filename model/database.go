package model

// NewDatabase creates a new database mode with th given name
func NewDatabase(n string) Database {
	return &database{
		name: n,
	}
}

func (d *database) isDatabase() bool {
	return true
}

func (d *database) ID() string {
	return "database#" + d.name
}

func (d *database) Name() string {
	return d.name
}

func (d *database) IsIfNotExists() bool {
	return d.ifnotexists
}

func (d *database) SetIfNotExists(v bool) Database {
	d.ifnotexists = v
	return d
}

func (t *database) AddOption(v DatabaseOption) Database {
	t.options = append(t.options, v)
	return t
}

func (t *database) Options() chan DatabaseOption {
	ch := make(chan DatabaseOption, len(t.options))
	for _, idx := range t.options {
		ch <- idx
	}
	close(ch)
	return ch
}

// NewDatabaseOption creates a new database option with the given name, value, and a flag indicating if quoting is necessary
func NewDatabaseOption(k, v string, q bool) DatabaseOption {
	return &databaseopt{
		key:        k,
		value:      v,
		needQuotes: q,
	}
}

func (t *databaseopt) ID() string       { return "databaseopt#" + t.key }
func (t *databaseopt) Key() string      { return t.key }
func (t *databaseopt) Value() string    { return t.value }
func (t *databaseopt) NeedQuotes() bool { return t.needQuotes }
