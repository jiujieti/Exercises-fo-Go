package weightconv

// KgToP converts a kilogram-weight to pounds
func KgToP(k Kilograms) Pounds { return Pounds(k * 2.204623) }

// PToKg converts a pound-weight to a kilograms
func PToKg(p Pounds) Kilograms { return Kilograms(p * 0.453592) }
