package flag

import (
	"flag"
	"fmt"
	"time"
)

func Standard()  {
	var period = flag.Duration("period", 1 * time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println()
}
