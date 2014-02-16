package datapackage

import "time"

// Definition according to http://dataprotocols.org/csv-dialect/
type CSVDialect struct {
	Header           bool   `json:"header"`
	Delimiter        string `json:"delimiter"`
	Doublequote      bool   `json:"doublequote"`
	LineTerminator   string `json:"lineTerminator"`
	QuoteChar        string `json:"quoteChar"`
	SkipInitialSpace bool   `json:"skipInitialSpace"`
}

// Definition according to http://dataprotocols.org/csv-dialect/
type CSVDDF struct {
	CsvddfVersion *string    `json:"csvddfVersion"`
	Dialect       CSVDialect `json:"dialect"`
}

// JSONTableSchema Field-specification as of http://dataprotocols.org/json-table-schema/
type FieldSpec struct {
	// Required
	Name *string `json:"name"`
	// Optional
	Title       *string `json:"title"`
	Description *string `json:"discription"`
	Format      *string `json:"format"`
	Type        *string `json:"type"`
}

var JSONTableTypes = []string{
	"string",
	"number",
	"integer",
	"date",
	"time",
	"datetime",
	"boolean",
	"binary",
	"object",
	"json",
	"geopoint",
	"geojson",
	"array",
	"any",
}

// Either: an array of strings or  a single string
type PKSpec struct {
	Keys []string
}

// JSONTableSchema specification as of http://dataprotocols.org/json-table-schema/
type JSONTableSchema struct {
	Fields     []FieldSpec `json:"fields"`
	PrimaryKey *PKSpec     `json:"primaryKey"`
}

type ISO8601 struct {
	Raw string
	time.Time
}

type Source struct {
	Name  *string `json:"name"`
	Web   *string `json:"web"`
	Email *string `json:"email"`
}

type Resource struct {
	// BEGIN:one of the following is required
	Url  *string     `json:"url"`
	Path *string     `json:"path"`
	Data interface{} `json:"data"`
	// END
	// Recommended fields
	Name *string `json:"name"`
	// Optional fields
	Format    *string     `json:"format"`
	Mediatype *string     `json:"mediatype"`
	Bytes     *int64      `json:"bytes"`
	Hash      *string     `json:"hash"`
	Modified  *ISO8601    `json:"modified"`
	Schema    interface{} `json:"schema"` // TODO: Describe here
	Sources   []Source    `json:"sources"`
	Licenses  []License   `json:"license"`
}

type License struct {
	ID  *string `json:"id"`
	Url *string `json:"url"`
}

// definition of a datapackage as specified in http://dataprotocols.org/data-packages/
type Datapackage struct {
	// Required Fields
	Name *string `json:"name"`
	// Fields which should show up
	Resources           []Resource `json:"resources"`
	Licenses            []License  `json:"license"`
	Datapackage_Version *string    `json:"datapackage_version"`
	// Recommended Fields
	Title            *string     `json:"title"`
	Description      *string     `json:"description"`
	Homepage         *string     `json:"homepage"`
	Version          *string     `json:"version"`
	Sources          []Source    `json:"sources"`
	Keywords         []string    `json:"keywords"`
	Last_Modified    *ISO8601    `json:"last_modified"`
	Image            *string     `json:"image"`
	Maintainers      []Source    `json:"maintainers"`
	Contributors     []Source    `json:"contributors"`
	Publisher        []Source    `json:"publisher"`
	Base             *string     `json:"base"`
	DataDependencies interface{} `json:"dataDependencies"`
}
