package version

import "fmt"

func Add(version int, path string) string {
	return path + fmt.Sprintf("?v=%d", version)
}

