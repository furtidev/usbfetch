package main

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

func main() {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()
	if runtime.GOOS == "linux" {
		data := genData()
		for _, v := range data.Devices {
			var deviceType string = v.ProductType
			var productName string = "(" + v.ProductName + ")"
			if deviceType == "unknown" && v.ProductName != "unknown" {
				deviceType = v.ProductName
			}
			if v.ProductName == "unknown" {
				productName = ""
			}
			fmt.Printf("%s\n    Vendor ID   %s %s\n    Product ID  %s %s\n    Speed       %s\n    Version    %s\n", green(deviceType), yellow(v.VendorId), cyan(v.VendorName), yellow(v.ProductId), cyan(productName), yellow(v.Speed+"Mbit/s"), yellow(v.Version))
		}
	} else {
		fmt.Printf("%s is %s a supported platform.\n", runtime.GOOS, red("not"))
	}
}
