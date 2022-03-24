/************Switc*********/
//*Switch is used to easily check multiple conditions
//**Alternative to many if...else if blocks
//*Switches execute from top-to-bottom
//*Optionally have a 'default' action

//*Having default allows something to be displayed
//if none of the cases get checked

/*Example*/

package main
import "fmt"

x := 3
switch x {
case 1:
	fmt.Println("1")
case 2:
	fmt.Println("2")
case 3:
	fmt.Println("3")
default:
	fmt.Println("other:", x)
}


url := "example.com"
switch url {
case "example.com":
	fmt.Println("test")
case "google.com"
	fmt.Println("live")
default:
	fmt.Println("dev")
}


/********Fallthrough********/
/*
switch letter {
case ' ':
case 'a', 'e', 'i', 'o', 'u':
	fmt.Println("Vowels are great")
default:
	fmt.Println("It's something else")
}