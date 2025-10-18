package api

import (
	"app-3/bins"
	"app-3/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetKey(cfg *config.Config) string {
	return cfg.Key
}

func Operation(operation string) {
	//var binlist = bins.BinList{}
	switch operation {
	case "create":
		//CreateBin()
	case "update":
	}
}

func CreateBin(filename string, binlist *bins.BinList) {
	binlist.ScanJson(filename)
	postBody, err := json.Marshal(binlist)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Создаём кастомный POST-запрос
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err)
	}
	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", GetKey(config.NewConfig("Key.env")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Ошибка: статус %d\nОтвет: %s", resp.StatusCode, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка чтения ответа:", err)
	}
	fmt.Println("✅ Bin успешно создан!")
	fmt.Println("Ответ сервера:", string(body))
}

func UpdateBin() {}

func DeleteBin() {}

func GetBin() {}
