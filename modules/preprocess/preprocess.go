/*Package preprocess has the functions to do preprocessing
of the data berfore compressing it. Files has to be categorized and
and transformed depending up on the format. Bigger Files has to be
broke into chunks. This package deals with these activities
*/
package preprocess

/*Chunkify converts the given bytes into
chunk of bytes. Bigger data can there by can be split into bytes
and processed each chunk seperately. The order or chunks should
be preserved to get accurate data while decompressing*/
func Chunkify(data []byte, chunkSize int) [][]byte {
	output := [][]byte{}
	size := len(data)
	if size <= chunkSize {
		return append(output, data)
	}
	for i := 0; i < size/chunkSize; i++ {
		output = append(output, data[i*chunkSize:(i+1)*chunkSize])
	}
	if size%chunkSize != 0 {
		output = append(output, data[(size/chunkSize)*chunkSize:])
	}
	return output
}

/*DeChunkify aggregates the chunks into one piece of data.append
It takes the chunk array as an argument and returns the unified chunk data
*/
func DeChunkify(data [][]byte) []byte {
	output := []byte{}
	size := len(data)
	for i := 0; i < size; i++ {
		output = append(output, data[i]...)
	}
	return output
}
