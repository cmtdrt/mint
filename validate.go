package mint

import "fmt"

// Validate checks internal integrity of the amount.
func (m Mint) Validate() error {
	if m.Amount == nil {
		return fmt.Errorf("mint: nil amount")
	}
	if m.Scale < 0 {
		return fmt.Errorf("mint: negative scale")
	}
	if !m.Currency.isValid() {
		return fmt.Errorf("mint: invalid currency")
	}
	return nil
}
