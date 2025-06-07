package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func CopyFiles(sourceDir string, destiDir string) {

	srcDir := sourceDir
	dstDir := destiDir

	// Buat folder tujuan jika belum ada
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		fmt.Println("Gagal membuat folder tujuan:", err)
	} else {
		// Baca isi folder sumber (pakai os.ReadDir)
		entries, err := os.ReadDir(srcDir)
		if err != nil {
			fmt.Println("Gagal membaca folder sumber:", err)
			return
		} else {
			for _, entry := range entries {
				if entry.IsDir() {
					continue // lewati subfolder
				}

				srcPath := filepath.Join(srcDir, entry.Name())
				dstPath := filepath.Join(dstDir, entry.Name())

				err := CopyF(srcPath, dstPath)
				if err != nil {
					fmt.Printf("Gagal menyalin %s: %v\n", entry.Name(), err)
				} else {
					fmt.Printf("Berhasil menyalin %s\n", entry.Name())
				}
			}
		}
	}

}
func CopyF(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func Deletefiles(foldername string) {
	folder := foldername

	// Baca semua isi folder
	entries, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Gagal membaca folder:", err)
	} else {
		// Loop dan hapus satu per satu
		for _, entry := range entries {
			path := filepath.Join(folder, entry.Name())

			// os.RemoveAll bisa hapus file atau folder
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Printf("Gagal menghapus %s: %v\n", path, err)
			} else {
				fmt.Printf("Berhasil menghapus: %s\n", path)
			}
		}
	}

}

func CopyFileTask() {
	var sourceF string
	var destinationF string

	fmt.Print("Masukkan nama folder asal : ")
	fmt.Scanln(&sourceF)

	fmt.Print("Masukkan nama folder tujuan : ")
	fmt.Scanln(&destinationF)

	if sourceF != "" && destinationF != "" {
		// COPY FILES
		CopyFiles(sourceF, destinationF)

		// DELETE FILES AFTER SUCCESS COPY
		Deletefiles(sourceF)
	}
}

func CronTask() {
	fmt.Println("Menjalankan tugas pada:", time.Now().Format("15:04:05"))
}
