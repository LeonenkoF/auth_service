package main

import (
	"auth_service/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Errorf("failed to init config: %v", err)
		return
	}

	fmt.Printf("port: %v", cfg.DBHost)

}
