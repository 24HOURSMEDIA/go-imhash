package imhash

import (
	"errors"
	"github.com/24HOURSMEDIA/go-imhash/implementations/tom64b_hasher"
	"github.com/24HOURSMEDIA/go-imhash/interfaces"
)

// CreateService is a factory method to create a new image hasher service
// by it's handle.
func CreateService(handle string) (interfaces.PerceptualHashImplementation, error) {
	switch handle {
	case "tom64b":
		return tom64b_hasher.Create(), nil
	}
	return nil, errors.New("Unrecognized image hasher handle " + handle)
}
