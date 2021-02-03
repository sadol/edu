// package to hadle basic statistical metrics for string of float64 values;
// for easy use in golang text templates
package metrics

import (
    "sort"
    "strings"
    "strconv"
    "math"
    "errors"
)

// base statistics storage; all fields are unexported
type Metrics struct {
    rawList string
    separators []string
    list []float64
    lenList int
    mean float64
    median float64
}

// factory function for statistics storage; DO-IT-ALL function.
func NewMetrics(rawString string, separators []string) (m *Metrics, err error) {
    m = new(Metrics)
    m.rawList = rawString
    m.separators = separators
    m.list = make([]float64, 0)
    if err = m.fillList(); err == nil {
        m.fillLenList()
        m.fillMean()
        m.fillMedian()
    }
    return
}

// private method; fills internal float64 buffer with proper vals
func (m *Metrics) fillList() (err error) {
    numOfSeparators, separatorIndex := 0, 0
    var justOne float64
    // check uniqueness of the separators in the raw string
    for i, separator := range m.separators {
        if strings.Contains(m.rawList, separator) {
            numOfSeparators++
            separatorIndex = i
        }
    }

    switch {
    case numOfSeparators == 1:
        almostProcessed := strings.Split(m.rawList, m.separators[separatorIndex])
        for _, v := range almostProcessed {
            if len(v) > 0 {                              // only non-empty vals
                if justOne, err = strconv.ParseFloat(v, 64); err == nil {
                    m.list = append(m.list, justOne)
                } else { break }                                       // error
            }
        }
    case numOfSeparators > 1:
        err = errors.New("Ambiguous separators in the input string!!!")
    default:                                             // no separators found
        if justOne, err = strconv.ParseFloat(m.rawList, 64); err == nil {
            m.list = append(m.list, justOne)
        } else {
            err = errors.New("Unknown separators in the input string!!!")
        }
    }

    return
}

// private method; fills lenList
func (m *Metrics) fillLenList() {
    m.lenList = len(m.list)
}

// private mathod;fills m.Mean val
func (m *Metrics) fillMean() {
    var mean float64 = 0.00

    for _, v := range m.list {
        mean += v
    }
    if m.lenList > 0 {
        m.mean = mean / float64(m.lenList)
    }
}

// private method; fills m.median val
func (m *Metrics) fillMedian() {
    sort.Float64s(m.list)
    if m.lenList == 0 {
        m.median = 0.00
    } else if m.lenList == 1 {
        m.median = m.list[0]
    } else if int(math.Mod(float64(m.lenList), 2)) == 0 {
        m.median = (m.list[m.lenList / 2] + m.list[(m.lenList / 2) - 1]) / 2.00
    } else if int(math.Mod(float64(m.lenList), 2)) == 1 {
        m.median = m.list[(m.lenList - 1) / 2]
    }
}

// TODO: double check returned results
// it's safer to use dedicated getters in case of Metrics struct storage;
// there are no setters allowed (besides initial factory function)
func (m *Metrics) GetRawList() string {
    return m.rawList
}

func (m *Metrics) GetLen() int {
    return m.lenList
}

func (m *Metrics) GetMean() float64 {
    return m.mean
}

func (m *Metrics) GetMedian() float64 {
    return m.median
}
