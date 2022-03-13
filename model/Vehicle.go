package model

type Vehicle struct {
	number string
	color  string
}

func (v *Vehicle) getVehicleNum() string {
	return v.number
}
