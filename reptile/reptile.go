package reptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var reQQEmail = `(\d+)@qq.com`

func FuncGetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	handleError(err, "http.Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	handleError(err, "ioutil.ReadAll")
	pageStr := string(pageBytes)
	create, err := os.Create("reptile.html")
	handleError(err,"os.Create")
	_, err = create.Write(pageBytes)
	handleError(err,"os.Write")
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)

	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}

}

func handleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
