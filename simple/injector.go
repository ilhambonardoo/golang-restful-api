//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}

/*
	Injector
	- Setelah kita membuat Provider untuk nanti kita gunakan, selanjutnya kita perlu membuat Injector
	- Injector sendiri adalah sebuah function constructor, namun isinya berupa konfigurasi yang kita beritahukan ke Google Wire
	- Injector ini sendiri sebenarnya tidak akan digunakan oleh kode program kita, Injector ini adalah function yang akan digunakan oleh Google Wire untuk melakukan auto generate kode Dependency Injection
	- Khusus ketika membuat Injector, pada file nya kita perlu tambahkan komentar penanda :
		//go:build wireinject
		// +build wireinject
*/

/*
	Dependency Injection
	- Setelah kita membuat Injector dan Provider, selanjutnya yang perlu kita lakukan adalah menggunakan aplikasi command line Google Wire untuk melakukan auto generate kode Dependency Injection
	- Kita bisa menggunakan perintah ini untuk melakukan auto generate kode dependency injection :
	- wire gen namapackage
	- Secara otomatis aplikasi Google Wire akan mencari kode Injector di package tersebut, lalu membuat file wire_gen.go yang isinya adalah kode otomatis dependency injection
*/
