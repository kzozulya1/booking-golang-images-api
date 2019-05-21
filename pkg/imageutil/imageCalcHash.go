package imageutil
import (
    "os"
    "crypto/md5"
    "encoding/hex"
    "encoding/binary"
    "io"
    "time"
)

/*
*  Calc md5 of file, append with unix timestamp
*/
func Md5sum(filePath string) (result string, err error) {
    file, err := os.Open(filePath)
    if err != nil {
        return
    }
    defer file.Close()

    hash := md5.New()
    _, err = io.Copy(hash, file)
    if err != nil {
        return
    }
    
    result = hex.EncodeToString(hash.Sum(nil)) + hex.EncodeToString(getUnixTimeStamp())
    return
}

/**
* Get unix timestamp in byte array
*/
func getUnixTimeStamp() []byte {
    unixTimeStamp := uint32(time.Now().Unix())
    bs := make([]byte, 4)
    binary.LittleEndian.PutUint32(bs, unixTimeStamp)
    return  bs
}