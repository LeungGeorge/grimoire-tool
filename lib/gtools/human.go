package gtools

import "fmt"

const (
	sizeGBits = 1 << 30
	sizeMBits = 1 << 20
	sizeKBits = 1 << 10
)

// HumanBytes TODO
// 单位转换：
// 1024B -> 1K
// 1024K -> 1M
// 1024M -> 1G
func HumanBytes(bytes uint64) string {
	suffix := "B"
	if bytes > sizeGBits {
		suffix = "G"
		bytes /= sizeGBits
	} else if bytes > sizeMBits {
		suffix = "M"
		bytes /= sizeMBits
	} else if bytes > sizeKBits {
		suffix = "K"
		bytes /= sizeKBits
	}

	return fmt.Sprintf("%d%s", bytes, suffix)
}
