package vision

import (
	"image"
	"log"

	"gocv.io/x/gocv"
)

// trackTarget is a function that tracks a target (e.g., a face) in a video feed.
func TrackTarget(camera int) {
    xmlFile:= "data/haarcascade_frontalface_default.xml"
    // Open a video capture device (webcam)
    webcam, err := gocv.OpenVideoCapture(camera)
    if err != nil {
        panic(err)
    }
    defer webcam.Close()

    // Open a window to display the video feed
    window := gocv.NewWindow("Target Tracking")
    defer window.Close()

    // Load the pre-trained face detection classifier
    classifier := gocv.NewCascadeClassifier()
    defer classifier.Close()
    if !classifier.Load(xmlFile) {
        log.Fatalln("could not load facial detection  data file")
        return
    }

    // Create a Mat to store the current frame
    frame := gocv.NewMat()
    defer frame.Close()

    // Initialize the cropping region
    cropRect := image.Rect(0, 0, 0, 0)

    // Capture and process frames in a loop
    for {
        if ok := webcam.Read(&frame); !ok {
            break
        }
        if frame.Empty() {
            continue
        }

        // Convert the frame to grayscale for face detection
        gray := gocv.NewMat()
        defer gray.Close()
        gocv.CvtColor(frame, &gray, gocv.ColorBGRToGray)

        // Detect faces in the frame
        faces := classifier.DetectMultiScale(gray)
        if len(faces) > 0 {
            // Get the first detected face
            face := faces[0]

            // Calculate the center of the detected face
            centerX := face.Min.X + (face.Max.X-face.Min.X)/2
            centerY := face.Min.Y + (face.Max.Y-face.Min.Y)/2

            // Calculate the new cropping region to keep the face centered
            cropRect = image.Rect(centerX-100, centerY-100, centerX+100, centerY+100)

            // Ensure the cropping region is within the frame bounds
         //   cropRect = cropRect.Intersect(frame.Bounds())
			
        }

        // Crop the frame to the defined region
        croppedFrame := frame.Region(cropRect)

        // Display the cropped frame with detected faces
        window.IMShow(croppedFrame)

        // Check for user input to exit
        key := window.WaitKey(1)
        if key == 27 { // 'Esc' key
            break
        }
    }
} 


// test is open cv is working 
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