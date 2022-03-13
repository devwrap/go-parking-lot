package model

type Vehicle struct {
	Number string
	Color  string
}

func (v *Vehicle) getVehicleNum() string {
	return v.Number
}
