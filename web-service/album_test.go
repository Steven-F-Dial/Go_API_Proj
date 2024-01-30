package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRouterRun(t *testing.T) {
	t.Parallel()
	router := gin.Default()
	go func(r *gin.Engine) {
		router.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		router.POST("/albums", addAlbum)
		routerRun(r, "localhost:8080")
	}(router)
}

func TestGetAlbums(t *testing.T) {
	t.Parallel()
	al, err := http.Get("http://localhost:8080/albums")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, al.Body)
	}
}

func TestAddAlbum(t *testing.T) {
	t.Parallel()
	data := []byte(`{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/albums", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("response Body:", string(body))

}

func TestGetAlbumByID(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Millisecond)
	al, err := http.Get("http://localhost:8080/albums/4")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, al.Body)
	}
}
