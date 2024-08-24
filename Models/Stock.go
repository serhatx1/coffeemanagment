package Models

import (
	"fmt"
	"reflect"
	"strings"
)

var StockSingelton *Stock

type StockI struct {
	createStock func()
	editStock   func()
}
type Stock struct {
	Milk            int
	MilkLactoseFree int
	Arabica         int
	Robusta         int
	Espresso        int
	Mugs            int
	VanillaSyrup    int
	CaramelSyrup    int
	NutSyrup        int
	Cream           int
}

func (c *CoffeeManagment) createStock(milk, milkLactoseFree, arabica, robusta, espresso, mugs, vanillaSyrup, caramelSyrup, nutSyrup, cream int) *Stock {
	if StockSingelton == nil {
		StockSingelton = &Stock{}
	} else {
		fmt.Println("Its already created")
	}
	return &Stock{
		Milk:            milk,
		MilkLactoseFree: milkLactoseFree,
		Arabica:         arabica,
		Robusta:         robusta,
		Espresso:        espresso,
		Mugs:            mugs,
		VanillaSyrup:    vanillaSyrup,
		CaramelSyrup:    caramelSyrup,
		NutSyrup:        nutSyrup,
		Cream:           cream,
	}

}
func (s *Stock) EditStock(EditedValue map[string]int) {
	value := reflect.ValueOf(s)
	typeOfS := value.Type()
	i := 0
	for k, v := range EditedValue {
		if strings.ToUpper(k) == strings.ToUpper(typeOfS.Field(i).Name) {
			fieldValue := value.Field(i)
			if fieldValue.IsValid() && fieldValue.CanSet() && fieldValue.Kind() == reflect.Int {
				fieldValue.SetInt(int64(v))
			}
		}
		i += 1
	}
}

func (s *Stock) CheckStock(c *Coffee) bool {
	value := reflect.ValueOf(s).Elem()

	for k, v := range c.recipe {
		fieldName := strings.ToUpper(k)
		found := false

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			if strings.ToUpper(field.Name) == fieldName {
				found = true
				fieldValue := value.Field(i).Interface()

				// Ensure the field value is of int type
				if fieldValue, ok := fieldValue.(int); ok {
					if fieldValue < v {
						return false
					}
				} else {
					return false
				}
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
