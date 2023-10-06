package vision

import (
	"gocv.io/x/gocv"
)


func CvTest(){
	webcam,_:= gocv.OpenVideoCapture(0)
	window:= gocv.NewWindow("hello")
	img:= gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}