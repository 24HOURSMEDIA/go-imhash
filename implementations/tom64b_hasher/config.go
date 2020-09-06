package tom64b_hasher

type Config struct {

	// UseImagickLib is a flag indicating wether to use the imagick binding or the convert utility
	UseImagickLib bool
	// ImagickExecutable is the path where the 'convert' utility of imagick can be found
	ImagickExecutable string
}

func NewConfig() Config {
	cfg := Config{
		UseImagickLib:     true,
		ImagickExecutable: "convert",
	}
	return cfg
}
