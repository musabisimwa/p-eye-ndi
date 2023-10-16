package vision

import (
	"gocv.io/x/gocv"
	"errors"
	"log"
)



func loadModel(modelPath string ,configFile string) (error, gocv.Net) {
	net := gocv.ReadNet(modelPath,configFile)
	if !net.Empty()  {
		return nil,net
	}
	return errors.New("couldn't load model from path"),net
}

func Initialize(modelPath string, confPath string){
	err,net := loadModel(modelPath,confPath)

	if err !=nil{
		//log error
		log.Fatalln("there was an error while loading model")
	}
	net.SetPreferableBackend(gocv.NetBackendDefault)


}


