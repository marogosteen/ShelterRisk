package sample

import(
	"fmt"
)

func main(){

}
type foo struct {
	x int
	y int
	var voo [5][10]foo
	
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			v:=foo{
				x: i+i,
				y: j+j,
			}
			voo[i][j] = v
	
		}
		fmt.Println(voo)
	}
}


