package main

import "Lab4/worker"

func main() {
	//for i:=0;i<10;i++{
	//	fmt.Println(generator.GenerateBirthYear())
	//	time.Sleep(1*time.Millisecond)
	//}
	//fmt.Println(generator.GeneratePassword())
	worker.StartHashing()
}
