//多种输出样式的进度条 参考：https://github.com/briandowns/spinner

package main

import (
    "github.com/briandowns/spinner"
    "time"
)

func main() {
    s := spinner.New(spinner.CharSets[70], 100*time.Millisecond)  // Build our new spinner
    s.Start()                                                    // Start the spinner
    time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
    s.Stop()
}
