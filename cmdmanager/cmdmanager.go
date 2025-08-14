package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter the prices. Confirm every price with ENTER:")

	var prices []string
	for {
		fmt.Println("Price: ")
		var price string
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
