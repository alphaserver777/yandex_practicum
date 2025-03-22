package main

import (
	"encoding/xml"
	"fmt"
)

const MockXMLDocument = `
<report>
  <competition>
    <location>РФ, Санкт-Петербург, Дворец творчества юных техников</location>
    <class>ТА-24</class>
  </competition>
  <racer global_id="100">
    <nick>RacerX</nick>
    <best_lap_ms>61012</best_lap_ms>
    <laps>52.3</laps>
  </racer>
  <racer global_id="127">
    <nick>Иван The Шумахер</nick>
    <best_lap_ms>61023</best_lap_ms>
    <laps>51</laps>
  </racer>
  <racer global_id="203">
    <nick>Петя Иванов</nick>
    <best_lap_ms>63000</best_lap_ms>
    <laps>49.9</laps>
    <!--Болид не соответствует техническому регламенту, 
    результат не учитывается в общем рейтинге-->
  </racer>
  <racer>
    <nick>Гость 1</nick>
    <best_lap_ms>123001</best_lap_ms>
    <laps>25.8</laps>
  </racer>
</report> 
`

func main() {
	output, err := FilterXML(MockXMLDocument, 50)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type (
	RaceReport struct {
		XMLName             xml.Name      `xml:"report"`
		CompetitionLocation string        `xml:"competition>location"`
		CompetitionClass    string        `xml:"competition>class"`
		Results             []RacerResult `xml:"racer"`
	}

	RacerResult struct {
		XMLName   xml.Name `xml:"racer"`
		GlobalId  *int     `xml:"global_id,attr,omitempty"` // Сделали указателем, чтобы отличать "0" от "не указан"
		Nick      string   `xml:"nick"`
		BestLapMs int64    `xml:"best_lap_ms"`
		Laps      float32  `xml:"laps"`
		Comment   string   `xml:",omitempty"` // Добавили omitempty
	}
)

// FilterXML оставляет в XML только тех гонщиков, у которых больше, чем laps, кругов.
func FilterXML(input string, laps float32) (output string, err error) {
	var rp RaceReport

	// десериализуем XML
	err = xml.Unmarshal([]byte(input), &rp)
	if err != nil {
		return
	}

	// создаём новый список гонщиков
	filter := make([]RacerResult, 0, len(rp.Results))
	for _, racer := range rp.Results {
		if racer.Laps > laps {
			filter = append(filter, racer)
		}
	}
	rp.Results = filter

	// сериализуем данные в XML с отступами
	var data []byte
	data, err = xml.MarshalIndent(rp, "", "   ")
	if err != nil {
		return
	}
	output = string(data)
	return
}
