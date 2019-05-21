package api

import (
    "net/http"
	"github.com/gin-gonic/gin"  
     logger "app/pkg/loggerutil"
     imageutil "app/pkg/imageutil"
     config "app/pkg/settings"
)

/*
*  Upload file
*/
func Upload(c *gin.Context) {
    //Apply headers
    ApplyHeaders(c)
    
    //Get `file` param
    file, err := c.FormFile("file")
    if err != nil {
        logger.Log("Error: " + err.Error(),"system.log")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return;
    }

    //Save uploaded file
    err = c.SaveUploadedFile(file, config.IMAGES_DIR+file.Filename)
    if err != nil {
        logger.Log("Error: " + err.Error(),"system.log")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return;
    }
    
    //Clear meta data from EXIF
    imageutil.CleanMeta(config.IMAGES_DIR+file.Filename)
    
    //Calculate hash of file
    hash, err :=  imageutil.Md5sum(config.IMAGES_DIR+file.Filename)
    if err != nil {
        logger.Log("Error calculate md5: " + err.Error(),"system.log")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return;
    }
    
    //Rename file
    hasedFilename, err := imageutil.RenameKeepExtension(config.IMAGES_DIR,file.Filename,hash)
    if err != nil {
        logger.Log("Rename error: " + err.Error(),"system.log")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return;
    }
    c.JSON(http.StatusCreated, gin.H{"id": hasedFilename})
}
