package iteration

func Repeart(charactor string, repeart int) string {
	var repearted string
	for i := 0; i < repeart; i++ {
		repearted += charactor
	}
	return repearted
}
