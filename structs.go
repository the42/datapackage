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
	Title       *string `json:"title,omitempty"`
	Description *string `json:"discription,omitempty"`
	Format      *string `json:"format,omitempty"`
	Type        *string `json:"type,omitempty"`
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
	PrimaryKey *PKSpec     `json:"primaryKey,omitempty"`
}

type ISO8601 struct {
	Raw string
	time.Time
}

type Source struct {
	Name  *string `json:"name,omitempty"`
	Web   *string `json:"web,omitempty"`
	Email *string `json:"email,omitempty"`
}

type Resource struct {
	// BEGIN:one of the following is required
	Url  *string     `json:"url,omitempty"`
	Path *string     `json:"path,omitempty"`
	Data interface{} `json:"data,omitempty"`
	// END
	// Recommended fields
	Name *string `json:"name,omitempty"`
	// Optional fields
	Format    *string     `json:"format,omitempty"`
	Mediatype *string     `json:"mediatype,omitempty"`
	Bytes     *int64      `json:"bytes,omitempty"`
	Hash      *string     `json:"hash,omitempty"`
	Modified  *ISO8601    `json:"modified,omitempty"`
	Schema    interface{} `json:"schema,omitempty"` // TODO: Describe here
	Sources   []Source    `json:"sources,omitempty"`
	Licenses  []License   `json:"license,omitempty"`
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
	Title            *string     `json:"title,omitempty"`
	Description      *string     `json:"description,omitempty"`
	Homepage         *string     `json:"homepage,omitempty"`
	Version          *string     `json:"version,omitempty"`
	Sources          []Source    `json:"sources,omitempty"`
	Keywords         []string    `json:"keywords,omitempty"`
	Last_Modified    *ISO8601    `json:"last_modified,omitempty"`
	Image            *string     `json:"image,omitempty"`
	Maintainers      []Source    `json:"maintainers,omitempty"`
	Contributors     []Source    `json:"contributors,omitempty"`
	Publisher        []Source    `json:"publisher,omitempty"`
	Base             *string     `json:"base,omitempty"`
	DataDependencies interface{} `json:"dataDependencies,omitempty"`
}
