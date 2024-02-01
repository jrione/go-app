package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var pilihan = [5]string{
	"Luas Persegi",
	"Luas Segitiga",
	"Luas Trapesium",
	"Volume Kubus",
	"Volume Balok",
}

type Itung struct {
	Angka1 int
	Angka2 int
}

func (i *Itung) SetValue(angka1, angka2 int) {
	i.Angka1 = angka1
	i.Angka2 = angka2
}

func (i *Itung) Persegi() {
	fmt.Println(i.Angka1 * i.Angka2)
}

func (i *Itung) Segitiga() {
	res := i.Angka1 * i.Angka2 * 1 / 2
	fmt.Println(res)
}

func (i *Itung) Testi() {
	fmt.Println(i.Angka1)
	fmt.Println(i.Angka2)
}

func clearscreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println("Kalkulator Bangun Datar & Bangun Ruang")
	fmt.Println("======================================")
	for k, v := range pilihan {
		fmt.Printf("%d.) %s \n", k+1, v)
	}

	var pilih int
	fmt.Printf("Inputkan angka: ")
	_, err := fmt.Scan(&pilih)
	if err != nil {
		fmt.Print("Masukkan angka yang benar!")
		panic(err)
	}

	fmt.Println("======================================")
	var props Itung
	goMath(props, pilih, pilihan[pilih-1])
	fmt.Println("======================================")

}

func goMath(props Itung, pilih int, judul string) {
	judul = cases.Title(language.English, cases.Compact).String(strings.ToLower(judul))
	words := strings.Split(judul, " ")
	fmt.Printf("Menghitung, %s\n", judul)
	getResult(props, words[len(words)-1])
}

func getResult(props Itung, methodName string) {
	angka := props

	v := [2]int{0, 0}
	for i := 0; i < len(v); i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		_, err := fmt.Scan(&v[i])
		if err != nil {
			panic(err)
		}
	}
	angka.SetValue(v[0], v[1])

	defer func() {
		fmt.Print("Hasil: ")
		reflect.ValueOf(&angka).MethodByName(methodName).Call([]reflect.Value{})
	}()

}
