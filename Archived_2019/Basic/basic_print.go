
package main

import (
	"fmt"
)

type trinitycore struct {
	id int
	version string
	gamebuild int

	os struct {
		os_family string
		compiler string
	}


}

func main() {
	var tc trinitycore
	tc.id = 1
	tc.version = "7.3.5"
	tc.gamebuild = 26792
	tc.os.os_family = "Ubuntu18"
	tc.os.compiler = "GCC 7.4.0"

	fmt.Println(tc)
	fmt.Println("Game Version: \t ", tc.version)
	fmt.Println("Game Build: \t", tc.gamebuild)
	fmt.Println("Server ENV: %d %d",tc.os.os_family, tc.os.compiler)
}

/*
Output:

```
{1 7.3.5 26792 {Ubuntu18 GCC 7.4.0}}
Game Version:     7.3.5
Game Build:      26792
Server ENV: %d %d Ubuntu18 GCC 7.4.0
```


 */