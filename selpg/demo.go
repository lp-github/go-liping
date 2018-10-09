package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	flag "github.com/spf13/pflag"
)

var progname string
var sa selpg_args

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int  /* default value, can be overriden by "-l number" on command line */
	page_type   bool /* 'l' for lines-delimited, 'f' for form-feed-delimited */
	/* default is 'l' */
	print_dest string
}

func Init() {
	flag.IntVarP(&(sa.start_page), "startPage", "s", -1, "-s int or -startPage int")
	flag.IntVarP(&(sa.end_page), "endPage", "e", -1, "-e int or -endPage int")
	flag.IntVarP(&(sa.page_len), "length", "l", 72, "-length int or -l int")
	flag.BoolVarP(&(sa.page_type), "formfeed", "f", false, "-formfeed bool or -f bool")
	flag.StringVarP(&(sa.print_dest), "dest", "d", "", "-dest string or -d string")
}

func main() {
	Init()
	flag.Parse()
	process_args()
	process_input()
}

/*

 */
func process_args() {
	if flag.NArg() > 1 {
		panic_and_print_usage()
	}
	if flag.NArg() == 1 {
		sa.in_filename = flag.Args()[0]
	}
	if sa.start_page == -1 || sa.end_page == -1 {
		fmt.Fprint(os.Stderr, "you must set your start page and end page with positive integers")
		panic_and_print_usage()
	}
	if sa.start_page > sa.end_page {
		fmt.Println("start large than end page")
		os.Exit(1)
	}
}
func usage() string {
	return progname + " -sNumber -eNumber [-lNumber] [-f] [-dDestination] [output file name]"

}

func process_input() {
	var file *os.File
	var cmd *exec.Cmd
	var cmdin io.WriteCloser
	defer file.Close()
	if sa.in_filename == "" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(sa.in_filename)
		if err == nil {
			fmt.Println("file open successfully")
		} else {
			fmt.Println(err)
		}
	}
	reader := bufio.NewReader(file)
	var writer *bufio.Writer
	if sa.print_dest == "" {
		writer = bufio.NewWriter(os.Stdout)
	} else {
		var err error
		cmd = exec.Command("lp", "-d"+sa.print_dest)
		//cmd = exec.Command("cat")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println("failed open cmd ", cmd.Args, "stdinpipe")
		}
	}
	chanb := make(chan byte)
	chanel := make(chan []byte, 10)
	quit := make(chan int)
	defer close(chanb)
	defer close(chanel)
	defer close(quit)
	//---------------here reading----------------
	go func() {
		//----line determined page------
		if sa.page_type == false {
			pgctr := 1
			lctr := 0
			for {
				line, _, err := reader.ReadLine()
				if err == nil {
					lctr++
					if lctr >= sa.page_len { //time to finish one page
						pgctr++
						lctr = 0
					}
					if pgctr < sa.start_page { //havn't reach the start page
						continue
					}
					if pgctr > sa.end_page { //finish work
						chanel <- line
						time.Sleep(time.Millisecond * 100)
						quit <- 0
						break
					}
					chanel <- line //work
				} else if err == io.EOF {
					time.Sleep(time.Millisecond * 100)
					quit <- 0
					break
				} else {
					time.Sleep(100 * time.Millisecond)
					quit <- 0
					break
				}

			}
		} else { //-------form feeded page-----------
			pgctr := 1
			for {
				b, err := reader.ReadByte()
				if err == nil { //no error
					if b == '\f' {
						pgctr++
					}
					if pgctr < sa.start_page { //not start
						continue
					} else if pgctr > sa.end_page { //already finish
						time.Sleep(100 * time.Millisecond)
						quit <- 0
						break
					}
					//work
					chanb <- b
				} else if err == io.EOF { //last letter of text
					time.Sleep(100 * time.Millisecond)
					quit <- 0
					break
				} else { //unknown problem
					time.Sleep(time.Millisecond * 100)
					quit <- 0
					break
				}
			}
		}
	}()
	//---------------here writing
	func() {
		for {
			select {
			case line := <-chanel:
				if sa.print_dest == "" {
					writer.Write(line)
					writer.WriteByte('\n')
					writer.Flush()
				} else { //write to the command
					fmt.Fprint(cmdin, (string)(line)+"\n")
				}
			case <-quit:
				if sa.print_dest != "" {
					cmd.Start()
					cmd.Wait()
				}
				break
			case bt := <-chanb:
				if sa.print_dest == "" {
					writer.WriteByte(bt)
					writer.Flush()
				} else {
					fmt.Fprint(cmdin, string(bt))
				}
			}
		}

	}()
}
func panic_and_print_usage() {
	panic(errors.New(usage()))
}
