package main

import "springboot-hand-stand/myutil"

func main() {

	str := "marvel_luobosi_xueliu_naitang"

	println(myutil.ToUpperUnderlineWordInitialsForField(str))
	println(myutil.ToUpperUnderlinedWordInitialsForClassName(str))

}
