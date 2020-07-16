package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
//网络请求超时控制
type Result struct{
	r *http.Response
	err error
}

func main(){
	process()
}

func process() {
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	//释放资源
	defer cancel()
	tr:=&http.Transport{}
	client:=&http.Client{Transport: tr}
	resultChan:=make(chan Result,1)
	//发起请求
	req,err:=http.NewRequest("Get","http://www.baidu.com",nil)
	if err!=nil{
		fmt.Println("http request failed,err:",err)
		return
	}
	go func(){
		resp,err:= client.Do(req)
		pack:=Result{r:resp,err:err}
		//将返回信息返回管道
		resultChan<-pack
	}()
	select{
		case<-ctx.Done():
			tr.CancelRequest(req)
			er:=<-resultChan
			fmt.Println("timeout!",er.err)
		case res:=<-resultChan:
			defer res.r.Body.Close()
			out,_:=ioutil.ReadAll(res.r.Body)
			fmt.Printf("server response:%s",out)

	}
	return


}






