package field

type operators string

func (o operators) String() string {
	return string(o)
}

// StringWithSpace return string with space in start
func (o operators) StringWithSpace() string {
	return " " + o.String()
}

const (
	EQ        operators = "="
	NEQ       operators = "!="
	GT        operators = ">"
	GTE       operators = ">="
	LT        operators = "<"
	LTE       operators = "<="
	IN        operators = "IN"
	NIN       operators = "NOT IN"
	Like      operators = "LIKE"
	NotLike   operators = "NOT LIKE"
	IsNull    operators = "IS NULL"
	IsNotNull operators = "IS NOT NULL"
)

type operationOptions uint

const (
	operationInBrackets operationOptions = 1 << iota
)

type operator struct {
	operator operators

	opts operationOptions
}

func newOperator(op operators, opts operationOptions) operator {
	return operator{
		operator: op,
		opts:     opts,
	}
}

func (o operator) applyOption(s string) string {
	if o.operator == NIN || o.operator == IN {
		return s + "(?)"
	}

	if o.operator == IsNull || o.operator == IsNotNull {
		return s
	}

	switch o.opts {
	case operationInBrackets:
		return s + " (?)"
	}

	return s + " ?"
}

func (o operator) String() string {
	return o.applyOption(o.operator.String())
}

func (o operator) StringWithSpace() string {
	return o.applyOption(o.operator.StringWithSpace())
}
