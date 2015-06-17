package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("run Server")

	http.Handle("/access/", http.StripPrefix("/access/", http.FileServer(http.Dir("."))))
	// http.HandleFunc("/thumbnail", thumbnailHandler)
	http.HandleFunc("/play/", playHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/files/", filesHandler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	//fmt.Fprintf(w, "[%s]", strings.Trim(r.URL.Path, "/files/"))
	nextPath := strings.TrimPrefix(r.URL.Path, "/files/")
	fmt.Printf("nextPath=%s\n", nextPath)

	list, err := ioutil.ReadDir("./" + nextPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "<html>")
	for _, finfo := range list {
		// if finfo.IsDir() || strings.Index(finfo.Name(), ".mp4") == -1 {
		// 	continue
		// }
		realPath := nextPath + "/" + finfo.Name()
		fmt.Printf("realPath=%s\n", realPath)
		if finfo.IsDir() {
			fmt.Fprintf(w, "<a href='/files/%s'>%s/</a><br>", realPath, finfo.Name())
		} else if strings.Index(finfo.Name(), ".mp4") > 0 {
			fmt.Fprintf(w, "<a href='/play/%s'>%s</a><br>", realPath, finfo.Name())
		}

	}

	// str := `
	// <form action="http://localhost:9999/thumbnail" method="post" enctype="multipart/form-data">
	// <label for="file">Filename:</label>
	// <input type="file" name="file" id="file">
	// <input type="submit" name="submit" value="Submit">
	// </form>
	// `
	//
	// fmt.Fprintf(w, "%s", str)
	//
	fmt.Fprintf(w, "</html>")

}

// カレントフォルダ内のmp4ファイルの一覧表示
func rootHandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("run rootHandler")
	list, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "<html>")
	for _, finfo := range list {
		// if finfo.IsDir() || strings.Index(finfo.Name(), ".mp4") == -1 {
		// 	continue
		// }
		if finfo.IsDir() {
			fmt.Fprintf(w, "<a href='/files/%s'>%s/</a><br>", finfo.Name(), finfo.Name())
		} else if strings.Index(finfo.Name(), ".mp4") > 0 {
			fmt.Fprintf(w, "<a href='/play/%s'>%s</a><br>", finfo.Name(), finfo.Name())
		}

	}

	// str := `
	// <form action="http://localhost:9999/thumbnail" method="post" enctype="multipart/form-data">
	// <label for="file">Filename:</label>
	// <input type="file" name="file" id="file">
	// <input type="submit" name="submit" value="Submit">
	// </form>
	// `
	//
	// fmt.Fprintf(w, "%s", str)
	//
	fmt.Fprintf(w, "</html>")

}

func playHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "run play handler %v", strings.Trim(r.URL.Path, "/play/"))
	filename := strings.TrimPrefix(r.URL.Path, "/play/")
	fmt.Printf("r.URL.Path=%s\n", r.URL.Path)
	fmt.Printf("filename=%s\n", filename)
	//				poster="http://video-js.zencoder.com/oceans-clip.png"

	playhtml := `
	<html>
		<head>
		<link href="//vjs.zencdn.net/4.12/video-js.css" rel="stylesheet">
		<script src="//vjs.zencdn.net/4.12/video.js"></script>
		</head>
		<body>
		<video id="example_video_1" class="video-js vjs-default-skin"
			controls preload="auto" width="1024" height="768"
				data-setup='{"example_option":true}'>
				<source src="/access/` + filename + `" type='video/mp4' />
				<p class="vjs-no-js">To view this video please enable JavaScript, and consider upgrading to a web browser that <a href="http://videojs.com/html5-video-support/" target="_blank">supports HTML5 video</a></p>
				</video>
		</body>
	</html>
	`
	fmt.Fprint(w, playhtml)

}

//
// func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
//
// 	if r.Method == "POST" {
// 		file, header, err := r.FormFile("file")
//
// 		if err != nil {
// 			fmt.Printf("error %v", err)
// 			return
// 		}
//
// 		defer file.Close()
//
// 		out, err := os.Create("test.jpg")
// 		if err != nil {
// 			fmt.Printf("error %v", err)
// 			return
// 		}
//
// 		defer out.Close()
//
// 		_, err = io.Copy(out, file)
// 		if err != nil {
// 			fmt.Printf("error %v", err)
// 			return
// 		}
//
// 		fmt.Printf("file upload success! %v", header.Filename)
//
// 	} else {
// 		fmt.Printf("Method is %v", r.Method)
// 	}
// }
