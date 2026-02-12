package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации: %v\n", err)
			continue
		}
		fmt.Println(info)
	}
}