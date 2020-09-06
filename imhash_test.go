package imhash

import "testing"

func TestCreateService(t *testing.T) {

	handles := []string{"tom64b"}
	for _, handle := range handles {
		_, err := CreateService(handle)
		if err != nil {
			t.Fail()
		}
	}
}
