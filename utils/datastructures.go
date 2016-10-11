package utils

import "strconv"

//ODSAData is the structure for the compressed data
type ODSAData struct {
	//pMap is the position map of the compressed data
	pMap map[byte]int
	//lPosition is the last position of the compressed data
	lPosition int
	//lLetter is the last letter in the transformed data
	lLetter byte
	//pCArray is the position-character mapping array
	pCArray []byte
	//pIArray is the position-indedx mapping array
	pIArray []int
	//nCArray is the character array of noise
	nCArray []byte
	//nIArray is the index mapping of the noise with the character
	nIArray []int
}

//GetLastPos returns the last position of the data
func (data *ODSAData) GetLastPos() int {
	return data.lPosition
}

//SetLastPos sets the last position of the data
func (data *ODSAData) SetLastPos(pos int) {
	data.lPosition = pos
}

//GetLastLetter returns the last letter of the data
func (data *ODSAData) GetLastLetter() byte {
	return data.lLetter
}

//SetLastLetter sets the last letter of the data
func (data *ODSAData) SetLastLetter(char byte) {
	data.lLetter = char
}

//ToString converts the data to plain text
func (data *ODSAData) ToString() string {
	output := "position -> ["
	length := len(data.pIArray)
	/*
		Iterating through the position arrays to get the value
	*/
	for i := 0; i < length; i++ {
		output += " " + string(data.pCArray[i]) + " : " + strconv.Itoa(data.pIArray[i])
		if i < length-1 {
			output += ","
		}
	}
	output += " ]\nnoise -> ["
	length = len(data.nIArray)
	/*
		Iterating through the noise arrays to get the value
	*/
	for i := 0; i < length; i++ {
		output += " " + string(data.nCArray[i]) + " : " + strconv.Itoa(data.nIArray[i])
		if i < length-1 {
			output += ","
		}
	}
	output += " ]\nlast letter -> " + strconv.Itoa(data.lPosition)
	//Returning the output text
	return output
}

//Data interface to protect the Data
type Data interface {
	AddData(byte) bool
	GetData() string
	ToString() string
	GetLastPos() int
	SetLastPos(int)
	GetLastLetter() byte
	SetLastLetter(byte)
}
