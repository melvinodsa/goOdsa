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
