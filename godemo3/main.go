package main

import "fmt"

func main() {
	bookmarks := map[string]string{}
	var action int

	for action >= 0 && action < 4 {
		fmt.Println("1 - Посмотреть закладки")
		fmt.Println("2 - Добавить закладку")
		fmt.Println("3 - Удалить закладку")
		fmt.Println("4 - Удалить закладку")
		fmt.Scan(&action)

		switch action {
		case 1:
			fmt.Println(bookmarks)
		case 2:
			var k, v string
			fmt.Scan(&k, &v)
			bookmarks[k] = v
		case 3:
			var k string
			fmt.Scan(&k)
			delete(bookmarks, k)
		}

	}
}
