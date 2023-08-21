package main

import (
	"github.com/gin-gonic/gin"
	"wgnalvian.com/test1/controller"
	"wgnalvian.com/test1/service"
)

func main() {
	r := gin.Default()
	productService := service.NewProductService()
	productController := controller.ProductController{
		ProductService: productService,
	}

	productController.Route(r)
	// Memantau perubahan pada file
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer watcher.Close()

	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			if event.Op&fsnotify.Write == fsnotify.Write {
	// 				if strings.HasSuffix(event.Name, ".go") {
	// 					log.Println("File diubah:", event.Name)
	// 					RestartApp()
	// 				}
	// 			}
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("Error:", err)
	// 		}
	// 	}
	// }()

	// Menambahkan file yang akan dipantau
	// err = watcher.Add(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	r.Run(":3000")
}

// func RestartApp() {
// 	log.Println("Menghentikan server...")
// 	cmd := exec.Command("pkill", "-f", "main.go")
// 	cmd.Run()

// 	log.Println("Menjalankan ulang server...")
// 	go runMain()
// }

// func runMain() {
// 	cmd := exec.Command("go", "run", "main.go")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }

// func init() {
// 	go runMain()
// }
