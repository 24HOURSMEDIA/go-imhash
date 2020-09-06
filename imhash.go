package imhash

import (
	"errors"
	"github.com/24HOURSMEDIA/go-imhash/imhash_interfaces"
	"github.com/24HOURSMEDIA/go-imhash/implementations/tom64b_hasher"
)

// CreateService is a factory method to create a new image hasher service
// by it's handle.
func CreateService(handle string) (imhash_interfaces.PerceptualHashImplementation, error) {
	switch handle {
	case "tom64b":
		return tom64b_hasher.Create(), nil
	}
	return nil, errors.New("Unrecognized image hasher handle " + handle)
}
