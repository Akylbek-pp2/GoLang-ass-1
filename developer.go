package developer


type Developer struct 
{
    position string
    salary   int
    address  string
}

func NewDeveloper(position string, salary int, address string) *Developer {
    return &Developer{position, salary, address}
}

func (d *Developer) GetPosition() string {
    return d.position
}

func (d *Developer) SetPosition(position string) {
    d.position = position
}

func (d *Developer) GetSalary() int {
    return d.salary
}

func (d *Developer) SetSalary(salary int) {
    d.salary = salary
}

func (d *Developer) GetAddress() string {
    return d.address
}

func (d *Developer) SetAddress(address string) {
    d.address = address
}