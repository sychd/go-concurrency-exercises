package main

//
//import (
//	"fmt"
//	"image"
//	"log"
//	"net/http"
//	"os"
//	"path/filepath"
//	"runtime"
//	"sync"
//	"time"
//
//	"github.com/disintegration/imaging"
//)
//
//type pathsType struct {
//	path string; err error
//}
//
//type processedImage struct {
//	image *image.NRGBA; path string; err error
//}
//
//// Image processing - sequential
//// Input - directory with images.
//// output - thumbnail images
//func main() {
//	if len(os.Args) < 2 {
//		log.Fatal("need to send directory path of images")
//	}
//	//runtime.GOMAXPROCS(runtime.NumCPU())
//	start := time.Now()
//
//	err := <- walkFiles(os.Args[1])
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Time taken: %s\n, routines: %v\n", time.Since(start), runtime.NumGoroutine())
//}
//
//// walfiles - take diretory path as input
//// does the file walk
//// generates thumbnail images
//// saves the image to thumbnail directory.
//func walkFiles(root string) <-chan error {
//	pathsCh := make(chan pathsType)
//	wg := sync.WaitGroup{}
//
//	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		wg.Add(1)
//		go func(e error)  {
//			defer wg.Done() // QUESTION: when I set it not here but near ch dispatch, it hangs
//							// ANSWER: because in case of crash (if -> return) it does not reach this line, so WG is not closed
//							// ot happens on fmt.Println("Walk out #2") :)
//			if err != nil {
//				fmt.Println("Walk out #1")
//				return
//			}
//
//			// check if it is file
//			if !info.Mode().IsRegular() {
//				fmt.Println("Walk out #2")
//				return
//			}
//
//			// check if it is image/jpeg
//			contentType, _ := getFileContentType(path)
//			if contentType != "image/jpeg" {
//				fmt.Println("Walk out #3")
//				return
//			}
//			pathsCh <- pathsType{path, err}
//		}(err)
//
//		return nil
//	})
//
//	go func() {
//		wg.Wait()
//		close(pathsCh)
//	}()
//	//go func() {
//	//	for _, p := range arr {
//	//		pathsCh <- pathsType{p, nil}
//	//	}
//	//
//	//	close(pathsCh)
//	//}()
//
//
//	imgsCh := processImage(pathsCh)
//	return saveThumbnail(imgsCh)
//}
//
//// processImage - takes image file as input
//// return pointer to thumbnail image in memory.
//func processImage(paths <-chan pathsType) <-chan processedImage{
//	out := make(chan processedImage)
//	wg := sync.WaitGroup{}
//
//	for p := range paths {
//		wg.Add(1)
//		go func (path pathsType) {
//			if path.err != nil {
//				out <- processedImage{nil, path.path, path.err}
//				return
//			}
//			srcImage, err := imaging.Open(path.path)
//			thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)
//
//			out <- processedImage{thumbnailImage, path.path, err}
//			wg.Done()
//		}(p)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
//
//// saveThumbnail - save the thumbnail image to folder
//func saveThumbnail(images <-chan processedImage) <-chan error {
//	out := make(chan error)
//
//	wg := sync.WaitGroup{}
//
//	for img := range images {
//		wg.Add(1)
//		go func(image processedImage) {
//			if image.err != nil {
//				fmt.Println("error")
//				return
//			}
//
//			filename := filepath.Base(image.path)
//			dstImagePath := "thumbnail/" + filename
//
//			err := imaging.Save(image.image, dstImagePath)
//			if err != nil {
//				return
//			}
//
//			fmt.Printf("%s -> %s\n", image.path, dstImagePath)
//			wg.Done()
//			}(img)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
//
//// getFileContentType - return content type and error status
//func getFileContentType(file string) (string, error) {
//
//	out, err := os.Open(file)
//	if err != nil {
//		return "", err
//	}
//	defer out.Close()
//
//	// Only the first 512 bytes are used to sniff the content type.
//	buffer := make([]byte, 512)
//
//	_, err = out.Read(buffer)
//	if err != nil {
//		return "", err
//	}
//
//	// Use the net/http package's handy DectectContentType function. Always returns a valid
//	// content-type by returning "application/octet-stream" if no others seemed to match.
//	contentType := http.DetectContentType(buffer)
//
//	return contentType, nil
//}
