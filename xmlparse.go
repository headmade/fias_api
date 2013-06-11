package main

import (
  // "bufio"
  "fmt"
  "os"
  "flag"
  "encoding/xml"
  "strings"
  // "regexp"
  "net/url"
  "math"
)

var inputFile = flag.String("infile", "objects.xml", "Input file path")
// var inputFile = flag.String("infile", "/media/DATA/fias_xml/AS_ADDROBJ_20130608_2ecf7d9a-6eab-43ec-a8fc-da85bfbf845a.XML", "Input file path")
var indexFile = flag.String("indexfile", "list.txt", "objects list output file")

type XmlObject struct {
  AOID string `xml:"AOID,attr"`
  AOGUID string `xml:"AOGUID,attr"`
  PARENTGUID string `xml:"PARENTGUID,attr"`
  FORMALNAME string `xml:"FORMALNAME,attr"`
  OFFNAME string `xml:"OFFNAME,attr"`
  SHORTNAME string `xml:"SHORTNAME,attr"`
  AOLEVEL string `xml:"AOLEVEL,attr"`
  REGIONCODE string `xml:"REGIONCODE,attr"`
  AREACODE string `xml:"AREACODE,attr"`
  AUTOCODE string `xml:"AUTOCODE,attr"`
  CITYCODE string `xml:"CITYCODE,attr"`
  CTARCODE string `xml:"CTARCODE,attr"`
  PLACECODE string `xml:"PLACECODE,attr"`
  STREETCODE string `xml:"STREETCODE,attr"`
  EXTRCODE string `xml:"EXTRCODE,attr"`
  SEXTCODE string `xml:"SEXTCODE,attr"`
  PLAINCODE string `xml:"PLAINCODE,attr"`
  CODE string `xml:"CODE,attr"`
  CURRSTATUS string `xml:"CURRSTATUS,attr"`
  ACTSTATUS string `xml:"ACTSTATUS,attr"`
  LIVESTATUS string `xml:"LIVESTATUS,attr"`
  CENTSTATUS string `xml:"CENTSTATUS,attr"`
  OPERSTATUS string `xml:"OPERSTATUS,attr"`
  IFNSFL string `xml:"IFNSFL,attr"`
  IFNSUL string `xml:"IFNSUL,attr"`
  TERRIFNSFL string `xml:"TERRIFNSFL,attr"`
  TERRIFNSUL string `xml:"TERRIFNSUL,attr"`
  OKATO string `xml:"OKATO,attr"`
  OKTMO string `xml:"OKTMO,attr"`
  POSTALCODE string `xml:"POSTALCODE,attr"`
  STARTDATE string `xml:"STARTDATE,attr"`
  ENDDATE string `xml:"ENDDATE,attr"`
  UPDATEDATE string `xml:"UPDATEDATE,attr"`
}

func CanonicalizeTitle(title string) string {
  can := strings.ToLower(title)
  can = strings.Replace(can, " ", "_", -1)
  can = url.QueryEscape(can)
  return can
}

func main() {
  flag.Parse()

  xmlFile, err := os.Open(*inputFile)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer xmlFile.Close()

  decoder := xml.NewDecoder(xmlFile)
  total := 0
  var inElement string
  for {
    // Read tokens from the XML document in a stream.
    t, _ := decoder.Token()
    if t == nil {
      break
    }
    // Inspect the type of the token just read.
    switch se := t.(type) {
    case xml.StartElement:
      // If we just read a StartElement token
      inElement = se.Name.Local
      // ...and its name is "Object"
      if inElement == "Object" {
        var p XmlObject
        // decode a whole chunk of following XML into the
        // variable p which is a Page (se above)
        decoder.DecodeElement(&p, &se)
        total++;
        if math.Mod(float64(total), 1000) == 0 {
          fmt.Println(total)
        }
        fmt.Println(p)

        // Do some stuff with the page.
        // p.Title = CanonicalizeTitle(p.Title)
        // m := filter.MatchString(p.Title)
      }
    default:
    }

  }

  fmt.Printf("Total objects: %d \n", total)
}