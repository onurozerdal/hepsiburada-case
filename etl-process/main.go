package main

import "hepsiburada-case/etl-process/bestseller"
import "hepsiburada-case/etl-process/product"

func main() {
	bestseller.Find()
	product.Process()
}
