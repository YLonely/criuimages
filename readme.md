# criu-images

A Go package helps reading data from criu's images

Create a `Image` struct from a criu image:
```
import criuimg "github.com/YLonely/criuimages"
import "fmt"

i, err := criuimg.New("path/to/image/file")
fmt.Printf("Main type:%s Sub type:%s", i.MainType(), i.SubType())
```
Read an entry from image:
```
import types "github.com/YLonely/criuimages/types"

e := &types.MntEntry{}
err := i.ReadOne(e)
```