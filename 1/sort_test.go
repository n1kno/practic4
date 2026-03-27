package main

import (
	"reflect"
	"sort"
	"testing"
)

// TestBubbleSortBasic тестирует базовую функциональность сортировки
func TestBubbleSortBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "один элемент",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "уже отсортированный",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "обратный порядок",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "случайный порядок",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "с отрицательными числами",
			input:    []int{-5, 3, -1, 0, -2, 4},
			expected: []int{-5, -2, -1, 0, 3, 4},
		},
		{
			name:     "с дубликатами",
			input:    []int{5, 2, 5, 1, 2, 5, 1},
			expected: []int{1, 1, 2, 2, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем копию входных данных для сравнения
			original := make([]int, len(tt.input))
			copy(original, tt.input)
			
			BubbleSort(tt.input)
			
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, ожидалось %v",
					original, tt.input, tt.expected)
			}
		})
	}
}

// TestBubbleSortWithRandomData тестирует сортировку на случайных данных
func TestBubbleSortWithRandomData(t *testing.T) {
	// Тестируем разные размеры слайсов
	sizes := []int{0, 1, 2, 10, 50, 100, 200}
	
	for _, size := range sizes {
		t.Run( fmt.Sprintf("размер_%d", size), func(t *testing.T) {
			// Создаем случайный слайс
			input := make([]int, size)
			for i := 0; i < size; i++ {
				input[i] = rand.Intn(2001) - 1000
			}
			
			// Создаем копию для сравнения с правильной сортировкой
			expected := make([]int, size)
			copy(expected, input)
			sort.Ints(expected) // Используем стандартную сортировку как эталон
			
			// Применяем нашу сортировку
			BubbleSort(input)
			
			// Сравниваем результаты
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("BubbleSort неверно отсортировала слайс размера %d\n"+
					"Получено: %v\nОжидалось: %v", size, input, expected)
			}
		})
	}
}

// TestBubbleSortStability тестирует, что сортировка изменяет исходный слайс
func TestBubbleSortStability(t *testing.T) {
	input := []int{3, 2, 1}
	original := []int{3, 2, 1}
	
	BubbleSort(input)
	
	// Проверяем, что исходный слайс был изменен
	if reflect.DeepEqual(input, original) {
		t.Errorf("BubbleSort не изменила исходный слайс: %v", input)
	}
}

// BenchmarkBubbleSort измеряет производительность сортировки
func BenchmarkBubbleSort(b *testing.B) {
	// Создаем тестовые данные разного размера
	sizes := []int{10, 50, 100, 500}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("размер_%d", size), func(b *testing.B) {
			// Создаем случайный слайс
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = rand.Intn(2001) - 1000
			}
			
			// Сбрасываем таймер и запускаем тест
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Создаем копию для каждого запуска
				testData := make([]int, size)
				copy(testData, data)
				BubbleSort(testData)
			}
		})
	}
}

// BenchmarkBubbleSortBestCase тестирует лучший случай (уже отсортированный)
func BenchmarkBubbleSortBestCase(b *testing.B) {
	size := 100
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testData := make([]int, size)
		copy(testData, data)
		BubbleSort(testData)
	}
}

// BenchmarkBubbleSortWorstCase тестирует худший случай (обратный порядок)
func BenchmarkBubbleSortWorstCase(b *testing.B) {
	size := 100
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = size - i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testData := make([]int, size)
		copy(testData, data)
		BubbleSort(testData)
	}
}
