package special_errors

import (
	"net/http"
)

/*
    #include <stdlib.h>
	#include <string.h>
	#include <pthread.h>


	# cgo LDFLAGS : - lm

    void causeSegmentationFault() {
        int *ptr = NULL;
        *ptr = 10;  // Trying to dereference a null pointer
    }

	void memotyLeak(){
		int *ptr = malloc(sizeof(int));
    	return 0;  // Утечка памяти, так как ptr не был освобожден
	}

	void bufferOverflow(){
		char buffer[10];
    	strcpy(buffer, "0123456789ABCDE");  // Переполнение буфера
    }

	void undefinedBehavior(){
		int a = 10, b = 0;
	    int result = a / b;  // Деление на ноль
    	printf("Result: %d\n", result);
	}

	int balance = 0;

	void *deposit(void *arg) {
		int amount = *(int *)arg;
		balance += amount;
		return NULL;
	}
	void adversarialConditions(){
	    pthread_t thread1, thread2;
		int amount1 = 100, amount2 = 200;
		pthread_create(&thread1, NULL, deposit, &amount1);
		pthread_create(&thread2, NULL, deposit, &amount2);
		pthread_join(thread1, NULL);
		pthread_join(thread2, NULL);
		printf("Final balance: %d\n", balance);
	}

	void incorrectMemoryUsage(){
		int *ptr = malloc(sizeof(int));
		free(ptr);
		*ptr = 42;  // Некорректное использование освобожденной памяти
	}
*/
import "C"

// Ошибка утечки памяти
func memory_leak() {
	C.memotyLeak()
}

// sermentation fault
func error_memory_read() {
	C.causeSegmentationFault()
}

// переполнение буфера
func buffer_overflow() {
	C.bufferOverflow()
}

// Неопределенное поведение
func undefined_behavior() {
	C.undefinedBehavior()
}

// Состязательные условия
func adversarial_conditions() {
	C.adversarialConditions()
}

// Некорректное использование памяти
func incorrect_memory_usage() {
	C.incorrectMemoryUsage()
}

// panic
func get_panic() {
	panic("Panic's founded by fuzzer!")
}

// panic or not_allowed_answer
// ToDo check!!!
func not_allowed_answer(w http.ResponseWriter, r *http.Request) {
	// В данном примере мы пытаемся получить доступ к несуществующему индексу в пустом срезе.
	var emptySlice []int
	_ = emptySlice[0] // Эта строка вызовет panic, потому что срез пустой и доступ к нулевому элементу не определён.
}
