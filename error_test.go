package div_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/Jumpaku/go-div"
)

func TestErrOverflow_Error(t *testing.T) {
	sut := div.NewOverflowError("error", 0, 0)
	got := sut.Error()
	want := `overflow`
	if !strings.Contains(strings.ToLower(got), want) {
		t.Errorf(`"overflow" not contained`)
	}
}

func TestErrOverflow_Is(t *testing.T) {
	sut := div.NewOverflowError("error", 0, 0)
	t.Run(`simple`, func(t *testing.T) {
		if !errors.Is(sut, div.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
	t.Run(`joined`, func(t *testing.T) {
		if !errors.Is(errors.Join(sut, errors.New(`some error`)), div.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
	t.Run(`wrapped`, func(t *testing.T) {
		if !errors.Is(fmt.Errorf(`some error : %w`, sut), div.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
}

func TestErrZeroDivision_Error(t *testing.T) {
	sut := div.NewZeroDivisionError("error", 0, 0)
	got := sut.Error()
	if !strings.Contains(strings.ToLower(got), `division`) {
		t.Errorf(`"division" not contained`)
	}
	if !strings.Contains(strings.ToLower(got), `zero`) {
		t.Errorf(`"zero" not contained`)
	}
}

func TestErrZeroDivision_Is(t *testing.T) {
	sut := div.NewZeroDivisionError("error", 0, 0)
	t.Run(`simple`, func(t *testing.T) {
		if !errors.Is(sut, div.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
	t.Run(`joined`, func(t *testing.T) {
		if !errors.Is(errors.Join(sut, errors.New(`some error`)), div.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
	t.Run(`wrapped`, func(t *testing.T) {
		if !errors.Is(fmt.Errorf(`some error : %w`, sut), div.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
}
