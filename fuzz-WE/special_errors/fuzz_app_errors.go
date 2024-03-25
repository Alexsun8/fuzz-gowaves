package special_errors

import (
	"fmt"
	"net/http"
)

// cgo LDFLAGS : - lm

/*
    #include <stdlib.h>
	#include <string.h>
	#include <pthread.h>


	// # cgo LDFLAGS : - lm

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
func Memory_leak() {
	fmt.Sprintf("memotyLeak was founded")
	C.memotyLeak()
}

// sermentation fault
func Sermentation_fault() {
	fmt.Sprintf("causeSegmentationFault was founded")
	C.causeSegmentationFault()
}

// переполнение буфера
func Buffer_overflow() {
	fmt.Sprintf("bufferOverflow was founded")
	C.bufferOverflow()
}

// Неопределенное поведение
func Undefined_behavior() {
	fmt.Sprintf("undefinedBehavior was founded")
	C.undefinedBehavior()
}

// Состязательные условия
func Adversarial_conditions() {
	fmt.Sprintf("adversarialConditions was founded")
	C.adversarialConditions()
}

// Некорректное использование памяти
func Incorrect_memory_usage() {
	fmt.Sprintf("incorrectMemoryUsage was founded")
	C.incorrectMemoryUsage()
}

// panic
func Get_panic() {
	fmt.Sprintf("Panic was founded")
	panic("Panic's founded by fuzzer!")
}

// panic or not_allowed_answer
// ToDo check!!!
func Not_allowed_answer(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintf("not_allowed_answer was founded")
	// В данном примере мы пытаемся получить доступ к несуществующему индексу в пустом срезе.
	var emptySlice []int
	_ = emptySlice[0] // Эта строка вызовет panic, потому что срез пустой и доступ к нулевому элементу не определён.
}
