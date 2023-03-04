package person

type Person struct {
	name   string
	age    int
	height float64
}

func NewPerson(name string, age int, height float64) *Person {
	return &Person{name, age, height}
}

func (p *Person) SetName(name string) {

	p.name = name

}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetAge(age int) {

	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetHeight(height float64) {
	p.height = height
}

func (p *Person) GetHeight() float64 {
	return p.height
}
