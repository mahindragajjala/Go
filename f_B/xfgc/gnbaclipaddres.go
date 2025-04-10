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


// Table Name : GnbACLIPAddres
type GnbACLIPAddres struct {
	Id int;
	IPAddres string;
	Block int;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type GnbACLIPAddres_list struct {
	TotalCount int
	Items []GnbACLIPAddres
	ErrorNo int;
	ErrorName string;
}




func GnbACLIPAddresDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter GnbACLIPAddres
var ctrGnbACLIPAddresInsertSuccess * common.Counter = &common.Counter{Name:"GnbACLIPAddresInsertSuccess"};
var ctrGnbACLIPAddresInsertError * common.Counter = &common.Counter{Name:"GnbACLIPAddresInsertError"};
var ctrGnbACLIPAddresUpdateSuccess * common.Counter = &common.Counter{Name:"GnbACLIPAddresUpdateSuccess"};
var ctrGnbACLIPAddresUpdateError * common.Counter = &common.Counter{Name:"GnbACLIPAddresUpdateError"};
var ctrGnbACLIPAddresSelectSuccess * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectSuccess"};
var ctrGnbACLIPAddresSelectError * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectError"};
var ctrGnbACLIPAddresSelectAllSuccess * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectAllSuccess"};
var ctrGnbACLIPAddresSelectAllError * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectAllError"};
var ctrGnbACLIPAddresSelectPageSuccess * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectPageSuccess"};
var ctrGnbACLIPAddresSelectPageError * common.Counter = &common.Counter{Name:"GnbACLIPAddresSelectPageError"};

func RegisterCounter_GnbACLIPAddres() {
	common.RegisterCounter( ctrGnbACLIPAddresInsertSuccess);
	common.RegisterCounter( ctrGnbACLIPAddresInsertError);
	common.RegisterCounter( ctrGnbACLIPAddresUpdateSuccess);
	common.RegisterCounter( ctrGnbACLIPAddresUpdateError);
	common.RegisterCounter( ctrGnbACLIPAddresSelectSuccess);
	common.RegisterCounter( ctrGnbACLIPAddresSelectError);
	common.RegisterCounter( ctrGnbACLIPAddresSelectAllSuccess);
	common.RegisterCounter( ctrGnbACLIPAddresSelectAllError);
	common.RegisterCounter( ctrGnbACLIPAddresSelectPageSuccess);
	common.RegisterCounter( ctrGnbACLIPAddresSelectPageError);
}



func GnbACLIPAddres_Add( dbid int, ptr * GnbACLIPAddres) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into gnbaclipaddres( ipaddres, block, isactive) values( $1, $2, $3) RETURNING id`
		err := dbConn.Db.QueryRow( insertStmt, ptr.IPAddres, ptr.Block, ptr.IsActive).Scan( &ptr.Id)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrGnbACLIPAddresInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql GnbACLIPAddres_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrGnbACLIPAddresInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbACLIPAddres_Add")
	}
	return false
}

func GnbACLIPAddres_Update( dbid int, ptr * GnbACLIPAddres) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update gnbaclipaddres set ipaddres=$1 , block=$2 , isactive=$3 where id=$4`
		_, err := dbConn.Db.Exec( updateStmt, ptr.IPAddres, ptr.Block, ptr.IsActive, ptr.Id)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrGnbACLIPAddresUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql GnbACLIPAddres_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrGnbACLIPAddresUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - GnbACLIPAddres_Update")
	}
	return false
}

func GnbACLIPAddres_Get( dbid int, ptr * GnbACLIPAddres) * GnbACLIPAddres {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbACLIPAddres_Get");
		return nil;
	}

	selectStmt := `select id, ipaddres, block, isactive from gnbaclipaddres where id=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.Id)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrGnbACLIPAddresSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql GnbACLIPAddres_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * GnbACLIPAddres = new(GnbACLIPAddres)

	var strIPAddres sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.Id, &strIPAddres, &obj.Block, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan GnbACLIPAddres_Get  %v+", err));
		}

		if strIPAddres.String != "" {
			obj.IPAddres = strIPAddres.String;
		}

	}
	rows.Close()

	ctrGnbACLIPAddresSelectSuccess.Increment();
	return obj
}

func GnbACLIPAddres_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * GnbACLIPAddres) * GnbACLIPAddres_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbACLIPAddres_GetByPaging");
		return nil;
	}

	var slist * GnbACLIPAddres_list = new(GnbACLIPAddres_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT id, ipaddres, block, isactive FROM gnbaclipaddres  where isactive=$1 ORDER BY id OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj GnbACLIPAddres

			var strIPAddres sql.NullString

			if rerr := rows.Scan( &obj.Id, &strIPAddres, &obj.Block, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan GnbACLIPAddres_GetByPaging  %v+", rerr));
			}

			if strIPAddres.String != "" {
				obj.IPAddres = strIPAddres.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM gnbaclipaddres where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql GnbACLIPAddres_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []GnbACLIPAddres{}
		}
		ctrGnbACLIPAddresSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrGnbACLIPAddresSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql GnbACLIPAddres_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func GnbACLIPAddres_GetAll( dbid int, ptr * GnbACLIPAddres) * GnbACLIPAddres_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - GnbACLIPAddres_GetAll");
		return nil;
	}

	var slist * GnbACLIPAddres_list = new(GnbACLIPAddres_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT id, ipaddres, block, isactive FROM gnbaclipaddres  ORDER BY id")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj GnbACLIPAddres

			var strIPAddres sql.NullString

			if rerr := rows.Scan( &obj.Id, &strIPAddres, &obj.Block, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan GnbACLIPAddres_GetAll  %v+", rerr));
			}

			if strIPAddres.String != "" {
				obj.IPAddres = strIPAddres.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []GnbACLIPAddres{};
		}
		ctrGnbACLIPAddresSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrGnbACLIPAddresSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql GnbACLIPAddres_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func GnbACLIPAddres_Delete( dbid int, ptr * GnbACLIPAddres) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update gnbaclipaddres set IsActive=10 where id=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.Id)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrGnbACLIPAddresUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql GnbACLIPAddres_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrGnbACLIPAddresUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrGnbACLIPAddresUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - GnbACLIPAddres_Update")
	}
	return false
}



// Table Name : GnbACLIPAddres

func addGnbACLIPAddres(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_GnbACLIPAddres GnbACLIPAddres

	if err := c.BindJSON(&new_GnbACLIPAddres); err != nil {
		common.AppLogError("BindJSON failed for addGnbACLIPAddres");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_GnbACLIPAddres.IPAddres == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"IPAddres is required"})
		return
	}


	//set custom feild values
	new_GnbACLIPAddres.IsActive = 1;

	GnbACLIPAddres_Add( 0, &new_GnbACLIPAddres)
	c.IndentedJSON( http.StatusOK, gin.H{"Id": new_GnbACLIPAddres.Id, "ErrorNo": new_GnbACLIPAddres.ErrorNo, "ErrorName": new_GnbACLIPAddres.ErrorName});
}


func updateGnbACLIPAddres(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbACLIPAddres GnbACLIPAddres

	if err := c.BindJSON( &obj_GnbACLIPAddres); err != nil {
		common.AppLogError("BindJSON failed for updateGnbACLIPAddres");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_GnbACLIPAddres.IPAddres == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"IPAddres is required"})
		return
	}


	//set custom feild values
	obj_GnbACLIPAddres.IsActive = 1;


	sts := GnbACLIPAddres_Update( 0, &obj_GnbACLIPAddres)
	c.IndentedJSON( http.StatusOK, gin.H{"Id": obj_GnbACLIPAddres.Id, "Status": sts, "ErrorNo": obj_GnbACLIPAddres.ErrorNo, "ErrorName": obj_GnbACLIPAddres.ErrorName});
}


func getGnbACLIPAddres(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbACLIPAddres GnbACLIPAddres;
	id, ok := c.GetQuery("id");
	if ok {
		obj_GnbACLIPAddres.Id,_ = strconv.Atoi(id);
	}

	if obj_GnbACLIPAddres.Id == 0 {
		common.AppLogError("id not found in query string - getGnbACLIPAddres");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := GnbACLIPAddres_Get( 0, &obj_GnbACLIPAddres);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteGnbACLIPAddres(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_GnbACLIPAddres GnbACLIPAddres;
	id, ok := c.GetQuery("id");
	if ok {
		obj_GnbACLIPAddres.Id,_ = strconv.Atoi(id);
	}

	if obj_GnbACLIPAddres.Id == 0 {
		common.AppLogError("id not found in query string - getGnbACLIPAddres");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := GnbACLIPAddres_Delete( 0, &obj_GnbACLIPAddres);
	c.IndentedJSON( http.StatusOK, res)
}


func getGnbACLIPAddresByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getGnbACLIPAddresList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj GnbACLIPAddres
	obj.IsActive = 1;

	list := GnbACLIPAddres_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllGnbACLIPAddres(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj GnbACLIPAddres

	list := GnbACLIPAddres_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : GnbACLIPAddres
func InitModule_GnbACLIPAddres( router * gin.Engine) {
	router.POST("/gnbaclipaddres/add", addGnbACLIPAddres);
	router.POST("/gnbaclipaddres/upd", updateGnbACLIPAddres);
	router.POST("/gnbaclipaddres/getgnbaclipaddresbypaging", getGnbACLIPAddresByPaging);
	router.GET("/gnbaclipaddres/getall", getAllGnbACLIPAddres);
	router.GET("/gnbaclipaddres/getbyid", getGnbACLIPAddres);
	router.GET("/gnbaclipaddres/delete", deleteGnbACLIPAddres);
}




