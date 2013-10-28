package main

import "errors"
import "fmt"

func devide( value int, by int ) (int, error) {
	if  by == 0 {
		return 0, errors.New( "Can't work with 0!")
	}

	return value/by, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf( "%d - %s", e.arg, e.prob )
}

func devide_2( value int, by int ) (int, error){
	if by == 0 {
		return 0, &argError{ by, "Can't work with 0!" }
	}
	return value/by, nil
}

func main(){

	for i := 0; i < 3; i++ {
		ret, err := devide(5,i)
		fmt.Println( 5, " devide ", i, " is ", ret, err, '\n' )
	}

	fmt.Println()
	for _, i := range []int{0,1,2, 3} {
		ret, err := devide_2( 5, i )
		fmt.Println( 5, " devide ", i, " is ", ret, err, '\n' )
	}
}
