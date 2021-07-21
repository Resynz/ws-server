/**
 * @Author: Resynz
 * @Date: 2021/7/21 11:56
 */
package tools

func CheckIsInStringArray(arr []string, a string) bool {
	b := false
	for _, v := range arr {
		if a == v {
			b = true
			break
		}
	}
	return b
}
