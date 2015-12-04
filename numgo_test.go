package numgo

import "testing"

func TestCreate(t *testing.T) {
	shp := []int{2, 3, 4}
	a := Create(shp...)
	if len(a.data) != 24 {
		t.Logf("Length %d, expected %d", len(a.data), 24)
		t.FailNow()
	}

	for _, v := range a.data {
		if v != 0 {
			t.Logf("Value %f, expected %d", v, 0)
		}
	}

	a = Full(1, shp...)
	if len(a.data) != 24 {
		t.Logf("Length %d, expected %d\n", len(a.data), 24)
		t.FailNow()
	}

	for _, v := range a.data {
		if v != 1 {
			t.Logf("Value %f, expected %d\n", v, 1)
			t.FailNow()
		}
	}
}

func TestShapes(t *testing.T) {
	shp := []int{3, 3, 4, 7}
	a := Create(shp...)
	for i, v := range a.shape {
		if uint64(shp[i]) != v {
			t.Log(a.shape, "!=", shp)
			t.FailNow()
		}
	}
}

func TestArange(t *testing.T) {
	a := Arange(24)
	if len(a.data) != 24 {
		t.Logf("Length %d.  Expected size %d\n", len(a.data), 24)
	}
	if len(a.shape) != 1 {
		t.Logf("Axis %d.  Expected %d\n", len(a.shape), 1)
	}
	for i, v := range a.data {
		if float64(i) != v {
			t.Logf("Value %f.  Expected %d\n", v, i)
		}
	}
}