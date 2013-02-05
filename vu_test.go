package vu

import "testing"

func TestCCW(t *testing.T) {
	p00 := Point{0, 0}
	p01, p10, p11 := Point{0, 1}, Point{1, 0}, Point{1, 1}
	p02, p22 := Point{0, 2}, Point{2, 2}
	if !CCW(p00, p10, p11) {
		t.Fail()
	}
	if CCW(p11, p10, p01) {
		t.Fail()
	}
	if CCW(p00, p01, p02) || CCW(p02, p01, p00) {
		t.Fail()
	}
	if CCW(p00, p11, p22) || CCW(p22, p11, p00) {
		t.Fail()
	}
	if CCW(p00, p22, p11) || CCW(p22, p00, p11) {
		t.Fail()
	}
}
