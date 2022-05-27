package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // 0.1 second
	boom := time.After(500 * time.Millisecond) // 0.5 second

	for {
		select {
		case <-tick:
			fmt.Println(" tick. ")
		case <-boom:
			fmt.Println(" BOOM! ")
			return
		default:
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond) // 0.05 second delay
		}
	}

	/*
		0.1초 전에 0.05초 딜레이가 두 번 		누적 second
		. 					// 0.05 second
		.						// 0.05 second
		tick. 			// 0.1 second
										딜레이 두 번 			0.1s
		. 					// 0.05 second
		.						// 0.05 second
		tick. 			// 0.1 second
										딜레이 두 번			0.2s
		. 					// 0.05 second
		.						// 0.05 second
		tick. 			// 0.1 second
										딜레이 두 번			0.3s
		. 					// 0.05 second
		.						// 0.05 second
		tick. 			// 0.1 second
										딜레이 두 번			0.4s
		. 					// 0.05 second
		.						// 0.05 second
		tick. 			// 0.1 second
										딜레이 두 번			0.5s
		.						// 0.05 second
		.						// 0.05 second
		BOOM!
	*/
}
