package div_test

import (
	"testing"
)

type testcaseDiv struct {
	a, b int64
	want int64
}

type testcaseDivErr struct {
	a, b int64
	want int64
	err  error
}

func doTestDiv(t *testing.T, sut func(int64, int64) int64, tc testcaseDiv) {
	t.Helper()
	got := sut(tc.a, tc.b)
	if got != tc.want {
		t.Errorf("Want: %v, Got: %v", tc.want, got)
	}
}

func doTestDivErr(t *testing.T, sut func(int64, int64) int64 /*,err*/, tc testcaseDivErr) {
	t.Helper()
	defer func() {
		_ = recover()
	}()
	got /*, err*/ := sut(tc.a, tc.b)
	/*if err != nil {
		if !errors.Is(err, tc.err) {
			t.Errorf("Want err: %v, Got err: %v", tc.err, err)
		}
	} else */if got != tc.want {
		t.Errorf("Want: %v, Got: %v", tc.want, got)
	}
}
