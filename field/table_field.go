package field

import "fmt"

// FieldWithTable return string with table name and field name
//
// Example:
//
//	field := `"Id"`
//	table := "EventTypes"
//	str := FieldWithTable(table, field)
//	print(str)
//
// Output:
//
//	`"EventTypes"."Id"`
func FieldWithTable(table string, field string) string {
	return fmt.Sprintf(`"%s".%s`, table, field)
}

type (
	TableField interface {
		WithTable() string
		String() string
	}

	TableFieldWithSubQuery interface {
		TableField

		// add brackets to operator
		SubQuery() TableField
	}

	TableFieldWithOperator interface {
		TableField
		EQ() TableFieldWithSubQuery
		NEQ() TableFieldWithSubQuery
		GT() TableFieldWithSubQuery
		GTE() TableFieldWithSubQuery
		LT() TableFieldWithSubQuery
		LTE() TableFieldWithSubQuery
		IN() TableFieldWithSubQuery
		NIN() TableFieldWithSubQuery
		Like() TableFieldWithSubQuery
		NotLike() TableFieldWithSubQuery
		IsNull() TableFieldWithSubQuery
		IsNotNull() TableFieldWithSubQuery
	}

	OperatorOpts interface {
		inBrackets() bool
	}
)

type tableField struct {
	field string
	table TableName
}

func (t tableField) String() string {
	return t.field
}

func (t tableField) WithTable() string {
	return FieldWithTable(t.table.String(), t.field)
}

type tableFieldWithOperator struct {
	tableField
	operator operator
}

func (t tableFieldWithOperator) String() string {
	if t.operator.operator == "" {
		return t.tableField.String()
	}

	return t.tableField.String() + t.operator.StringWithSpace()
}

func (t tableFieldWithOperator) WithTable() string {
	if t.operator.operator == "" {
		return t.tableField.WithTable()
	}

	return t.tableField.WithTable() + t.operator.StringWithSpace()
}

func (t tableFieldWithOperator) SubQuery() TableField {
	t.operator.opts = operationInBrackets
	return t
}

func (t tableFieldWithOperator) EQ() TableFieldWithSubQuery {
	t.operator.operator = EQ
	return t
}

func (t tableFieldWithOperator) NEQ() TableFieldWithSubQuery {
	t.operator.operator = NEQ
	return t
}

func (t tableFieldWithOperator) GT() TableFieldWithSubQuery {
	t.operator.operator = GT
	return t
}

func (t tableFieldWithOperator) GTE() TableFieldWithSubQuery {
	t.operator.operator = GTE
	return t
}

func (t tableFieldWithOperator) LT() TableFieldWithSubQuery {
	t.operator.operator = LT
	return t
}

func (t tableFieldWithOperator) LTE() TableFieldWithSubQuery {
	t.operator.operator = LTE
	return t
}

func (t tableFieldWithOperator) IN() TableFieldWithSubQuery {
	t.operator.operator = IN
	return t
}

func (t tableFieldWithOperator) NIN() TableFieldWithSubQuery {
	t.operator.operator = NIN
	return t
}

func (t tableFieldWithOperator) Like() TableFieldWithSubQuery {
	t.operator.operator = Like
	return t
}

func (t tableFieldWithOperator) NotLike() TableFieldWithSubQuery {
	t.operator.operator = NotLike
	return t
}

func (t tableFieldWithOperator) IsNull() TableFieldWithSubQuery {
	t.operator.operator = IsNull
	return t
}

func (t tableFieldWithOperator) IsNotNull() TableFieldWithSubQuery {
	t.operator.operator = IsNotNull
	return t
}

func NewTableField(table TableName, field string) TableFieldWithOperator {
	return tableFieldWithOperator{
		tableField: tableField{
			field: field,
			table: table,
		},
		operator: operator{},
	}
}
