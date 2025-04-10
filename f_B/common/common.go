package common

import (
	"sync"
	"strconv"
)


type FPageInformation struct {
	Search string;
	SortColumn string;
	SortOrder string;
	CurrentPage int;
	PageSize int;
}

type FPageInfo struct {
	Info FPageInformation;
}

type Counter struct {
	mu sync.Mutex;
	Name string;
	Count int64;
	LastCount int64;	
}

func (counter * Counter) Increment() {
	counter.mu.Lock();
	counter.Count++;
	counter.mu.Unlock();
}

func IntConv(arg interface{}) int {
	switch x := arg.(type) {
    case int:
        return x
    case string:
        {
			xy,_ := strconv.Atoi(x)
			return xy;
		}
    }
	return 0;
}

func (counter * Counter) GetCounts() (int64,int64) {
	
	var current int64 = 0;
	var total int64 = 0;
	
	current = counter.Count - counter.LastCount
	total = counter.Count
	
	counter.LastCount = counter.Count;
	
	return current, total;
}


var counters []*Counter;

func RegisterCounter( counter * Counter) {
	counters = append( counters, counter);
}

func PerfLogger() {

	
}
