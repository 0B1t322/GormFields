package field_test

import (
	"testing"

	"github.com/0B1t322/GormFields/field"
	"github.com/stretchr/testify/require"
)

func TestFunc_TableField(t *testing.T) {
	tableField := field.NewTableField(
		field.TableName("Table"),
		`"Field"`,
	)

	require.Equal(
		t,
		`"Field" = ?`,
		tableField.EQ().String(),
	)

	require.Equal(
		t,
		`"Table"."Field" = ?`,
		tableField.EQ().WithTable(),
	)
}
