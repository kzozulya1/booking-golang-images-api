package imageutil
import (
    "os"
    "image"
	"image/jpeg"
    logger "app/pkg/loggerutil"
)

/*
*  Remove EXIF Meta data from file
*/
func CleanMeta(filePath string) {
    inputFile, err := os.Open(filePath)
    if err != nil {
        logger.Log(err.Error(),"system.log")
        panic(err.Error())
    }
    defer inputFile.Close()

    img, _, err := image.Decode(inputFile)
    if err != nil {
        logger.Log(err.Error(),"system.log")
        panic(err.Error())
    }

    outfile, err := os.Create(filePath)
    if err != nil {
        logger.Log(err.Error(),"system.log")
        panic(err.Error())
    }
    defer outfile.Close()

    if err := jpeg.Encode(outfile, img, nil); err != nil {
        logger.Log(err.Error(),"system.log") 
        panic(err.Error())
    }
}