package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"

	"log"
	"os"

	ath "./athletes"
	fn "./function"
)

type Employee struct {
	id   int
	name string
	age  int
}

var waitG sync.WaitGroup
var m sync.Mutex
var cpuUsed = 1
var maxRandomNums = 20
var counter = 0
var atomicCounter uint64

/* Enable when running Parrallelism() */
// func init() {
// 	maxCPU := runtime.NumCPU() // max is 16 on this PC

// 	cpuUsed = 4
// 	runtime.GOMAXPROCS(cpuUsed)

// 	fmt.Printf("Number of CPUs (Total=%d - Used=%d \n", maxCPU, cpuUsed)
// }

func main() {
	//main1()

	//main2()

	//main3()

	//interface1()

	//emptyInterface()

	//sort1()

	//interfaceInterface()

	//fnClosure()

	//fnCallback()

	//Recursion()

	//Fibonacci()

	//Defer()

	//PanicRecover()

	//Structs()

	//GoRoutines()

	//WaitGroup()

	//Parrallelism()

	//RaceCondition()

	//Mutex()

	//Atomic()

	//Channel_Semaphore()

	//Channel_MultipleReceivers()

	//Channel_Direction()

	//Channel_Multiplexing()

	//Channel_Buffered()

	// var gcd = Channel_GreatestCommonDivisor(5, 5)
	// fmt.Println(gcd)

	// Channel_Assignment_Pipeline1()

	// Channel_Assignment_Pipeline2()

	//Packages()

	//Error()

	//Error1()

	//Error2()

	//nonRefType()

	//refType()\

	//Json()

	//WorkingWithFiles1()

	//WorkingWithFiles2()

	//StringManipulation()

	//Reflection1()
}

/* Reflection */
func Reflection1() {
	fmt.Println(GetType1(10))
	fmt.Println(GetType1(true))
	fmt.Println(GetType1("hello"))
	fmt.Println(GetType1(10.31))
	fmt.Println(GetType1('r'))

	fmt.Println()

	bar := []string{"test"}
	fmt.Println(GetType1(bar))
	fmt.Println((reflect.TypeOf(bar)))
	fmt.Println((reflect.TypeOf(bar).Elem()))
	fmt.Println((reflect.ValueOf(bar)))
	fmt.Println((reflect.ValueOf(bar).Kind()))
	fmt.Println((reflect.ValueOf(bar).Interface()))
	fmt.Println((reflect.ValueOf(bar).String()))

	fmt.Println()

	fmt.Println((reflect.TypeOf(Reflection1)))
	fmt.Println((reflect.ValueOf(Reflection1)))
	fmt.Println((reflect.ValueOf(Reflection1).Kind()))
	fmt.Println((reflect.ValueOf(Reflection1).String()))
	fmt.Println((reflect.ValueOf(Reflection1).Type()))

	fmt.Println()

	fmt.Println((reflect.TypeOf(nil)))
	fmt.Println((reflect.ValueOf(nil)))
	fmt.Println((reflect.ValueOf(nil).Kind()))
	fmt.Println((reflect.ValueOf(nil).String()))

	fmt.Println()

	fmt.Println(GetType2(10))
	fmt.Println(GetType2(true))
	fmt.Println(GetType2("hello"))
	fmt.Println(GetType2(10.31))
	fmt.Println(GetType2('r'))
	fmt.Println(GetType2(Reflection1))

	fmt.Println()

	name := runtime.FuncForPC(reflect.ValueOf(Reflection1).Pointer()).Name()
	fmt.Println("Name of function : " + name)
}

func GetType1(t interface{}) string {
	result := ""
	switch t := t.(type) {
	case int:
		result = strconv.Itoa(t) + "/int"

	case float64:
		result = strconv.FormatFloat(t, 'f', 6, 64) + "/float64"

	case bool:
		if t {
			result = "true/bool"
		} else {
			result = "false/bool"
		}

	case string:
		result = t + "/string"

	case rune:
		result = string(t) + "/rune"

	default:
		result = fmt.Sprintf("%v/unknown", t)
	}

	return result
}

func GetType2(value interface{}) string {

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"

	case reflect.Int, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Uint, reflect.Uint8:
		return strconv.FormatUint(v.Uint(), 10)

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 6, 64)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.String:
		return strconv.Quote(v.String())

	case reflect.Func:
		return v.Type().String() + " value"

	default:
		return v.Type().String() + " value"
	}
}

/* String Manipulation*/
func StringManipulation() {
	fmt.Println(strings.Compare("at", "amber")) //1
	fmt.Println(strings.Compare("amber", "at")) //-1
	fmt.Println(strings.Compare("at", "at"))    //0

	fmt.Println(strings.Contains("Any", "ny")) //true
	fmt.Println(strings.Count("asdasd", "sd")) //2

	fmt.Println(strings.Fields(" foi wwe baz bax foi"))
	fmt.Println(strings.HasPrefix("Golang", "Go"))
	fmt.Println(strings.HasSuffix("Golang", "lang"))

	fmt.Println(strings.Index("booklet", "let"))
	fmt.Println(strings.LastIndex("boola", "la"))
	fmt.Println(strings.Repeat("ssad", 5))

	shiftOne := func(r rune) rune {
		return r + 1
	}

	fmt.Println(strings.Map(shiftOne, "abcde 12345"))

	fmt.Println(strings.Replace("asdasdasdasd", "asd", "bb", 10))
	fmt.Println(strings.Split("sss,ddd,ttt,eee,qqq,aaa", ","))
	fmt.Println(strings.Trim(" asdasdasd ", " "))
}

/* Working with files */
func WorkingWithFiles2() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		io.Copy(os.Stdout, f)
	}

	wd, err1 := os.Getwd()
	if err1 != nil {
		fmt.Println("Error: ", err1)
	} else {
		fmt.Println(wd)
	}
}

type team []string

func WorkingWithFiles1() {
	const filename = "names.txt"
	const newFilename = "new.txt"

	players := team{"_a", "_b", "_c", "_d", "_e", "_f", "_g"}

	// write
	err := ioutil.WriteFile(filename, []byte(players.toString()), 0666)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Wrote the following to file: %v \n", players.toString())
	}

	// read
	bs, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		fmt.Println("Error: ", err1)
	} else {
		s := strings.Split(string(bs), ",")
		fmt.Println(s)
	}

	// rename
	err2 := os.Rename(filename, newFilename)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	} else {
		fmt.Println("File renamed.")
	}

	// remove
	err3 := os.Remove(newFilename)
	if err3 != nil {
		fmt.Println("Error: ", err3)
	} else {
		fmt.Println("File removed.")
	}
}

func (t team) toString() string {
	return strings.Join([]string(t), ",")
}

/* Json */
type SoccerClubs struct {
	Name    string
	Country string
	Value   float32 `json:"Dollar Value(B)"`
	Players []string
}

var teams = []SoccerClubs{
	{Name: "C1", Country: "_A", Value: 3.689, Players: []string{"a", "b", "c"}},
	{Name: "C2", Country: "_B", Value: 3.689, Players: []string{"_a", "_b", "_c"}},
	{Name: "C3", Country: "_C", Value: 3.689, Players: []string{"_a_", "_b_", "_c_"}},
}

func Json() {
	data, err := json.Marshal(teams)

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Println("\n---------------------------------------------------------")

	fmt.Printf("%v\n", teams)

	fmt.Printf("%s\n", data)

	fmt.Println("\n---------------------------------------------------------")

	var names []struct{ Name string }

	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(names)
}

/* Non reference type */
const aLen = 7

func nonRefType() {

	langs := [aLen]string{1: "Go", 2: "C", 3: "Java"}

	fmt.Println(langs)

	processA1(langs)
	fmt.Println(langs) // Nothing modified

	processA2(&langs)
	fmt.Println(langs) // modified
}

func processA1(lang [aLen]string) {
	lang[1] = "GoLang"
}

func processA2(lang *[aLen]string) {
	lang[2] = "C++"
}

/* Reference type */
func refType() {
	sal := map[string]float64{
		"Joseph": 60000.00,
		"Maiku":  70000.00,
		"Alley":  80000.00,
	}

	fmt.Println(sal)
	f4(sal)
	fmt.Println(sal) // modified
}

func f4(sal map[string]float64) {
	sal["Maiku"] += 9999
}

/* Error2 - Error Interface */
type myError struct {
	errType string
	err     error
}

func (e *myError) Error() string {
	return fmt.Sprintf("[%s : %s]", e.errType, e.err)
}

func Error2() {
	_, _, err := produceError2(false)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No Error...")
	}
}

func produceError2(b bool) (int, int, error) {
	if !b {
		return 0, 0, nil
	}

	errMsg1 := errors.New("New error...")
	return 0, 0, &myError{"critical", errMsg1}
}

/* Error1 */
func Error1() {
	_, _, err1, err2 := produceError1(true)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("No Error...")
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("No Error...")
	}
}

func produceError1(b bool) (int, int, error, error) {

	if !b {
		return 0, 0, nil, nil
	}

	errMsg1 := errors.New("New error...")
	errMsg2 := fmt.Errorf("New error...")

	fmt.Printf("errMsg1 is %T \n", errMsg1)
	fmt.Printf("errMsg2 is %T \n", errMsg2)

	return 0, 0, errMsg1, errMsg2
}

/* Error */
func Error() {
	var f = "./non_existing.txt"

	_, err := os.Open(f)

	if err != nil {
		log.Println("Error: ", err)
		panic(err)
	}
}

/* Packages */
func Packages() {
	fn.Like()

	player1 := ath.Player{"Ronaldo", "Football", 43, ath.Info{"Spain", "Black"}}

	fmt.Println(player1)

	fmt.Println(*player1.ToLowerCase())
}

/* Channel Assignment (pipeline2) */
func Channel_Assignment_Pipeline2() {
	words := []string{"one", "two", "three"}
	c1 := make(chan string, 3)
	c1a := make(chan string, 3)
	c2 := make(chan string, 3)
	c3 := make(chan rune)

	go func() {
		for _, w := range words {
			c1 <- w
			c1a <- w
		}
		close(c1)
		close(c1a)
	}()

	go func() {
		for w := range c1 {
			c2 <- w + " -> " + strings.ToUpper(w)
		}
		close(c2)
	}()

	go func() {
		for w := range c1a {
			for i, c := range w {
				if i%2 == 0 {
					c3 <- unicode.ToUpper(c)
				} else {
					c3 <- unicode.ToLower(c)
				}
			}
		}
		close(c3)
	}()

	for nw := range c2 {
		fmt.Println((nw))
	}

	for c := range c3 {
		fmt.Printf("%v ", string(c))
	}
}

/* Channel Assignment (pipeline1) */
func Channel_Assignment_Pipeline1() {
	nums := []int{1, 2, 4}
	c := gen(nums...)
	out := sq(c)

	for range nums {
		fmt.Printf("%4d ", <-out)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

/* Channel Assignment */
func Channel_GreatestCommonDivisor(x int, y int) int {
	c := make(chan int)

	go func() {
		for y != 0 {
			x, y = y, (x % y)
		}
		c <- x
		close(c)
	}()

	return <-c
}

/* Channel - Buffered */
func Channel_Buffered() {
	c := make(chan int, 1)

	fmt.Printf("capacity: %d length: %d \n", cap(c), len(c))

	for i := 1; i <= 3; i++ {
		go PrintMsg(c, i)
	}

	time.Sleep(10 * time.Second)
}

func PrintMsg(c chan int, id int) {
	fmt.Printf("ooo %d is waiting for a channel space...\n", id)

	c <- id
	fmt.Printf("=== %d has a channel space\n", id)

	time.Sleep(600 * time.Millisecond)
	fmt.Printf("xxx %d has release the channel space\n", id)

	<-c
}

/* Channel - Multiplexing */
func Channel_Multiplexing() {
	stats := make(map[string]int)

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i := 0; i < 20; i++ {
		// go func() {
		// 	time.Sleep(10 * time.Second)
		// 	c1 <- "Hello from customer service #1"
		// 	switch1(c1, "Peter", stats)
		// }()

		// go func() {
		// 	time.Sleep(8 * time.Second)
		// 	c2 <- "Hello from customer service #2"
		// 	switch1(c2, "Nick", stats)
		// }()

		// go func() {
		// 	time.Sleep(8 * time.Second)
		// 	c3 <- "Hello from customer service #3"
		// 	switch1(c3, "John", stats)
		// }()

		// go func() {
		// 	time.Sleep(2 * time.Second)
		// 	switch1(nil, "Customer Waiting", stats)
		// }()

		go func() {
			time.Sleep(10 * time.Second)
			c1 <- "Hello from customer service #1"
		}()

		go func() {
			time.Sleep(8 * time.Second)
			c2 <- "Hello from customer service #2"
		}()

		go func() {
			time.Sleep(8 * time.Second)
			c3 <- "Hello from customer service #3"
		}()

		select {
		case msg1 := <-c1:
			stats["Peter"]++
			time.Sleep(time.Second)
			fmt.Println(msg1)
		case msg2 := <-c2:
			stats["Nick"]++
			time.Sleep(time.Second)
			fmt.Println(msg2)
		case msg3 := <-c3:
			stats["John"]++
			time.Sleep(time.Second)
			fmt.Println(msg3)
		default:
			stats["Customer Waiting"]++
			time.Sleep(2 * time.Second)
			fmt.Println("No customer service is available at this time!")
		}
	}
	fmt.Printf("\n***Customer Service***\n%v", stats)

	close(c1)
	close(c2)
	close(c3)
}

func switch1(c chan string, provider string, stats map[string]int) {
	select {
	case msg1 := <-c:
		stats[provider]++
		time.Sleep(time.Second)
		fmt.Println(msg1)
	default:
		stats[provider]++
		time.Sleep(2 * time.Second)
		fmt.Println("No customer service is available at this time!")
	}
}

/* Channel - Direction */
var currentBalance = 200

func Channel_Direction() {
	fmt.Println("Press Enter to stop program ...")
	rand.Seed(time.Now().UnixNano())

	var c = make(chan int)

	go credit(c)
	go debit(c)
	go balance(c)

	var input string
	fmt.Scanln(&input)
}

func credit(c chan int) { // a 'bi-directional' channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) + 1
	}
}

func debit(c chan<- int) { //a 'send-only' channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) - 10
	}
}

func balance(c <-chan int) { // a 'receive-only' channel
	for {
		num, ok := <-c

		if ok == false {
			fmt.Println("Error!")
			break
		}

		oldBalance := currentBalance
		currentBalance += num

		fmt.Printf("=> %d + (%d) = %d \n", oldBalance, num, currentBalance)
		time.Sleep(1 * time.Second)
	}
}

/* Channel - Multiple Receivers */
var wordSet = []string{"red", "green", "blue", "yellow", "white"}

func Channel_MultipleReceivers() {

	var c = make(chan string)
	var b = make(chan bool)

	go sender(c)

	receivers := []interface{}{receiver1, receiver2}

	for i := range receivers {
		go receivers[i].(func(chan bool, chan string))(b, c)
	}

	for range receivers {
		fmt.Println(<-b)
	}
}

func sender(c chan string) {
	for _, w := range wordSet {
		t := strings.ToUpper(w)
		fmt.Printf("(s) %v \n", t)
		c <- t
	}
	close(c)
}

func receiver1(b chan bool, c chan string) {
	for w := range c {
		w = strings.ToLower(w)
		fmt.Printf("(r1) %v \n", w)
		time.Sleep((2 * time.Second))
	}

	b <- true
}

func receiver2(b chan bool, c chan string) {
	for w := range c {
		w = strings.ToLower(w)
		fmt.Printf("(r2) %v \n", w)
		time.Sleep(time.Second)
	}

	b <- true
}

/* Channel - Semaphore */
var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "blue", "green", "yellow"}

func Channel_Semaphore() {

	var c = make(chan string)
	var b = make(chan bool)

	f := []interface{}{sender1, sender2}

	for i := range f {
		go f[i].(func(chan string, chan bool))(c, b)
	}

	go closeSenders(c, b)

	fmt.Println("Before getting to the 'channel for range loop'.")
	for val := range c {
		fmt.Println(val)
	}
}

func sender1(c chan string, b chan bool) {
	for _, w := range wordSet1 {
		c <- w
	}

	b <- true
}

func sender2(c chan string, b chan bool) {
	for _, w := range wordSet2 {
		c <- w
	}

	b <- true
}

func closeSenders(c chan string, b chan bool) {
	<-b
	<-b
	close(c)
}

/* Concurrency - Atomic */
func Atomic() {
	waitG.Add(2)
	go numbers5(1)
	go numbers5(2)
	waitG.Wait()

	fmt.Printf("\ncounter: %d", atomicCounter)
}

func numbers5(callID int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {

		time.Sleep(200 * time.Millisecond)

		atomic.AddUint64(&atomicCounter, 1)
		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, atomic.LoadUint64(&atomicCounter))
	}

	waitG.Done()
}

/* Concurrency - Mutex (Mutual Exclusion) */
func Mutex() {
	waitG.Add(2)
	go numbers4(1)
	go numbers4(2)
	waitG.Wait()

	fmt.Printf("\ncounter: %d", counter)
}

func numbers4(callID int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {

		time.Sleep(200 * time.Millisecond)

		m.Lock()
		counter++
		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, counter)
		m.Unlock()
	}

	waitG.Done()
}

/* Concurrency - Race condition */
func RaceCondition() {
	waitG.Add(2)
	go numbers3(1)
	go numbers3(2)
	waitG.Wait()

	fmt.Printf("\ncounter: %d", counter)
}

func numbers3(callID int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		tmpCounter := counter
		tmpCounter++

		time.Sleep(200 * time.Millisecond)
		counter = tmpCounter

		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, counter)
	}

	waitG.Done()
}

/* Concurrency - Parrallelism */
func Parrallelism() {
	start := time.Now()
	ids := []string{"#", "!", "^", "*"}

	waitG.Add(4)
	for i := range ids {
		go numbers2(ids[i])
	}
	waitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nprogram took %s. \n", elapsed)
}

func numbers2(id string) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= maxRandomNums; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s%d ", id, rand.Intn(20)+20)
	}

	waitG.Done()
}

/* Concurrency - wait group */
func WaitGroup() {
	waitG.Add(2)
	go numbers()
	go alphabets()
	waitG.Wait()

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nMain terminated")
}

/* Concurrency - Go Routines*/
func GoRoutines() {
	go numbers()
	go alphabets()

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nMain terminated")
}

func numbers() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}

	waitG.Done()
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}

	waitG.Done()
}

/* Structs */
func Structs() {
	type vehicle struct {
		name string
	}

	a := vehicle{"alpha"}

	fmt.Println(a)
	p := &a

	p.name = "romeo"

	fmt.Println(a)
	fmt.Println(*p)

	b := &vehicle{"mazerrati"}
	fmt.Println(*b)

	b1 := new(vehicle)
	fmt.Println(*b1)
	*b1 = vehicle{"ferrari"}
	fmt.Println(*b1)

	type vehicle1 struct {
		string
		int
	}

	p1 := vehicle1{"sdfsdf", 1999}
	fmt.Println(p1.int)
	fmt.Println(p1.string)

	p2 := &p1

	fmt.Println((*p2).string)
	fmt.Println((*p2).int)

	type makemodel struct {
		make  string
		model string
	}

	mm1 := makemodel{"m1", "m2"}

	type car struct {
		v  vehicle1
		mm makemodel
	}

	c1 := car{v: p1, mm: mm1}
	fmt.Println(c1)

	c4 := &c1

	c4.v.string = "adadad"
	fmt.Println(*c4)
	fmt.Println(c1)

	c4.mm.make = "b1"
	fmt.Println(*c4)
	fmt.Println(c1)

	w1 := worker{"Tesla", 100000.00}
	fmt.Println(w1)

	w2 := &w1
	fmt.Println(*w2)
	decrement(w2)
	fmt.Println(*w2)
	fmt.Println(w1)
}

type worker struct {
	name   string
	salary float64
}

func decrement(w *worker) {
	w.salary *= 1.15
}

/* Panic & Recover*/
func PanicRecover() {

	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()

	panic("Boo!!!")
}

/* Defer */
func Defer() {
	fmt.Println(square2(2))
	fmt.Println(square2(4))
	fmt.Println(square2(5))
}

func square2(n int) (result int) {
	result = n * n

	defer func() {
		if n == 2 || n == 4 {
			result += n
		} else {
			result += 1
		}
	}()

	return
}

/* Fibonacci */
func Fibonacci() {
	fmt.Printf("\n\n%d", fibo(20))
}

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

/* Recursion */
func Recursion() {
	fmt.Println("=>", factorial(1))
	fmt.Println("=>", factorial(2))
	fmt.Println("=>", factorial(3))
	fmt.Println("=>", factorial(7))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(n, " ")
	return n * factorial(n-1)
}

/* Function - Callback */
func fnCallback() {
	mult := func(nums ...int) int {
		total := 1
		for _, num := range nums {
			total *= num
		}
		return total
	}

	nums := []int{1, 2, 3, 4}
	fmt.Println(mult(nums...))
	fmt.Println(calc2(mult, nums...))

	add := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println(add(nums...))
	fmt.Println(calc2a(add, nums))

	add2 := func(nums []int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println(add2(nums))
	fmt.Println(calc2b(add2, nums))
}

func calc2(f func(...int) int, x ...int) int {
	return f(x...)
}

func calc2a(f func(...int) int, x []int) int {
	return f(x...)
}

func calc2b(f func([]int) int, x []int) int {
	return f(x)
}

/* Function - Closure */

func fnClosure() {
	addCounter, multCounter := addBy(), multBy()

	fmt.Print(addCounter(2), " ")
	fmt.Print(addCounter(3), " ")
	fmt.Print(addCounter(-1), " ")
	fmt.Println()
	fmt.Print(multCounter(3), " ")
	fmt.Print(multCounter(4), " ")
	fmt.Print(multCounter(-2), " ")
}

func addBy() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func multBy() func(int) int {
	total := 1
	return func(i int) (ret int) {
		total *= i
		ret = total
		return
	}
}

/* Interface - Interface */

func interfaceInterface() {
	p := []player{
		{"Roger Federer", "Switzerland"},
		{"Lionel Messi", "Argentina"},
		{"Michael Jordan", "USA"},
	}

	fmt.Println(p[0])
	fmt.Println(p)

	p1 := "Lionel Messi"
	p2 := "Lionel Messia"

	fmt.Println(p1, " exists: ", AnyByName(p, p1))
	fmt.Println(p2, " exists: ", AnyByName(p, p2))

	p11 := player{"Lionel Messi", "Argentina"}
	p12 := player{"Lionel Messi", "Argentinaa"}

	fmt.Println(p11, " exists: ", AnyByNameCountry(p, p11.name, p11.country))
	fmt.Println(p12, " exists: ", AnyByNameCountry(p, p12.name, p12.country))

	pp := byCountry{
		{"Roger Federer", "Switzerland"},
		{"Lionel Messi", "Argentina"},
		{"Michael Jordan", "USA"},
	}

	fmt.Println(p1, " exists: ", pp.AnyByName(p1))
	fmt.Println(p2, " exists: ", pp.AnyByName(p2))
}

type player struct {
	name    string
	country string
}

type byCountry []player

func (c byCountry) Len() int           { return len(c) }
func (c byCountry) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCountry) Less(i, j int) bool { return c[i].country < c[j].country }

func (players byCountry) AnyByName(i string) bool {
	for _, v := range players {
		if v.name == i {
			return true
		}
	}

	return false
}

func AnyByName(players []player, i string) bool {
	for _, v := range players {
		if v.name == i {
			return true
		}
	}

	return false
}

func AnyByNameCountry(players []player, i string, j string) bool {
	for _, v := range players {
		if v.name == i && v.country == j {
			return true
		}
	}

	return false
}

func (p player) String() string {
	return fmt.Sprintf("toString() - %s: %s\n", p.name, p.country)
}

/* Interface - sort */
func sort1() {
	var i = []int{2, 7, 8, 56, 345, 12, 1}
	fmt.Println(i)
	//sort.Ints(i) //sort by int
	sort.Sort(sort.IntSlice(i)) // uses an interface
	fmt.Println(i)
	sort.Sort(sort.Reverse(sort.IntSlice(i))) // reverse
	fmt.Println(i)

	var str = []string{"Susan", "Diana", "Cynthia", "Cheryl", "Zebra"}
	fmt.Println(str)
	//sort.Strings(str) //sort by string
	sort.Sort(sort.StringSlice(str)) // uses an interface
	fmt.Println(str)
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println(str)
}

/* Empty Interface */
func emptyInterface() {
	messi := football{}
	pele := football{}
	federer := tennis{}
	nadal := tennis{}

	favAthletes := []athletes{messi, pele, federer, nadal}

	for k, v := range favAthletes {
		fmt.Println(k, " - ", v)
	}

	messi = football{athlete{"Leo Messi", "Argentina"}, "Attcker"}
	federer = tennis{athlete{"Roger Federer", "Switzerland"}, true}
	playerInfo(messi)
	playerInfo(federer)

	pele = football{athlete{"Pele", "Brazil"}, "Attcker"}
	nadal = tennis{athlete{"Rafael Nadal", "Spain"}, false}

	// using empty interface directly
	favAthletes2 := []interface{}{messi, pele, federer, nadal}

	fmt.Println(favAthletes2)
}

type athlete struct {
	name    string
	country string
}

type football struct {
	athlete
	position string
}

type tennis struct {
	athlete
	rightHanded bool
}

type athletes interface{}

func playerInfo(a interface{}) {
	fmt.Println(a)
}

/* Interface */
func interface1() {
	r1 := rectangle{2, 3}
	fmt.Printf("area()=%d\n", r1.area())
	fmt.Printf("perim=()=%d\n", r1.perim())

	r2 := rectangle{20, 30}
	fmt.Printf("area()=%d\n", r2.area())
	fmt.Printf("perim=()=%d\n", r2.perim())

	s1 := square{1}
	fmt.Printf("area()=%d\n", s1.area())
	fmt.Printf("perim=()=%d\n", s1.perim())

	s2 := square{2}
	fmt.Printf("area()=%d\n", s2.area())
	fmt.Printf("perim=()=%d\n", s2.perim())

	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

	fmt.Printf("Total Area=%d\n", totalArea(&r1, &r2, &s1, &s2))
}

type rectangle struct {
	w, l int // width & length
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.l + c.w)
}

type square struct {
	s int // side
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

type shape interface {
	area() int
	perim() int
	//info(s shape)
	//totalArea(shapes ...shape) int
}

func info(s shape) {
	fmt.Printf("area()=%d perim=()=%d\n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}

/**/

func main3() {
	executionTime := 3 * time.Second
	start := time.Now()

	fmt.Printf("Program will end in about %v.\n", executionTime)
	fmt.Print("Operation in progress ...")

	s := `\|/-`
	i := 0

	for {
		printSpinningSymbol(string(s[i]))

		if time.Since(start) > executionTime {
			fmt.Println("Done")
			fmt.Println("Elapsed Time: ", time.Since(start))
			break
		}

		i = (i + 1) % 4
	}
}

var delayTime = 100 * time.Millisecond

func printSpinningSymbol(symbol string) {
	fmt.Print(symbol)
	time.Sleep(delayTime)
	fmt.Print("\b")
}

var players = make(map[string]map[string]bool)
var players2 = make(map[string]bool)

func main2() {
	addPlayer("Leo", "Messi")
	addPlayer2("Leo", "Messi")

	fmt.Println(hasPlayer("Leo", "Messi"))
	fmt.Println(hasPlayer2("Leo", "Messi"))

	fmt.Println(hasPlayer("Leo", "Messi2"))
	fmt.Println(hasPlayer2("Leo", "Messi2"))
}

func addPlayer(fName, lName string) {
	n := players[fName]
	if n == nil {
		n = make(map[string]bool)
		players[fName] = n
	}
	n[lName] = true
}

func hasPlayer(fName, lName string) bool {
	return players[fName][lName]
}

func addPlayer2(fName, lName string) {
	player := fName + " " + lName
	n := players2[player]
	if !n {
		players2[player] = true
	}
}

func hasPlayer2(fName, lName string) bool {
	player := fName + " " + lName
	return players2[player]
}

func main1() {
	x1 := "Hello World"
	fmt.Println(x1)
	fmt.Println(len(x1))

	var a = 1
	var b = 2

	fmt.Printf("%+v %+v", a, b)

	var e1 = Employee{1, "Alex", 45}
	fmt.Printf("%+v", e1)

	x := [...]int{10, 20, 30}
	for i, val := range x {
		fmt.Println(i, val)
	}

	fmt.Println("------------------------------")

	var seasonNo = 3
	switch seasonNo {
	case 3:
		fmt.Println("summer - ", seasonNo)
		fallthrough
	case 4:
		fmt.Println("winter - ", seasonNo)
		fallthrough
	default:
		fmt.Println("a new season - ", seasonNo)
	}

	fmt.Println("------------------------------")

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Println("%v", days[:5])

	fmt.Println("------------------------------")

	var f []float32
	f = append(f, 1.2)
	f = append(f, f...)
	fmt.Println(f)

	fmt.Println("------------------------------")
	var team []string
	team = append(team, goalkeeper())
	team = append(team, midfielders()...)
	team = append(team, strikers()...)
	fmt.Println(team)

	fmt.Println(cap(team), len(team))

	var opponent []string
	opponent = append(opponent, team...)

	team[0] = "Casilas"
	fmt.Println(team)
	fmt.Println(opponent)

	fmt.Println(&team)
	fmt.Println(&opponent)

	fmt.Println("------------------------------")

	sal := map[string]float64{
		"a1": 70000.0,
		"b1": 80000.50,
		"c1": 90000.0,
	}

	names := make([]string, 0, len(sal))
	for n := range sal {
		//salString := strconv.FormatInt(int64(sal[n]), 10)
		//names = append(names, n+":"+salString)
		names = append(names, n)
	}

	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}

	fmt.Println("------------------------------")

	fmt.Printf("%.2f", avg([]float32{1, 2, 3}))
}

func strikers() []string {
	return []string{"a1", "b1", "c1"}
}

func midfielders() []string {
	return []string{"a", "b", "c"}
}

func goalkeeper() string {
	return "Bufon"
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}

	return total / float32(len(scores))
}
