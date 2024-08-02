package utils

import "os"

func ApplicationArguments() []string {
	if os.Args == nil || len(os.Args) != 2 {
		return make([]string, 0)
	}
	return os.Args[1:]
}
