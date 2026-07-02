package main

/* Создать приложение, которое сначала выдает меню:
 * - 1. Посмотреть закладки
 * - 2. Добавить закладку
 * - 3. Удалить закладку
 * - 4. Выход
 * При выборе 1 - выводит список закладок
 * При выборе 2 - вводит 2 поля (название, URL) и добавляет новую закладку
 * При выборе 3 - вводит название закладки и удаляет ее
 * При выборе 4 - выходит из приложения
 *
 * После пунктов 1, 2, 3 возвращается в меню
 */
import (
	"fmt"
)

func main() {
	bookmarks := make(map[string]string)
	for {
		choice := getMenuChoice()
		switch choice {
		case 1:
			displayBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}

func getMenuChoice() (choice int) {
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")

	fmt.Scanln(&choice)
	return
}

func displayBookmarks(bookmarks map[string]string) {
	fmt.Println("Закладки:")
	for name, url := range bookmarks {
		fmt.Printf("- %s: %s\n", name, url)
	}
}

func addBookmark(bookmarks map[string]string) {
	var name, url string
	fmt.Println("Введите название:")
	fmt.Scanln(&name)
	fmt.Println("Введите URL:")
	fmt.Scanln(&url)
	bookmarks[name] = url
}

func deleteBookmark(bookmarks map[string]string) {
	var name string
	fmt.Println("Введите название закладки:")
	fmt.Scanln(&name)
	delete(bookmarks, name)
}
