package main

import (
	"bufio"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/anjun/.bash_profile")
	defer f.Close()

	/*
	   直接读取
	   	for {
	   		n, _ := f.Read(buf)
	   		if n == 0 {
	   			break
	   		}
	   		os.Stdout.Write(buf[:n])
	   	}
	*/
	//buffio 读写
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
}
