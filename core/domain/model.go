package domain

//Struct representation of a pokemon created on excel
type Pokemon struct {
	ID                                                                                  int
	Name, Types, Ability                                                                string
	Stage                                                                               int
	GrowRatio                                                                           string
	TotalStats                                                                          int
	Category, Nature                                                                    string
	HpPhases, AttackPhases, DefensePhases, SpAttackPhases, SpDefensePhases, SpeedPhases string
	//Each pokemon should have min 1 movement and max 4 movements
	Movements []Movement
}

type Movement struct {
	Name            string
	Power, Accuracy int
	MoveType, Class string
}
