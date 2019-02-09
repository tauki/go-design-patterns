package builder

import "testing"

type vehicleTester struct {
	t *testing.T
}


func TestBuilderPattern(t *testing.T) {
	vehicleTester := &vehicleTester{
		t:t,
	}

	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	vehicleTester.checkCar(car)

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorBike := bikeBuilder.GetVehicle()
	vehicleTester.checkBike(motorBike)

	motorBike.Seats = 1
	if motorBike.Seats != 1 {

	}

}

func (t *vehicleTester) checkCar(product VehicleProduct) {

	if product.Seats != 5 {
		t.t.Errorf("Expected the Seats of a Car to be 5, got %d", product.Seats)
	}

	if product.Wheels != 4 {
		t.t.Errorf("Expected the Wheels of a Car to be 4, got %d", product.Wheels)
	}

	if product.Structure != "Car" {
		t.t.Errorf("Expected the Structure of a Car to be Car, got %s", product.Structure)
	}
}

func (t *vehicleTester) checkBike(product VehicleProduct) {

	if product.Seats != 2 {
		t.t.Errorf("Expected the Seats of a MotorBike to be 2, got %d", product.Seats)
	}

	if product.Wheels != 2 {
		t.t.Errorf("Expected the Wheels of a Motorbike to be 2, got %d", product.Wheels)
	}

	if product.Structure != "MotorBike" {
		t.t.Errorf("Expected the Structure of a Motorbike to be MotorBike, got %s", product.Structure)
	}
}