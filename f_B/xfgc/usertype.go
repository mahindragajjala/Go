package xfgc

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"database/sql"

	"xius.com/xfgc/common"
	"xius.com/xfgc/postgres"

	"github.com/gin-gonic/gin"
)


// Table Name : UserType
type UserType struct {
	UserTypeId int;
	UserTypeName string;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type UserType_list struct {
	TotalCount int
	Items []UserType
	ErrorNo int;
	ErrorName string;
}




func UserTypeDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter UserType
var ctrUserTypeInsertSuccess * common.Counter = &common.Counter{Name:"UserTypeInsertSuccess"};
var ctrUserTypeInsertError * common.Counter = &common.Counter{Name:"UserTypeInsertError"};
var ctrUserTypeUpdateSuccess * common.Counter = &common.Counter{Name:"UserTypeUpdateSuccess"};
var ctrUserTypeUpdateError * common.Counter = &common.Counter{Name:"UserTypeUpdateError"};
var ctrUserTypeSelectSuccess * common.Counter = &common.Counter{Name:"UserTypeSelectSuccess"};
var ctrUserTypeSelectError * common.Counter = &common.Counter{Name:"UserTypeSelectError"};
var ctrUserTypeSelectAllSuccess * common.Counter = &common.Counter{Name:"UserTypeSelectAllSuccess"};
var ctrUserTypeSelectAllError * common.Counter = &common.Counter{Name:"UserTypeSelectAllError"};
var ctrUserTypeSelectPageSuccess * common.Counter = &common.Counter{Name:"UserTypeSelectPageSuccess"};
var ctrUserTypeSelectPageError * common.Counter = &common.Counter{Name:"UserTypeSelectPageError"};

func RegisterCounter_UserType() {
	common.RegisterCounter( ctrUserTypeInsertSuccess);
	common.RegisterCounter( ctrUserTypeInsertError);
	common.RegisterCounter( ctrUserTypeUpdateSuccess);
	common.RegisterCounter( ctrUserTypeUpdateError);
	common.RegisterCounter( ctrUserTypeSelectSuccess);
	common.RegisterCounter( ctrUserTypeSelectError);
	common.RegisterCounter( ctrUserTypeSelectAllSuccess);
	common.RegisterCounter( ctrUserTypeSelectAllError);
	common.RegisterCounter( ctrUserTypeSelectPageSuccess);
	common.RegisterCounter( ctrUserTypeSelectPageError);
}



func UserType_Add( dbid int, ptr * UserType) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into usertype( usertypename, isactive) values( $1, $2) RETURNING usertypeid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.UserTypeName, ptr.IsActive).Scan( &ptr.UserTypeId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrUserTypeInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql UserType_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrUserTypeInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrUserTypeInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - UserType_Add")
	}
	return false
}

func UserType_Update( dbid int, ptr * UserType) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update usertype set usertypename=$1 , isactive=$2 where usertypeid=$3`
		_, err := dbConn.Db.Exec( updateStmt, ptr.UserTypeName, ptr.IsActive, ptr.UserTypeId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrUserTypeUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql UserType_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrUserTypeUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrUserTypeUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - UserType_Update")
	}
	return false
}

func UserType_Get( dbid int, ptr * UserType) * UserType {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrUserTypeSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - UserType_Get");
		return nil;
	}

	selectStmt := `select usertypeid, usertypename, isactive from usertype where usertypeid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.UserTypeId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrUserTypeSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql UserType_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * UserType = new(UserType)

	var strUserTypeName sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.UserTypeId, &strUserTypeName, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan UserType_Get  %v+", err));
		}

		if strUserTypeName.String != "" {
			obj.UserTypeName = strUserTypeName.String;
		}

	}
	rows.Close()

	ctrUserTypeSelectSuccess.Increment();
	return obj
}

func UserType_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * UserType) * UserType_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrUserTypeSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - UserType_GetByPaging");
		return nil;
	}

	var slist * UserType_list = new(UserType_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT usertypeid, usertypename, isactive FROM usertype  where isactive=$1 ORDER BY usertypeid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj UserType

			var strUserTypeName sql.NullString

			if rerr := rows.Scan( &obj.UserTypeId, &strUserTypeName, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan UserType_GetByPaging  %v+", rerr));
			}

			if strUserTypeName.String != "" {
				obj.UserTypeName = strUserTypeName.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM usertype where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql UserType_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []UserType{}
		}
		ctrUserTypeSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrUserTypeSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql UserType_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func UserType_GetAll( dbid int, ptr * UserType) * UserType_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrUserTypeSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - UserType_GetAll");
		return nil;
	}

	var slist * UserType_list = new(UserType_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT usertypeid, usertypename, isactive FROM usertype  ORDER BY usertypeid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj UserType

			var strUserTypeName sql.NullString

			if rerr := rows.Scan( &obj.UserTypeId, &strUserTypeName, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan UserType_GetAll  %v+", rerr));
			}

			if strUserTypeName.String != "" {
				obj.UserTypeName = strUserTypeName.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []UserType{};
		}
		ctrUserTypeSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrUserTypeSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql UserType_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func UserType_Delete( dbid int, ptr * UserType) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update usertype set IsActive=10 where usertypeid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.UserTypeId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrUserTypeUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql UserType_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrUserTypeUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrUserTypeUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - UserType_Update")
	}
	return false
}



// Table Name : UserType

func addUserType(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_UserType UserType

	if err := c.BindJSON(&new_UserType); err != nil {
		common.AppLogError("BindJSON failed for addUserType");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_UserType.UserTypeName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"UserTypeName is required"})
		return
	}


	//set custom feild values
	new_UserType.IsActive = 1;

	UserType_Add( 0, &new_UserType)
	c.IndentedJSON( http.StatusOK, gin.H{"UserTypeId": new_UserType.UserTypeId, "ErrorNo": new_UserType.ErrorNo, "ErrorName": new_UserType.ErrorName});
}


func updateUserType(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_UserType UserType

	if err := c.BindJSON( &obj_UserType); err != nil {
		common.AppLogError("BindJSON failed for updateUserType");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_UserType.UserTypeName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"UserTypeName is required"})
		return
	}


	//set custom feild values
	obj_UserType.IsActive = 1;


	sts := UserType_Update( 0, &obj_UserType)
	c.IndentedJSON( http.StatusOK, gin.H{"UserTypeId": obj_UserType.UserTypeId, "Status": sts, "ErrorNo": obj_UserType.ErrorNo, "ErrorName": obj_UserType.ErrorName});
}


func getUserType(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_UserType UserType;
	id, ok := c.GetQuery("id");
	if ok {
		obj_UserType.UserTypeId,_ = strconv.Atoi(id);
	}

	if obj_UserType.UserTypeId == 0 {
		common.AppLogError("id not found in query string - getUserType");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := UserType_Get( 0, &obj_UserType);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteUserType(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_UserType UserType;
	id, ok := c.GetQuery("id");
	if ok {
		obj_UserType.UserTypeId,_ = strconv.Atoi(id);
	}

	if obj_UserType.UserTypeId == 0 {
		common.AppLogError("id not found in query string - getUserType");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := UserType_Delete( 0, &obj_UserType);
	c.IndentedJSON( http.StatusOK, res)
}


func getUserTypeByPaging(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var pageInfo common.FPageInfo

	if perr := c.BindJSON( &pageInfo); perr != nil {
		common.AppLogError("BindJSON failed for getUserTypeList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj UserType
	obj.IsActive = 1;

	list := UserType_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllUserType(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj UserType

	list := UserType_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : UserType
func InitModule_UserType( router * gin.Engine) {
	router.POST("/usertype/add", addUserType);
	router.POST("/usertype/upd", updateUserType);
	router.POST("/usertype/getusertypebypaging", getUserTypeByPaging);
	router.GET("/usertype/getall", getAllUserType);
	router.GET("/usertype/getbyid", getUserType);
	router.GET("/usertype/delete", deleteUserType);
}




