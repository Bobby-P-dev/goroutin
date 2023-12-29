package code

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// time.After() adalah Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
// Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)
func TestTimeAfter(t *testing.T) {
	channel := time.After(10 * time.Second)

	tick := <-channel
	fmt.Println(tick)

}

/*
Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil
ketika Timer mengirim kejadiannya
*/
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("execute after 1 second")
		group.Done() // Saya sudah selesai
	})
	group.Wait()
}

/*
Ticker adalah representasi kejadian yang berulang
Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
*/

func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		time.Sleep(1 * time.Second)
		fmt.Println(tick)
		ticker.Stop()
	}
}

/*
Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
mengembalikan Ticker, hanya mengembalikan channel timer nya saja
*/

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}

/*
Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah
function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau
mengambil jumlah thread
Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/

func TestCpu(t *testing.T) {

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

}
