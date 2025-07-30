package pizza

type Pizza struct {
	size        string
	cheese      bool
	pepperoni   bool
	mushroom    bool
	spinach     bool
	prosciutto  bool
	pesto       bool
	basil       bool
	tomatoSauce bool
}
type Builder struct {
	size        string
	cheese      bool
	pepperoni   bool
	mushroom    bool
	spinach     bool
	prosciutto  bool
	pesto       bool
	basil       bool
	tomatoSauce bool
}

func NewPizzaBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Size(size string) *Builder {
	b.size = size
	return b
}

func (b *Builder) AddCheese() *Builder {
	b.cheese = true
	return b
}

func (b *Builder) AddPepperoni() *Builder {
	b.pepperoni = true
	return b
}

func (b *Builder) AddMushroom() *Builder {
	b.mushroom = true
	return b
}

func (b *Builder) AddSpinach() *Builder {
	b.spinach = true
	return b
}

func (b *Builder) AddProsciutto() *Builder {
	b.prosciutto = true
	return b
}

func (b *Builder) AddPesto() *Builder {
	b.pesto = true
	return b
}

func (b *Builder) AddBasil() *Builder {
	b.basil = true
	return b
}

func (b *Builder) AddTomatoSauce() *Builder {
	b.tomatoSauce = true
	return b
}

func (b *Builder) Build() Pizza {
	return Pizza{b.size, b.cheese, b.pepperoni, b.mushroom, b.spinach, b.prosciutto, b.pesto, b.basil, b.tomatoSauce}
}
