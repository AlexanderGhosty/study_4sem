package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Storage interface {
	GetAll() ([]Item, error)
	AddItem(item Item) (Item, error)
	UpdateItem(id int, updated Item) error
}

type JSONStorage struct {
	filename string
	mu       sync.Mutex
	items    []Item
}

func NewJSONStorage(filename string) (*JSONStorage, error) {
	js := &JSONStorage{
		filename: filename,
		items:    []Item{},
	}

	if err := js.loadData(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return js, nil
}

// GetAll возвращает копию всех элементов
func (js *JSONStorage) GetAll() ([]Item, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	result := make([]Item, len(js.items))
	copy(result, js.items)
	return result, nil
}

// AddItem добавляет новый элемент в хранилище
func (js *JSONStorage) AddItem(item Item) (Item, error) {
	js.mu.Lock()
	defer js.mu.Unlock()

	maxID := 0
	for _, it := range js.items {
		if it.ID > maxID {
			maxID = it.ID
		}
	}
	item.ID = maxID + 1
	js.items = append(js.items, item)

	if err := js.saveData(); err != nil {
		return Item{}, err
	}
	return item, nil
}

// UpdateItem обновляет элемент с заданным ID
func (js *JSONStorage) UpdateItem(id int, updated Item) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	found := false
	for i, it := range js.items {
		if it.ID == id {
			js.items[i].Name = updated.Name
			js.items[i].Description = updated.Description
			found = true
			break
		}
	}
	if !found {
		return errors.New("элемент не найден")
	}
	return js.saveData()
}

// Загружает данные из файла
func (js *JSONStorage) loadData() error {
	file, err := os.Open(js.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&js.items); err != nil {
		return err
	}
	return nil
}

// Сохраняет данные в файл
func (js *JSONStorage) saveData() error {
	file, err := os.Create(js.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(js.items)
}
