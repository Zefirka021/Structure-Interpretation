package main

import (
	"fmt"
	"os"
	"time"
)

// Чайничек
type kettle struct {
	isFull   bool
	isBoiled bool
}

// Кран
type tap struct {
	isOpen bool
}

func (t *tap) open() {
	if !t.isOpen {
		t.isOpen = true
		fmt.Println("Кран открыт")
	} else {
		fmt.Println("Кран и так открыт")
	}

}

func (t *tap) close() {
	if t.isOpen {
		t.isOpen = false
		fmt.Println("Кран закрыт")
	} else {
		fmt.Println("Кран и так закрыт")
	}
}

func (t *tap) fill(k *kettle) {
	if t.isOpen {
		k.isFull = true
		fmt.Println("Наполнение чайника...")
		for i := 0; i < 10; i++ {
			fmt.Printf("Наполнение чайника - %v%%\n", (i+1)*10)
			time.Sleep(1)
		}
		fmt.Println("Чайник наполнен")
	} else {
		fmt.Println("Не удалось наполнить чайник")
	}

}

func (t *tap) empty(k *kettle) {
	k.isFull = false
	fmt.Println("Чайник пустой")
}

// Спички детям...
type matches struct {
	isLight bool
}

func (m *matches) light() {
	if !m.isLight {
		fmt.Println("Спичка зажглась")
		m.isLight = true
	} else {
		fmt.Println("Спичка и так горит")
	}
}

func (m *matches) putOut() {
	if m.isLight {
		fmt.Println("Спичка потухла")
		m.isLight = true
	} else {
		fmt.Println("Спичка и так не горит")
	}
}

// Газовая плита
type gasStove struct {
	isBurning bool
}

func (g *gasStove) lightStove(m *matches) {
	if m.isLight {
		g.isBurning = true
		fmt.Println("Газовая плита зажглась")
	} else {
		fmt.Println("Газовая плита не зажглась")
	}

}

func (g *gasStove) offStove() {
	if g.isBurning {
		g.isBurning = false
		fmt.Println("Газовая плита потухла")
	} else {
		fmt.Println("Газовая плита и так не горит")
	}
}

func (g *gasStove) heatKettle(k *kettle) {
	if k.isFull {
		fmt.Println("Кипечение воды...")
		for i := 0; i < 10; i++ {
			fmt.Printf("Кипечение воды - %v%%\n", (i+1)*10)
			time.Sleep(2)
		}
		k.isBoiled = true
		fmt.Println("Вода вскипела!")
	} else {
		fmt.Println("Воды в чайнике нет!")
		fmt.Println("Пожар!!!")
		os.Exit(0)
	}
}

func main() {
	var kettle1 kettle
	var tap1 tap
	var matches1 matches
	var gasStove1 gasStove

	fmt.Println("Первый случай:")
	tap1.open()                     // Открываем кран
	tap1.fill(&kettle1)             // Наполняем чайник
	tap1.close()                    //Закрываем кран
	matches1.light()                // Зажигаем спичку
	gasStove1.lightStove(&matches1) // Зажигаем газ
	matches1.putOut()               // Не забываем потушить спичку
	gasStove1.heatKettle(&kettle1)  // Кипятим чайник
	gasStove1.offStove()            // Выключаем плиту
	fmt.Println()

	var kettle2 kettle
	kettle2.isFull = true
	var tap2 tap
	var matches2 matches
	var gasStove2 gasStove

	fmt.Println("Второй случай:")
	tap2.empty(&kettle2)            // Выливаем чайник
	tap2.open()                     // Открываем кран
	tap2.fill(&kettle2)             // Наполняем чайник
	tap2.close()                    //Закрываем кран
	matches2.light()                // Зажигаем спичку
	gasStove2.lightStove(&matches2) // Зажигаем газ
	matches2.putOut()               // Не забываем потушить спичку
	gasStove2.heatKettle(&kettle2)  // Кипятим чайник
	gasStove2.offStove()            // Выключаем плиту
	fmt.Println()

}
