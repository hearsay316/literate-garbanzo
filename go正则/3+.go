package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str:="3.14 123 123.321 .68 haha 1.0 abc 7. ab.3 66.6 123."
	str := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Title</title>
	</head>
	<body>
	<div>
		wswswswsw
	</div>
	<script>

	</script>

	<script>
		console.log(a);
	console.log(b);
	</script>
	</body>
	</html>`
	ret := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	all := ret.FindAllStringSubmatch(str, -1)

	fmt.Println(all[0][1])
}
