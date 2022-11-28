package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func ping_check(ii int, time_set int, service_data *Services_slice, failed_data *string) {
	withtimeout := net.Dialer{Timeout: time.Duration(time_set) * time.Millisecond}
	conn, err := withtimeout.Dial("tcp", service_data.Services[ii].Ip+":"+service_data.Services[ii].Port)

	if err != nil {
		*failed_data = *failed_data + service_data.Services[ii].Ip + ":" + service_data.Services[ii].Port + " - (" + service_data.Services[ii].Service + ")" + "\n"
	} else {
		defer conn.Close()
	}
}

func Check_service_status() string {
	service_data := Parser_services()
	failed_data := ""
	for ii := range service_data.Services {
		ping_check(ii, 500, &service_data, &failed_data)
	}
	if len(failed_data) > 0 {
		failed_data = "↻ Check Status ...... 🔴FAILED! \n - \n" + failed_data + "-"
	} else {
		failed_data = "↻ Check Status ...... 🟢PASS! \n" + failed_data
	}

	return failed_data
}

func List_service_status() string {
	service_data := Parser_services()
	services_info := "-\n"

	for ii := range service_data.Services {
		services_info = services_info + service_data.Services[ii].Ip + ":" + service_data.Services[ii].Port + " - (" + service_data.Services[ii].Service + ")" + "\n"
	}
	services_info = services_info + "-\n"

	return services_info
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
