package sales

type Sales struct 
{
    position string
    salary   int
    address  string
}

func NewSales(position string, salary int, address string) *Sales {
    return &Sales{position, salary, address}
}

func (s *Sales) GetPosition() string {
    return s.position
}

func (s *Sales) SetPosition(position string) {
    s.position = position
}

func (s *Sales) GetSalary() int {
    return s.salary
}

func (s *Sales) SetSalary(salary int) {
    s.salary = salary
}

func (s *Sales) GetAddress() string {
    return s.address
}

func (s *Sales) SetAddress(address string) {
    s.address = address
}