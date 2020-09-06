package imhash

// PerceptualHash represents the hash of an image
type PerceptualHash interface {
	// String returns the hash as a string, usually a hexadecimal representation
	String() string
}

// HammingDistance represents the distance between two hashes
type HammingDistance int

type PerceptualHashImplementation interface {

	// Handle returns a canonical name for the hashing service
	// use it so you can refer to the specific algorithm later
	Handle() string

	// HashFromPath creates a hash from an image file at the given path
	HashFromPath(path string) (PerceptualHash, error)

	// HashFromString recreates a hash from a hash string
	// use it to recreate a hash from a stored value
	HashFromString(hashAsString string) (PerceptualHash, error)

	// Distance calculates the hamming distance between two hashes
	Distance(hash1 PerceptualHash, hash2 PerceptualHash) (HammingDistance, error)
}
