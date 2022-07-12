package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type APPConfig struct {
	Port   string
	SectorId   float64
}

func (c *APPConfig) Read() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	c.Port=os.Getenv("APP_PORT")
	s:=os.Getenv("SECTOR_ID")
	sectorIdStr ,err := strconv.ParseFloat(s,64)

	if err!=nil {

		log.Fatalf("Fail to convert")
	}
	c.SectorId=sectorIdStr
}