package models

type City struct {
	CityName string  `json:"city_name"`
	State    *State  `json:"state"`
	States   []State `json:"states"`
}


func NewCity(cityName string) *City {
	return &City{
		CityName: cityName,
		States:   []State{},
	}
}


func (c *City) AddState(state State) {
	c.States = append(c.States, state)
}


func (c *City) GetStates() []State {
	return c.States
}

func (c *City) SetState(state State) {
	c.State = &state
}