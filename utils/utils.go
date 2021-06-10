package utils

import (
	"errors"
	"math"
	"sort"

	"github.com/charlesonunze/busha-test/model"
)

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ConvertToFeet(height int) (feet, inch float64) {
	length := float64(height) / 2.54
	feet = math.Floor(length / float64(12))
	inch = math.Floor((length-float64(12*feet))*100) / 100
	return
}

func Sort(sortBy, sortOrder string, characters []model.Character) ([]model.Character, error) {
	switch sortBy {
	case "name":
		sort.SliceStable(characters, func(i, j int) bool {
			if sortOrder == "desc" {
				return characters[i].Name < characters[j].Name
			}
			return characters[i].Name > characters[j].Name
		})

	case "gender":
		sort.SliceStable(characters, func(i, j int) bool {
			if sortOrder == "desc" {
				return characters[i].Gender < characters[j].Gender
			}
			return characters[i].Gender > characters[j].Gender
		})

	case "height":
		sort.SliceStable(characters, func(i, j int) bool {
			if sortOrder == "desc" {
				return characters[i].HeightInCm > characters[j].HeightInCm
			}
			return characters[i].HeightInCm < characters[j].HeightInCm
		})

	default:
		return characters, errors.New("invalid sort param")
	}

	return characters, nil
}

func FilterByGender(c []model.Character, gender string) ([]model.Character, error) {
	result := []model.Character{}

	switch gender {
	case "male":
	case "female":
		break

	default:
		return result, errors.New("invalid gender")
	}

	for i := range c {
		if c[i].Gender == gender {
			result = append(result, c[i])
		}
	}

	return result, nil
}
