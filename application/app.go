package application

func Fibbonacci(a int) int {
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else if a == 2 {
		return 1
	} else {
		return Fibbonacci(a-1) + Fibbonacci(a-2)
	}
}
