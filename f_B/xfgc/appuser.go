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


// Table Name : AppUser
type AppUser struct {
	AppUserId int;
	AppUserName string;
	Password string;
	IsAllowed int;
	UserTypeId int;
	CreatedBy int;
	UpdatedBy int;
	RowVersion int;
	CreatedDate string;
	UpdatedDate string;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type AppUser_list struct {
	TotalCount int
	Items []AppUser
	ErrorNo int;
	ErrorName string;
}




func AppUserDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter AppUser
var ctrAppUserInsertSuccess * common.Counter = &common.Counter{Name:"AppUserInsertSuccess"};
var ctrAppUserInsertError * common.Counter = &common.Counter{Name:"AppUserInsertError"};
var ctrAppUserUpdateSuccess * common.Counter = &common.Counter{Name:"AppUserUpdateSuccess"};
var ctrAppUserUpdateError * common.Counter = &common.Counter{Name:"AppUserUpdateError"};
var ctrAppUserSelectSuccess * common.Counter = &common.Counter{Name:"AppUserSelectSuccess"};
var ctrAppUserSelectError * common.Counter = &common.Counter{Name:"AppUserSelectError"};
var ctrAppUserSelectAllSuccess * common.Counter = &common.Counter{Name:"AppUserSelectAllSuccess"};
var ctrAppUserSelectAllError * common.Counter = &common.Counter{Name:"AppUserSelectAllError"};
var ctrAppUserSelectPageSuccess * common.Counter = &common.Counter{Name:"AppUserSelectPageSuccess"};
var ctrAppUserSelectPageError * common.Counter = &common.Counter{Name:"AppUserSelectPageError"};
var ctrAppUser_SelectByUserNameSuccess * common.Counter = &common.Counter{Name:"AppUser_SelectByUserNameSuccess"};
var ctrAppUser_SelectByUserNameError * common.Counter = &common.Counter{Name:"AppUser_SelectByUserNameError"};

func RegisterCounter_AppUser() {
	common.RegisterCounter( ctrAppUserInsertSuccess);
	common.RegisterCounter( ctrAppUserInsertError);
	common.RegisterCounter( ctrAppUserUpdateSuccess);
	common.RegisterCounter( ctrAppUserUpdateError);
	common.RegisterCounter( ctrAppUserSelectSuccess);
	common.RegisterCounter( ctrAppUserSelectError);
	common.RegisterCounter( ctrAppUserSelectAllSuccess);
	common.RegisterCounter( ctrAppUserSelectAllError);
	common.RegisterCounter( ctrAppUserSelectPageSuccess);
	common.RegisterCounter( ctrAppUserSelectPageError);
	common.RegisterCounter( ctrAppUser_SelectByUserNameSuccess);
	common.RegisterCounter( ctrAppUser_SelectByUserNameError);
}



func AppUser_Add( dbid int, ptr * AppUser) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into appuser( appusername, password, isallowed, usertypeid, createdby, updatedby, rowversion, createddate, updateddate, isactive) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING appuserid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.AppUserName, ptr.Password, ptr.IsAllowed, ptr.UserTypeId, ptr.CreatedBy, ptr.UpdatedBy, ptr.RowVersion, ptr.CreatedDate, ptr.UpdatedDate, ptr.IsActive).Scan( &ptr.AppUserId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrAppUserInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql AppUser_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrAppUserInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAppUserInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - AppUser_Add")
	}
	return false
}

func AppUser_Update( dbid int, ptr * AppUser) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update appuser set appusername=$1 , password=$2 , isallowed=$3 , usertypeid=$4 , createdby=$5 , updatedby=$6 , rowversion=$7 , isactive=$8 where appuserid=$9`
		_, err := dbConn.Db.Exec( updateStmt, ptr.AppUserName, ptr.Password, ptr.IsAllowed, ptr.UserTypeId, ptr.CreatedBy, ptr.UpdatedBy, ptr.RowVersion, ptr.IsActive, ptr.AppUserId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrAppUserUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql AppUser_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrAppUserUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAppUserUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - AppUser_Update")
	}
	return false
}

func AppUser_Get( dbid int, ptr * AppUser) * AppUser {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAppUserSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - AppUser_Get");
		return nil;
	}

	selectStmt := `select appuserid, appusername, password, isallowed, usertypeid, createdby, updatedby, rowversion, isactive from appuser where appuserid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.AppUserId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrAppUserSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql AppUser_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * AppUser = new(AppUser)

	var strAppUserName sql.NullString
	var strPassword sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.AppUserId, &strAppUserName, &strPassword, &obj.IsAllowed, &obj.UserTypeId, &obj.CreatedBy, &obj.UpdatedBy, &obj.RowVersion, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan AppUser_Get  %v+", err));
		}

		if strAppUserName.String != "" {
			obj.AppUserName = strAppUserName.String;
		}

		if strPassword.String != "" {
			obj.Password = strPassword.String;
		}

	}
	rows.Close()

	ctrAppUserSelectSuccess.Increment();
	return obj
}

func AppUser_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * AppUser) * AppUser_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAppUserSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - AppUser_GetByPaging");
		return nil;
	}

	var slist * AppUser_list = new(AppUser_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT appuserid, appusername, password, isallowed, usertypeid, createdby, updatedby, rowversion, isactive FROM appuser  where isactive=$1 ORDER BY appuserid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj AppUser

			var strAppUserName sql.NullString
			var strPassword sql.NullString

			if rerr := rows.Scan( &obj.AppUserId, &strAppUserName, &strPassword, &obj.IsAllowed, &obj.UserTypeId, &obj.CreatedBy, &obj.UpdatedBy, &obj.RowVersion, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan AppUser_GetByPaging  %v+", rerr));
			}

			if strAppUserName.String != "" {
				obj.AppUserName = strAppUserName.String;
			}

			if strPassword.String != "" {
				obj.Password = strPassword.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM appuser where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql AppUser_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []AppUser{}
		}
		ctrAppUserSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrAppUserSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql AppUser_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func AppUser_GetAll( dbid int, ptr * AppUser) * AppUser_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAppUserSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - AppUser_GetAll");
		return nil;
	}

	var slist * AppUser_list = new(AppUser_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT appuserid, appusername, password, isallowed, usertypeid, createdby, updatedby, rowversion, isactive FROM appuser  ORDER BY appuserid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj AppUser

			var strAppUserName sql.NullString
			var strPassword sql.NullString

			if rerr := rows.Scan( &obj.AppUserId, &strAppUserName, &strPassword, &obj.IsAllowed, &obj.UserTypeId, &obj.CreatedBy, &obj.UpdatedBy, &obj.RowVersion, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan AppUser_GetAll  %v+", rerr));
			}

			if strAppUserName.String != "" {
				obj.AppUserName = strAppUserName.String;
			}

			if strPassword.String != "" {
				obj.Password = strPassword.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []AppUser{};
		}
		ctrAppUserSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrAppUserSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql AppUser_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func AppUser_Delete( dbid int, ptr * AppUser) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update appuser set IsActive=10 where appuserid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.AppUserId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrAppUserUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql AppUser_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrAppUserUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAppUserUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - AppUser_Update")
	}
	return false
}

func AppUser_SelectByUserName( dbid int, ptr * AppUser) * AppUser_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAppUser_SelectByUserNameError.Increment();
		common.DbLogError("DBConnection Allocation failed - AppUser_SelectByUserName");
		return nil;
	}

	selectStmt := `select appuserid, appusername, password, isallowed, usertypeid, createdby, updatedby, rowversion, isactive from appuser where appusername=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.AppUserName)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrAppUser_SelectByUserNameError.Increment();
		common.DbLogError( fmt.Sprintf( "error in AppUser_SelectByUserName-sql  %s %v+", selectStmt, err));
		return nil;
	}

	var slist * AppUser_list = new(AppUser_list)
	slist.TotalCount = 0

	for rows.Next() {
		var obj AppUser

		var strAppUserName sql.NullString
		var strPassword sql.NullString

		err = rows.Scan( &obj.AppUserId, &strAppUserName, &strPassword, &obj.IsAllowed, &obj.UserTypeId, &obj.CreatedBy, &obj.UpdatedBy, &obj.RowVersion, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in AppUser_SelectByUserName-Scan  %v+", err));
		}

		if strAppUserName.String != "" {
			obj.AppUserName = strAppUserName.String;
		}

		if strPassword.String != "" {
			obj.Password = strPassword.String;
		}

		slist.Items = append( slist.Items, obj)
	}
	rows.Close()

	slist.TotalCount = len(slist.Items)
	ctrAppUser_SelectByUserNameSuccess.Increment();
	return slist
}



// Table Name : AppUser

func addAppUser(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_AppUser AppUser

	if err := c.BindJSON(&new_AppUser); err != nil {
		common.AppLogError("BindJSON failed for addAppUser");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_AppUser.AppUserName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"AppUserName is required"})
		return
	}

	if new_AppUser.Password == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"Password is required"})
		return
	}


	//set custom feild values
	new_AppUser.CreatedBy = user.UserId;
	new_AppUser.UpdatedBy = user.UserId;
	new_AppUser.RowVersion = 1;
	new_AppUser.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05");
	new_AppUser.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05");
	new_AppUser.IsActive = 1;

	AppUser_Add( 0, &new_AppUser)
	c.IndentedJSON( http.StatusOK, gin.H{"AppUserId": new_AppUser.AppUserId, "ErrorNo": new_AppUser.ErrorNo, "ErrorName": new_AppUser.ErrorName});
}


func updateAppUser(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AppUser AppUser

	if err := c.BindJSON( &obj_AppUser); err != nil {
		common.AppLogError("BindJSON failed for updateAppUser");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_AppUser.AppUserName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"AppUserName is required"})
		return
	}

	if obj_AppUser.Password == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"Password is required"})
		return
	}


	//set custom feild values
	obj_AppUser.CreatedBy = user.UserId;
	obj_AppUser.UpdatedBy = user.UserId;
	obj_AppUser.RowVersion = 1;
	obj_AppUser.IsActive = 1;


	sts := AppUser_Update( 0, &obj_AppUser)
	c.IndentedJSON( http.StatusOK, gin.H{"AppUserId": obj_AppUser.AppUserId, "Status": sts, "ErrorNo": obj_AppUser.ErrorNo, "ErrorName": obj_AppUser.ErrorName});
}


func getAppUser(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AppUser AppUser;
	id, ok := c.GetQuery("id");
	if ok {
		obj_AppUser.AppUserId,_ = strconv.Atoi(id);
	}

	if obj_AppUser.AppUserId == 0 {
		common.AppLogError("id not found in query string - getAppUser");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := AppUser_Get( 0, &obj_AppUser);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteAppUser(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AppUser AppUser;
	id, ok := c.GetQuery("id");
	if ok {
		obj_AppUser.AppUserId,_ = strconv.Atoi(id);
	}

	if obj_AppUser.AppUserId == 0 {
		common.AppLogError("id not found in query string - getAppUser");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := AppUser_Delete( 0, &obj_AppUser);
	c.IndentedJSON( http.StatusOK, res)
}


func getAppUserByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getAppUserList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj AppUser
	obj.IsActive = 1;

	list := AppUser_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllAppUser(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj AppUser

	list := AppUser_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAppUserByUserName(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AppUser AppUser

	if err := c.BindJSON( &obj_AppUser); err != nil {
		common.AppLogError("BindJSON failed for updateAppUser");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	list := AppUser_SelectByUserName( 0, &obj_AppUser)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : AppUser
func InitModule_AppUser( router * gin.Engine) {
	router.POST("/appuser/add", addAppUser);
	router.POST("/appuser/upd", updateAppUser);
	router.POST("/appuser/getappuserbypaging", getAppUserByPaging);
	router.GET("/appuser/getall", getAllAppUser);
	router.GET("/appuser/getbyid", getAppUser);
	router.GET("/appuser/delete", deleteAppUser);
}




