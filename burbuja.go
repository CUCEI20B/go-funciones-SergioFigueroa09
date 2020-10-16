package main

func Burbuja(s []int64)  {
	var aux int64
	for i:= len(s); i > 0; i--{
		for j:=1; j < i; j++{
			if s[j-1] > s[j]{
				aux = s[j]
				s[j] = s[j-1]
				s[j-1] = aux
			}
		}
	}
}

func main()  {
	
}