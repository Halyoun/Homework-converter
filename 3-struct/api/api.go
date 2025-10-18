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

type jsonBinResponse struct {
	Record   bins.BinList `json:"record"`
	Metadata any          `json:"metadata"`
}

func GetKey(cfg *config.Config) string {
	return cfg.Key
}

func CreateBin(filename, nameBin string, binlist *bins.BinList) {
	binlist.ScanJson(filename)
	postBody, err := json.Marshal(binlist)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Создаём кастомный POST-запрос
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/"+"", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err)
	}
	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", GetKey(config.NewConfig("Key.env")))
	req.Header.Set("X-Bin-Name", nameBin)

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

func UpdateBin(filename, id string, binlist *bins.BinList) {
	if id == "" {
		fmt.Println("ID НЕ БЫЛ ВВЕДЕН")
		return
	}
	binlist.ScanJson(filename)
	postBody, err := json.Marshal(binlist)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Println("✅ Bin успешно обновлен!")
	fmt.Println("Ответ сервера:", string(body))
}

func DeleteBin(id string) {
	if id == "" {
		fmt.Println("ID НЕ БЫЛ ВВЕДЕН")
		return
	}
	req, err := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Println("✅ Bin успешно удален!")
	fmt.Println("Ответ сервера:", string(body))
}

func GetBin(id string) (binlist bins.BinList) {
	req, err := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}

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
	var response jsonBinResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal("Ошибка Unmarshal:", err)
	}

	return response.Record
}
