package tom64b_hasher

import (
	"fmt"
	"github.com/24HOURSMEDIA/go-imhash/image_driver"
	"testing"
)

func TestImplementation_HashFromPath(t *testing.T) {
	img1 := "./../../resources/test1.jpg"
	for _, driver := range image_driver.SupportedImageDrivers {
		service := Create()
		service.Config.ImageDriver = driver
		hash, err := service.HashFromPath(img1)
		if err != nil {
			t.Fatal(err)
		}
		expectedHash, _ := service.HashFromString("8c1e07e8f86864f8")
		distance, _ := service.Distance(hash, expectedHash)
		if hash.String() != expectedHash.String() {
			t.Fatalf("Invalid hash result for driver %d, expected %s, got %s. Distance = %d", driver, expectedHash.String(), hash.String(), distance)
		}
	}
}

func ExampleDistance() {
	img1 := "./../../resources/test1.jpg"
	img2 := "./../../resources/test1_modified.jpg"
	img3 := "./../../resources/test1_modified2.jpg"

	service := Create()

	hash1, err := service.HashFromPath(img1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hash1.String())
	}

	hash2, err := service.HashFromPath(img2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hash2.String())
	}

	hash3, err := service.HashFromPath(img3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hash3.String())
	}

	distance, err := service.Distance(hash1, hash2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(distance)
	}

	distance13, err := service.Distance(hash1, hash3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(distance13)
	}

	distance23, err := service.Distance(hash2, hash3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(distance23)
	}

	// Output:
	// 8c1e07e8f86864f8
	// 8c1e07e8f86860f8
	// 8c1e07e8f86864f8
	// 1
	// 0
	// 1
}

func ExamplePopCount() {
	img1 := "./../../resources/test1.jpg"

	service := Create()

	_hash1, err := service.HashFromPath(img1)
	if err != nil {
		fmt.Println(err)
	} else {
		hash1 := _hash1.(tom64bHash)
		fmt.Println(hash1.String())
		fmt.Println(hash1.PopCount())
	}

	// Output:
	// 8c1e07e8f86864f8
	// 30
}

func ExampleDeserialize() {
	service := Create()

	hash, err := service.HashFromString("8c1e07e8f86864f8")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hash.String())
		fmt.Println(hash.(tom64bHash).PopCount())
	}

	// Output:
	// 8c1e07e8f86864f8
	// 30
}

func BenchmarkImplementation_HashFromPath_WithConvert(b *testing.B) {
	img1 := "./../../resources/large.jpg"
	service := Create()
	service.Config.ImageDriver = image_driver.ImageDriverImagickExecutable
	for i := 0; i < b.N; i++ {
		_, err := service.HashFromPath(img1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkImplementation_HashFromPath_WithImaging(b *testing.B) {
	img1 := "./../../resources/large.jpg"
	service := Create()
	service.Config.ImageDriver = image_driver.ImageDriverImaging
	for i := 0; i < b.N; i++ {
		_, err := service.HashFromPath(img1)
		if err != nil {
			b.Fatal(err)
		}
	}
}
