package criuimages

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

const (
	ImageTypeInventory string = "inventory"
	ImageTypeNormal    string = "normal"
	ImageTypeAuxiliary string = "auxiliary"
)

func New(path string) (*Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	i := &Image{
		file: f,
	}
	if err = i.checkType(); err != nil {
		return nil, errors.Wrap(err, "error create image")
	}
	return i, nil
}

type Image struct {
	mainType string
	subType  string
	file     *os.File
}

func (i *Image) MainType() string {
	return i.mainType
}

func (i *Image) SubType() string {
	return i.subType
}

func (i *Image) ReadOne(obj proto.Message) error {
	data := make([]byte, 4)
	n, err := i.file.Read(data)
	if n == 0 && err == io.EOF {
		return io.EOF
	}
	if n != len(data) {
		return errors.New("error read data")
	}
	size := binary.LittleEndian.Uint32(data)
	msgData := make([]byte, int(size))
	n, err = i.file.Read(msgData)
	if n != len(msgData) || err != nil {
		return errors.New("error read message data")
	}
	if err = proto.Unmarshal(msgData, obj); err != nil {
		return errors.Wrap(err, "failed to unmarshal message")
	}
	return nil
}

func (i *Image) Close() error {
	if i.file != nil {
		return nil
	}
	return i.file.Close()
}

// File returns the file under the image
func (i *Image) File() *os.File {
	return i.file
}

func (i *Image) checkType() error {
	data := make([]byte, 4)
	n, err := i.file.Read(data)
	if n != len(data) || err != nil {
		return errors.New("error read image")
	}
	t := binary.LittleEndian.Uint32(data)
	typeName, exists := imageTypes[t]
	if !exists {
		return errors.New("error image type")
	}
	i.mainType = typeName
	if typeName == ImageTypeInventory {
		return nil
	}
	n, err = i.file.Read(data)
	if n != len(data) || err != nil {
		return errors.New("error read image's subtype")
	}
	t = binary.LittleEndian.Uint32(data)
	typeName, exists = imageTypes[t]
	if !exists {
		return errors.New("error image subtype")
	}
	i.subType = typeName
	return nil
}
