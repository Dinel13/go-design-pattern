package proxy

const successDriven = "car is being driven"
const failDriven = "driver too old"

type Driven interface {
	Drive() string
}

type Car struct {
}

func (c Car) Drive() string {
	return successDriven
}

type Driver struct {
	Age int
}

// CarProxy will become proxy to add some functionality like validation
// CarProxy must implement same interface of Car, so it can do all the car can be doing
// CarProxy must call the function of the car in the end

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() string {
	if c.driver.Age >= 16 {
		return c.car.Drive()
	} else {
		return failDriven
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{driver: driver}
}
