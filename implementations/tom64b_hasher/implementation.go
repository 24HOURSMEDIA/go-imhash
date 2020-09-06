// Package tom64b_hasher is a perceptual image hashing Implementation based on
// the php Implementation at https://github.com/Tom64b/dHash
package tom64b_hasher

import (
	"errors"
	"fmt"
	"github.com/24HOURSMEDIA/go-imhash/environment"
	"github.com/24HOURSMEDIA/go-imhash/imhash_interfaces"
	"github.com/24HOURSMEDIA/go-imhash/util"
	"github.com/google/uuid"
	"github.com/tmthrgd/go-popcount"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// distanceCounts is a global required to calculate distances between hashes
var distanceCounts = [16]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

// tom64bHash is the specific Implementation of the hash
type tom64bHash struct {
	hash     string
	popCount uint
}

// String returns the hash as a string, usually a hexadecimal representation
func (hash tom64bHash) String() string {
	return hash.hash
}

// PopCount returns the number of '1' bits in the hash
// It may be used to quickly select other 'near' hashes in a database
// See https://github.com/Tom64b/dHash/issues/1
func (hash tom64bHash) PopCount() uint {
	return hash.popCount
}

// Implementation is the specific Implementation of the hashing service
type Implementation struct {
}

// Create creates a new instance of the hashing service
func Create() Implementation {
	return Implementation{}
}

func (imp Implementation) GetHandle() string {
	return "tom64b"
}

// HashFromPath creates a hash from an image file at the given path
func (imp Implementation) HashFromPath(path string) (imhash_interfaces.PerceptualHash, error) {
	return imp.newHashFromFileWithImagick(path)
}

// HashFromString recreates a hash from a hash string
// use it to recreate a hash from a stored value
func (imp Implementation) HashFromString(hashAsString string) (imhash_interfaces.PerceptualHash, error) {
	decimal, err := strconv.ParseUint(hashAsString, 16, 64)
	if err != nil {
		return nil, err
	}
	slice := make([]uint64, 1)
	slice[0] = decimal
	return tom64bHash{hash: hashAsString, popCount: uint(popcount.CountSlice64(slice))}, nil
}

// Distance calculates the hamming distance between two hashes
func (imp Implementation) Distance(h1 imhash_interfaces.PerceptualHash, h2 imhash_interfaces.PerceptualHash) (imhash_interfaces.HammingDistance, error) {
	hash1 := h1.(tom64bHash)
	hash2 := h2.(tom64bHash)
	result := 0
	for i := 0; i < 16; i++ {
		if hash1.hash[i] != hash2.hash[i] {
			v1, _ := strconv.ParseUint(string(hash1.hash[i]), 16, 8)
			v2, _ := strconv.ParseUint(string(hash2.hash[i]), 16, 8)
			result += distanceCounts[v1^v2]
		}
	}
	return imhash_interfaces.HammingDistance(result), nil
}

// newHashFromGrayScaleMatrix creates a DHash from a 9x8 matrix object
// with grayscale values
func (imp Implementation) newHashFromGrayScaleMatrix(matrix util.GrayScaleMatrix) (imhash_interfaces.PerceptualHash, error) {
	var hash = uint64(0)
	var bit = uint64(1)
	bitCount := uint(0)
	for y := 0; y < 8; y++ {
		previous := matrix.GetAt(0, y)
		for x := 1; x < 9; x++ {
			current := matrix.GetAt(x, y)
			if previous > current {
				hash = hash | bit
				bitCount++
			}
			bit = bit * 2
			previous = current
		}
	}
	return tom64bHash{hash: fmt.Sprintf("%x", hash), popCount: bitCount}, nil
}

// newHashFromPreparedImage creates an image hash from a 9x8 image object in colour.
func (imp Implementation) newHashFromPreparedImage(image image.Image) (imhash_interfaces.PerceptualHash, error) {
	matrix := util.CreateGrayscaleMap(image, util.NewBT601Weights())
	return imp.newHashFromGrayScaleMatrix(matrix)
}

// newHashFromPreparedFile creates an image hash from a prepared file at the specified path
// The prepared file must be a 9x8 png image in colour
func (imp Implementation) newHashFromPreparedFile(path string) (imhash_interfaces.PerceptualHash, error) {
	// Read image from file that already exists
	imageFile, err := os.Open(path)
	if err != nil {
		return tom64bHash{}, err
	}
	defer imageFile.Close()
	img, err := png.Decode(imageFile)
	if err != nil {
		return nil, err
	}
	return imp.newHashFromPreparedImage(img)
}

// newHashFromFileWithImagick creates a hash from an image file using the image magick convert utility
func (imp Implementation) newHashFromFileWithImagick(sourcePath string) (imhash_interfaces.PerceptualHash, error) {
	targetPath := filepath.Join(environment.WorkDir, uuid.New().String()+".png")
	// see: http://www.imagemagick.org/Usage/filter/
	output, err := exec.Command("convert", sourcePath, "-filter", "box", "-resize", "9x8!", targetPath).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}
	defer os.Remove(targetPath)
	return imp.newHashFromPreparedFile(targetPath)
}
