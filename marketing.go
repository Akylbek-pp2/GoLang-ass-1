package marketing

type Marketing struct 
{
    position string // unexported field
    salary   int
    address  string
}

func NewMarketing(position string, salary int, address string) *Marketing {
    return &Marketing{position, salary, address}
}

func (m *Marketing) GetPosition() string {
    return m.position
}

func (m *Marketing) SetPosition(position string) {
    m.position = position
}

func (m *Marketing) GetSalary() int {
    return m.salary
}

func (m *Marketing) SetSalary(salary int) {
    m.salary = salary
}

func (m *Marketing) GetAddress() string {
    return m.address
}

func (m *Marketing) SetAddress(address string) {
    m.address = address
}