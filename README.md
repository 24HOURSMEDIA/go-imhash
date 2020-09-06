# GO-DHASH

Calculate fast image distance hash in Go.

Requires Imagick to be installed (with the 'convert' command line tool).

Adapted from: https://github.com/Tom64b/dHash,
see: 
* http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
* https://github.com/jenssegers/imagehash/

Foto door Jeremy Bishop via Pexels

## Example

```go
package main

import (
	"fmt"
    "github.com/24HOURSMEDIA/go-dhash/implementations/tom64b_hasher"
)

func main() {
    // show hashed values and calculate distance hasher := tom64b_hasher.Create()

    imgPath1 := "./image1.jpg"
    imgPath2 := "./image2.jpg"
   
    hash1, _ := hasher.HashFromPath(imgPath1)
    hash2, _ := hasher.HashFromPath(imgPath2)
    fmt.Println(hash1.String())
    fmt.Println(hash2.String())
    fmt.Println(hasher.Distance(hash1, hash2))
    
    // show the population counts
    fmt.Println(hash1.(tom64b_hasher.Hash).PopCount())

}
```

## Tests

```
go test github.com/24HOURSMEDIA/go-imhash/implementations/tom64b_hasher
```