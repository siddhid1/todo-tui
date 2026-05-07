package main

func main() {
	todos := Todos{}
	todos.add("Task1")
	todos.add("Task2")
	todos.add("Task3")

	todos.toggle(0)
	todos.delete(1)
	todos.print()
}
