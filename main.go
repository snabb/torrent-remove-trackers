package main

import (
	"fmt"
	"os"

	"github.com/anacrolix/torrent/metainfo"
)

func processFile(fn string) error {
	mi, err := metainfo.LoadFromFile(fn)
	if err != nil {
		return fmt.Errorf("error loading torrent file: %w", err)
	}

	mi.Announce = ""
	mi.AnnounceList = nil

	newFn := fn + ".new"

	f, err := os.Create(newFn)
	if err != nil {
		return fmt.Errorf("error creating new torrent file: %w", err)
	}

	err = mi.Write(f)
	if err != nil {
		os.Remove(newFn)
		return fmt.Errorf("error writing new torrent file: %w", err)
	}

	err = f.Close()
	if err != nil {
		os.Remove(newFn)
		return fmt.Errorf("error closing new torrent file: %w", err)
	}

	err = os.Rename(newFn, fn)
	if err != nil {
		os.Remove(newFn)
		return fmt.Errorf("error renaming torrent file %v -> %v: %w", newFn, fn, err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: torrent-remove-trackers foo.torrent ..")
		os.Exit(2)
	}
	var err error
	for _, fn := range os.Args[1:] {
		err2 := processFile(fn)
		if err2 != nil {
			err = err2
			fmt.Fprintln(os.Stderr, fn+":", err)
		}
	}
	if err != nil {
		os.Exit(1)
	}
}
