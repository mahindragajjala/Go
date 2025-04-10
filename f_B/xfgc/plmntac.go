package xfgc

import (
	"fmt"
	"time"
	"strconv"
	"net/http"

	"xius.com/xfgc/common"
	"xius.com/xfgc/postgres"

	"github.com/gin-gonic/gin"
)


// Table Name : PlmnTac
type PlmnTac struct {
	PlmnTacId int;
	MCC int;
	MNC int;
	TAC int;
	Allowed int;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type PlmnTac_list struct {
	TotalCount int
	Items []PlmnTac
	ErrorNo int;
	ErrorName string;
}




func PlmnTacDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter PlmnTac
var ctrPlmnTacInsertSuccess * common.Counter = &common.Counter{Name:"PlmnTacInsertSuccess"};
var ctrPlmnTacInsertError * common.Counter = &common.Counter{Name:"PlmnTacInsertError"};
var ctrPlmnTacUpdateSuccess * common.Counter = &common.Counter{Name:"PlmnTacUpdateSuccess"};
var ctrPlmnTacUpdateError * common.Counter = &common.Counter{Name:"PlmnTacUpdateError"};
var ctrPlmnTacSelectSuccess * common.Counter = &common.Counter{Name:"PlmnTacSelectSuccess"};
var ctrPlmnTacSelectError * common.Counter = &common.Counter{Name:"PlmnTacSelectError"};
var ctrPlmnTacSelectAllSuccess * common.Counter = &common.Counter{Name:"PlmnTacSelectAllSuccess"};
var ctrPlmnTacSelectAllError * common.Counter = &common.Counter{Name:"PlmnTacSelectAllError"};
var ctrPlmnTacSelectPageSuccess * common.Counter = &common.Counter{Name:"PlmnTacSelectPageSuccess"};
var ctrPlmnTacSelectPageError * common.Counter = &common.Counter{Name:"PlmnTacSelectPageError"};

func RegisterCounter_PlmnTac() {
	common.RegisterCounter( ctrPlmnTacInsertSuccess);
	common.RegisterCounter( ctrPlmnTacInsertError);
	common.RegisterCounter( ctrPlmnTacUpdateSuccess);
	common.RegisterCounter( ctrPlmnTacUpdateError);
	common.RegisterCounter( ctrPlmnTacSelectSuccess);
	common.RegisterCounter( ctrPlmnTacSelectError);
	common.RegisterCounter( ctrPlmnTacSelectAllSuccess);
	common.RegisterCounter( ctrPlmnTacSelectAllError);
	common.RegisterCounter( ctrPlmnTacSelectPageSuccess);
	common.RegisterCounter( ctrPlmnTacSelectPageError);
}



func PlmnTac_Add( dbid int, ptr * PlmnTac) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into plmntac( mcc, mnc, tac, allowed, isactive) values( $1, $2, $3, $4, $5) RETURNING plmntacid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.MCC, ptr.MNC, ptr.TAC, ptr.Allowed, ptr.IsActive).Scan( &ptr.PlmnTacId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrPlmnTacInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql PlmnTac_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrPlmnTacInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrPlmnTacInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - PlmnTac_Add")
	}
	return false
}

func PlmnTac_Update( dbid int, ptr * PlmnTac) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update plmntac set mcc=$1 , mnc=$2 , tac=$3 , allowed=$4 , isactive=$5 where plmntacid=$6`
		_, err := dbConn.Db.Exec( updateStmt, ptr.MCC, ptr.MNC, ptr.TAC, ptr.Allowed, ptr.IsActive, ptr.PlmnTacId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrPlmnTacUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql PlmnTac_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrPlmnTacUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrPlmnTacUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - PlmnTac_Update")
	}
	return false
}

func PlmnTac_Get( dbid int, ptr * PlmnTac) * PlmnTac {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrPlmnTacSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - PlmnTac_Get");
		return nil;
	}

	selectStmt := `select plmntacid, mcc, mnc, tac, allowed, isactive from plmntac where plmntacid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.PlmnTacId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrPlmnTacSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql PlmnTac_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * PlmnTac = new(PlmnTac)


	for rows.Next() {
		err = rows.Scan( &obj.PlmnTacId, &obj.MCC, &obj.MNC, &obj.TAC, &obj.Allowed, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan PlmnTac_Get  %v+", err));
		}

	}
	rows.Close()

	ctrPlmnTacSelectSuccess.Increment();
	return obj
}

func PlmnTac_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * PlmnTac) * PlmnTac_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrPlmnTacSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - PlmnTac_GetByPaging");
		return nil;
	}

	var slist * PlmnTac_list = new(PlmnTac_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT plmntacid, mcc, mnc, tac, allowed, isactive FROM plmntac  where isactive=$1 ORDER BY plmntacid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj PlmnTac


			if rerr := rows.Scan( &obj.PlmnTacId, &obj.MCC, &obj.MNC, &obj.TAC, &obj.Allowed, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan PlmnTac_GetByPaging  %v+", rerr));
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM plmntac where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql PlmnTac_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []PlmnTac{}
		}
		ctrPlmnTacSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrPlmnTacSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql PlmnTac_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func PlmnTac_GetAll( dbid int, ptr * PlmnTac) * PlmnTac_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrPlmnTacSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - PlmnTac_GetAll");
		return nil;
	}

	var slist * PlmnTac_list = new(PlmnTac_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT plmntacid, mcc, mnc, tac, allowed, isactive FROM plmntac  ORDER BY plmntacid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj PlmnTac


			if rerr := rows.Scan( &obj.PlmnTacId, &obj.MCC, &obj.MNC, &obj.TAC, &obj.Allowed, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan PlmnTac_GetAll  %v+", rerr));
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []PlmnTac{};
		}
		ctrPlmnTacSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrPlmnTacSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql PlmnTac_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func PlmnTac_Delete( dbid int, ptr * PlmnTac) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update plmntac set IsActive=10 where plmntacid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.PlmnTacId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrPlmnTacUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql PlmnTac_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrPlmnTacUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrPlmnTacUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - PlmnTac_Update")
	}
	return false
}



// Table Name : PlmnTac

func addPlmnTac(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_PlmnTac PlmnTac

	if err := c.BindJSON(&new_PlmnTac); err != nil {
		common.AppLogError("BindJSON failed for addPlmnTac");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation


	//set custom feild values
	new_PlmnTac.IsActive = 1;

	PlmnTac_Add( 0, &new_PlmnTac)
	c.IndentedJSON( http.StatusOK, gin.H{"PlmnTacId": new_PlmnTac.PlmnTacId, "ErrorNo": new_PlmnTac.ErrorNo, "ErrorName": new_PlmnTac.ErrorName});
}


func updatePlmnTac(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_PlmnTac PlmnTac

	if err := c.BindJSON( &obj_PlmnTac); err != nil {
		common.AppLogError("BindJSON failed for updatePlmnTac");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation


	//set custom feild values
	obj_PlmnTac.IsActive = 1;


	sts := PlmnTac_Update( 0, &obj_PlmnTac)
	c.IndentedJSON( http.StatusOK, gin.H{"PlmnTacId": obj_PlmnTac.PlmnTacId, "Status": sts, "ErrorNo": obj_PlmnTac.ErrorNo, "ErrorName": obj_PlmnTac.ErrorName});
}


func getPlmnTac(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_PlmnTac PlmnTac;
	id, ok := c.GetQuery("id");
	if ok {
		obj_PlmnTac.PlmnTacId,_ = strconv.Atoi(id);
	}

	if obj_PlmnTac.PlmnTacId == 0 {
		common.AppLogError("id not found in query string - getPlmnTac");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := PlmnTac_Get( 0, &obj_PlmnTac);
	c.IndentedJSON( http.StatusOK, res)
}


func deletePlmnTac(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_PlmnTac PlmnTac;
	id, ok := c.GetQuery("id");
	if ok {
		obj_PlmnTac.PlmnTacId,_ = strconv.Atoi(id);
	}

	if obj_PlmnTac.PlmnTacId == 0 {
		common.AppLogError("id not found in query string - getPlmnTac");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := PlmnTac_Delete( 0, &obj_PlmnTac);
	c.IndentedJSON( http.StatusOK, res)
}


func getPlmnTacByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getPlmnTacList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj PlmnTac
	obj.IsActive = 1;

	list := PlmnTac_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllPlmnTac(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj PlmnTac

	list := PlmnTac_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : PlmnTac
func InitModule_PlmnTac( router * gin.Engine) {
	router.POST("/plmntac/add", addPlmnTac);
	router.POST("/plmntac/upd", updatePlmnTac);
	router.POST("/plmntac/getplmntacbypaging", getPlmnTacByPaging);
	router.GET("/plmntac/getall", getAllPlmnTac);
	router.GET("/plmntac/getbyid", getPlmnTac);
	router.GET("/plmntac/delete", deletePlmnTac);
}




