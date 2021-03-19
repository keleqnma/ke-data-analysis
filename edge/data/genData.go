package main

import (
	"os"
	"strconv"
	"time"

	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
	"github.com/sjwhitworth/golearn/knn"
)

const NODE_TYPES = 5
var cls *knn.KNNClassifier

type message struct {
	Msg string
}

func (msg *message) String() string {
	return msg.Msg
}

func main() {
	//source := ext.NewChanSource(tickerChan(time.Second * 1))
	pwd,_ := os.Getwd()
	path = pwd+"/edge/data"
	source := ext.NewFileSource(path+"/data.csv")
	flowMap := flow.NewMap(mapp, 1)
	sink := ext.NewStdoutSink()
	cls = knn.NewKnnClassifier("manhattan","kdtree",1)
	cls.Load(path+"/datasetclassifier")

	source.Via(flowMap).To(sink)

	select {}
}

var mapp = func(in interface{}) interface{} {
	msg := in.(*message)
	msg.Msg += "-UTC"
	return msg
}

func tickerChan(repeat time.Duration) chan interface{} {
	ticker := time.NewTicker(repeat)
	oc := ticker.C
	nc := make(chan interface{})
	go func() {
		for range oc {
			nc <- &message{strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
		}
	}()
	return nc
}