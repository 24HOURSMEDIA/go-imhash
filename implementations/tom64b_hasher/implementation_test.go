package tom64b_hasher

import "fmt"

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
