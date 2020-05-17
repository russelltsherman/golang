package patterns

// BuildProcess interface
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector struct
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct ManufacturingDirector
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder ManufacturingDirector
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct struct
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// Builder of type car
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels setter
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats setter
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure setter
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle getter
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

//A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels setter
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats setter
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure setter
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

// GetVehicle getter
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// BusBuilder struct
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels setter
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

// SetSeats setter
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

// SetStructure setter
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

// GetVehicle getter
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
