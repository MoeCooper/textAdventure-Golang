package main

import "fmt"

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

func (node* storyNode) executeCommand (cmd string)  *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd)
		currentChoice = currentChoice.nextChoice
	}
}

func main() {

}
