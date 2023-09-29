package cmd

import "github.com/joho/godotenv"

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}
