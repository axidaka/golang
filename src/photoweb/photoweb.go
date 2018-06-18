package main
import (
	"os"
	"io"
	"path"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"runtime/debug"
	"fmt"
)

const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
	}

	templatePath = TEMPLATE_DIR + "/" + templateName
	log.Println("Loading template:", templatePath)
	t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
		if e, ok := recover().(error); ok {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		// 或者输出自定义的 50x 错误页面
		// w.WriteHeader(http.StatusInternalServerError)
		// renderHtml(w, "error", e)
		// logging
		log.Println("WARN: panic in %v - %v", fn, e)
		log.Println(string(debug.Stack()))
		}
		}()
		fn(w, r)
	}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// t, err := template.ParseFiles("view/upload.html")
		// if err != nil {
		// 	http.Error(w, err.Error(),http.StatusInternalServerError)
		// 	return
		// }
		// t.Execute(w, nil)

		templates["upload.html"].Execute(w, nil)
		return
	}

	if r.Method == "POST" {

		fmt.Println("Post:", *r)
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		// if err != nil {
		// 	http.Error(w, err.Error(),
		// 	http.StatusInternalServerError)
		// 	return
		// }
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		// if err != nil {
		// 	http.Error(w, err.Error(),
		// 	http.StatusInternalServerError)
		// 	return
		// }
		http.Redirect(w, r, "/view?id="+filename,
		http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view:", *r)
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	// if err != nil {
	// 	http.Error(w, err.Error(),
	// 	http.StatusInternalServerError)
	// 	return
	// }

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images 
	// t, err:= template.ParseFiles("view/list.html")
	// if err != nil {
	// 	http.Error(w, err.Error(),
	// 	http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w, locals)
	templates["list.html"].Execute(w, locals)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}