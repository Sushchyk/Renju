package model


type Player struct{
	name string
	symbol string
}

func (p *Player) CreatePlayer(plName string, plSymbol string) {
	p.name = plName;
	p.symbol = plSymbol;
}

func (p Player) GetSymbol() string{
	return p.symbol;
}

func (p Player) GetName() string{
	return p.name;
}
