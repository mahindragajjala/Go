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


// Table Name : IMSIRange
type IMSIRange struct {
	IMSIRangeId int;
	IMSIBegin int;
	IMSIEnd int;
	SubK string;
	OPc string;
	MCC int;
	MNC int;
	PLMN string;
	IsActive int;
	CreatedDate string;
	CreatedUserId int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type IMSIRange_list struct {
	TotalCount int
	Items []IMSIRange
	ErrorNo int;
	ErrorName string;
}




func IMSIRangeDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter IMSIRange
var ctrIMSIRangeInsertSuccess * common.Counter = &common.Counter{Name:"IMSIRangeInsertSuccess"};
var ctrIMSIRangeInsertError * common.Counter = &common.Counter{Name:"IMSIRangeInsertError"};
var ctrIMSIRangeUpdateSuccess * common.Counter = &common.Counter{Name:"IMSIRangeUpdateSuccess"};
var ctrIMSIRangeUpdateError * common.Counter = &common.Counter{Name:"IMSIRangeUpdateError"};
var ctrIMSIRangeSelectSuccess * common.Counter = &common.Counter{Name:"IMSIRangeSelectSuccess"};
var ctrIMSIRangeSelectError * common.Counter = &common.Counter{Name:"IMSIRangeSelectError"};
var ctrIMSIRangeSelectAllSuccess * common.Counter = &common.Counter{Name:"IMSIRangeSelectAllSuccess"};
var ctrIMSIRangeSelectAllError * common.Counter = &common.Counter{Name:"IMSIRangeSelectAllError"};
var ctrIMSIRangeSelectPageSuccess * common.Counter = &common.Counter{Name:"IMSIRangeSelectPageSuccess"};
var ctrIMSIRangeSelectPageError * common.Counter = &common.Counter{Name:"IMSIRangeSelectPageError"};

func RegisterCounter_IMSIRange() {
	common.RegisterCounter( ctrIMSIRangeInsertSuccess);
	common.RegisterCounter( ctrIMSIRangeInsertError);
	common.RegisterCounter( ctrIMSIRangeUpdateSuccess);
	common.RegisterCounter( ctrIMSIRangeUpdateError);
	common.RegisterCounter( ctrIMSIRangeSelectSuccess);
	common.RegisterCounter( ctrIMSIRangeSelectError);
	common.RegisterCounter( ctrIMSIRangeSelectAllSuccess);
	common.RegisterCounter( ctrIMSIRangeSelectAllError);
	common.RegisterCounter( ctrIMSIRangeSelectPageSuccess);
	common.RegisterCounter( ctrIMSIRangeSelectPageError);
}



func IMSIRange_Add( dbid int, ptr * IMSIRange) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into imsirange( imsibegin, imsiend, subk, opc, mcc, mnc, plmn, isactive, createddate, createduserid) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING imsirangeid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.IMSIBegin, ptr.IMSIEnd, ptr.SubK, ptr.OPc, ptr.MCC, ptr.MNC, ptr.PLMN, ptr.IsActive, ptr.CreatedDate, ptr.CreatedUserId).Scan( &ptr.IMSIRangeId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrIMSIRangeInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql IMSIRange_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrIMSIRangeInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - IMSIRange_Add")
	}
	return false
}

func IMSIRange_Update( dbid int, ptr * IMSIRange) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update imsirange set imsibegin=$1 , imsiend=$2 , subk=$3 , opc=$4 , mcc=$5 , mnc=$6 , plmn=$7 , isactive=$8 , createduserid=$9 where imsirangeid=$10`
		_, err := dbConn.Db.Exec( updateStmt, ptr.IMSIBegin, ptr.IMSIEnd, ptr.SubK, ptr.OPc, ptr.MCC, ptr.MNC, ptr.PLMN, ptr.IsActive, ptr.CreatedUserId, ptr.IMSIRangeId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrIMSIRangeUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql IMSIRange_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrIMSIRangeUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - IMSIRange_Update")
	}
	return false
}

func IMSIRange_Get( dbid int, ptr * IMSIRange) * IMSIRange {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - IMSIRange_Get");
		return nil;
	}

	selectStmt := `select imsirangeid, imsibegin, imsiend, subk, opc, mcc, mnc, plmn, isactive, createduserid from imsirange where imsirangeid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.IMSIRangeId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrIMSIRangeSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql IMSIRange_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * IMSIRange = new(IMSIRange)

	var strSubK sql.NullString
	var strOPc sql.NullString
	var strPLMN sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.IMSIRangeId, &obj.IMSIBegin, &obj.IMSIEnd, &strSubK, &strOPc, &obj.MCC, &obj.MNC, &strPLMN, &obj.IsActive, &obj.CreatedUserId)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan IMSIRange_Get  %v+", err));
		}

		if strSubK.String != "" {
			obj.SubK = strSubK.String;
		}

		if strOPc.String != "" {
			obj.OPc = strOPc.String;
		}

		if strPLMN.String != "" {
			obj.PLMN = strPLMN.String;
		}

	}
	rows.Close()

	ctrIMSIRangeSelectSuccess.Increment();
	return obj
}

func IMSIRange_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * IMSIRange) * IMSIRange_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - IMSIRange_GetByPaging");
		return nil;
	}

	var slist * IMSIRange_list = new(IMSIRange_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT imsirangeid, imsibegin, imsiend, subk, opc, mcc, mnc, plmn, isactive, createduserid FROM imsirange  where isactive=$1 ORDER BY imsirangeid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj IMSIRange

			var strSubK sql.NullString
			var strOPc sql.NullString
			var strPLMN sql.NullString

			if rerr := rows.Scan( &obj.IMSIRangeId, &obj.IMSIBegin, &obj.IMSIEnd, &strSubK, &strOPc, &obj.MCC, &obj.MNC, &strPLMN, &obj.IsActive, &obj.CreatedUserId); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan IMSIRange_GetByPaging  %v+", rerr));
			}

			if strSubK.String != "" {
				obj.SubK = strSubK.String;
			}

			if strOPc.String != "" {
				obj.OPc = strOPc.String;
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM imsirange where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql IMSIRange_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []IMSIRange{}
		}
		ctrIMSIRangeSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrIMSIRangeSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql IMSIRange_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func IMSIRange_GetAll( dbid int, ptr * IMSIRange) * IMSIRange_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - IMSIRange_GetAll");
		return nil;
	}

	var slist * IMSIRange_list = new(IMSIRange_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT imsirangeid, imsibegin, imsiend, subk, opc, mcc, mnc, plmn, isactive, createduserid FROM imsirange  ORDER BY imsirangeid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj IMSIRange

			var strSubK sql.NullString
			var strOPc sql.NullString
			var strPLMN sql.NullString

			if rerr := rows.Scan( &obj.IMSIRangeId, &obj.IMSIBegin, &obj.IMSIEnd, &strSubK, &strOPc, &obj.MCC, &obj.MNC, &strPLMN, &obj.IsActive, &obj.CreatedUserId); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan IMSIRange_GetAll  %v+", rerr));
			}

			if strSubK.String != "" {
				obj.SubK = strSubK.String;
			}

			if strOPc.String != "" {
				obj.OPc = strOPc.String;
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []IMSIRange{};
		}
		ctrIMSIRangeSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrIMSIRangeSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql IMSIRange_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func IMSIRange_Delete( dbid int, ptr * IMSIRange) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update imsirange set IsActive=10 where imsirangeid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.IMSIRangeId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrIMSIRangeUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql IMSIRange_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrIMSIRangeUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrIMSIRangeUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - IMSIRange_Update")
	}
	return false
}



// Table Name : IMSIRange

func addIMSIRange(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_IMSIRange IMSIRange

	if err := c.BindJSON(&new_IMSIRange); err != nil {
		common.AppLogError("BindJSON failed for addIMSIRange");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_IMSIRange.SubK == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"SubK is required"})
		return
	}

	if new_IMSIRange.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"OPc is required"})
		return
	}

	if new_IMSIRange.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}


	//set custom feild values
	new_IMSIRange.IsActive = 1;
	new_IMSIRange.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05");
	new_IMSIRange.CreatedUserId = user.UserId;

	IMSIRange_Add( 0, &new_IMSIRange)
	c.IndentedJSON( http.StatusOK, gin.H{"IMSIRangeId": new_IMSIRange.IMSIRangeId, "ErrorNo": new_IMSIRange.ErrorNo, "ErrorName": new_IMSIRange.ErrorName});
}


func updateIMSIRange(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_IMSIRange IMSIRange

	if err := c.BindJSON( &obj_IMSIRange); err != nil {
		common.AppLogError("BindJSON failed for updateIMSIRange");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_IMSIRange.SubK == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"SubK is required"})
		return
	}

	if obj_IMSIRange.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"OPc is required"})
		return
	}

	if obj_IMSIRange.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}


	//set custom feild values
	obj_IMSIRange.IsActive = 1;
	obj_IMSIRange.CreatedUserId = user.UserId;


	sts := IMSIRange_Update( 0, &obj_IMSIRange)
	c.IndentedJSON( http.StatusOK, gin.H{"IMSIRangeId": obj_IMSIRange.IMSIRangeId, "Status": sts, "ErrorNo": obj_IMSIRange.ErrorNo, "ErrorName": obj_IMSIRange.ErrorName});
}


func getIMSIRange(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_IMSIRange IMSIRange;
	id, ok := c.GetQuery("id");
	if ok {
		obj_IMSIRange.IMSIRangeId,_ = strconv.Atoi(id);
	}

	if obj_IMSIRange.IMSIRangeId == 0 {
		common.AppLogError("id not found in query string - getIMSIRange");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := IMSIRange_Get( 0, &obj_IMSIRange);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteIMSIRange(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_IMSIRange IMSIRange;
	id, ok := c.GetQuery("id");
	if ok {
		obj_IMSIRange.IMSIRangeId,_ = strconv.Atoi(id);
	}

	if obj_IMSIRange.IMSIRangeId == 0 {
		common.AppLogError("id not found in query string - getIMSIRange");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := IMSIRange_Delete( 0, &obj_IMSIRange);
	c.IndentedJSON( http.StatusOK, res)
}


func getIMSIRangeByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getIMSIRangeList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj IMSIRange
	obj.IsActive = 1;

	list := IMSIRange_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllIMSIRange(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj IMSIRange

	list := IMSIRange_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : IMSIRange
func InitModule_IMSIRange( router * gin.Engine) {
	router.POST("/imsirange/add", addIMSIRange);
	router.POST("/imsirange/upd", updateIMSIRange);
	router.POST("/imsirange/getimsirangebypaging", getIMSIRangeByPaging);
	router.GET("/imsirange/getall", getAllIMSIRange);
	router.GET("/imsirange/getbyid", getIMSIRange);
	router.GET("/imsirange/delete", deleteIMSIRange);
}




