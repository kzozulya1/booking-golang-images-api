package api

import (
    "net/http"
	"github.com/gin-gonic/gin"  
    "github.com/nfnt/resize"
    "os"
    "io"
    "strconv"
    "image/jpeg"
    config "app/pkg/settings"
    logger "app/pkg/loggerutil"
)

/*
*  Download file
*/
func Download(c *gin.Context) {
    image := c.Query("id")
    
    //Check ID Param is ok
    if len(image) < 35 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or incorrect param: id"})
        return;
    }
    
    fullImagePath := getFullPath(image)
    //Check File Exists
    if !checkFileExists(fullImagePath){
        c.JSON(http.StatusNotFound, gin.H{"error": "File " +  image + " not found"})
        return;
    }
    
    //Check if consumer want to resize file:
    _width, _  := strconv.ParseUint(c.DefaultQuery("w", "0"),10,32)
    width := uint(_width)
    _height, _ := strconv.ParseUint(c.DefaultQuery("h", "0"),10,32)
    height := uint(_height)
    
    //Check file already has resized version    
    if _cachedFilepath := resizeImage(config.IMAGES_DIR+filenameSubdir(image), image, width, height); _cachedFilepath != "" {
        fullImagePath = _cachedFilepath
    }
    
    //Open image file and send its content
    imageData, err := os.Open(fullImagePath)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Can't open " +  fullImagePath})
        return;
    }
    defer imageData.Close()
    
    c.Header("Content-Type", "image/jpeg")
    io.Copy(c.Writer, imageData)
}

/*
* Resize image, if already exists - then return resized
* Return the path to new image
*/
func resizeImage(imgDir, imgFilename string, w, h uint) string {  
    
    wStr :=  strconv.FormatUint(uint64(w), 10)
    hStr :=  strconv.FormatUint(uint64(h), 10)

    resizedFilename := wStr + "x" + hStr + "_" + imgFilename
    
    //If resized file already exists - then render it
    if checkFileExists(imgDir + resizedFilename) {
        return imgDir + resizedFilename
    }
    
    //Open image file
	file, err := os.Open(imgDir + imgFilename)
	if err != nil {
		logger.Log(err.Error(), "resizer.error.log")
        return ""
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		logger.Log(err.Error(), "resizer.error.log")
        return ""
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	resizedImage := resize.Resize(w, h, img, resize.Lanczos3) 

	//Store resized image
    out, err := os.Create(imgDir + resizedFilename)
	if err != nil {
		logger.Log(err.Error(), "resizer.error.log")
        return ""
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, resizedImage, nil)
    return imgDir + resizedFilename
}


/*
* Get subdir, where file is located
*/
func filenameSubdir(filename string) string {
    firstSubfolder := filename[:1]
    secondSubfolder := filename[1:2] 
    return firstSubfolder + "/" +  secondSubfolder + "/"
}

/*
* Check file exists in FS
*/
func checkFileExists(fullpath string) bool {
    _, err := os.Stat(fullpath);    
    return !os.IsNotExist(err)
}

/*
* Get full path for filename
*/
func getFullPath(filename string) string {
    return config.IMAGES_DIR +  filenameSubdir(filename) +  filename
}