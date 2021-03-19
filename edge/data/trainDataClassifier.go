package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

var(
	trainData, testData base.FixedDataGrid
	rawData *base.DenseInstances
	maxAccuracy         int64
	path                string
)

func main() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	pwd,_ := os.Getwd()
	path = pwd+"/edge/data/dataset"
	maxAccuracy = 0

	rawData, err := base.ParseCSVToInstances(path+"/data.csv", false)
	if err != nil {
		panic(err)
	}
	fmt.Println("file read", time.Now())
	//Do a training-test split
	trainData, testData = getMiniData(rawData,6)

	distFuncs := []string{"euclidean","manhattan"}
	algorithms := []string{"kdtree"}
	knnNeighbors := []int{1}

	var wg sync.WaitGroup
	for _, distFunc := range distFuncs{
		for _, algorithm := range algorithms {
			for _, neighbor := range knnNeighbors{
				wg.Add(1)
				go testKNNClassifier(distFunc,algorithm,neighbor,&wg)
			}
		}
	}

	wg.Wait()

	fmt.Println("maxAccu",float64(maxAccuracy)/1000)
}

func getMiniData(rawData *base.DenseInstances,times int)(base.FixedDataGrid,base.FixedDataGrid){
	var trainData, testData,data base.FixedDataGrid
	data = rawData
	for i:=0;i<times;i++{
		trainData, testData = base.InstancesTrainTestSplit(data, 0.50)
		data = testData
	}
	return trainData, testData
}

func testKNNClassifier(distFunc string, algorithm string,neighbor int,wg *sync.WaitGroup ){
	defer wg.Done()
	//Initialises a new KNN classifier
	fmt.Println(distFunc,algorithm,neighbor,"begin train", time.Now())
	cls := knn.NewKnnClassifier(distFunc, algorithm, neighbor)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(distFunc,algorithm,neighbor,"\n",evaluation.GetSummary(confusionMat))

	curAccuracy := int64(1000*evaluation.GetAccuracy(confusionMat))
	if curAccuracy > atomic.LoadInt64(&maxAccuracy){
		fmt.Println("maxAccu",distFunc,algorithm,neighbor,evaluation.GetAccuracy(confusionMat))
		atomic.StoreInt64(&maxAccuracy,curAccuracy)
		cls.Save(path+"classifier")
	}
}

