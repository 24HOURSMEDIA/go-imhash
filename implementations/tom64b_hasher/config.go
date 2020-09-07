package tom64b_hasher

import (
	"github.com/24HOURSMEDIA/go-imhash/image_driver"
)

type Config struct {

	// UseImagickLib is a flag indicating wether to use the imagick binding or the convert utility
	ImageDriver int
	// ImagickExecutable is the path where the 'convert' utility of imagick can be found
	ImagickExecutable string
}

func NewConfig() Config {
	cfg := Config{
		// ImageDriver is the driver used for preprocessing images
		ImageDriver: image_driver.ImageDriverImagickExecutable,
		// ImagickExecutable is the location of the Imagick executable
		ImagickExecutable: "convert",
	}
	return cfg
}
