package main

import "fmt"

// Flyweight factory
const (
    //TerroristDressType terrorist dress type
    TerroristDressType = "tDress"
    //CounterTerrroristDressType terrorist dress type
    CounterTerrroristDressType = "ctDress"
)

var (
    dressFactorySingleInstance = &DressFactory{
        dressMap: make(map[string]Dress),
    }
)

type DressFactory struct {
    dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
    if d.dressMap[dressType] != nil {
        return d.dressMap[dressType], nil
    }

    if dressType == TerroristDressType {
        d.dressMap[dressType] = newTerroristDress()
        return d.dressMap[dressType], nil
    }
    if dressType == CounterTerrroristDressType {
        d.dressMap[dressType] = newCounterTerroristDress()
        return d.dressMap[dressType], nil
    }

    return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
    return dressFactorySingleInstance
}

// Flyweight interface
type Dress interface {
    getColor() string
}

// Concrete flyweight object
type TerroristDress struct {
    color string
}

func (t *TerroristDress) getColor() string {
    return t.color
}

func newTerroristDress() *TerroristDress {
    return &TerroristDress{color: "red"}
}

// Concrete flyweight object
type CounterTerroristDress struct {
    color string
}

func (c *CounterTerroristDress) getColor() string {
    return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
    return &CounterTerroristDress{color: "green"}
}

// Context
type Player struct {
    dress      Dress
    playerType string
    lat        int
    long       int
}

func newPlayer(playerType, dressType string) *Player {
    dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
    return &Player{
        playerType: playerType,
        dress:      dress,
    }
}

func (p *Player) newLocation(lat, long int) {
    p.lat = lat
    p.long = long
}

// client code
type game struct {
    terrorists        []*Player
    counterTerrorists []*Player
}

func newGame() *game {
    return &game{
        terrorists:        make([]*Player, 1),
        counterTerrorists: make([]*Player, 1),
    }
}

func (c *game) addTerrorist(dressType string) {
    player := newPlayer("T", dressType)
    c.terrorists = append(c.terrorists, player)
    return
}

func (c *game) addCounterTerrorist(dressType string) {
    player := newPlayer("CT", dressType)
    c.counterTerrorists = append(c.counterTerrorists, player)
    return
}

func main() {
    game := newGame()

    //Add Terrorist
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)

    //Add CounterTerrorist
    game.addCounterTerrorist(CounterTerrroristDressType)
    game.addCounterTerrorist(CounterTerrroristDressType)
    game.addCounterTerrorist(CounterTerrroristDressType)

    dressFactoryInstance := getDressFactorySingleInstance()

    for dressType, dress := range dressFactoryInstance.dressMap {
        fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
    }
}