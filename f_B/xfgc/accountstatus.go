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


// Table Name : AccountStatus
type AccountStatus struct {
	AccountStatusId int;
	AccountStatusName string;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type AccountStatus_list struct {
	TotalCount int
	Items []AccountStatus
	ErrorNo int;
	ErrorName string;
}




func AccountStatusDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter AccountStatus
var ctrAccountStatusInsertSuccess * common.Counter = &common.Counter{Name:"AccountStatusInsertSuccess"};
var ctrAccountStatusInsertError * common.Counter = &common.Counter{Name:"AccountStatusInsertError"};
var ctrAccountStatusUpdateSuccess * common.Counter = &common.Counter{Name:"AccountStatusUpdateSuccess"};
var ctrAccountStatusUpdateError * common.Counter = &common.Counter{Name:"AccountStatusUpdateError"};
var ctrAccountStatusSelectSuccess * common.Counter = &common.Counter{Name:"AccountStatusSelectSuccess"};
var ctrAccountStatusSelectError * common.Counter = &common.Counter{Name:"AccountStatusSelectError"};
var ctrAccountStatusSelectAllSuccess * common.Counter = &common.Counter{Name:"AccountStatusSelectAllSuccess"};
var ctrAccountStatusSelectAllError * common.Counter = &common.Counter{Name:"AccountStatusSelectAllError"};
var ctrAccountStatusSelectPageSuccess * common.Counter = &common.Counter{Name:"AccountStatusSelectPageSuccess"};
var ctrAccountStatusSelectPageError * common.Counter = &common.Counter{Name:"AccountStatusSelectPageError"};

func RegisterCounter_AccountStatus() {
	common.RegisterCounter( ctrAccountStatusInsertSuccess);
	common.RegisterCounter( ctrAccountStatusInsertError);
	common.RegisterCounter( ctrAccountStatusUpdateSuccess);
	common.RegisterCounter( ctrAccountStatusUpdateError);
	common.RegisterCounter( ctrAccountStatusSelectSuccess);
	common.RegisterCounter( ctrAccountStatusSelectError);
	common.RegisterCounter( ctrAccountStatusSelectAllSuccess);
	common.RegisterCounter( ctrAccountStatusSelectAllError);
	common.RegisterCounter( ctrAccountStatusSelectPageSuccess);
	common.RegisterCounter( ctrAccountStatusSelectPageError);
}



func AccountStatus_Add( dbid int, ptr * AccountStatus) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into accountstatus( accountstatusname, isactive) values( $1, $2) RETURNING accountstatusid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.AccountStatusName, ptr.IsActive).Scan( &ptr.AccountStatusId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrAccountStatusInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql AccountStatus_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrAccountStatusInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAccountStatusInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - AccountStatus_Add")
	}
	return false
}

func AccountStatus_Update( dbid int, ptr * AccountStatus) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update accountstatus set accountstatusname=$1 , isactive=$2 where accountstatusid=$3`
		_, err := dbConn.Db.Exec( updateStmt, ptr.AccountStatusName, ptr.IsActive, ptr.AccountStatusId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrAccountStatusUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql AccountStatus_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrAccountStatusUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAccountStatusUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - AccountStatus_Update")
	}
	return false
}

func AccountStatus_Get( dbid int, ptr * AccountStatus) * AccountStatus {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAccountStatusSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - AccountStatus_Get");
		return nil;
	}

	selectStmt := `select accountstatusid, accountstatusname, isactive from accountstatus where accountstatusid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.AccountStatusId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrAccountStatusSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql AccountStatus_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * AccountStatus = new(AccountStatus)

	var strAccountStatusName sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.AccountStatusId, &strAccountStatusName, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan AccountStatus_Get  %v+", err));
		}

		if strAccountStatusName.String != "" {
			obj.AccountStatusName = strAccountStatusName.String;
		}

	}
	rows.Close()

	ctrAccountStatusSelectSuccess.Increment();
	return obj
}

func AccountStatus_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * AccountStatus) * AccountStatus_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAccountStatusSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - AccountStatus_GetByPaging");
		return nil;
	}

	var slist * AccountStatus_list = new(AccountStatus_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT accountstatusid, accountstatusname, isactive FROM accountstatus  where isactive=$1 ORDER BY accountstatusid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj AccountStatus

			var strAccountStatusName sql.NullString

			if rerr := rows.Scan( &obj.AccountStatusId, &strAccountStatusName, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan AccountStatus_GetByPaging  %v+", rerr));
			}

			if strAccountStatusName.String != "" {
				obj.AccountStatusName = strAccountStatusName.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM accountstatus where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql AccountStatus_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []AccountStatus{}
		}
		ctrAccountStatusSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrAccountStatusSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql AccountStatus_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func AccountStatus_GetAll( dbid int, ptr * AccountStatus) * AccountStatus_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrAccountStatusSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - AccountStatus_GetAll");
		return nil;
	}

	var slist * AccountStatus_list = new(AccountStatus_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT accountstatusid, accountstatusname, isactive FROM accountstatus  ORDER BY accountstatusid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj AccountStatus

			var strAccountStatusName sql.NullString

			if rerr := rows.Scan( &obj.AccountStatusId, &strAccountStatusName, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan AccountStatus_GetAll  %v+", rerr));
			}

			if strAccountStatusName.String != "" {
				obj.AccountStatusName = strAccountStatusName.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []AccountStatus{};
		}
		ctrAccountStatusSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrAccountStatusSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql AccountStatus_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func AccountStatus_Delete( dbid int, ptr * AccountStatus) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update accountstatus set IsActive=10 where accountstatusid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.AccountStatusId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrAccountStatusUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql AccountStatus_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrAccountStatusUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrAccountStatusUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - AccountStatus_Update")
	}
	return false
}



// Table Name : AccountStatus

func addAccountStatus(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_AccountStatus AccountStatus

	if err := c.BindJSON(&new_AccountStatus); err != nil {
		common.AppLogError("BindJSON failed for addAccountStatus");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_AccountStatus.AccountStatusName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"AccountStatusName is required"})
		return
	}


	//set custom feild values
	new_AccountStatus.IsActive = 1;

	AccountStatus_Add( 0, &new_AccountStatus)
	c.IndentedJSON( http.StatusOK, gin.H{"AccountStatusId": new_AccountStatus.AccountStatusId, "ErrorNo": new_AccountStatus.ErrorNo, "ErrorName": new_AccountStatus.ErrorName});
}


func updateAccountStatus(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AccountStatus AccountStatus

	if err := c.BindJSON( &obj_AccountStatus); err != nil {
		common.AppLogError("BindJSON failed for updateAccountStatus");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_AccountStatus.AccountStatusName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"AccountStatusName is required"})
		return
	}


	//set custom feild values
	obj_AccountStatus.IsActive = 1;


	sts := AccountStatus_Update( 0, &obj_AccountStatus)
	c.IndentedJSON( http.StatusOK, gin.H{"AccountStatusId": obj_AccountStatus.AccountStatusId, "Status": sts, "ErrorNo": obj_AccountStatus.ErrorNo, "ErrorName": obj_AccountStatus.ErrorName});
}


func getAccountStatus(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AccountStatus AccountStatus;
	id, ok := c.GetQuery("id");
	if ok {
		obj_AccountStatus.AccountStatusId,_ = strconv.Atoi(id);
	}

	if obj_AccountStatus.AccountStatusId == 0 {
		common.AppLogError("id not found in query string - getAccountStatus");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := AccountStatus_Get( 0, &obj_AccountStatus);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteAccountStatus(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_AccountStatus AccountStatus;
	id, ok := c.GetQuery("id");
	if ok {
		obj_AccountStatus.AccountStatusId,_ = strconv.Atoi(id);
	}

	if obj_AccountStatus.AccountStatusId == 0 {
		common.AppLogError("id not found in query string - getAccountStatus");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := AccountStatus_Delete( 0, &obj_AccountStatus);
	c.IndentedJSON( http.StatusOK, res)
}


func getAccountStatusByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getAccountStatusList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj AccountStatus
	obj.IsActive = 1;

	list := AccountStatus_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllAccountStatus(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj AccountStatus

	list := AccountStatus_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : AccountStatus
func InitModule_AccountStatus( router * gin.Engine) {
	router.POST("/accountstatus/add", addAccountStatus);
	router.POST("/accountstatus/upd", updateAccountStatus);
	router.POST("/accountstatus/getaccountstatusbypaging", getAccountStatusByPaging);
	router.GET("/accountstatus/getall", getAllAccountStatus);
	router.GET("/accountstatus/getbyid", getAccountStatus);
	router.GET("/accountstatus/delete", deleteAccountStatus);
}




