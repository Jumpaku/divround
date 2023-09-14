package divround_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	dr "github.com/Jumpaku/divround"
)

func TestErrOverflow_Error(t *testing.T) {
	sut := dr.NewOverflowError("error", 0, 0)
	got := sut.Error()
	want := `overflow`
	if !strings.Contains(strings.ToLower(got), want) {
		t.Errorf(`"overflow" not contained`)
	}
}

func TestErrOverflow_Is(t *testing.T) {
	sut := dr.NewOverflowError("error", 0, 0)
	t.Run(`simple`, func(t *testing.T) {
		if !errors.Is(sut, dr.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
	t.Run(`joined`, func(t *testing.T) {
		if !errors.Is(errors.Join(sut, errors.New(`some error`)), dr.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
	t.Run(`wrapped`, func(t *testing.T) {
		if !errors.Is(fmt.Errorf(`some error : %w`, sut), dr.ErrOverflow) {
			t.Errorf(`is not div.ErrOverflow`)
		}
	})
}

func TestErrZeroDivision_Error(t *testing.T) {
	sut := dr.NewZeroDivisionError("error", 0, 0)
	got := sut.Error()
	if !strings.Contains(strings.ToLower(got), `division`) {
		t.Errorf(`"division" not contained`)
	}
	if !strings.Contains(strings.ToLower(got), `zero`) {
		t.Errorf(`"zero" not contained`)
	}
}

func TestErrZeroDivision_Is(t *testing.T) {
	sut := dr.NewZeroDivisionError("error", 0, 0)
	t.Run(`simple`, func(t *testing.T) {
		if !errors.Is(sut, dr.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
	t.Run(`joined`, func(t *testing.T) {
		if !errors.Is(errors.Join(sut, errors.New(`some error`)), dr.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
	t.Run(`wrapped`, func(t *testing.T) {
		if !errors.Is(fmt.Errorf(`some error : %w`, sut), dr.ErrZeroDivision) {
			t.Errorf(`is not div.ErrZeroDivision`)
		}
	})
}
