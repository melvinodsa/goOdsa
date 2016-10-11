package reverset

import "github.com/melvinodsa/goOdsa/utils"

/*RTransform will do the reverse transformation of the
data that was passed to it as the argument data. It returns a string
as reversed transformed data
*/
func RTransform(data utils.Data) string {
	return data.GetData()
}
