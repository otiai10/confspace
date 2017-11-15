package confspace

import "fmt"

func ExampleSpace_Open() {

	space, _ := New("myproject")
	f, _ := space.Open("config.json")
	defer f.Close()
	f.Write([]byte("192.168.0.1"))

	fmt.Println(f.Name())
	// Output:
	// /Users/otiai10/.myproject/config.json
}
