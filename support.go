package support

type Support struct
{
    position string
    salary   int
    address  string
}

func NewSupport(position string, salary int, address string) *Support {
    return &Support{position, salary, address}
}

func (s *Support) GetPosition() string {
    return s.position
}

func (s *Support) SetPosition(position string) {
    s.position = position
}

func (s *Support) GetSalary() int {
    return s.salary
}

func (s *Support) SetSalary(salary int) {
    s.salary = salary
}

func (s *Support) GetAddress() string {
    return s.address
}

func (s *Support) SetAddress(address string) {
    s.address = address
}