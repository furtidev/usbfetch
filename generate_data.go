package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Device struct {
	ProductType string
	VendorName  string
	ProductName string
	ProductId   string
	VendorId    string
	Speed       string
	Version     string
}

type Data struct {
	Devices []Device
}

var alreadyVisitedIds []string

// stolen from: https://stackoverflow.com/a/75049947
func getAbs(path string, locationPath string) string {
	var new_abs_path string
	new_file_path := path
	if filepath.IsAbs(new_file_path) {
		new_abs_path = new_file_path
	} else {
		new_abs_path, _ = filepath.Abs(filepath.Join(locationPath, path))
	}
	return new_abs_path
}

func getBasicValues(location string) string {
	var result string
	value, err := os.ReadFile(location)
	if err != nil {
		result = "unknown"
	} else {
		result = string(value)[:len(value)-1]
	}
	return result
}

func genData() Data {
	usb_interface := "/sys/bus/usb/devices"
	data := Data{}
	filepath.WalkDir(usb_interface, func(path string, d fs.DirEntry, err error) error {
		if path != usb_interface {
			if parts := strings.Split(path, ":"); len(parts) < 2 {
				walkFolder(&data, path)
			}
		}
		return nil
	})
	return data
}

func walkFolder(data *Data, location string) {
	linkPath, _ := os.Readlink(location)
	location = getAbs(linkPath, "/sys/bus/usb/devices")
	filepath.WalkDir(location, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "idProduct" {
			dev := Device{}
			dev.ProductId = getBasicValues(path)
			if slices.Contains(alreadyVisitedIds, dev.ProductId) {
				return nil
			}

			dev.VendorId = getBasicValues(location + "/idVendor")

			dev.ProductType = getBasicValues(location + "/product")

			dev.Speed = getBasicValues(location + "/speed")

			dev.Version = getBasicValues(location + "/version")

			// nasty consequences of using interfaces
			if x, ok := Ids[dev.VendorId].(map[string]interface{}); ok {
				dev.VendorName = x["name"].(string)
				devices := x["devices"].(map[string]string)
				name, ok := devices[dev.ProductId]
				if !ok {
					name = "unknown"
				}
				dev.ProductName = name
			} else {
				dev.VendorName, dev.ProductName = "unknown", "unknown"
			}
			data.Devices = append(data.Devices, dev)
			alreadyVisitedIds = append(alreadyVisitedIds, dev.ProductId)
		}
		return nil
	})
}
