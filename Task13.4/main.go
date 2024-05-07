package main

import (
	"encoding/xml"
	"fmt"
)

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}
	d := contracts{}
	d.List = append(d.List, c)
	res, err := xml.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(xml.Header, string(res))
}
