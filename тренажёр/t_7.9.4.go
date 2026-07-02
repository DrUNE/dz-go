package main

import "fmt"

// **Описание**: Создайте программу, которая использует type alias для типа map
// и изменяет значение существующего ключа в этой map, после чего выводит обновлённое значение.
//
// **Входные данные**: Готовый type alias и map с настройками приложения (встроены в код):
//   type SettingsMap = map[string]string
//   settings := SettingsMap{
//       "theme": "light",
//       "language": "en",
//       "timezone": "UTC",
//   }
// Ключ для изменения: "theme"
// Новое значение: "dark"
//
// **Выходные данные**: Одна строка в формате:
//   новое значение theme: <значение>
//
// **Ограничения**:
// - Обязательно использовать type alias SettingsMap для объявления переменной
// - Ключ и новое значение заданы как константы в коде
// - Работать с исходной map, не создавать новую
//
// **Примеры**:
// Input: settings типа SettingsMap с ключом "theme": "light", новое значение "dark"
// Output: новое значение theme: dark
//
// Input: settings типа SettingsMap с ключом "theme": "light", новое значение "blue"
// Output: новое значение theme: blue

func main() {
	type SettingsMap = map[string]string
	settings := SettingsMap{
		"theme":    "light",
		"language": "en",
		"timezone": "UTC",
	}
	key, value := "theme", "dark"
	settings[key] = value
	fmt.Printf("новое значение %s: %s\n", key, settings[key])
}
