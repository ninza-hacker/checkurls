package main
import ("fmt"
       "bufio"
	"os"
	"sync"
	"net/http"
	//"github.com/fatih/color"
	//"github.com/gookit/color"
	 )



var wg sync.WaitGroup



func main(){

work := make(chan string)
go func(){


		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			wg.Add(1)
			defer wg.Done()
			url:=scanner.Text()
			work <- url
		}
		 close(work)
	}()


	for i := 0; i < 200 ; i++ {
            wg.Add(1)
            //fmt.Println("hello")
            go dowork(work)
        }
        wg.Wait()
    }


func dowork(work chan string){
//c := color.New(color.FgCyan).Add(color.Underline)
defer wg.Done()
	for val2 := range work {
		resp, err := http.Get(val2)
		if err != nil {

			continue

}


fmt.Println(http.StatusText(resp.StatusCode),"  " , resp.StatusCode," "val2)	

	}
	
}
