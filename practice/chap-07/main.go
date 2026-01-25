package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Student struct {
	Roll    int
	Name    string
	Subject string
}

type Pipeline struct {
	Name    string
	inpChan chan Student
	proChan chan Student

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

/* ---------------- Consumer ---------------- */
func (p *Pipeline) Consumer() {
	defer func() {
		p.wg.Done()
		close(p.proChan) // Close input channel [Consumer is the owner]
	}()

	i := 0
	for std := range p.inpChan {
		i++
		fmt.Println(i, ":- data recieved...", std)
		p.proChan <- std
		time.Sleep(2 * time.Second)
	}
	fmt.Println("All data read successfully now returning...")
}

/* ---------------- Consumer ---------------- */
func (p *Pipeline) Processor() {
	defer p.wg.Done()

	i := 0
	for std := range p.proChan {
		i++
		fmt.Println(i, ":- data processing...", std)
		time.Sleep(5 * time.Second)
	}
	fmt.Println("All data proccessed successfuly...")
}

/* ---------------- Generator ---------------- */
func (p *Pipeline) Generator() {
	defer func() {
		p.wg.Done()
		close(p.inpChan) // Close input channel only on Generator
	}()

	i := 0
	for {
		select {
		case p.inpChan <- Student{Roll: 100 + i, Name: fmt.Sprintf("Abhi-%d", i), Subject: fmt.Sprintf("Python-%d", i)}:
			i++
			fmt.Printf("%d :- data sent\n", i)
			time.Sleep(100 * time.Millisecond)

		case <-p.ctx.Done():
			fmt.Println("Context is cancelled now returning...")
			return
		}
	}
}

/* ---------------- Lifecycle ---------------- */
func NewPipeline(name string) *Pipeline {
	ctx, cancel := context.WithCancel(context.Background())
	inpChan := make(chan Student, 1024)
	proChan := make(chan Student, 1024)

	return &Pipeline{
		Name:    name,
		inpChan: inpChan,
		proChan: proChan,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (p *Pipeline) Start() {
	fmt.Println("Start called...")
	p.wg.Add(3)
	go p.Consumer()
	go p.Processor()

	go p.Generator()
}

func (p *Pipeline) Stop() {
	fmt.Println("Stop called...")
	p.cancel()
	p.wg.Wait()
}

func main() {
	// Stoping process in windows
	// tasklist|findstr go.exe
	// taskkill /pid 22960 /F

	fmt.Println("---------- Main Start ----------")
	pipeline := NewPipeline("Master Pipeline")

	// Start the pipeline
	pipeline.Start()

	// Handle termination
	sigs := make(chan os.Signal, 1)
	fmt.Println("Waiting for interrupt...")

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Stopping pipeline...")
	pipeline.Stop()
	fmt.Println("---------- Main End ----------")
}
