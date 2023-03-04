package person

type Person struct {
	name   string
	age    int
	height float64
}

func NewPerson(name string, age int, height float64) *Person {
	return &Person{name, age, height}
}

func (p *Person) setName(name string) {

	p.name = name

}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) setAge(age int) {

	p.age = age
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) setHeight(height float64) {
	p.height = height
}

func (p *Person) getHeight() float64 {
	return p.height
}
