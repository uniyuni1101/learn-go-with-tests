package iteration

func Repeart(charactor string) string {
	var repearted string
	for i := 0; i < 5; i++ {
		repearted += charactor
	}
	return repearted
}
