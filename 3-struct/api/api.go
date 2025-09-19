package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gobasics/bins"
	"gobasics/config"
	"gobasics/file"
	"io"
	"net/http"
	"slices"
	"time"
)

type RespStruct struct {
	Metadata struct {
		Id        string    `json:"id"`
		Private   bool      `json:"private"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
}

func CreateBin(c *config.Config, binListWithDb *bins.BinListWithDb, filePath *string, name *string) ([]byte, error) {
	if !file.IsJson(*filePath) {
		return []byte{}, errors.New("not json file")
	}

	fileBody, err := file.ReadSomeFile(*filePath)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl, bytes.NewBuffer(fileBody))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", c.Key)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return []byte{}, err
	}

	var respJson RespStruct
	json.Unmarshal(body, &respJson)
	newBin := bins.NewBin(respJson.Metadata.Id, respJson.Metadata.Private, respJson.Metadata.CreatedAt, *name)

	binListWithDb.Bins = append(binListWithDb.Bins, *newBin)
	content, _ := json.Marshal(binListWithDb.BinList.Bins)

	binListWithDb.Db.WriteStorage(content)
	fmt.Println("Ответ сервера:", string(body))
	return body, nil
}

func GetBin(c *config.Config, id *string) ([]byte, error) {
	url := c.BaseUrl + "/" + *id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", c.Key)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return []byte{}, err
	}
	return body, nil
}

func DeleteBin(c *config.Config, binListWithDb *bins.BinListWithDb, id *string) ([]byte, error) {
	url := c.BaseUrl + "/" + *id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", c.Key)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return []byte{}, err
	}

	idx := slices.IndexFunc(binListWithDb.Bins, func(b bins.Bin) bool { return b.Id == *id })
	binListWithDb.Bins = append(binListWithDb.Bins[:idx], binListWithDb.Bins[idx+1:]...)
	content, _ := json.Marshal(binListWithDb.BinList.Bins)
	binListWithDb.Db.WriteStorage(content)

	return body, nil
}

func UpdateBin(c *config.Config, binListWithDb *bins.BinListWithDb, id *string, filePath *string) ([]byte, error) {
	if !file.IsJson(*filePath) {
		return []byte{}, errors.New("not json file")
	}

	fileBody, err := file.ReadSomeFile(*filePath)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	url := c.BaseUrl + "/" + *id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(fileBody))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", c.Key)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return []byte{}, err
	}

	idx := slices.IndexFunc(binListWithDb.Bins, func(b bins.Bin) bool { return b.Id == *id })
	binListWithDb.Bins[idx].CreatedAt = time.Now()
	content, _ := json.Marshal(binListWithDb.BinList.Bins)
	binListWithDb.Db.WriteStorage(content)

	return body, nil
}
