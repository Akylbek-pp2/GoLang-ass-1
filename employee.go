package employee

type Employee interface
{
	GetPosition() string
    SetPosition(string)
    GetSalary() int
    SetSalary(int)
    GetAddress() string
    SetAddress(string)
}