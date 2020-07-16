package main

import "fmt"

//总结一下；就是多层slice复制，一定从最里层形成一个slice,然后由里向外逐层append即可
func main(){
	//mapResults:=make(map[int]string)
	var arrResults [][]string
	count:=5
	for i:=0;i<count;i++{
		//valueStr:=fmt.Sprintf("this is %d",i)
		//mapResults[i]=valueStr
		var tmpArr []string
		for j:=0;j<15;j++{
			tmpArr=append(tmpArr,"a")

		}
		arrResults=append(arrResults,tmpArr)
	}
	//fmt.Println(mapResults)
	fmt.Println(arrResults)
}
