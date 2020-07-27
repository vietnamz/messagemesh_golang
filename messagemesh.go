package main

import (
	"fmt"

	"container/list"

	"github.com/golang-collections/go-datastructures/queue"
)

/*
  network = {
    'Min'     : ['William', 'Jayden', 'Omar'],
    'William' : ['Min', 'Noam'],
    'Jayden'  : ['Min', 'Amelia', 'Ren', 'Noam'],
    'Ren'     : ['Jayden', 'Omar'],
    'Amelia'  : ['Jayden', 'Adam', 'Miguel'],
    'Adam'    : ['Amelia', 'Miguel', 'Sofia', 'Lucas'],
    'Miguel'  : ['Amelia', 'Adam', 'Liam', 'Nathan'],
    'Noam'    : ['Nathan', 'Jayden', 'William'],
    'Omar'    : ['Ren', 'Min', 'Scott'],
    ...
}
*/

func constructRoute(aMap map[string]string, start string, end string) list.List {
	var result list.List
	tmp := end
	for k, v := range aMap {
		if k == tmp {
			tmp = v
			result.PushFront(tmp)
		}
	}
	return result
}

func main() {
	// we want to use a queue to traverse a graph.
	var q queue.Queue

	// input
	m := make(map[string][]string)

	nodeWeAldreadySee := make(map[string]string)

	m["Min"] = []string{"William", "Jayden", "Omar"}
	m["William"] = []string{"Min", "Noam"}
	m["Jayden"] = []string{"Min", "Amelia", "Ren", "Noam"}
	m["Ren"] = []string{"Jame", "Omar"}
	m["Amelia"] = []string{"Jayden", "Adam", "Miguel"}
	m["Adam"] = []string{"Amelia", "Miguel", "Sofia", "Lucas"}
	m["Miguel"] = []string{"Amelia", "Adam", "Liam", "Nathan"}
	m["Noam"] = []string{"Nathan", "Jayden", "William"}
	m["Omar"] = []string{"Ren", "Min", "Scott"}

	// start algorithm DFS

	var startnode string = "Min"

	var endnode string = "Omar"

	q.Put(startnode)

	for !q.Empty() {

		v, e := q.Get(1)
		if e != nil {
			fmt.Println("There was a error")
			return
		}
		if v[0] == endnode {
			return constructRoute(nodeWeAldreadySee, startnode, endnode)
		}
	}
}
