package main

import (
	"example/user/hello/business"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, world.")
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar3k.Print())
	fmt.Println(solar2k.Print())
	fmt.Println(windwest.Print())

	fmt.Printf("Solar: %s\n", business.PrintGeneric[business.Solar](solar3k))
	fmt.Printf("Wind: %s\n", business.PrintGeneric[business.Wind](windwest))
	fmt.Printf("Solar: %s\n", business.PrintGeneric(solar2k))
	fmt.Printf("Wind: %s\n", business.PrintGeneric(windwest))

	// Print functions for 01_04
	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss)
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

	// business.Cost(10, solar2k.Netto)
	// business.Cost(0.46, 10)
	// business.Cost(0.46, float64(10))
	// business.Cost[float64](0.46, 10)

	// Print functions for 03_02
	ss2 := business.SolarSlice{solar3k, solar2k}
	business.PrintSlice(ss2)
	business.PrintSlice2(ss2)

	// Print functions for 04_02
	fmt.Printf("index: %d\n", slices.Index(ss, solar2k))
	business.SortByCost(ss)
	fmt.Printf("index: %d\n", slices.Index(ss, solar2k))
	business.SortByCost(ss)
}
