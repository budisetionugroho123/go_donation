package migrations

import (
	"fmt"

	"github.com/budisetionugroho123/go_donation/internal/config"
	"github.com/budisetionugroho123/go_donation/internal/models"
)

func RunMigration() {
	err := config.InitDB().AutoMigrate(&models.Role{}, &models.User{}, &models.Organization{}, &models.Donation{}, &models.Transaction{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success run migration")
}
