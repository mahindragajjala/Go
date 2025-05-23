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


// Table Name : Dnn
type Dnn struct {
	DnnId int;
	DNName string;
	SupportedPduSessionType int;
	AllowedPduSessionType int;
	SupportedSSCMode1 int;
	AllowedSSCMode1 int;
	SupportedSSCMode2 int;
	AllowedSSCMode2 int;
	SupportedSSCMode3 int;
	AllowedSSCMode3 int;
	QosProfilePriorityLevel int;
	QosProfile5qi int;
	QosProfileArpPriorityLevel int;
	QosProfileArpPreemptCap int;
	QosProfileArpPreemptVuln int;
	Uplink int;
	Downlink int;
	PLMN string;
	IsActive int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type Dnn_list struct {
	TotalCount int
	Items []Dnn
	ErrorNo int;
	ErrorName string;
}




func DnnDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter Dnn
var ctrDnnInsertSuccess * common.Counter = &common.Counter{Name:"DnnInsertSuccess"};
var ctrDnnInsertError * common.Counter = &common.Counter{Name:"DnnInsertError"};
var ctrDnnUpdateSuccess * common.Counter = &common.Counter{Name:"DnnUpdateSuccess"};
var ctrDnnUpdateError * common.Counter = &common.Counter{Name:"DnnUpdateError"};
var ctrDnnSelectSuccess * common.Counter = &common.Counter{Name:"DnnSelectSuccess"};
var ctrDnnSelectError * common.Counter = &common.Counter{Name:"DnnSelectError"};
var ctrDnnSelectAllSuccess * common.Counter = &common.Counter{Name:"DnnSelectAllSuccess"};
var ctrDnnSelectAllError * common.Counter = &common.Counter{Name:"DnnSelectAllError"};
var ctrDnnSelectPageSuccess * common.Counter = &common.Counter{Name:"DnnSelectPageSuccess"};
var ctrDnnSelectPageError * common.Counter = &common.Counter{Name:"DnnSelectPageError"};

func RegisterCounter_Dnn() {
	common.RegisterCounter( ctrDnnInsertSuccess);
	common.RegisterCounter( ctrDnnInsertError);
	common.RegisterCounter( ctrDnnUpdateSuccess);
	common.RegisterCounter( ctrDnnUpdateError);
	common.RegisterCounter( ctrDnnSelectSuccess);
	common.RegisterCounter( ctrDnnSelectError);
	common.RegisterCounter( ctrDnnSelectAllSuccess);
	common.RegisterCounter( ctrDnnSelectAllError);
	common.RegisterCounter( ctrDnnSelectPageSuccess);
	common.RegisterCounter( ctrDnnSelectPageError);
}



func Dnn_Add( dbid int, ptr * Dnn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into dnn( dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING dnnid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.DNName, ptr.SupportedPduSessionType, ptr.AllowedPduSessionType, ptr.SupportedSSCMode1, ptr.AllowedSSCMode1, ptr.SupportedSSCMode2, ptr.AllowedSSCMode2, ptr.SupportedSSCMode3, ptr.AllowedSSCMode3, ptr.QosProfilePriorityLevel, ptr.QosProfile5qi, ptr.QosProfileArpPriorityLevel, ptr.QosProfileArpPreemptCap, ptr.QosProfileArpPreemptVuln, ptr.Uplink, ptr.Downlink, ptr.PLMN, ptr.IsActive).Scan( &ptr.DnnId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrDnnInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql Dnn_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrDnnInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrDnnInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - Dnn_Add")
	}
	return false
}

func Dnn_Update( dbid int, ptr * Dnn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update dnn set dnname=$1 , supportedpdusessiontype=$2 , allowedpdusessiontype=$3 , supportedsscmode1=$4 , allowedsscmode1=$5 , supportedsscmode2=$6 , allowedsscmode2=$7 , supportedsscmode3=$8 , allowedsscmode3=$9 , qosprofileprioritylevel=$10 , qosprofile5qi=$11 , qosprofilearpprioritylevel=$12 , qosprofilearppreemptcap=$13 , qosprofilearppreemptvuln=$14 , uplink=$15 , downlink=$16 , plmn=$17 , isactive=$18 where dnnid=$19`
		_, err := dbConn.Db.Exec( updateStmt, ptr.DNName, ptr.SupportedPduSessionType, ptr.AllowedPduSessionType, ptr.SupportedSSCMode1, ptr.AllowedSSCMode1, ptr.SupportedSSCMode2, ptr.AllowedSSCMode2, ptr.SupportedSSCMode3, ptr.AllowedSSCMode3, ptr.QosProfilePriorityLevel, ptr.QosProfile5qi, ptr.QosProfileArpPriorityLevel, ptr.QosProfileArpPreemptCap, ptr.QosProfileArpPreemptVuln, ptr.Uplink, ptr.Downlink, ptr.PLMN, ptr.IsActive, ptr.DnnId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrDnnUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql Dnn_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrDnnUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrDnnUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - Dnn_Update")
	}
	return false
}

func Dnn_Get( dbid int, ptr * Dnn) * Dnn {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrDnnSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - Dnn_Get");
		return nil;
	}

	selectStmt := `select dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive from dnn where dnnid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.DnnId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrDnnSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql Dnn_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * Dnn = new(Dnn)

	var strDNName sql.NullString
	var strPLMN sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.DnnId, &strDNName, &obj.SupportedPduSessionType, &obj.AllowedPduSessionType, &obj.SupportedSSCMode1, &obj.AllowedSSCMode1, &obj.SupportedSSCMode2, &obj.AllowedSSCMode2, &obj.SupportedSSCMode3, &obj.AllowedSSCMode3, &obj.QosProfilePriorityLevel, &obj.QosProfile5qi, &obj.QosProfileArpPriorityLevel, &obj.QosProfileArpPreemptCap, &obj.QosProfileArpPreemptVuln, &obj.Uplink, &obj.Downlink, &strPLMN, &obj.IsActive)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan Dnn_Get  %v+", err));
		}

		if strDNName.String != "" {
			obj.DNName = strDNName.String;
		}

		if strPLMN.String != "" {
			obj.PLMN = strPLMN.String;
		}

	}
	rows.Close()

	ctrDnnSelectSuccess.Increment();
	return obj
}

func Dnn_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * Dnn) * Dnn_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrDnnSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - Dnn_GetByPaging");
		return nil;
	}

	var slist * Dnn_list = new(Dnn_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive FROM dnn  where isactive=$1 ORDER BY dnnid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := dbConn.Db.Query( selQuery, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj Dnn

			var strDNName sql.NullString
			var strPLMN sql.NullString

			if rerr := rows.Scan( &obj.DnnId, &strDNName, &obj.SupportedPduSessionType, &obj.AllowedPduSessionType, &obj.SupportedSSCMode1, &obj.AllowedSSCMode1, &obj.SupportedSSCMode2, &obj.AllowedSSCMode2, &obj.SupportedSSCMode3, &obj.AllowedSSCMode3, &obj.QosProfilePriorityLevel, &obj.QosProfile5qi, &obj.QosProfileArpPriorityLevel, &obj.QosProfileArpPreemptCap, &obj.QosProfileArpPreemptVuln, &obj.Uplink, &obj.Downlink, &strPLMN, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan Dnn_GetByPaging  %v+", rerr));
			}

			if strDNName.String != "" {
				obj.DNName = strDNName.String;
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM dnn where 1=1  and isactive=$1", ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql Dnn_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []Dnn{}
		}
		ctrDnnSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrDnnSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql Dnn_GetByPaging  %s %v+", selQuery, err));
	}

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func Dnn_GetAll( dbid int, ptr * Dnn) * Dnn_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrDnnSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - Dnn_GetAll");
		return nil;
	}

	var slist * Dnn_list = new(Dnn_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive FROM dnn  ORDER BY dnnid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj Dnn

			var strDNName sql.NullString
			var strPLMN sql.NullString

			if rerr := rows.Scan( &obj.DnnId, &strDNName, &obj.SupportedPduSessionType, &obj.AllowedPduSessionType, &obj.SupportedSSCMode1, &obj.AllowedSSCMode1, &obj.SupportedSSCMode2, &obj.AllowedSSCMode2, &obj.SupportedSSCMode3, &obj.AllowedSSCMode3, &obj.QosProfilePriorityLevel, &obj.QosProfile5qi, &obj.QosProfileArpPriorityLevel, &obj.QosProfileArpPreemptCap, &obj.QosProfileArpPreemptVuln, &obj.Uplink, &obj.Downlink, &strPLMN, &obj.IsActive); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan Dnn_GetAll  %v+", rerr));
			}

			if strDNName.String != "" {
				obj.DNName = strDNName.String;
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []Dnn{};
		}
		ctrDnnSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrDnnSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql Dnn_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func Dnn_Delete( dbid int, ptr * Dnn) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update dnn set IsActive=10 where dnnid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.DnnId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrDnnUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql Dnn_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrDnnUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrDnnUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - Dnn_Update")
	}
	return false
}



// Table Name : Dnn

func addDnn(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_Dnn Dnn

	if err := c.BindJSON(&new_Dnn); err != nil {
		common.AppLogError("BindJSON failed for addDnn");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_Dnn.DNName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"DNName is required"})
		return
	}

	if new_Dnn.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}


	//set custom feild values
	new_Dnn.IsActive = 1;

	Dnn_Add( 0, &new_Dnn)
	c.IndentedJSON( http.StatusOK, gin.H{"DnnId": new_Dnn.DnnId, "ErrorNo": new_Dnn.ErrorNo, "ErrorName": new_Dnn.ErrorName});
}


func updateDnn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_Dnn Dnn

	if err := c.BindJSON( &obj_Dnn); err != nil {
		common.AppLogError("BindJSON failed for updateDnn");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_Dnn.DNName == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"DNName is required"})
		return
	}

	if obj_Dnn.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}


	//set custom feild values
	obj_Dnn.IsActive = 1;


	sts := Dnn_Update( 0, &obj_Dnn)
	c.IndentedJSON( http.StatusOK, gin.H{"DnnId": obj_Dnn.DnnId, "Status": sts, "ErrorNo": obj_Dnn.ErrorNo, "ErrorName": obj_Dnn.ErrorName});
}


func getDnn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_Dnn Dnn;
	id, ok := c.GetQuery("id");
	if ok {
		obj_Dnn.DnnId,_ = strconv.Atoi(id);
	}

	if obj_Dnn.DnnId == 0 {
		common.AppLogError("id not found in query string - getDnn");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := Dnn_Get( 0, &obj_Dnn);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteDnn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_Dnn Dnn;
	id, ok := c.GetQuery("id");
	if ok {
		obj_Dnn.DnnId,_ = strconv.Atoi(id);
	}

	if obj_Dnn.DnnId == 0 {
		common.AppLogError("id not found in query string - getDnn");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := Dnn_Delete( 0, &obj_Dnn);
	c.IndentedJSON( http.StatusOK, res)
}


func getDnnByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getDnnList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj Dnn
	obj.IsActive = 1;

	list := Dnn_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllDnn(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj Dnn

	list := Dnn_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : Dnn
func InitModule_Dnn( router * gin.Engine) {
	router.POST("/dnn/add", addDnn);
	router.POST("/dnn/upd", updateDnn);
	router.POST("/dnn/getdnnbypaging", getDnnByPaging);
	router.GET("/dnn/getall", getAllDnn);
	router.GET("/dnn/getbyid", getDnn);
	router.GET("/dnn/delete", deleteDnn);
}




