package main

import "fmt"

/*
and example Dependency Inversion Principle

High Level Modules should not depend on Low Level Modules
Both should depend on abstractions
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

type Research struct {
	// relationships Relationships // low-level moodule dependency
	browser RelationshipBrowser // low-level dependency abstraction
}

func (r *Research) Investigate(name string) {
	// tight coupling of relationships module
	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "John" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	//}

	// loose coupling using an interface
	for _, p := range r.browser.FindAllChildrenOf(name) {
		fmt.Println(name, "has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate("John")
}
