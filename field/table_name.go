package field

type TableName string

func (t TableName) String() string {
	return string(t)
}
