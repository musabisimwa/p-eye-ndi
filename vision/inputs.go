package vision

import (
	"fmt"

	"gocv.io/x/gocv"
)
const (
	DEFAULT_CAM =0
	SECONDARY_CAM=1
)

// ability to select any input device including a screen

type InputDevice struct{
	id string
	device *gocv.VideoCapture
}

func (device *InputDevice)setInputDevice(id int){
	var err error
	device.device,err =gocv.OpenVideoCapture(device)
	
	fmt.Printf("%s", err)


}