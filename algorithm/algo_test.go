package algorithm

// func TestString(t *testing.T) {
// 	for _, i := range AlgoTypes {
// 		switch i {
// 		case ContinuedFraction:
// 			if i.String() != "continued_fraction" {
// 				t.Log("unexpected algorithm name:", i.String(), "expected continued_fraction")
// 				t.Fail()
// 			}
// 		case NestedRadical:
// 			if i.String() != "nested_radical" {
// 				t.Log("unexpected algorithm name:", i.String(), "expected nested_radical")
// 				t.Fail()
// 			}
// 		case RationalFunc:
// 			if i.String() != "rational_func" {
// 				t.Log("unexpected algorithm name:", i.String(), "expected rational_func")
// 				t.Fail()
// 			}
// 		}
// 	}
// }

// func TestGet(t *testing.T) {
// 	var fn Algorithm
// 	var err error

// 	fn, err = Get(ContinuedFraction)
// 	if fn == nil || err != nil {
// 		t.Log("expected ContinuedFraction function")
// 		t.Fail()
// 	}

// 	fn, err = Get(NestedRadical)
// 	if fn == nil || err != nil {
// 		t.Log("expected NestedRadical function")
// 		t.Fail()
// 	}

// 	fn, err = Get(RationalFunc)
// 	if fn == nil || err != nil {
// 		t.Log("expected RationalFunc function")
// 		t.Fail()
// 	}

// 	fn, err = Get(math.MaxInt16)
// 	if fn != nil || err == nil {
// 		t.Log("expected error")
// 		t.Fail()
// 	}
// }
