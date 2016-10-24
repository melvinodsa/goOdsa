package forwardt

import "github.com/melvinodsa/goOdsa/utils"

//FTransform does the forward transform of a given string
func FTransform(iString []byte) utils.Data {
	result := utils.NewODSAData()
	if len(iString) == 0 {
		return result
	}
	result.SetLastLetter(0)
	result.SetLastPos(-1)
	sLen := len(iString)
	//Algorithm starts
	for i := 0; i < sLen; i++ {
		result.AddData(iString[i])
		if iString[i] > result.GetLastLetter() {
			result.SetLastLetter(iString[i])
		}
		result.SetLastPos(i)
	}
	return result
}

/*fTransformWrapper is a wrapper function for concuurent
transformation of the input data to be processed in chunks
*/
func fTransformWrapper(channel chan utils.ChanData, iString []byte, index int) {
	result := FTransform(iString)
	channel <- utils.ChanData{Data: result, Index: index}
}

/*fTransformFactory creates the wrappers function necessary for the chunk processing
of data once the maxChunk limit is crossed. endChannel is used to communicate the finished
data array.
*/
func fTransformFactory(channel chan utils.ChanData, endChannel chan []utils.ChanData, strings [][]byte, size, startIndex int) {
	dataArray := []utils.ChanData{}
	currentSize := 0
	for {
		//Getting each data from the channel
		data := <-channel
		dataArray = append(dataArray, data)
		currentSize++
		if startIndex < size {
			//Creating a wrapper function to process next chunk
			go fTransformWrapper(channel, strings[startIndex], startIndex)
			startIndex++
		} else if currentSize == size {
			endChannel <- dataArray
			<-channel
			return
		}
	}
}

/*FTransformChunk will transform the chunk data with
go routines handling the FTransform strings is the
array of byte array containing the data to be compressed in
chunks. maxChunks has the maximum chunks to be processed concurrently.
endChannel will passed to the factory function to communicate the end result
 so that other can use the data.
*/
func FTransformChunk(strings [][]byte, maxChunks int, channel chan utils.ChanData, endChannel chan []utils.ChanData) {
	size := len(strings)
	useFactory := false
	if maxChunks < size {
		size = maxChunks
		useFactory = true
	}
	//Cuncurrently processing the data to be transformed
	for i := 0; i < size; i++ {
		go fTransformWrapper(channel, strings[i], i)
	}
	//Once the max Limit is reached then using the factory function
	// to proocess the rest of the chunks
	if useFactory {
		size = len(strings)
		go fTransformFactory(channel, endChannel, strings, size, maxChunks)
	}
}
