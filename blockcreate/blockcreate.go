package blockcreate

import (
	"fmt"
)

func Empty(...any) any {

	return 0
}
func Runner(){
	BlockCreatePassive()
}

func BlockCreatePassive() {

	gridV2 := make(map[int]map[int]map[int]int)

	for i := 0; i < 3; i++ {
		gridV2[i] = make(map[int]map[int]int)
		for j := 0; j < 5; j++ {
			gridV2[i][j] = make(map[int]int)
			for k := 0; k < 10; k++ {
				gridV2[i][j][k] = k
			}
		}
	}

	fmt.Println("")

	for i := 0; i < 3; i++ {

		fmt.Printf("%d th BLOCK", i)
		fmt.Println("\n----------")
		fmt.Println("")
		for j := 0; j < 5; j++ {
			fmt.Printf("%d th LINE\t", j)
			for k := 0; k < 10; k++ {
				fmt.Printf("%d ", gridV2[i][j][k])
			}
			fmt.Println("")

		}
		fmt.Println("")
	}
}