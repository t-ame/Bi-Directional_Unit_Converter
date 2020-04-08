package main

import (
	"fmt"
	"math"
)

//Converter is a Bi-Directional unit converter
type Converter struct {
}

type Feet float64
type Centimeter float64
type Minutes float64
type Seconds float64
type Milliseconds float64
type Celcius float64
type Fahrenheit float64
type Radian float64
type Degree float64
type Kilogram float64
type Pounds float64

const PI = math.Pi

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f * 30.48)
}

func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	return Seconds(m * 60)
}

func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

func (cvr Converter) SecondsToMilliseconds(s Seconds) Milliseconds {
	return Milliseconds(s * 1000)
}

func (cvr Converter) MillisecondsToSeconds(m Milliseconds) Seconds {
	return Seconds(m / 1000)
}

func (cvr Converter) CelciusToFahrenheit(c Celcius) Fahrenheit { //(0°C × 9/5) + 32
	return Fahrenheit((c * (9.0 / 5)) + 32)
}

func (cvr Converter) FahrenheitToCelcius(f Fahrenheit) Celcius { //(0°F − 32) × 5/9
	return Celcius((f - 32) * (5.0 / 9))
}

func (cvr Converter) RadianToDegree(r Radian) Degree { //1rad × 180/π
	return Degree(r * (180 / PI))
}

func (cvr Converter) DegreeToRadian(d Degree) Radian { //1° × π/180
	return Radian(d * (PI / 180))
}

func (cvr Converter) KilogramToPounds(k Kilogram) Pounds {
	return Pounds(k * 2.205)
}

func (cvr Converter) PoundsToKilogram(p Pounds) Kilogram {
	return Kilogram(p / 2.205)
}

func main() {
	converter := new(Converter)

	fmt.Println(converter.CentimeterToFeet(13.50))
	fmt.Println(converter.FeetToCentimeter(5))
	fmt.Println(converter.MinutesToSeconds(2))
	fmt.Println(converter.SecondsToMinutes(120))
	fmt.Println(converter.SecondsToMilliseconds(120))
	fmt.Println(converter.MillisecondsToSeconds(120000))
	fmt.Println(converter.CelciusToFahrenheit(28))
	fmt.Println(converter.FahrenheitToCelcius(80))
	fmt.Println(converter.RadianToDegree(0.5))
	fmt.Println(converter.DegreeToRadian(10))
	fmt.Println(converter.KilogramToPounds(100))
	fmt.Println(converter.PoundsToKilogram(241))

}
