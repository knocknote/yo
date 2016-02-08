// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// Column represents PostgreSQL class (ie, table, view, etc) attributes.
type Column struct {
	ColumnName       string // column_name
	TableName        string // table_name
	DataType         string // data_type
	FieldOrdinal     int    // field_ordinal
	IsNullable       bool   // is_nullable
	IsIndex          bool   // is_index
	IsUnique         bool   // is_unique
	IsPrimaryKey     bool   // is_primary_key
	IsForeignKey     bool   // is_foreign_key
	IndexName        string // index_name
	ForeignIndexName string // foreign_index_name
	HasDefault       bool   // has_default
	DefaultValue     string // default_value
	Field            string // field
	Type             string // type
	NilType          string // nil_type
	Tag              string // tag
	Len              int    // len
	Comment          string // comment
}

// ColumnsByRelkindSchema runs a custom query, returning results as Column.
func ColumnsByRelkindSchema(db XODB, relkind string, schema string) ([]*Column, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.attname, ` + // ::varchar AS column_name
		`c.relname, ` + // ::varchar AS table_name
		`format_type(a.atttypid, a.atttypmod), ` + // ::varchar AS data_type
		`a.attnum, ` + // ::integer AS field_ordinal
		`(NOT a.attnotnull), ` + // ::boolean AS is_nullable
		`COALESCE(i.oid <> 0, false), ` + // ::boolean AS is_index
		`COALESCE(ct.contype = 'u' OR ct.contype = 'p', false), ` + // ::boolean AS is_unique
		`COALESCE(ct.contype = 'p', false), ` + // ::boolean AS is_primary_key
		`COALESCE(cf.contype = 'f', false), ` + // ::boolean AS is_foreign_key
		`COALESCE(i.relname, ''), ` + // ::varchar AS index_name
		`COALESCE(cf.conname, ''), ` + // ::varchar AS foreign_index_name
		`a.atthasdef, ` + // ::boolean AS has_default
		`COALESCE(pg_get_expr(ad.adbin, ad.adrelid), ''), ` + // ::varchar AS default_value
		`'', ` + // ::varchar AS field
		`'', ` + // ::varchar AS type
		`'', ` + // ::varchar AS nil_type
		`'', ` + // ::varchar AS tag
		`0, ` + // ::integer AS len
		`'' ` + // ::varchar AS comment
		`FROM pg_attribute a ` +
		`JOIN ONLY pg_class c ON c.oid = a.attrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`LEFT JOIN pg_constraint ct ON ct.conrelid = c.oid AND a.attnum = ANY(ct.conkey) AND ct.contype IN('p', 'u') ` +
		`LEFT JOIN pg_constraint cf ON cf.conrelid = c.oid AND a.attnum = ANY(cf.conkey) AND cf.contype IN('f') ` +
		`LEFT JOIN pg_attrdef ad ON ad.adrelid = c.oid AND ad.adnum = a.attnum ` +
		`LEFT JOIN pg_index ix ON a.attnum = ANY(ix.indkey) AND c.oid = a.attrelid AND c.oid = ix.indrelid ` +
		`LEFT JOIN pg_class i ON i.oid = ix.indexrelid ` +
		`WHERE c.relkind = $1 AND a.attnum > 0 AND n.nspname = $2 ` +
		`ORDER BY c.relname, a.attnum`

	// run query
	q, err := db.Query(sqlstr, relkind, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Column{}
	for q.Next() {
		c := Column{}

		// scan
		err = q.Scan(&c.ColumnName, &c.TableName, &c.DataType, &c.FieldOrdinal, &c.IsNullable, &c.IsIndex, &c.IsUnique, &c.IsPrimaryKey, &c.IsForeignKey, &c.IndexName, &c.ForeignIndexName, &c.HasDefault, &c.DefaultValue, &c.Field, &c.Type, &c.NilType, &c.Tag, &c.Len, &c.Comment)
		if err != nil {
			return nil, err
		}

		res = append(res, &c)
	}

	return res, nil
}
