package coursera

import (
	"fmt"
	"strconv"
)

func main() {
	var x float32
	var truncInt int
	fmt.Scan(&x)
	fmt.Scan(&truncInt)
	fmt.Printf("%."+strconv.Itoa(truncInt)+"f", x)
}
