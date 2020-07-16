package main
//goroutine调度器的工作流程
func main() {

	// 程序启动时的初始化代码
	//........
	var N int
	for i := 0; i < N; i++ { // 创建N个操作系统线程执行schedule函数
		create_os_thread(schedule) // 创建一个操作系统线程执行schedule函数
	}
}
	//schedule函数实现调度逻辑
	func schedule() {
		for { //调度循环
			// 根据某种算法从M个goroutine中找出一个需要运行的goroutine
			g := find_a_runnable_goroutine_from_M_goroutines()
			run_g(g) // CPU运行该goroutine，直到需要调度其它goroutine才返回
			save_status_of_g(g) // 保存goroutine的状态，主要是寄存器的值
		}
	}

func save_status_of_g(g interface{}) {

}

func run_g(g interface{}) {

}

func find_a_runnable_goroutine_from_M_goroutines() interface{} {
	return nil
}
func create_os_thread(schedule interface{}) {

}
