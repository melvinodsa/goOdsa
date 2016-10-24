package reverset

import "github.com/melvinodsa/goOdsa/utils"

/*RTransform will do the reverse transformation of the
data that was passed to it as the argument data. It returns a string
as reversed transformed data
*/
func RTransform(data utils.Data) []byte {
	return data.GetData()
}

/*rTransformWrapper is a wrapper function for concuurent
transformation of the input data to be processed in chunks
*/
func rTransformWrapper(channel chan utils.ChanByte, iString utils.ChanData) {
	result := utils.ChanByte{Output: RTransform(iString), Index: iString.Index}
	channel <- result
}

/*fTransformFactory creates the wrappers function necessary for the chunk processing
of data once the maxChunk limit is crossed. endChannel is used to communicate the finished
data array.
*/
func rTransformFactory(channel chan utils.ChanByte, endChannel chan [][]byte, data []utils.ChanData, maxChunk, startIndex, size int) {
	dataArray := make([][]byte, size)
	currentSize := 0
	for {
		//Getting each data from the channel
		chanOut := <-channel
		dataArray[chanOut.Index] = chanOut.Output
		currentSize++
		if startIndex < size {
			//Creating a wrapper function to process next chunk
			go rTransformWrapper(channel, data[startIndex])
			startIndex++
		} else if currentSize == size {
			endChannel <- dataArray
			<-channel
			return
		}
	}
}

/*RTransformChunk will reverse transform the chunk data with
go routines handling the RTransform. maxChunks has the maximum
 chunks to be processed concurrently.
endChannel will passed to the factory function to communicate the end result
 so that other can use the data.
*/
func RTransformChunk(maxChunks int, channel chan []utils.ChanData, dataChan chan utils.ChanByte, endChannel chan [][]byte) {
	for {
		data := <-channel
		size := len(data)
		startIndex := size
		if maxChunks < size {
			size = maxChunks
			startIndex = maxChunks
		}
		go rTransformFactory(dataChan, endChannel, data, maxChunks, startIndex, len(data))
		//Cuncurrently processing the data to be transformed
		for i := 0; i < size; i++ {
			go rTransformWrapper(dataChan, data[i])
		}
		return
	}
}
