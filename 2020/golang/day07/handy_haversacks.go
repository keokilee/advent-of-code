package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagData struct {
	Bags map[string]*Bag
}

func NewBagData() *BagData {
	return &BagData{Bags: make(map[string]*Bag)}
}

var inputRegex = regexp.MustCompile(`^(?P<color>.+) bags contain (?P<children>.+)\.$`)
var childRegex = regexp.MustCompile(`(?P<count>\d+) (?P<color>.+) bags?`)

func (data *BagData) AddBag(text string) {
	matches := inputRegex.FindStringSubmatch(text)

	color, children := matches[1], matches[2]
	parentBag := data.getOrCreateBag(color)

	if children == "no other bags" {
		return
	}

	childStrings := strings.Split(children, ",")
	for _, childText := range childStrings {
		matches = childRegex.FindStringSubmatch(childText)
		childBag := data.getOrCreateBag(matches[2])
		childBag.Parents = append(childBag.Parents, parentBag)

		parentBag.Children[childBag], _ = strconv.Atoi(matches[1])
	}
}

func (data *BagData) getOrCreateBag(color string) *Bag {
	_, ok := data.Bags[color]
	if !ok {
		data.Bags[color] = NewBag(color)
	}

	return data.Bags[color]
}

type Bag struct {
	Color    string
	Parents  []*Bag
	Children map[*Bag]int
}

func NewBag(color string) *Bag {
	return &Bag{Color: color, Parents: make([]*Bag, 0), Children: make(map[*Bag]int)}
}

func (bag *Bag) UniqueColors() (colors []string) {
	colorSet := make(map[string]bool)
	bag.uniqueParentColors(colorSet)

	colors = make([]string, 0)
	for color, _ := range colorSet {
		colors = append(colors, color)
	}

	return colors
}

func (bag *Bag) uniqueParentColors(colors map[string]bool) {
	bag.PrintBag()
	for _, parent := range bag.Parents {
		colors[parent.Color] = true
		parent.uniqueParentColors(colors)
	}
}

func (bag *Bag) InnerBagCount() int {
	total := 0
	for child, count := range bag.Children {
		total += count + (count * child.InnerBagCount())
	}

	return total
}

func (bag *Bag) PrintBag() {
	parents := make([]string, 0)
	for _, parent := range bag.Parents {
		parents = append(parents, parent.Color)
	}

	log.Printf("color=%s parents=%s", bag.Color, strings.Join(parents, ", "))
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	data := NewBagData()
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		data.AddBag(text)
	}

	bag := data.Bags["shiny gold"]
	colors := bag.UniqueColors()
	log.Printf("unique color count %d", len(colors))
	log.Printf("inner bag count %d", bag.InnerBagCount())
}
