package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Printf("json2csv: \n\n")
	if len(os.Args) > 2 {
		os.Exit(0)
	}
	var rawMessages []json.RawMessage
	var rawMessage json.RawMessage
	var jsondata interface{}
	csv := make(map[string][]string)
	var isArr bool
	flag.Func("In", "input json `file path`", func(s string) error {
		if s == "-" {
			return nil
		}
		f, err := os.Open(s)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		defer f.Close()
		err = json.Unmarshal([]byte(data), &rawMessages)
		if err == nil {
			isArr = true
		} else {
			json.Unmarshal([]byte(data), &rawMessage)
		}
		return nil
	})
	flag.Func("S", "input json chars string `json chars`", func(s string) error {
		err := json.Unmarshal([]byte(s), &rawMessages)
		if err == nil {
			isArr = true
		} else {
			json.Unmarshal([]byte(s), &rawMessage)
		}
		return nil
	})
	flag.Parse()
	if isArr {
		for _, jsonraw := range rawMessages {
			_ = json.Unmarshal([]byte(jsonraw), &jsondata)
			parseJson(jsondata, csv)
		}
		tocsv(csv)
	} else {
		_ = json.Unmarshal([]byte(rawMessage), &jsondata)
		parseJson(jsondata, csv)
		tocsv(csv)
	}
}

func parseJson(data interface{}, csv map[string][]string) {
	for k, v := range data.(map[string]interface{}) {
		_, ok := csv[k]
		if ok {
			csv[k] = append(csv[k], v.(string))
		} else {
			csv[k] = []string{v.(string)}
		}
	}
}

func tocsv(csv map[string][]string) {
	var title []string
	sli := make(map[int][]string)
	for key, vals := range csv {
		title = append(title, key)
		for i, val := range vals {
			index := strings.Index(val, ",")
			if index != -1 {
				val = fmt.Sprintf("\"%s\"", val)
			}
			sli[i] = append(sli[i], val)
		}
	}
	newcsv := strings.Join(title, ",")
	for _, v := range sli {
		newv := strings.Join(v, ",")
		newcsv = newcsv + "\n" + newv
	}
	fmt.Fprintf(os.Stdout, "%s\n\n", newcsv)
}
