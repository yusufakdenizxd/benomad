package core

import (
	"fmt"
	"os"
)

// The utils file
// here mostly sub functions and other little things

func CheckDir() {
	if _, err := os.Stat(BenDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(BenDir, 0755); err != nil {
				fmt.Println("Failed to create benomad directory!", err)
				os.Exit(1)
			}
		}
	}
}
