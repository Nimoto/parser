package main

import (
	"os"
	"os/signal"
	"parser/configs"
	"parser/internal/filemanager"
	"parser/internal/reader"
	"sync"
	"time"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}

	fm := filemanager.InitFileManager()
	defer fm.Write(cfg.Source.Dir, cfg.Source.Metafile, fm.GetUpdateDate(cfg.Source.Dir, cfg.Source.File).Format(time.RFC3339Nano))

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)

	go func() {
		<-signals
		fm.Write(cfg.Source.Dir, cfg.Source.Metafile, fm.GetUpdateDate(cfg.Source.Dir, cfg.Source.File).Format(time.RFC3339Nano))
		os.Exit(0)
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	r := reader.InitDataReader(fm)
	mr := reader.InitMetaReader(fm)

	t1, _ := mr.Read(cfg.Source.Dir, cfg.Source.Metafile)

	var t2 time.Time

	go func() {
		for true {
			t2 = t1
			t1 = fm.GetUpdateDate(cfg.Source.Dir, cfg.Source.File)

			if t1.IsZero() || t1.After(t2) {
				d := r.Read(cfg.Source.Dir, cfg.Source.File)
				for _, line := range d {
					println(line.GetId())
				}
			}

			time.Sleep(time.Duration(cfg.Period) * time.Second)
		}
	}()

	wg.Wait()
}
