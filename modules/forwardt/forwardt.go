package forwardt

import "github.com/melvinodsa/goOdsa/utils"

//FTransform does the forward transform of a given string
func FTransform(iString string) utils.Data {
	result := utils.NewODSAData()
	if len(iString) == 0 {
		return result
	}
	result.SetLastLetter(iString[0] - 1)
	result.SetLastPos(-1)
	sLen := len(iString)
	//Algorithm starts
	for i := 0; i < sLen; i++ {
		result.AddData(iString[i])
		result.SetLastLetter(iString[i])
		result.SetLastPos(i)
	}
	return result
}
