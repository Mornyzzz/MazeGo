# Название проекта
PROJECT_NAME := maze

# Директории
DIST_DIR := ../dist
BUILD_DIR := ../build
TESTS_DIR := tests

ARCHIVE_NAME := maze_v1.0.tar.gz
all: clean install tests

dvi:
	@echo "Создание документации..."
	@echo "Перейдите на http://localhost:6060/pkg/maze/entity/"
	@cd entity && godoc -http :6060

install:
	@echo "Установка проекта..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(PROJECT_NAME)
	@mv $(PROJECT_NAME) $(BUILD_DIR)
	@echo "Проект установлен"

uninstall:
	@echo "Удаление проекта..."
	@rm -rf $(BUILD_DIR)/$(PROJECT_NAME)
	@echo "Проект удален."

clean:
	@echo "Очистка сборок..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(DIST_DIR)
	@rm tests/maze.txt
	@echo "Очистка завершена."

tests:
	@echo "Запуск тестов..."
	@go test $(TESTS_DIR)/*.go
	@echo "Тесты завершены."

dist:
	@echo "Создание дистрибутива..."
	@mkdir -p $(DIST_DIR)
	@go build -o $(PROJECT_NAME)
	@cp ../README_RUS.md .
	@tar -cvf $(ARCHIVE_NAME) $(PROJECT_NAME) README_RUS.md
	@mv $(ARCHIVE_NAME) $(DIST_DIR)
	@rm README_RUS.md maze
	@echo "Дистрибутив создан."

.PHONY: all install uninstall clean dvi dist tests