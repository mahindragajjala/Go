package common

import (
  "fmt"
  "sync"
  "time"
  //"github.com/gin-contrib/sessions"
  //"github.com/gin-gonic/gin"
)

type CacheItem struct {
	Key string
	Data interface{};
	Loader func() interface{};
	HoldSeconds int;
	CurrentTick int;
	DataLock sync.Mutex;
	LazyReload bool;
}

type Cache struct {
	cacheMap map[string]*CacheItem;
	cacheMapLock sync.Mutex;	
}

var appCache * Cache;


func init(){
	appCache = new(Cache);
	appCache.cacheMap = make(map[string]*CacheItem);
}


func Cache_SetCache( key string, val * CacheItem) {
	appCache.cacheMapLock.Lock();
	defer appCache.cacheMapLock.Unlock();
	appCache.cacheMap[key] = val;
}

func Cache_GetCache( key string) * CacheItem {
	appCache.cacheMapLock.Lock();
	defer appCache.cacheMapLock.Unlock();
	return appCache.cacheMap[key];
}

func Cache_LazyReload( key string) {
	ci := Cache_GetCache( key)
	
	if ci != nil {
		ci.LazyReload = true;
	}
}


func Cache_DeleteCache( key string){
	appCache.cacheMapLock.Lock();
	defer appCache.cacheMapLock.Unlock();
	
	delete( appCache.cacheMap, key);
}

func Cache_LoadCache( cacheItem * CacheItem) {
	data					:= cacheItem.Loader();
	
	cacheItem.DataLock.Lock();
	cacheItem.Data 			= data;
	cacheItem.DataLock.Unlock();
	
	cacheItem.CurrentTick 	= 0;
	cacheItem.LazyReload 	= false
	//fmt.Printf("Key=%s\nData=%+v\n", cacheItem.Key, cacheItem.Data);
}

func Cache_Register( key string, Loader func() interface{}, HoldSeconds int) {
	
	var cacheItem * CacheItem = new(CacheItem)

	cacheItem.Key 			= key;
	cacheItem.Loader 		= Loader;
	cacheItem.HoldSeconds 	= HoldSeconds;
	cacheItem.CurrentTick 	= 0;
	
	Cache_LoadCache( cacheItem);
	Cache_SetCache( key, cacheItem);
	//fmt.Printf("Cache_Register() completed\n");
}


func CacheSyncRoutine() {

	for {		
		for _, cacheItem := range appCache.cacheMap {
			if cacheItem.HoldSeconds > 0 {
				cacheItem.CurrentTick++;
				
				
				//fmt.Printf("HoldSeconds=%d CurrentTick=%d \n", cacheItem.HoldSeconds, cacheItem.CurrentTick);
				
				if cacheItem.HoldSeconds == cacheItem.CurrentTick || cacheItem.LazyReload == true {
					Cache_LoadCache( cacheItem);
				}
			}
		}
		time.Sleep( time.Second * 1)
	}
}


func InitalizeCache() {
	fmt.Printf("cache.InitalizeCache()\n");
	go CacheSyncRoutine();
}













