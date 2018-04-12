package main

import (
	"math/rand"
	"os"
	"log"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"github.com/logrusorgru/aurora"
)

func calculateSpeed(speed int, burn int, gravity int) int {
	return speed + gravity - burn
}

func help() {
	fmt.Print("Lunar Lander version 1.1\n")
	fmt.Print("Made by Lukas Fülling and Nicolai Süper\n\n")
	fmt.Print("The following arguments are possible (only one):\n")
	fmt.Print("-d [1/2/3]\tDefine difficulty. 1 is easy 3 is hard.\n")
	fmt.Print("--info\tShow different intro.")
	fmt.Print("--help\tPrint this help and exit.\n")
}

func windowCleaner(step int) int {
	if step >= 24 {
		fmt.Print("\nTime\t")
		fmt.Print("Speed\t\t")
		fmt.Print("Fuel\t\t")
		fmt.Print("Height\t\t")
		fmt.Print("Burn\n")
		fmt.Print("----\t")
		fmt.Print("-----\t\t")
		fmt.Print("----\t\t")
		fmt.Print("------\t\t")
		fmt.Print("----\n")
		step = 1
	} else if step < 32 {
		step++
	}
	return step
}

func randomHeight() int {
	return rand.Intn(15000)%15000 + 4000
}

func main() {
	const gravity = 100 /* The rate in which the spaceship descents in free fall (in ten seconds) */
	reader := bufio.NewReader(os.Stdin)

	var speed int
	var height int
	var fuel int
	var tensec int
	var burn int
	var prevheight int
	var step int

	const version = "1.1-go" /* The Version of the program */
	const dead = "\nThere were no survivors."
	const crashed = "\nThe Spaceship crashed. Good luck getting back home."
	const success = "\nYou made it! Good job!"
	const emptyFuel = "\nThere is no fuel left. You're floating around like Wheatley."
	fmt.Println("\nLunar Lander - Version " + version)
	fmt.Println("This is a computer simulation of an Apollo lunar landing capsule.")
	fmt.Println("The on-board computer has failed so you have to land the capsule manually.")
	fmt.Println("Set burn rate of retro rockets to any value between 0 (free fall) and 200")
	fmt.Println("(maximum burn) kilo per second. Set burn rate every 10 seconds.")
	fmt.Println("Good Luck!")

	/* Set initial height, time, fuel, burn, prevheight, step and speed according to difficulty. */

	if len(os.Args) == 1 { /* If there is only one argument (which is the program's name) */
		speed = 1000 /* Default to easy (and randomize the height) */
		height = randomHeight()
		fuel = 12000
		tensec = 0
		burn = 0
		prevheight = height
		step = 1
	} else { /* If there are more arguments (or less) */
		if os.Args[1] == "-d" { /* If the "first" Argument is -d, check for the second argument*/
			if os.Args[2] == "1" { /* Easy */
				speed = 1000
				height = randomHeight()
				fuel = 12000
				tensec = 0
				burn = 0
				prevheight = height
				step = 1
			}
			if os.Args[2] == "2" { /* Medium */
				speed = 1000
				height = randomHeight()
				fuel = 1000
				tensec = 0
				burn = 0
				prevheight = height
				step = 1
			}
			if os.Args[2] == "3" { /* Hard */
				speed = 2000
				height = randomHeight() - 2000
				fuel = 900
				tensec = 0
				burn = 0
				prevheight = height
				step = 1
			} else { /* If argv[1] is not -d, default to Easy */
				speed = 1000
				height = randomHeight()
				fuel = 12000
				tensec = 0
				burn = 0
				prevheight = height
				step = 1
			}
		} else if os.Args[1] == "--info" {
			fmt.Println("\nLunar Lander - Version " + version)
			fmt.Println("Made by Lukas Fülling and Nicolai Süper")
			fmt.Println("\n\nContact us at http://k40s.net")
		} else if os.Args[1] == "--help" {
			help()
		} else { /* If the first Argument is something else, default to Easy */
			speed = 1000
			height = randomHeight()
			fuel = 12000
			tensec = 0
			burn = 0
			prevheight = height
			step = 1
		}
	}

	fmt.Print("\nTime\t")
	fmt.Print("Speed\t\t")
	fmt.Print("Fuel\t\t")
	fmt.Print("Height\t\t")
	fmt.Print("Burn\n")
	fmt.Print("----\t")
	fmt.Print("-----\t\t")
	fmt.Print("----\t\t")
	fmt.Print("------\t\t")
	fmt.Print("----\n")
	for height > 0 {

		step = windowCleaner(step)
		fmt.Print(strconv.Itoa(tensec) + "0\t", )
		fmt.Print(strconv.Itoa(speed) + "\t\t")
		fmt.Print(strconv.Itoa(fuel) + "\t\t")
		if height < prevheight {
			fmt.Print(aurora.Red(strconv.Itoa(height) + "\t\t"))
		} else if height == prevheight {
			fmt.Print(strconv.Itoa(height) + "\t\t")
		} else if height > prevheight {
			fmt.Print(aurora.Green(strconv.Itoa(height) + "\t\t"))
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		burn, err = strconv.Atoi(strings.Replace(input, "\n", "", -1))
		if err != nil {
			log.Fatal(err)
		}

		if burn < 0 || burn > 200 { /* If there is a wrong entry */
			fmt.Println(aurora.Red("The burn rate rate must be between 0 and 200."))
			continue
		}

		prevheight = height
		speed = calculateSpeed(speed, burn, gravity)
		height = height - speed
		fuel = fuel - burn
		if fuel <= 0 {
			break
		}

		tensec++
	}

	if height <= 0 {
		if speed > 10 {
			fmt.Println(dead)
		} else if speed < 10 && speed > 3 {
			fmt.Println(crashed)
		} else if speed < 3 {
			fmt.Println(success)
		}
	} else if height > 0 {
		fmt.Println(emptyFuel)
	}
}
