package mat

func CekGanjilGenap(num ...int) string{
	result := "" 
	for _, i := range num{
		if i%2==0 {
			result += "genap, "
		}else{
			result += "ganjil, "
		} 
	}
	return result
	
}