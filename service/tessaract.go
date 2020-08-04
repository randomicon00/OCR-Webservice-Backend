//step1 go get gopkg.in/GeertJohan/go.tesseract.v1
//step2 sudo apt-get install -t testing libtesseract3 libtesseract-dev
//step3 install language files e.g. sudo apt-get install -t testing tesseract-ocr-nld

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/GeertJohan/go.leptonica"
	"github.com/GeertJohan/go.tesseract"
)

func main() {
	//get the image to try
	flag.Parse()
	image := flag.Arg(0)
	
	//print the version
	fmt.Println(tessaract.Version())
	

	//create new tess instance and point it to the tessdata location. Set language to english...
	tessdata_prefix := os.Getenv("TESSDATA_PREFIX")
	if tessdata_prefix == "" {
		tessdata_prefix = "/usr/local/share"
	}
	//TODO add 20 different languages
	t, err := tessaract.NewTess(filepath.Join(tessadata_prefix, "tessadata"), "eng")
	if err != nil {
		log.Fatalf("Error while initializing Tess: %s", err)
	}
	defer t.Close()
	
	//open new pix from file with leptonica
	pix, err := leptonica.NewPixFromFile(image)
	if err != nil {
		log.Fatalf("Error while getting pix from file %s", err)
	}
	defer pix.Close() //remember to cleanup
	
	//set the page seg mode to autodetect 
	t.SetPageSegMode(tessaract.PSM_AUTO_OSD)	
	
	//setup a whitelist of all basic ascii
	err = t.SetVariable("tessedit_char_whitelist", ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_abcdefghijklmnopqrstuvwxyz{|}~`+"`")
	if err != nil {
		log.Fatalf("failed to setvariable: %s", err)
	}
	
	//set the image to the tessaract instance
	t.SetImagePix(pix)
	
	//retrieve text from the tessaract instance 
	fmt.Println(t.Text())
	
	//retrieve text from the tessaract instance
	fmt.Println(t.HOCRText(0))
	
	//retrieve text from the tessaract instance
	fmt.Println(t.BoxText(0))
	
	//now select just the first two columns if using FelixScan.jpg
	t.SetRectange(30, 275, 1120, 1380)
	fmt.Println(t.Text())
	fmt.Println(t.BoxText(0))
}


