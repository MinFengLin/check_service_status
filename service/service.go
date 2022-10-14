package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net"
	"time"
)

// https://stackoverflow.com/questions/64693710/parse-json-file-in-golang

type Services_slice struct {
	Services []Services `json:"Services"`
}

type Services struct {
	Ip        string `json:"Ip"`
	Service   string `json:"Service"`
	Port      string `json:"Port"`
}

func Check_service_status(ii int, time_set int, service_data *Services_slice, failed_data *string) {
	withtimeout := net.Dialer{Timeout: time.Duration(time_set)*time.Millisecond}
	conn, err := withtimeout.Dial("tcp", service_data.Services[ii].Ip+":"+service_data.Services[ii].Port)

	if err != nil {
		*failed_data = *failed_data + service_data.Services[ii].Ip+":"+service_data.Services[ii].Port + " - (" +service_data.Services[ii].Service + ")" + "\n"
	} else {
		defer conn.Close()
	}
}

func Parser_services() Services_slice {

	filename := "./service_data.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := Services_slice{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	// fmt.Printf("%+v\n", data)

	return data
}
