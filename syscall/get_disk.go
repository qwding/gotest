package main

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func main() {
	disk := DiskUsage("/")
	fmt.Printf("disk %v\n", disk)
	fmt.Println("disk all:", disk.All)
	all := disk.All / (1024 * 1024)
	fmt.Println("all :", all, "M")
	fmt.Println("all:", all/1024, "G")
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
