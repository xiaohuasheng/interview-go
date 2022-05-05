package main

import "fmt"

func main() {
	fmt.Println("dd")
}

type linkedList struct {
}

func (l *linkedList) offer(n int) {
}

func (l *linkedList) peek() int {
	return 0
}
func (l *linkedList) poll() int {
	return 0
}
func (l *linkedList) isEmpty() bool {
	return true
}

type monotonicQueue struct {
	Data linkedList
}

func (m *monotonicQueue) offer(n int) {
	for !m.Data.isEmpty() {
		if m.Data.peek() < n {
			m.Data.poll()
		} else {
			break
		}
	}
	m.Data.offer(n)
}
