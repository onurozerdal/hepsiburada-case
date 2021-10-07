package main

import "github.com/onurozerdal/hepsiburada-case/etl-process/bestseller"
import "github.com/onurozerdal/hepsiburada-case/etl-process/product"

func main() {
	bestseller.Find()
	product.Process()
}
