package main

import "fmt"

func main(){
	var arrResults [][]string
	count:=5
	for i:=0;i<count;i++{
		var tmp []string
		for j:=0;j<15;j++{
			//arrResults[i][j]="a"
			//tmp[j]="a"
			tmp=append(tmp,"a")
		}
		arrResults=append(arrResults,tmp)
	}
	fmt.Println(arrResults)
}
