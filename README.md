# goc2l

light-well cli Cyrillic translator for web (utf8). Part of OwnCDN utils


## Install

go install github.com/gadost/goc2l

## Usage

```bash
$ goc2l "какое-то название файла.mp4"

kakoe_to_nazvanie_fayla.mp4
```

## As module

```golang
import (
    "fmt"
    "github.com/gadost/goc2l"
)

func main() {
    goc2l.Print("какое-то название файла.mp4")
    x := goc2l.Convert("какое-то название файла.mp4")
    y := goc2l.ConvertWithSpecials("какое-то название файла.mp4")
    fmt.Println(x)
    fmt.Println(y)
}
```
