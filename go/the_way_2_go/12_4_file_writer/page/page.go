package page

import (
	"bufio"
	"fmt"
	"os"
)

const BYTESBATCH = 1024 // how many bytes are processed as one batch

type Page struct {
	Title string
	Body  []byte
}

// returns file name suitable for reading from and writting to.
func (p *Page) GetName() string {
	return p.Title + ".txt"
}

// save contents of the page to the file
func (page *Page) Save() (err error) {
	var outputFile *os.File // in order to `Close after main IF
    lenBody := len(page.Body)
	if outputFile, err = os.OpenFile(page.GetName(), os.O_WRONLY|os.O_CREATE, 0666); err == nil {
		defer outputFile.Close()
		outputWriter := bufio.NewWriter(outputFile)
        start, end := 0, 0
        for {
            end += BYTESBATCH
            if end >= lenBody { end = lenBody }
            if _, err = outputWriter.Write(page.Body[start:end]); err == nil {
                outputWriter.Flush()
            }
            if start == lenBody { break }
            start = end
        }
	}
	return
}

// stringer
func (p *Page) String() string {
	return fmt.Sprintf("Title:%v\nBody:\n%v.", p.Title, p.Body)
}

// factory function for Page structs
// WARNING: remember to UNPACK bytes array(s)!!!
func NewPage(title string, body ...byte) (page *Page) {
	page = new(Page)
	page.Title = title
	page.Body = make([]byte, 0)

	page.Body = append(page.Body, body...) // unpacking
	return
}

// load page contents from the file
func Load(fileName string) (output *Page, err error) {
	partialBuffer := make([]byte, BYTESBATCH) // 1k buffer for partial reads
	retPage := NewPage(fileName[:len(fileName)-4])
	var inputFile *os.File // in order to close after main IF

	if inputFile, err = os.Open(retPage.GetName()); err == nil {
		defer inputFile.Close()
		inputReader := bufio.NewReader(inputFile)
		for { // buffer by buffer reading
			if noOfBytes, err := inputReader.Read(partialBuffer); noOfBytes > 0 && err == nil {
				retPage.Body = append(retPage.Body, partialBuffer[:noOfBytes]...)
			} else { break } // in case of read error
		}
	}
    if err == nil { output = retPage }
	return
}
