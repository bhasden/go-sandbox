package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeStamp struct {
	Created time.Time
}
func main() {
	ts := TimeStamp{Created: time.Now()}
	marshal, err := json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Json: %s\n", marshal)
	fmt.Printf("Timestamp: %+v\n", ts)

	marshal, err = json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	myStamp := TimeStamp{}
	// s := []byte("{ \"created\": \"2021-06-09T18:34:23.765944-05:00\" }")
	s := []byte("{ \"created\": \"May 15, 2014\" }")
	json.Unmarshal(s, &myStamp)
	fmt.Printf("Json: %s\n", marshal)
	fmt.Printf("Timestamp: %+v\n", myStamp)
}

func (d *TimeStamp) MarshalJSON() ([]byte, error) {
	type Alias TimeStamp
	return json.Marshal(&struct {
		*Alias
		Stamp string `json:"stamp"`
	}{
		Alias: (*Alias)(d),
		Stamp: d.Created.Format("Mon Jan _2"),
	})
}