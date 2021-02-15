package utils

import (
	"fmt"
	"strings"
	"time"
)

/*MakeCommonID Change different channel IDs to common ID for processing*/
func MakeCommonID(body []byte) []byte {
    text := string(body)
    if !strings.Contains(text, "ProgramId") { // Empty EPG response
        return nil
    }
    text = `{"listings":{"ID"` + text[strings.Index(text, `:[{"ProgramId"`):]

    return []byte(text)
}

/*XmltvTime Change timestamp according to xmltv standard*/
func XmltvTime(tm, timeZone string) string{
    timestamp, err := time.Parse("2006-01-02T15:04:05", tm)
    Check(err)

    return timestamp.Format("20060102150405") + " " + timeZone
}

/*Check Error check*/
func Check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}
