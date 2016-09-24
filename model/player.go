package model

import "fmt"

type Player struct{
	Name string
	Symbol string
}


func (p Player) DrawPlayer(){
	fmt.Print(p.Symbol)
}