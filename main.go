package main

import (
	"io/ioutil"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	body := "本日は晴天なり"

	bin, _ := qrcode.Encode(body, qrcode.High, 256)
	ioutil.WriteFile("test.png", bin, 0644)
}
