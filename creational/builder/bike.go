package builder

type BikeBuilder struct{
	vehicleProduct VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.vehicleProduct.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.vehicleProduct.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.vehicleProduct.Structure = "MotorBike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.vehicleProduct
}
