package xml

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"

	"golang.org/x/net/html/charset"

	"github.com/zguillez/go-tools/system"
)

func Decoder(filePath string) *xml.Decoder {

	xmlFile, err := os.Open(filePath)
	system.CheckError(err)
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	system.CheckError(err)

	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	return decoder
}
