package main

import "time"

func main() {
	for {
		println("hello")
		println(".")
		time.Sleep(3 * time.Second)
	}
}

//stty -F /dev/ttyACM0 cs8 9600 ignbrk -brkint -imaxbel -opost -onlcr -isig -icanon -iexten -echo -echoe -echok -echoctl -echoke noflsh -ixon -crtscts
