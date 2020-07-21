package main

import "fmt"

func main() {
	states := make(map[string] string)
	fmt.Println(states)
	states["KA"] = "BLR"
	states["TN"] = "CHN"
	states["KE"] = "TRN"

	fmt.Println(states)
	for k,v := range states{
		fmt.Printf("%v %v \n",k,v)
	}

	delete(states, "KA")

	for k,v := range states{
		fmt.Printf("%v %v \n",k,v)
	}

	keys:= make([]string, 5, 5)
	values := make([]string, 5, 5)
	i := 0
	for k,v := range states{
		keys[i] = k
		values[i] = v
		i++
	}
	fmt.Println(keys)

	fmt.Println(values)
}
