package main

import "fmt"

func calSquares(number int, squchan chan int) {
	fmt.Println("write the square to the channel.")
	square := number * number
	squchan <- square // write data to channel
}
func calCube(number int, cubechan chan int) {
	fmt.Println("write the cube to the channel.")
	cube := number * number * number
	cubechan <- cube // write data to channel
}

func main() {
	square_chan := make(chan int)
	cube_chan := make(chan int)
	number := 2
	fmt.Printf("calcute the number %d's square and cube, and then plus them.\n", number)
	go calSquares(number, square_chan)
	go calCube(number, cube_chan)
	square_num, cube_num := <-square_chan, <-cube_chan // get data from the two channel
	fmt.Printf("the number %d's square and cube is %d %d, and plus then is %d", number, square_num, cube_num, square_num+cube_num)
}
