package main

import (
	adapter "github.com/404th/hex_arch/internal/adapters/core/arithmetic"
	"github.com/404th/hex_arch/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPorts = adapter.NewArithmeticAdapter()
}
