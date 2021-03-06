package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	tagRegex          = regexp.MustCompile(`^\(([[:alnum:]]{4}),([[:alnum:]]{4})\)$`)
	multiplicityRegex = regexp.MustCompile(`([0-9])(-([0-9n]))?`)
)

// XML structs:

type Cell struct {
	Value   string `xml:",chardata"`
	EmValue string `xml:"emphasis"`
}

func (c Cell) GetValue() string {
	if strings.TrimSpace(c.Value) != "" {
		return c.Value
	}

	return c.EmValue
}

type Row struct {
	Cells []Cell `xml:"td>para"`
}

func (r Row) GetValues() []string {
	s := make([]string, len(r.Cells))
	for _, cell := range r.Cells {
		s = append(s, cell.GetValue())
	}
	return s
}

type Table struct {
	Rows []Row `xml:"tbody>tr"`
}

type Section struct {
	Tables []Table `xml:"table"`
}

type Chapter struct {
	Label    string    `xml:"label,attr"`
	Tables   []Table   `xml:"table"`
	Sections []Section `xml:"section"`
}

func (c Chapter) GetTables() []Table {
	tables := make([]Table, 10)

	if c.Tables != nil {
		tables = append(tables, c.Tables...)
	}

	if c.Sections != nil {
		for _, section := range c.Sections {
			if section.Tables != nil {
				tables = append(tables, section.Tables...)
			}
		}
	}

	return tables
}

type Book struct {
	Chapters []Chapter `xml:"chapter"`
}

// output prep:
type Element struct {
	Tag     string
	Keyword string
	VR      string
	VM      string
	Retired bool
	Desc    string
}

// Get the tag value for the element with "x" nibbles replaced with the given
// hex character.  Will presumably produce either a high or low boundary on
// the range that this tag can support.
func (e Element) GetTagBoundary(replacement string) *uint32 {
	matches := tagRegex.FindStringSubmatch(e.Tag)
	if matches == nil {
		return nil
	}

	hex := matches[1] + matches[2]
	hex = strings.Replace(hex, "x", replacement, -1)

	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return nil
	}

	value32 := uint32(value)
	return &value32
}

// Get the lowest tag the element can have, which will normally be the only
// one.
func (e Element) GetTagLowValue() *uint32 {
	return e.GetTagBoundary("0")
}

// Get the highest tag the element can have, which will normally be the same as
// the lowest, except for retired elements that are represented by a range of
// tags.
func (e Element) GetTagHighValue() *uint32 {
	return e.GetTagBoundary("F")
}

func (e Element) GetKeyword() string {
	return regexp.MustCompile(`[^[:alnum:]]`).ReplaceAllString(e.Keyword, "")
}

// Unimplemented, not sure if useful
func (e Element) GetLowMultiplicity() *uint32 {
	return nil
}

// Unimplemented, not sure if useful
func (e Element) GetHighMultiplicity() *uint32 {
	return nil
}

func (e Element) GetVR() string {
	tag := e.GetTagLowValue()

	if tag == nil || *tag == 0xFFFEE000 || *tag == 0xFFFEE00D || *tag == 0xFFFEE0DD {
		return "UN"
	}

	// TODO: parse better, validate known possible vrs
	if len(e.VR) > 1 {
		return e.VR[:2]
	}

	return "UN"
}

func (e Element) String() string {
	base := e.GetKeyword()

	if e.GetTagLowValue() != nil {
		base = base + fmt.Sprintf(" (%08X)", *e.GetTagLowValue())
	}

	if e.GetTagHighValue() != nil && e.GetTagHighValue() != e.GetTagLowValue() {
		base = base + fmt.Sprintf("..(%08X)", *e.GetTagHighValue())
	}

	return base
}

func NewElement(row Row) Element {
	return Element{
		Tag:     row.Cells[0].GetValue(),
		Keyword: row.Cells[2].GetValue(),
		VR:      row.Cells[3].GetValue(),
		VM:      row.Cells[4].GetValue(),
		Retired: row.Cells[4].GetValue() == "RET",
		Desc:    row.Cells[1].GetValue(),
	}
}

type elements map[uint32]Element

func NewUID(row Row) UID {
	return UID{
		uidValue: row.Cells[0].GetValue(),
		uidName:  row.Cells[1].GetValue(),
	}
}

type uids map[string]UID

type UID struct {
	uidValue string
	uidName  string
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// TODO: flags to trigger only specific go files data dictionary generation
// TODO: flag to automatically fetch latest spec
func main() {

	fmt.Printf("Downloading part06.xml...\n")
	if err := DownloadFile("part06.xml", "http://dicom.nema.org/medical/dicom/current/source/docbook/part06/part06.xml"); err != nil {
		panic(err)
	}
	fmt.Printf("Downloading part07.xml...\n")
	if err := DownloadFile("part07.xml", "http://dicom.nema.org/medical/dicom/current/source/docbook/part07/part07.xml"); err != nil {
		panic(err)
	}

	elements := make(elements)
	fmt.Printf("Parsing part06...\n")
	parseStandard("part06.xml", elements, "6", "7", "8")
	fmt.Printf("Parsing part07...\n")
	parseStandard("part07.xml", elements, "E")

	fmt.Printf("Writing tags...\n")
	writeTags(elements)
	//writeDictionary(elements)
	fmt.Printf("Writing vrs...\n")
	writeVRs(elements)

	// fmt.Printf("Writing sample test data...\n")
	// writeSampleTestData("GENECG.dcm")
	//
	// uids := make(uids)
	// fmt.Printf("Parsing part06 for UIDs...\n")
	// parseStandardForUIDs("part06.xml", uids, "A")
	// fmt.Printf("Writing uids...\n")
	// writeUIDs(uids)

	fmt.Printf("Done.\n")
}

func parseStandard(filename string, elements elements, chapters ...string) {
	stream, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	decoder := xml.NewDecoder(stream)
	var book Book
	err = decoder.Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	tagChapters := NewSet(chapters...)

	for _, chapter := range book.Chapters {
		if !tagChapters.Contains(chapter.Label) {
			continue
		}

		for _, table := range chapter.GetTables() {
			for _, row := range table.Rows {
				element := NewElement(row)

				if element.GetTagLowValue() != nil {
					elements[*element.GetTagLowValue()] = element
				} else {
					log.Printf("WARNING: no tag value for %s\n", element)
				}
			}
		}
	}
}

func parseStandardForUIDs(filename string, uids uids, chapters ...string) {
	stream, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	decoder := xml.NewDecoder(stream)
	var book Book
	err = decoder.Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	tagChapters := NewSet(chapters...)

	for _, chapter := range book.Chapters {
		if !tagChapters.Contains(chapter.Label) {
			continue
		}

		for _, table := range chapter.GetTables() {
			for _, row := range table.Rows {
				uid := NewUID(row)
				uids[uid.uidValue] = uid
				//				break
			}
		}
	}
}

func forEach(elements elements, f func(element Element)) {
	keys := make(tagSlice, 0, len(elements))
	for key := range elements {
		keys = append(keys, key)
	}
	sort.Sort(keys)

	for _, tag := range keys {
		f(elements[tag])
	}
}

func writeTags(elements elements) {
	out, err := os.Create("../tmp/tags.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintln(out, "package dcm4go")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "// auto-generated, do not edit")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "const (")

	forEach(elements, func(element Element) {
		if element.GetKeyword() != "" {
			fmt.Fprintf(out, "// %sTag is tag for %s\n%sTag = 0x%08X\n", element.GetKeyword(), element.Desc, element.GetKeyword(), *element.GetTagLowValue())
		}
	})

	fmt.Fprintln(out, ")")
}

// func writeTags(elements elements) {
// 	out, err := os.Create("tags.go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer out.Close()
//
// 	fmt.Fprintln(out, "package main")
// 	fmt.Fprintln(out, "")
// 	fmt.Fprintln(out, "// auto-generated, do not edit")
// 	fmt.Fprintln(out, "")
// 	fmt.Fprintln(out, "const (")
//
// 	maxLen := 0
// 	forEach(elements, func(element Element) {
// 		newLen := len(element.GetKeyword())
// 		if newLen > maxLen {
// 			maxLen = newLen
// 		}
// 	})
//
// 	maxLenStr := fmt.Sprintf("%d", maxLen)
// 	forEach(elements, func(element Element) {
// 		// TODO: add comments with other attributes
// 		if element.GetKeyword() != "" {
// 			fmt.Fprintf(out, "\t%-"+maxLenStr+"s = Tag(0x%08X)\n",
// 				element.GetKeyword(), *element.GetTagLowValue())
// 		}
// 	})
//
// 	fmt.Fprintln(out, ")")
// }

func writeVRs(elements elements) {
	out, err := os.Create("../tmp/vrs.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintln(out, "package dcm4go")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "// auto-generated, do not edit")
	fmt.Fprintln(out)
	fmt.Fprint(out, "var vrs = map[uint32]string{\n")

	i := 0
	forEach(elements, func(element Element) {
		// TODO: add comments with other attributes
		if element.GetKeyword() != "" {
			fmt.Fprintf(out, "0x%08X:%q,",
				*element.GetTagLowValue(),
				element.GetVR())
		}
		i++
		if i%4 == 0 {
			fmt.Fprintf(out, "\n")
		}
	})

	fmt.Fprintln(out, "\n}")
}

func forEachUID(uids uids, f func(uid UID)) {
	keys := make(uidSlice, 0, len(uids))
	for key := range uids {
		keys = append(keys, key)
	}
	sort.Sort(keys)

	for _, tag := range keys {
		f(uids[tag])
	}
}

func writeUIDs(uids uids) {
	out, err := os.Create("../tmp/uids.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintln(out, "package dcm4go")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "// auto-generated, do not edit")
	fmt.Fprintln(out)
	fmt.Fprint(out, "const (\n")

	forEachUID(uids, func(uid UID) {
		if uid.uidName != "" && !strings.Contains(uid.uidName, "(Retired)") {
			name := strings.ReplaceAll(uid.uidName, " ", "")
			name = strings.ReplaceAll(name, "-", "")
			name = strings.ReplaceAll(name, "/", "")
			name = strings.ReplaceAll(name, ".", "")
			value := strings.ReplaceAll(uid.uidValue, "\u200b", "")
			fmt.Fprintf(out, "// %sUID is uid for %s\n"+"%sUID = %q\n", name, value, name, value)
		}
	})

	fmt.Fprintln(out, ")")
}

// func writeDictionary(elements elements) {
// 	out, err := os.Create("stddict.go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer out.Close()
//
// 	fmt.Fprintf(out, stddictHeader)
//
// 	forEach(elements, func(element Element) {
// 		// TODO: support for multi-tag elements
// 		// - either define a single value and map it multiple times
// 		// - or enhance datadict.go to search for these somehow...
// 		fmt.Fprintf(out, "\t\t"+elementSpecPattern+"\n",
// 			*element.GetTagLowValue(),
// 			*element.GetTagLowValue(),
// 			*element.GetTagHighValue(),
// 			element.GetVR(),
// 			element.Retired,
// 			element.Desc,
// 			element.GetKeyword(),
// 		)
// 	})
//
// 	fmt.Fprintf(out, "\t})\n")
// }
//
// const stddictHeader = `package main
// // auto-generated, do not edit
// var stddict = NewDataDictionary("",
// 	map[Tag]ElementSpec{
// `
//
// const elementSpecPattern = `Tag(0x%08X): {tag: Tag(0x%08X), ` +
// 	`maxValue: Tag(0x%08X), vr: %s, retired: %t, desc: "%s", keyword: "%s"},`

// dumb util for set membership check
type Set map[string]*struct{}

func NewSet(items ...string) Set {
	set := make(Set)
	for _, item := range items {
		set[item] = nil
	}
	return set
}
func (set Set) Contains(item string) bool {
	_, ok := set[item]
	return ok
}

// for sorting of uint32s
type tagSlice []uint32

func (p tagSlice) Len() int           { return len(p) }
func (p tagSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p tagSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// for sorting of uids
type uidSlice []string

func (p uidSlice) Len() int           { return len(p) }
func (p uidSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p uidSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func writeSampleTestData(path string) {
	out, err := os.Create("../tmp/sample.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintln(out, "package dcm4go")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "// auto-generated, do not edit")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "var sample []byte{")

	in, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	var buffer [1024]byte
	for {
		num, err := in.Read(buffer[:])
		for i := 0; i < num; i++ {
			fmt.Fprintf(out, "0x%02X,", buffer[i])
		}
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	fmt.Fprintln(out, "}")
}
