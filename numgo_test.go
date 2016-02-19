package numgo

import "testing"

func init() {
	debug = true
}

func TestNewArray64(t *testing.T) {
	shp := []int{2, 3, 4}
	a := NewArray64(nil, shp...)
	if len(a.data) != 24 {
		t.Logf("Length %d, expected %d", len(a.data), 24)
		t.FailNow()
	}

	for _, v := range a.data {
		if v != 0 {
			t.Logf("Value %f, expected %d", v, 0)
			t.Fail()
		}
	}
	a = NewArray64(nil)
	if e := a.GetErr(); e != nil {
		t.Log("Unexpected error:", e)
		t.Fail()
	}

	a = NewArray64([]float64{0, 1, 2, 3, 4})
	if e := a.Equals(Arange(5)); !e.All().At(0) {
		t.Log("Slice Assignment Failed", a.GetErr(), e)
		t.Fail()
	}

	a = NewArray64([]float64{0, 1, 2, 3, 4}, 3)
	if e := a.Equals(Arange(3)); !e.All().At(0) {
		t.Log("Slice Assignment Failed", a.GetErr(), e)
		t.Fail()
	}

	a = NewArray64([]float64{0, 1, 2, 3, 4, 5}, 2, -1, 3)
	if e := a.GetErr(); e != NegativeAxis {
		t.Log("Expected NegativeAxis, got:", e)
		t.Fail()
	}

	a = NewArray64(nil, 1, 2, 5, 9)
	if e := a.Equals(newArray64(1, 2, 5, 9)); !e.All().At(0) {
		t.Log("Creation has different results:", e)
		t.Fail()
	}
}
func TestFull(t *testing.T) {
	shp := []int{2, 3, 4}
	a := Full(1, shp...)
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
	a := NewArray64(nil, shp...)
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

func TestIdent(t *testing.T) {
	tmp := Identity(0)
	if len(tmp.shape) != 2 {
		t.Log("Incorrect identity shape.", tmp.shape)
		t.Fail()
	}
	if tmp.shape[0] != 0 || tmp.shape[1] != 0 {
		t.Log("Incorrect shape values.", tmp.shape)
		t.Fail()
	}
	if len(tmp.data) > 0 {
		t.Log("Data array incorrect.", tmp.data)
		t.Fail()
	}
}

func TestSubArray(t *testing.T) {
	a := Arange(100).Reshape(2, 5, 10)
	b := Arange(50).Reshape(5, 10)
	c := a.SubArr(0)
	if !c.Equals(b).All().At(0) {
		t.Log("Subarray incorrect. Expected\n", b, "\nReceived\n", c)
		t.Fail()
	}

	b = Arange(50, 100).Reshape(5, 10)
	c = a.SubArr(1)
	if !c.Equals(b).All().At(0) {
		t.Log("Subarray incorrect. Expected\n", b, "\nReceived\n", c)
		t.Fail()
	}
}
