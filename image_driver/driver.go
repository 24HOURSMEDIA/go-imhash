package image_driver

const (
	ImageDriverNone = iota
	// use the convert command line tool
	ImageDriverImagickExecutable
	// use imagick lib driver
	ImageDriverImagickLib
	// use disintegration driver (default)
	ImageDriverImaging
)

var SupportedImageDrivers [3]int

func init() {
	SupportedImageDrivers[0] = ImageDriverImagickExecutable
	SupportedImageDrivers[1] = ImageDriverImagickLib
	SupportedImageDrivers[2] = ImageDriverImaging
}
