package tempconv

func CToF(c Celsius) Fahreneit {
	return Fahreneit(c*9/5 + 32)
}

func FToC(f Fahreneit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
