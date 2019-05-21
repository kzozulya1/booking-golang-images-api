package imageutil
import (
    "path/filepath"
    "os"
)

/*
*  Rename file and keep old extension
*/
func RenameKeepExtension(dir string, fileName string, newFilename string) (hasedFilename string, err error) {
    
    firstSubfolder := newFilename[:1]
    secondSubfolder := newFilename[1:2] 
    
    touchDir(dir + firstSubfolder)
    touchDir(dir + firstSubfolder + "/" + secondSubfolder)
    
    extension := filepath.Ext(fileName)
    
    err = os.Rename(dir + fileName,  dir + firstSubfolder + "/" + secondSubfolder + "/" + newFilename + extension )
    hasedFilename = newFilename + extension
    return
}

/*
* Create dir if not exists
*/
func touchDir(path string){
    if _, err := os.Stat(path); os.IsNotExist(err) {
        err = os.MkdirAll(path, 0777)
          if err != nil {
              panic(err)
          }
    }
}
