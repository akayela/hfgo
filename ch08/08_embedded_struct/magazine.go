package magazine 

type Subscriber struct {
  Name string
  Rate float64
  Active bool
  Address // An embedded struct ie a struct within a struct
}

type Employee struct {
  Name string
  Salary float64
  Address
}

// The struct below can be accessed from the outer structs just like any other field
type Address struct {
  Street string
  City string
  State string
  PostalCode string
}
