package main

func main() {
	print(entityParser("&&&&&amp"))
}

func entityParser(text string) string {

	hashMap := map[string]string{
		"&quot;" : "\"",
		"&apos;" : "'",
		"&amp;" : "&",
		"&gt;" : ">",
		"&lt;" : "<",
		"&frasl;" : "/",
	}
	result := ""
	tmpString := ""
	flag := false

	for  _, letter := range text {
		if string(letter) == "&" {
			if flag {
				result += tmpString
				tmpString = ""
			}else {
				flag = true
			}
		}

		if flag {
			tmpString += string(letter)
		} else {
			result += string(letter)
		}

		if string(letter) == ";" {
			if value, ok := hashMap[tmpString]; ok {
				result += value
			} else  {
				result += tmpString
			}
			tmpString = ""
			flag = false
		}
	}

	return result+tmpString
}
