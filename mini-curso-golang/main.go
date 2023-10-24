// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type ID uint16

// var (
// 	name    string
// 	salario float32
// )

// func main() {
//Salario
// fmt.Printf("Insira o valor do salario: ")
// fmt.Scanf("%f", &salario)

// fmt.Printf("Salario %.2f", salario)

//Nome

// fmt.Printf("Insira seu nome: ")
// fmt.Scanf("%s", &name)

// fmt.Printf("Nome %", name)

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Insira o seu nome: ")
// 	name, err := reader.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}

// 	name = strings.Trim(name, "\n")
// 	fmt.Println(name)

// 	var id ID = 5

// 	fmt.Println(id)
// }

// ---------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// var (
// 	position int = 3
// 	value    int = 25
// )

// func main() {
// 	my_slice := []int{0, 1, 2, 3, 4, 5, 6, 8}

// 	fmt.Println(my_slice)

// 	my_slice = append(my_slice[:position], append([]int{value}, my_slice[position:]...)...)
// 	fmt.Println(my_slice)
// }

//--------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// var (
// 	my_map = map[string]string{}
// )

// func main() {
// 	my_map["Name"] = "Pedro"
// 	my_map["Age"] = "25"
// 	my_map["Height"] = "1.93"

// 	str, ok := my_map["Name"]
// 	if ok {
// 		fmt.Printf(str)
// 		return
// 	}

// 	fmt.Printf("Valor Inexistente!")
// 	return
// }

//--------------------------------------------------------

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	valor, err := MaiorQueLimite(15)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(valor)
// }

// func MaiorQueLimite(y float32) (float32, error) {
// 	if y > 10 {
// 		err := errors.New("Queda maior do que o esperado")
// 		return 0.0, err
// 	}
// 	return y, nil
// }

//------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(Soma(1, 24, 23, 242, 354, 654232))
// }

// func Soma(numbers ...int) int {
// 	total := 0
// 	for _, value := range numbers {
// 		total += value
// 	}

// 	return total
// }

//----------------------------------------------------------

package main

import "fmt"

func main() {
	soma_um := incrementa()

	fmt.Println(soma_um())
	fmt.Println(soma_um())
	fmt.Println(soma_um())
}

func incrementa() func() int {
	total := 0

	result := func() int {
		total += 1
		return total
	}
	return result
}
