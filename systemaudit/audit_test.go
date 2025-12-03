package systemaudit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemAudit(t *testing.T) {
	result, err := DoSystemAudit()
	require.NoError(t, err)
	require.NotNil(t, result)
}
