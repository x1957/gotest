package main

import (
	"github.com/x1957/chaoyang/storage"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

const BASE_URL string = "http://xqcx.bjchyedu.cn/search2016.jsp?t=p&n="

// ignore error
func httpGet(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func parseAndInsert(context []byte, store *storage.Store) {
	var jsonBody map[string]interface{}
	json.Unmarshal(context, &jsonBody)
	obj := jsonBody["data"].([]interface{})
	for _, data := range obj {
		store.Insert(data.(string))
	}
}

func getList(store *storage.Store, str string, ch chan bool) {
	url := BASE_URL + str
	context := httpGet(url)
	parseAndInsert(context, store)
	ch <- true // Done
}
func main() {
	store := storage.NewStorage()
	s := "abcdefghijklmnopqrstuvwzyx"
	chans := make([]chan bool, 26)
	for i := 0; i < 26; i++ {
		chans[i] = make(chan bool)
		go getList(store, s[i: i + 1], chans[i])
	}
	for _, chanTobeDone := range chans {
		<- chanTobeDone
	}
	store.GetList()
	return
}
