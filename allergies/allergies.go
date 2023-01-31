package allergies

const (
	eggs         = 1
	peanuts      = 2
	shellfish    = 4
	strawberries = 8
	tomatoes     = 16
	chocolate    = 32
	pollen       = 64
	cats         = 128
)

var allItems = map[uint]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats",
}

func Allergies(allergies uint) []string {
	items := make([]string, 0)
	num := getMaxRange(allergies)

	for num != 0 {
		if num <= allergies {
			allergies -= num
			if item, ok := allItems[num]; ok {
				items = append(items, item)
			}
		}
		num /= 2
	}
	return items
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, item := range Allergies(allergies) {
		if item == allergen {
			return true
		}
	}
	return false
}

func getMaxRange(allergies uint) uint {
	var i uint = 1
	for i < allergies {
		i = i * 2
	}
	return i
}
