package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogError(errorMessage string) {
	filePath := "/home/log/snmp_server_error.log"

	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		log.Printf("Zaman dilimi yüklenirken hata oluştu: %s\n", err)
		return
	}
	logTime := time.Now().In(loc).Format(time.RFC3339)

	logMessage := fmt.Sprintf("[%s] %s\n", logTime, errorMessage)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Hata günlüğü dosyasına yazma hatası: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(logMessage)
	if err != nil {
		log.Printf("Hata günlüğü dosyasına yazma hatası: %s\n", err)
	}
}
