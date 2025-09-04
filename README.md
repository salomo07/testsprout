==How to get Project==
- Ensure git was installed on yout computer
- git clone https://github.com/salomo07/testsprout.git
- cd testsprout


==How to Running & Play==
- Ensure Golang was installed
- go run main.go
- sistem akan meminta inputan, ketikkan titik awal bidak + spasi + titik tujuan (misal "b2 b3")
- lihat pada board, apakah bidak berubah posisi
- jika bidak tidak berpindah sesuai rule standar catur akan muncul error


==USECASE CATCHING KING==
- White (g1 f3)
- Black (e7 e6)
- White (f3 e5) 
- Black (f8 e7)
- White (d2 d3)
- Black (e7 b4)
- White (h2 h3)
- Black (b4 e1)
Black captured the King. Black wins!


==Unit Tests==
- Test board initialization : Ketika go run main.go board akan muncul dengan kondisi awal
- Test valid and invalid moves for different pieces : Coba lakukan langkah yang illegal ("f8 b2") akan ada error seperti "Move error: illegal move for piece"
- Test win condition (king capture) : Usecase diatas sudah cukup men-simulasikan sampai King tertangkap. Atau bisa lakukan unit test yang lain, yang saya siapkan. Jalankan go test ./chess_test -run TestWinByKingCapture -v