package validation_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wallacemachado/api-bank-transfers/src/utils/validation"
)

func TestValidateCPF(t *testing.T) {
	for i, v := range []struct {
		name     string
		expected error
		value    string
	}{
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "445111023156132153615"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), ""},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "nfhtgbdfcbc"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "000.000.000-00"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "111.111.111-11"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "222.222.222-22"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "333.333.333-33"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "444.444.444-44"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "555.555.555-55"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "666666.666-66"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "777.777.777-77"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "88888888888"},
		{"Data invalid - Should Return error: invalid CPF", errors.New("invalid CPF"), "999.999.999-99"},
		{"Invalid Digits - Should Return error: invalid CPF", errors.New("invalid CPF"), "797.674.610-34"},
		{"Invalid Digits - Should Return error: invalid CPF", errors.New("invalid CPF"), "797.674.610-25"},
		{"sucess", nil, "797.674.610-35"},
		{"sucess", nil, "79767461035"},
		{"sucess", nil, "  797.674.610-35   "},
	} {

		t.Run(fmt.Sprintf("% d - % s", i, v.name), func(t *testing.T) {
			_, err := validation.ValidateCPF(v.value)
			assert.Equal(t, err, v.expected)

		})
	}
}
