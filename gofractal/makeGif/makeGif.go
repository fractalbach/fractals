package main

import ( 
    "fmt"
    "os"
    "image"
    "image/gif"
    _ "image/png"
    "strconv"
)


func main() {

    // File Naming.
    // Input the number of files are present in the folder.
    // Assumes that those files' names are just numbers.

    numOfFiles := 9
    files := make([]string, numOfFiles)

    for i:=0; i<numOfFiles; i++ {
        files[i] =  "series/" + strconv.Itoa(i+2) + ".png"
    }


    


    // load static image and construct outGif
    outGif := &gif.GIF{}

    for _, name := range files {

        f, _ := os.Open(name)
        inputImage, formatName, _ := image.Decode(f)
        f.Close()

        outGif.Image = append(outGif.Image, inputImage)
        outGif.Delay = append(outGif.Delay, 0)
    }


    // save to out.gif
    f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, outGif)
}