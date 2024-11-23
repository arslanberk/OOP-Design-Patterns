package main

import "fmt"

// Collection
type Collection interface {
    createIterator() Iterator
}

// Concrete collection
type UserCollection struct {
    users []*User
}

func (u *UserCollection) createIterator() Iterator {
    return &UserIterator{
        users: u.users,
    }
}

// Iterator
type Iterator interface {
    hasNext() bool
    getNext() *User
}

// Concrete iterator
type UserIterator struct {
    index int
    users []*User
}

func (u *UserIterator) hasNext() bool {
    if u.index < len(u.users) {
        return true
    }
    return false

}
func (u *UserIterator) getNext() *User {
    if u.hasNext() {
        user := u.users[u.index]
        u.index++
        return user
    }
    return nil
}

// client code
type User struct {
    name string
    age  int
}

func main() {

    user1 := &User{
        name: "a",
        age:  30,
    }
    user2 := &User{
        name: "b",
        age:  20,
    }

    userCollection := &UserCollection{
        users: []*User{user1, user2},
    }

    iterator := userCollection.createIterator()

    for iterator.hasNext() {
        user := iterator.getNext()
        fmt.Printf("User is %+v\n", user)
    }
}