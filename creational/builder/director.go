package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetWheels().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
