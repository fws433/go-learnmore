package main

import "fmt"

type Map map[string]string

func main() {
	mp :=make(map[string]string, 10)
	mp["hi"] = "tata"

	//mp与ma有相同对底层内型map[string]string,并且mp是未命名类型
	//所以mp可以直接复制给ma
	var ma Map = mp
	fmt.Println(ma)
}
