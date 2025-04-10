package common

import (
	"fmt"
	"sync"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type Session struct {
	userMap map[string]*UserInfo;
	userMapLock sync.Mutex;
}

var userSessions * Session;

func init(){
	userSessions = new(Session);
	userSessions.userMap = make(map[string]*UserInfo);
}

func Dummy() {
	fmt.Printf("")
}

func Session_SetUserSession( key string, userInfo * UserInfo)  {
	userSessions.userMapLock.Lock();
	defer userSessions.userMapLock.Unlock();
	userSessions.userMap[key] = userInfo;
}

func Session_GetUserSession( key string) * UserInfo {
	userSessions.userMapLock.Lock();
	defer userSessions.userMapLock.Unlock();
	
	userInfo, exists := userSessions.userMap[key];
	
	if exists {
		return userInfo;
	}
	
	return nil;
}

func Session_DeleteUserSession( key string) {
	userSessions.userMapLock.Lock();
	defer userSessions.userMapLock.Unlock();

	delete( userSessions.userMap, key)
}





func Session_GetID(c *gin.Context) string {
	session := sessions.Default(c)
	return session.ID()
}


func Session_Set( c *gin.Context, key interface{}, val interface{}) {
	session := sessions.Default(c)
	session.Set( key, val)
}


func Session_Save( c *gin.Context) {
	session := sessions.Default(c)
	session.Save()
}


func Session_Get( c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	return session.Get( key);
}


func Session_DeleteKey( c *gin.Context, key interface{}) {
	session := sessions.Default(c)
	session.Delete(key)
}

func Session_Delete( c *gin.Context) {
	session := sessions.Default(c)
	session.Clear();
}


func Session_SaveUser( c * gin.Context, user * UserInfo) {
	session := sessions.Default(c)
	session.Set( "User", user);
}

func Session_GetUser( c * gin.Context) * UserInfo {

	// AuthorizationId := c.Request.Header["Authorization"][0];
	
	// if AuthorizationId != "" {
	
		// userInfo :=  Session_GetUserSession( AuthorizationId);
		
		// if userInfo != nil {
			// //fmt.Printf("Found Session From Session Context AuthorizationId=%s\n", AuthorizationId);
			// return userInfo;
		// } 
		// // else {
			// // fmt.Printf("NotFound Session From Session Context AuthorizationId=%s\n", AuthorizationId);
		// // }
	// } 
	// // else {
		// // fmt.Printf("AuthorizationId=%s Is NIL\n", AuthorizationId);
	// // }
	

	session := sessions.Default(c)
	obj := session.Get( "User")
	
	if obj != nil {
		return obj.(*UserInfo)
	}		
	
	ui := new(UserInfo);
	ui.UserId = 1;
	ui.OrganizationId = 1;
	
	return ui
}
