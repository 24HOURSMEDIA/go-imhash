package image_driver

const (
	ImageDriverNone = iota
	// use the convert command line tool
	ImageDriverImagickExecutable
	// use disintegration driver (default)
	ImageDriverImaging
)

var SupportedImageDrivers [2]int

func init() {
	SupportedImageDrivers[0] = ImageDriverImagickExecutable
	SupportedImageDrivers[1] = ImageDriverImaging
}
