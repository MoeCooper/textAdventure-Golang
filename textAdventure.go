package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// linked list
type choices struct {
	command     string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode) addChoice(cmd string, desc string, nextNode *storyNode) {
	choice := &choices{cmd, desc, nextNode, nil}

	if node.choices == nil {
		node.choices = choice
	} else {
		// loop down linked list til we find nil choice
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *storyNode) display() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.command, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCommand(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.command) == strings.ToLower(currentChoice.command) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("I didn't get that.")
	return node
}

// used to take in input
var scanner *bufio.Scanner

func (node *storyNode) playGame() {
	node.display()
	if node.choices != nil {
		scanner.Scan()
		node.executeCommand(scanner.Text()).playGame()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
		You are in a large room underground, 
		carrying a large lit lantern. 
		There are three passages in different directions 
		that seem to lead out. 
		North takes you to darkness
		South appears to head to upstairs
		West is seemingly a safe choice, with no darkness and 
		goes straight onwards
	`}

	darkRoom := storyNode{text: `
		It is pitch black, and cannot see anything. 
	`}

	darkRoomLit := {text: `
		The dark passageway is lit by your lantern. 
		You can continue North or head back South.
	`}

	bigTroll := storyNode{text: `
		While walking in the dark, you are eaten by a monster troll!
	`}

	trap := storyNode{text: `
		You are heading down the path when suddenly, a trap door
		falls open and you fell inside!
	`}

	treasureRoom := storyNode{text: `
		You arrive to a small chamber, filled with treasure!!
	`}

	
}
