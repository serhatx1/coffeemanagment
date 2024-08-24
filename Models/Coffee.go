package Models

type Coffee struct {
	temperature string
	drinkName   string
	price       float32
	calories    float32
	imageUrl    string
	recipe      map[string]int
}
type CoffeeI struct {
	createCoffee func()
	createRecipe func()
}

func (c *CoffeeManagment) createCoffee(temperature, drinkname, imageUrl string, calories, price float32, recipe map[string]int) *Coffee {
	return &Coffee{
		temperature: temperature,
		drinkName:   drinkname,
		price:       price,
		calories:    calories,
		imageUrl:    imageUrl,
		recipe:      recipe,
	}

}
func (c *Coffee) createRecipe(r *Coffee, recipe map[string]int) *Coffee {
	for k, v := range recipe {
		r.recipe[k] = v
	}
	return r
}
