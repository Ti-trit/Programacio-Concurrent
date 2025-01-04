
package main 
//SOLUCIÓ D'ALGORISME LR DEL SÒPAR DE FILOSOFS
import(
	"fmt"
	"runtime"
	"time"
)

const(
	Procs  = 4
	FILOSOFS = 5
    EatCount     = 100
)

type Empty struct{}
type Fork chan Empty
type Forks [FILOSOFS]Fork

func think (){
	time.Sleep(50 * time.Millisecond)
}

func right (i int ) int {
	return (i+1)%FILOSOFS
}

func eat (id int){
	fmt.Printf("%d start eat\n", id)
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("%d end eat\n", id)
}

func relearForks (id int, forks Forks){
	forks[id] <-Empty{} //equivalent  a "release()"
	forks[right(id)]<-Empty{}
}

func pickForks (id int, forks Forks){
	if id<right(id){
		<- forks[id]
		<- forks[right(id)]
	}else{
		<- forks[right(id)]
		<- forks[id]
	}
}

func filosof (id int, done chan Empty, forks Forks){

	for i:= 0; i<EatCount; i++{
		think()
		pickForks(id, forks)
		eat(id)
		relearForks(id,forks)
	}
	done <- Empty{}
}

func fork(ch chan Empty){
	for{
		ch<- Empty{}  // fork  es marca com disponible enviant un valor Empty al canal
		<- ch         // Després espera que sigui agafada per un filòsof
	}
}
func main(){
	runtime.GOMAXPROCS(Procs) //nº de processos a executar en //
	var forks Forks
	done:= make (chan Empty, 1)
	for i:= range forks{
		forks[i]= make (chan Empty)
		go fork(forks[i])
	}
	for i:= 0; i<FILOSOFS; i++{
		go filosof(i, done, forks)
	}
	

	for i:=0 ; i<FILOSOFS; i++{
		<- done
	}
	fmt.Printf("Simulació acabada")
}