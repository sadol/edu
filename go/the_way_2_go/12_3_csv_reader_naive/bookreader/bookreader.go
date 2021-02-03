package bookreader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const SEPARATOR rune = ';'

// removes double quotes from around the title field
func ExtractTitle(input string) (output string) {
	if strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"") {
		output = input[1:len(input)-1]
	} else {
		output = input
	}
	return
}

// private check function
func isSemicolon(someRune rune) (ret bool) {
	if someRune == SEPARATOR {
		ret = true
	} else {
		ret = false
	}
	return
}

type CsvBookLine struct {
	title      string  "book title"
	price      float64 "price of the book"
	quantity   float64 "available stock amount of the certain title"
}

// factory function for a book record, extracts fields from csv and puts  them
// into a struct
func NewCsvBookLine(csvLine string) (book CsvBookLine) {
	var fields []string = strings.FieldsFunc(csvLine, isSemicolon)
    book = *(new(CsvBookLine))

	book.title = ExtractTitle(fields[0])
	book.price, _ = strconv.ParseFloat(fields[1], 64)
	book.quantity, _ = strconv.ParseFloat(fields[2], 64)
	return
}

// stringificator, just to differentiate between raw data and processed data
func (book *CsvBookLine) String() string {
	return fmt.Sprintf("Title:%v,\tPrice:%v$,\tQuantity:%v",
		ExtractTitle(book.title), book.price, book.quantity)
}

// array of records
type Books []CsvBookLine

// loader function
func LoadBooks(fileName string) (books Books, err error) {
	var inputFile *os.File
    books = *(new(Books))
	var inputString string

	if inputFile, err = os.Open(fileName); err == nil {
		inputReader := bufio.NewReader(inputFile)
		for {
			if inputString, err = inputReader.ReadString('\n'); err != io.EOF {
				// discard the last byte of \n character
				books = append(books, NewCsvBookLine(inputString[:len(inputString)-1]))
			} else { break }
		}
	}
	defer inputFile.Close()
	if err == io.EOF { err = nil } // this is not an error afterall
	return
}
