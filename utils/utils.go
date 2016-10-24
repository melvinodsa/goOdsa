/*
Package utils has all the utility functions and data structures required for the
ODSA algorithm to work.
*/
package utils

/*findPosition returns the position and insert position of the given character in the PCArray.
-> If character found, then its position and insert position -1 is returned
-> If character not found, then its position would be -1 and its insert position is returned
    -> insert Position would be -1 if the character has to be inserted at the end of the array
-> In error case like the array and map is not in sync it will return -1 and -2
*/
func (data *ODSAData) findPosition(char byte) (int, int) {
	/*Checking whether the character exists in the PMap.
	  If missing in the PMap, then it won't exist in PCArray as PCArray would be
	  in sync with PMap
	*/
	if val, ok := data.pMap[char]; ok {
		return val, -1
	}
	/*The given character is not existing in the PCArray.
	Now searching for the position to insert the character
	*/
	if data.lLetter < char {
		//Character has to be inserted at the end
		return -1, -1
	}
	if len(data.pCArray) == 0 || char < data.pCArray[0] {
		return -1, -3
	}
	var length = len(data.pCArray)
	for i := 0; i < length-1; i++ {
		//Character has to be inserted in the sorted manner
		if char > data.pCArray[i] && char < data.pCArray[i+1] {
			return -1, i + 1
		}
	}
	//Error state
	return -1, -2
}

/*AddData adds a character to the ODSAData
It returns true if character addition was sucessfull
*/
func (data *ODSAData) AddData(char byte) bool {
	pos, iPos := data.findPosition(char)
	//Error state
	if iPos == -2 {
		return false
	}
	//Character not present in the Array. So adding it to the ODSA data
	if pos == -1 && iPos == -1 {
		data.pCArray = append(data.pCArray, char)             //Updating the PCArray
		data.pIArray = append(data.pIArray, data.lPosition+1) //Updating the PIArray with the index of the entry
		data.pMap[char] = len(data.pCArray) - 1               ///Syncing up the PMap w.r.t. the arrays
		return true
	}
	/*Character is present in the array. So put the offset of 1 to all
	the characters after it as the array is in sorted state
	*/
	if pos != -1 {
		data.updatePos(pos+1, 1, false, true) //Updating the indices
		if data.lLetter > char {
			//Recording the noise information
			data.nCArray = append(data.nCArray, char)
			data.nIArray = append(data.nIArray, data.lPosition+1)
		}
		return true
	}
	//Inserting the character at the begining of the position array
	if iPos == -3 {
		data.pCArray = append([]byte{char}, data.pCArray[:]...)
		data.pIArray = append([]int{0}, data.pIArray[:]...)
		data.pMap[char] = 0
		data.updatePos(1, 1, true, true)
		//Recording the noise information
		data.nCArray = append(data.nCArray, char)
		data.nIArray = append(data.nIArray, data.lPosition+1)
		return true
	}
	//appending in the given insert position``
	data.pCArray = append(data.pCArray[0:iPos], append([]byte{char}, data.pCArray[iPos:]...)...)
	//Updating the PIArray with insert and updated indices that comes after the same
	data.pIArray = append(data.pIArray[0:iPos], append([]int{data.pIArray[iPos]}, data.pIArray[iPos:]...)...)
	//Updating the trailing indicies with new value after the addition of new element
	data.updatePos(iPos+1, 1, true, true)
	//Updating the map to sync up the data with the arrays
	data.pMap[char] = iPos
	//Recording the noise information
	data.nCArray = append(data.nCArray, char)
	data.nIArray = append(data.nIArray, data.lPosition+1)
	return true
}

/*GetData returns the original data that was transformed
and compressed
*/
func (data *ODSAData) GetData() []byte {
	//Initializing the sorted variable
	sorted := []byte{}
	length := len(data.pCArray) - 1
	/*Building the sorted string by picking the psoiton index array and
	position index array
	*/
	for i := 0; i < length; i++ {
		for j := 0; j < (data.pIArray[i+1] - data.pIArray[i]); j++ {
			sorted = append(sorted, data.pCArray[i])
		}
	}
	//Handling the last character in the psotion array with lastPosition
	for j := 0; j < (data.lPosition + 1 - data.pIArray[length]); j++ {
		sorted = append(sorted, data.pCArray[length])
	}
	//Rebuilding the string from the noise fingerprint
	length = len(data.nCArray) - 1
	pLen := len(sorted)
	for i := length; i >= 0; i-- {
		/*Repositioning the noise character to its original
		postion from the sorted position
		*/
		splitPos := data.pIArray[data.pMap[data.nCArray[i]]]
		//Note the nIArray[i] will always be greater than splitPos
		sorted = append(sorted[0:splitPos], sorted[splitPos+1:]...)
		if data.nIArray[i] < pLen {
			sorted = append(sorted[0:data.nIArray[i]], append([]byte{data.nCArray[i]}, sorted[data.nIArray[i]:]...)...)
		} else {
			sorted = append(sorted, data.nCArray[i]) //appending to the end if position is last
		}
		/*/Although the positions pIArray[i] gets updated
		it won't affect the remaining alogirthm as the positions past
		pIArray[i] are not referenced again
		*/
		data.updatePos(data.pMap[data.nCArray[i]]+1, -1, false, true)
	}
	return sorted
}

//NewODSAData is the constructor function for ODSAData
func NewODSAData() Data {
	data := new(ODSAData)
	data.lPosition = 0
	data.pMap = make(map[byte]int)
	return data
}

/*updatePos updates the position of the
data. The arguments startValue denotes the position from which the updation
has to start from. Increament denotes the amount with which the data has to
be updated. mMap and pIArray boolean variables indicate whether the map/pIArray
of the data has to be updated
*/
func (data *ODSAData) updatePos(startValue, increament int, pMap, pIArray bool) {
	length := len(data.pIArray)
	for i := startValue; i < length; i++ {
		//Updating the map/array witn the increament
		if pIArray {
			data.pIArray[i] += increament
		}
		if pMap {
			data.pMap[data.pCArray[i]] += increament
		}
	}
}
