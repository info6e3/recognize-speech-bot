package download_file

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, path string) {
	out, _ := os.Create(path)
	res, _ := http.Get(url)
	defer func() {
		res.Body.Close()
		out.Close()
	}()

	io.Copy(out, res.Body)
}
