# persist

This package provides functions for saving and loading data from files.

## Functionality

The package provides two functions:

* `Save(data []byte, filename string) error`: This function saves a byte slice to a file.
* `Load(path string) ([]byte, error)`: This function loads a byte slice from a file.

## Usage

```go
import (
	"fmt"
	"persist"
)

func main() {
	data := []byte("Hello, world!")
	err := persist.Save(data, "data.txt")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	loadedData, err := persist.Load("data.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	fmt.Println(string(loadedData)) // Output: Hello, world!
}
```

## License

The bitarray package is licensed under the MIT License. See the LICENSE file for details.
