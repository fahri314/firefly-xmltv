package main

import (
	"encoding/json"
	"encoding/xml"
	"firefly-xmltv/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"xmltv/models"
)

const (
    timeZone = "+0300"
    generatorInfoName = "firefly-xmltv"
    generatorInfoURL = "github.com/fahri314/firefly-xmltv"
)

func main() {
    var xmlTV models.Tv
    var digiturkChannelIndex models.DigiturkChannelIndex
    var indexURL = "https://www.digiturk.com.tr/yayin-akisi/api/kanal/index/"
    
    // Get channel index
    client := &http.Client{}
    req, err := http.NewRequest("GET", indexURL, nil)
    utils.Check(err)

    res, err := client.Do(req)
    utils.Check(err)
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    utils.Check(err)

    // err = ioutil.WriteFile("output.json", body, 0644)
    // utils.Check(err)

    err = json.Unmarshal(body, &digiturkChannelIndex)
    utils.Check(err)

    xmlTV.GeneratorInfoName = generatorInfoName
    xmlTV.GeneratorInfoURL = generatorInfoURL

    // Channel loop
    for _, item := range digiturkChannelIndex.Init.Channels{
        var icon = "http://contentlibrary.digiturk.com.tr/Channel/" + strconv.Itoa(item.ID) + "/Image/" + item.FileName

        // Create channel
        channel := models.Channel{
            ID: item.Title,
            DisplayName: struct {
                Text string `xml:",chardata"`
                Lang string `xml:"lang,attr"`
            }{
                Text: item.Title,
                Lang: "tr",   // fmt.Sprintf("%v", epg.ProgramLanguage)
            },
            Icon: struct {
                Text string `xml:",chardata"`
                Src string `xml:"src,attr"`
            }{
                Src: icon,
            },
            URL: "https://www.digiturk.com.tr",
        }

        xmlTV.Channel = append(xmlTV.Channel, channel)
        
        // Day loop
        for day := 0; day < 7; day++ {
            var digiturkEPG models.DigiturkEPG
            var epgURL = "https://www.digiturk.com.tr/yayin-akisi/api/program/kanal/" + strconv.Itoa(item.ID) + "/" + time.Now().AddDate(0, 0, day).Format("2006-01-02") + "/0"

            client := &http.Client{}
            req, err := http.NewRequest("GET", epgURL, nil)
            utils.Check(err)

            res, err := client.Do(req)
            utils.Check(err)
            defer res.Body.Close()

            body, err := ioutil.ReadAll(res.Body)
            utils.Check(err)

            if body = utils.MakeCommonID(body); body == nil {
                continue
            }

            err = json.Unmarshal(body, &digiturkEPG)
            utils.Check(err)

            // EPG loop
            for _, epg := range digiturkEPG.Channel.EPGList {
                // Create EPG
                description := epg.Synopsis
                if epg.LongDescription != "" {
                    description += " (" + epg.LongDescription+")"
                }
                if epg.ProductionCountries != "" {
                    description += " ProductionCountries: " + epg.ProductionCountries
                }
        
                programme := models.Programme{
                    Start: utils.XmltvTime(epg.BroadcastStart, timeZone),
                    Stop: utils.XmltvTime(epg.BroadcastEnd, timeZone),
                    Channel: item.Title,
                    Title: struct {
                        Text string `xml:",chardata"`
                        Lang string `xml:"lang,attr"`
                    }{
                        Text: epg.ProgramName,
                        Lang: "tr",
                    },
                    Desc: struct {
                        Text string `xml:",chardata"`
                        Lang string `xml:"lang,attr"`
                    }{
                        Text: description,
                        Lang: "tr",
                    },
                }
        
                xmlTV.Programme = append(xmlTV.Programme, programme)
            }
        }
    }
 
	xmlstring, _ := xml.MarshalIndent(xmlTV, "", "  ")
    xmlstring = []byte(xml.Header + string(xmlstring))

    // Write file
	err = ioutil.WriteFile("xmltv.xml", xmlstring, 0644)
    utils.Check(err)
}
