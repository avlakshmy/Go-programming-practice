<<<<<<< HEAD
package furniture

type Furniture struct {
	Name string
	Material string
	Price int
}

func (f *Furniture) RetName () string {
	return f.Name
}

func (f *Furniture) RetPrice () int {
	return f.Price
}

func (f *Furniture) RetMaterial () string {
	return f.Material
}
=======
//This program illustrates a simple example of how interfaces are used in Go

package main

import "fmt"

type Furniture interface {
	Material () string
	Price () string
}

type chair struct {	
}

type table struct{
}

type sofa struct {
}

type bed struct {
}

func (c chair) Material() string {
    return "Teak wood"
}

func (c chair) Price() string {
    return "$50"
}

func (t table) Material() string {
    return "Oak wood"
}

func (t table) Price() string {
    return "$100"
}

func (s sofa) Material() string {
    return "Rosewood"
}

func (s sofa) Price() string {
    return "$300"
}

func (b bed) Material() string {
    return "Mahogany"
}

func (b bed) Price() string {
    return "$350"
}

func main() {
    furniture := []Furniture{chair{}, table{}, sofa{}, bed{}}
    for _, furn := range furniture {
	fmt.Printf("%T\n", furn)
        fmt.Println(furn.Material())
	fmt.Println(furn.Price())
	fmt.Println()
    }
}

>>>>>>> a1ce7d18a0f23696c45dc353ae81ad47b1e3c32f
