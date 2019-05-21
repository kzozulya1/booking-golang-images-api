package loggerutil
import (
    "os"
    "log"
)

func Log(msg string, filename string) {
    f, _ := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    defer f.Close()
    log.SetOutput(f)
    log.Print(msg)
}
