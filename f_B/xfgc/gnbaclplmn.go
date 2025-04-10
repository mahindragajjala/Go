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


// Table Name : GnbAclPlmn
type GnbAclPlmn struct {
	Id int;
	MCC int;
	MNC int;
	Block int;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type GnbAclPlmn_list struct {
	TotalCount int
	Items []GnbAclPlmn
	ErrorNo int;
	ErrorName string;
}




func GnbAclPlmnDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter GnbAclPlmn
var ctrGnbAclPlmnInsertSuccess * common.Counter = &common.Counter{Name:"GnbAclPlmnInsertSuccess"};
var ctrGnbAclPlmnInsertError * common.Counter = &common.Counter{Name:"GnbAclPlmnInsertError"};
var ctrGnbAclPlmnUpdateSuccess * common.Counter = &common.Counter{Name:"GnbAclPlmnUpdateSuccess"};
var ctrGnbAclPlmnUpdateError * common.Counter = &common.Counter{Name:"GnbAclPlmnUpdateError"};
var ctrGnbAclPlmnSelectSuccess * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectSuccess"};
var ctrGnbAclPlmnSelectError * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectError"};
var ctrGnbAclPlmnSelectAllSuccess * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectAllSuccess"};
var ctrGnbAclPlmnSelectAllError * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectAllError"};
var ctrGnbAclPlmnSelectPageSuccess * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectPageSuccess"};
var ctrGnbAclPlmnSelectPageError * common.Counter = &common.Counter{Name:"GnbAclPlmnSelectPageError"};

func RegisterCounter_GnbAclPlmn() {
	common.RegisterCounter( ctrGnbAclPlmnInsertSuccess);
	common.RegisterCounter( ctrGnbAclPlmnInsertError);
	common.RegisterCounter( ctrGnbAclPlmnUpdateSuccess);
	common.RegisterCounter( ctrGnbAclPlmnUpdateError);
	common.RegisterCounter( ctrGnbAclPlmnSelectSuccess);
	common.RegisterCounter( ctrGnbAclPlmnSelectError);
	common.RegisterCounter( ctrGnbAclPlmnSelectAllSuccess);
	common.RegisterCounter( ctrGnbAclPlmnSelectAllError);
	common.RegisterCounter( ctrGnbAclPlmnSelectPageSuccess);
	common.RegisterCounter( ctrGnbAclPlmnSelectPageError);
}



func GnbAclPlmn_Add( dbid int, ptr * GnbAclPlmn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into gnbaclplmn( mcc, mnc, block, isactive) values( $1, $2, $3, $4) RETURNING id`
		err := dbConn.Db.QueryRow( insertStmt, ptr.MCC, ptr.MNC, ptr.Block, ptr.IsActive).Scan( &ptr.Id)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrGnbAclPlmnInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql GnbAclPlmn_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrGnbAclPlmnInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbAclPlmn_Add")
	}
	return false
}

func GnbAclPlmn_Update( dbid int, ptr * GnbAclPlmn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update gnbaclplmn set mcc=$1 , mnc=$2 , block=$3 , isactive=$4 where id=$5`
		_, err := dbConn.Db.Exec( updateStmt, ptr.MCC, ptr.MNC, ptr.Block, ptr.IsActive, ptr.Id)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrGnbAclPlmnUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql GnbAclPlmn_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrGnbAclPlmnUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - GnbAclPlmn_Update")
	}
	return false
}

func GnbAclPlmn_Get( dbid int, ptr * GnbAclPlmn) * GnbAclPlmn {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbAclPlmn_Get");
		return nil;
	}

	selectStmt := `select id, mcc, mnc, block, isactive from gnbaclplmn where id=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.Id)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrGnbAclPlmnSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql GnbAclPlmn_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * GnbAclPlmn = new(GnbAclPlmn)


	for rows.Next() {
		err = rows.Scan( &obj.Id, &obj.MCC, &obj.MNC, &obj.Block, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan GnbAclPlmn_Get  %v+", err));
		}

	}
	rows.Close()

	ctrGnbAclPlmnSelectSuccess.Increment();
	return obj
}

func GnbAclPlmn_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * GnbAclPlmn) * GnbAclPlmn_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbAclPlmn_GetByPaging");
		return nil;
	}

	var slist * GnbAclPlmn_list = new(GnbAclPlmn_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT id, mcc, mnc, block, isactive FROM gnbaclplmn  where isactive=$1 ORDER BY id OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj GnbAclPlmn


			if rerr := rows.Scan( &obj.Id, &obj.MCC, &obj.MNC, &obj.Block, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan GnbAclPlmn_GetByPaging  %v+", rerr));
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM gnbaclplmn where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql GnbAclPlmn_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []GnbAclPlmn{}
		}
		ctrGnbAclPlmnSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrGnbAclPlmnSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql GnbAclPlmn_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func GnbAclPlmn_GetAll( dbid int, ptr * GnbAclPlmn) * GnbAclPlmn_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbAclPlmn_GetAll");
		return nil;
	}

	var slist * GnbAclPlmn_list = new(GnbAclPlmn_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT id, mcc, mnc, block, isactive FROM gnbaclplmn  ORDER BY id")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj GnbAclPlmn


			if rerr := rows.Scan( &obj.Id, &obj.MCC, &obj.MNC, &obj.Block, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan GnbAclPlmn_GetAll  %v+", rerr));
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []GnbAclPlmn{};
		}
		ctrGnbAclPlmnSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrGnbAclPlmnSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql GnbAclPlmn_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func GnbAclPlmn_Delete( dbid int, ptr * GnbAclPlmn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update gnbaclplmn set IsActive=10 where id=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.Id)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrGnbAclPlmnUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql GnbAclPlmn_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrGnbAclPlmnUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbAclPlmnUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - GnbAclPlmn_Update")
	}
	return false
}



// Table Name : GnbAclPlmn

func addGnbAclPlmn(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_GnbAclPlmn GnbAclPlmn

	if err := c.BindJSON(&new_GnbAclPlmn); err != nil {
		common.AppLogError("BindJSON failed for addGnbAclPlmn");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation


	//set custom feild values
	new_GnbAclPlmn.IsActive = 1;

	GnbAclPlmn_Add( 0, &new_GnbAclPlmn)
	c.IndentedJSON( http.StatusOK, gin.H{"Id": new_GnbAclPlmn.Id, "ErrorNo": new_GnbAclPlmn.ErrorNo, "ErrorName": new_GnbAclPlmn.ErrorName});
}


func updateGnbAclPlmn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbAclPlmn GnbAclPlmn

	if err := c.BindJSON( &obj_GnbAclPlmn); err != nil {
		common.AppLogError("BindJSON failed for updateGnbAclPlmn");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation


	//set custom feild values
	obj_GnbAclPlmn.IsActive = 1;


	sts := GnbAclPlmn_Update( 0, &obj_GnbAclPlmn)
	c.IndentedJSON( http.StatusOK, gin.H{"Id": obj_GnbAclPlmn.Id, "Status": sts, "ErrorNo": obj_GnbAclPlmn.ErrorNo, "ErrorName": obj_GnbAclPlmn.ErrorName});
}


func getGnbAclPlmn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbAclPlmn GnbAclPlmn;
	id, ok := c.GetQuery("id");
	if ok {
		obj_GnbAclPlmn.Id,_ = strconv.Atoi(id);
	}

	if obj_GnbAclPlmn.Id == 0 {
		common.AppLogError("id not found in query string - getGnbAclPlmn");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := GnbAclPlmn_Get( 0, &obj_GnbAclPlmn);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteGnbAclPlmn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbAclPlmn GnbAclPlmn;
	id, ok := c.GetQuery("id");
	if ok {
		obj_GnbAclPlmn.Id,_ = strconv.Atoi(id);
	}

	if obj_GnbAclPlmn.Id == 0 {
		common.AppLogError("id not found in query string - getGnbAclPlmn");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := GnbAclPlmn_Delete( 0, &obj_GnbAclPlmn);
	c.IndentedJSON( http.StatusOK, res)
}


func getGnbAclPlmnByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getGnbAclPlmnList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj GnbAclPlmn
	obj.IsActive = 1;

	list := GnbAclPlmn_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllGnbAclPlmn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj GnbAclPlmn

	list := GnbAclPlmn_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : GnbAclPlmn
func InitModule_GnbAclPlmn( router * gin.Engine) {
	router.POST("/gnbaclplmn/add", addGnbAclPlmn);
	router.POST("/gnbaclplmn/upd", updateGnbAclPlmn);
	router.POST("/gnbaclplmn/getgnbaclplmnbypaging", getGnbAclPlmnByPaging);
	router.GET("/gnbaclplmn/getall", getAllGnbAclPlmn);
	router.GET("/gnbaclplmn/getbyid", getGnbAclPlmn);
	router.GET("/gnbaclplmn/delete", deleteGnbAclPlmn);
}




