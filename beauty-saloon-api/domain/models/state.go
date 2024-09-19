package models


type State struct {
	StateName string `json:"state_name"`
	UF        string `json:"uf"`
}


func NewState(stateName, uf string) *State {
	return &State{
		StateName: stateName,
		UF:        uf,
	}
}


func (s *State) SetStateName(stateName string) {
	s.StateName = stateName
}


func (s *State) SetUF(uf string) {
	s.UF = uf
}