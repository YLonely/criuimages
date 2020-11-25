package criuimages

import "os"

const (
	ImageTypeInventory string = "Inventory"
	ImageTypeNormal    string = "Normal"
	ImageTypeAuxiliary string = "Auxiliary"
)

func New(path string) (*Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

}

type Image struct {
	mainType string
	subType  string
	file     *os.File
}
