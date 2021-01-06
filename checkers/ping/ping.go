package ping

// import (
// 	"fmt"

// 	ping "github.com/go-ping/ping"
// )

// func main() {
// 	pinger, err := ping.NewPinger("www.google.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	pinger.SetPrivileged(true)
// 	pinger.Count = 3
// 	err = pinger.Run() // Blocks until finished.
// 	if err != nil {
// 		panic(err)
// 	}
// 	stats := pinger.Statistics() // get send/receive/rtt stats
// 	fmt.Println(stats)
// }
func check() (bool, error) {
	return true, nil
}
