package main
import(
	"p-eye-ndi/vision"
	"p-eye-ndi/ndi"
	"fmt"
)
func main() {
vision.CvTest()
fmt.Printf("%s",ndi.Version())

}

