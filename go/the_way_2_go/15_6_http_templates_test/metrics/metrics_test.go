// tests for metrics.go package
package metrics

import (
    "testing"
    "errors"
)

// object for testing in bulk
type testCase struct {
    RawList string
    LenList int
    Mean float64
    Median float64
    Err error
}

// every 'gotcha' case can be tested in the test below only (other tests test
// getters only)
func TestNewMetrics(t *testing.T) {
    knownSeparators := []string{",", " "}
    // test cases slice
    cases := []testCase{
        {"1 2 3 4", 4, 2.50, 2.50, nil},
        {"    1 2     4   3", 4, 2.50, 2.50, nil},
        {"1,2,3,4", 4, 2.50, 2.50, nil},
        {"1,2,3,5,4", 5, 3.00, 3.00, nil},
        {"1, 2, 3, 4", 0, 0, 0, errors.New("Ambiguous separators in the input string!!!")},
        {"1;2;3;4", 0, 0, 0, errors.New("Unknown separators in the input string!!!")},
        {"1,2;3 4", 0, 0, 0, errors.New("Ambiguous separators in the input string!!!")},
        {"a, b, c, d", 0, 0, 0, errors.New("Ambiguous separators in the input string!!!")},
        {"a,b,c,d", 0, 0, 0, errors.New(`strconv.ParseFloat: parsing "a": invalid syntax`)},
    }

    for _, c := range cases {
        if _, err := NewMetrics(c.RawList, knownSeparators); err != nil { // metrics objects ignored completly
            if c.Err == nil {                   // Houston, we've had a problem
                t.Errorf("Unexpected error : %v; for raw string: %v\n", err.Error(), c.RawList)
            } else {                                 // check type of the error
                if err.Error() != c.Err.Error() {                   // not good
                    t.Errorf("Error should be: %v, BUT is: %v\n", c.Err.Error(), err.Error())
                }
            }
        } else {                                         // no error encoutered
            if c.Err != nil {                   // Houston, we've had a problem
                t.Errorf("Should be error : %v; for raw string: %v\n", c.Err.Error(), c.RawList)
            }
        }
    }
}

func TestGetLen(t *testing.T) {
    m1, _ := NewMetrics("1,2,3,4", []string{",", " "})

    l := m1.GetLen()
    if l != 4 {
        t.Errorf("GetLen should give: %v; but gives: %v instead.\n", 4, l)
    }
}

func TestGetMean(t *testing.T) {
    m1, _ := NewMetrics("1,2,3,4", []string{",",})
    m2, _ := NewMetrics("1 2 3 4 5", []string{" ",})
    testCases := map[*Metrics]float64 {m1: 2.50, m2: 3.00,}

    for i, v := range testCases {
        m := i.GetMean()
        if m != v {
            t.Errorf("GetMean should give: %v; but gives: %v instead.\n", m, v)
        }
    }
}

func TestGetMedian(t *testing.T) {
    m1, _ := NewMetrics("1,2,3,4", []string{",",})
    m2, _ := NewMetrics("1 2 3 4 5", []string{" ",})
    m3, _ := NewMetrics("1,2,3,4,10,11,12", []string{",",})
    testCases := map[*Metrics]float64 {m1: 2.50, m2: 3.00, m3: 4.00,}

    for i, v := range testCases {
        m := i.GetMedian()
        if m != v {
            t.Errorf("GetMedian should give: %v; but gives: %v instead.\n", m, v)
        }
    }
}

func TestGetRawList(t *testing.T) {
    r1, r2, r3 := "1,2,3,4", "1 2 3 4 5", "1,2,3,4,10,11,12"
    m1, _ := NewMetrics(r1, []string{",",})
    m2, _ := NewMetrics(r2, []string{" ",})
    m3, _ := NewMetrics(r3, []string{",",})
    testCases := map[*Metrics]string {m1: r1, m2: r2, m3: r3,}

    for i, v := range testCases {
        m := i.GetRawList()
        if m != v {
            t.Errorf("GetRawList should give: %v; but gives: %v instead.\n", m, v)
        }
    }
}
