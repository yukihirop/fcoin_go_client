package orders

import (
	"errors"
	"fmt"
)

func contains(s []string, e string) (ret bool) {
	ret = false
	for _, v := range s {
		if e == v {
			ret = true
		}
	}
	return
}

func slice2map(data []map[string]string) (ret map[string]string) {
	ret = make(map[string]string)
	if data == nil {
		ret = map[string]string{}
	} else {
		for _, el := range data {
			for k, v := range el {
				ret[k] = v
			}
		}
	}
	return
}

func includesErrorMessage(target interface{}, targetType string, list []string) (ret error) {
	var targetValue string
	if target != nil {
		targetValue = fmt.Sprint(target)
	} else {
		targetValue = "nil"
	}
	errorMessageValue := fmt.Sprintf("%s is %s. %s is not included in the %s.", targetType, targetValue, targetType, list)
	errorMessage := fmt.Sprintf("{%s: %s}", targetType, errorMessageValue)
	ret = errors.New(errorMessage)
	return
}

func betweenErrorMessage(target interface{}, targetType string, min float64, max float64) (ret error) {
	var targetValue string
	if target != nil {
		targetValue = fmt.Sprint(target)
	} else {
		targetValue = "nil"
	}
	errorMessageValue := fmt.Sprintf("%s is %s. %s is not betweeen %s and %s.", targetType, targetValue, targetType, fmt.Sprint(min), fmt.Sprint(max))
	errorMessage := fmt.Sprintf("{%s: %s}", targetType, errorMessageValue)
	ret = errors.New(errorMessage)
	return
}

func presenceErrorMessage(target interface{}, targetType string) (ret error) {
	var targetValue string
	if target != nil {
		targetValue = fmt.Sprint(target)
	} else {
		targetValue = "nil"
	}
	errorMessageValue := fmt.Sprintf("%s is %s. %s can't be blank.", targetType, targetValue, targetType)
	errorMessage := fmt.Sprintf("{%s: %s}", targetType, errorMessageValue)
	ret = errors.New(errorMessage)
	return
}
