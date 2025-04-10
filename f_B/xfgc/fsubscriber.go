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


// Table Name : FSubscriber
type FSubscriber struct {
	SubscriberId int;
	IMSI int;
	PLMN string;
	SKey string;
	OPc string;
	AMF string;
	SEQ string;
	AccountStatus int;
	DefSST int;
	DefSD string;
	Uplink int;
	Downlink int;
	DnnId1 int;
	DnnId2 int;
	DnnId3 int;
	CreatedDate string;
	UpdatedDate string;
	CreatedBy int;
	UpdatedBy int;
	IsActive int;
	RowVersion int;
	MSISDN string;
	MCC int;
	MNC int;

	Search string;
	ErrorNo int;
	ErrorName string;
}

type FSubscriber_list struct {
	TotalCount int
	Items []FSubscriber
	ErrorNo int;
	ErrorName string;
}




func FSubscriberDummy(){
    fmt.Printf( "%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"));
    
}

// Counter FSubscriber
var ctrFSubscriberInsertSuccess * common.Counter = &common.Counter{Name:"FSubscriberInsertSuccess"};
var ctrFSubscriberInsertError * common.Counter = &common.Counter{Name:"FSubscriberInsertError"};
var ctrFSubscriberUpdateSuccess * common.Counter = &common.Counter{Name:"FSubscriberUpdateSuccess"};
var ctrFSubscriberUpdateError * common.Counter = &common.Counter{Name:"FSubscriberUpdateError"};
var ctrFSubscriberSelectSuccess * common.Counter = &common.Counter{Name:"FSubscriberSelectSuccess"};
var ctrFSubscriberSelectError * common.Counter = &common.Counter{Name:"FSubscriberSelectError"};
var ctrFSubscriberSelectAllSuccess * common.Counter = &common.Counter{Name:"FSubscriberSelectAllSuccess"};
var ctrFSubscriberSelectAllError * common.Counter = &common.Counter{Name:"FSubscriberSelectAllError"};
var ctrFSubscriberSelectPageSuccess * common.Counter = &common.Counter{Name:"FSubscriberSelectPageSuccess"};
var ctrFSubscriberSelectPageError * common.Counter = &common.Counter{Name:"FSubscriberSelectPageError"};
var ctrFSubscriber_SelectByIMSISuccess * common.Counter = &common.Counter{Name:"FSubscriber_SelectByIMSISuccess"};
var ctrFSubscriber_SelectByIMSIError * common.Counter = &common.Counter{Name:"FSubscriber_SelectByIMSIError"};

func RegisterCounter_FSubscriber() {
	common.RegisterCounter( ctrFSubscriberInsertSuccess);
	common.RegisterCounter( ctrFSubscriberInsertError);
	common.RegisterCounter( ctrFSubscriberUpdateSuccess);
	common.RegisterCounter( ctrFSubscriberUpdateError);
	common.RegisterCounter( ctrFSubscriberSelectSuccess);
	common.RegisterCounter( ctrFSubscriberSelectError);
	common.RegisterCounter( ctrFSubscriberSelectAllSuccess);
	common.RegisterCounter( ctrFSubscriberSelectAllError);
	common.RegisterCounter( ctrFSubscriberSelectPageSuccess);
	common.RegisterCounter( ctrFSubscriberSelectPageError);
	common.RegisterCounter( ctrFSubscriber_SelectByIMSISuccess);
	common.RegisterCounter( ctrFSubscriber_SelectByIMSIError);
}



func FSubscriber_Add( dbid int, ptr * FSubscriber) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		insertStmt := `insert into fsubscriber( imsi, plmn, skey, opc, amf, seq, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, createddate, updateddate, createdby, updatedby, isactive, rowversion, msisdn, mcc, mnc) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) RETURNING subscriberid`
		err := dbConn.Db.QueryRow( insertStmt, ptr.IMSI, ptr.PLMN, ptr.SKey, ptr.OPc, ptr.AMF, ptr.SEQ, ptr.AccountStatus, ptr.DefSST, ptr.DefSD, ptr.Uplink, ptr.Downlink, ptr.DnnId1, ptr.DnnId2, ptr.DnnId3, ptr.CreatedDate, ptr.UpdatedDate, ptr.CreatedBy, ptr.UpdatedBy, ptr.IsActive, ptr.RowVersion, ptr.MSISDN, ptr.MCC, ptr.MNC).Scan( &ptr.SubscriberId)
		if err != nil {
			ptr.ErrorNo = 5002;
			ctrFSubscriberInsertError.Increment();
			common.DbLogError( fmt.Sprintf( "error in insert-sql FSubscriber_Add  %s %v+", insertStmt, err));
			return false;
		}

		ctrFSubscriberInsertSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrFSubscriberInsertError.Increment();
		common.DbLogError("DBConnection Allocation failed - FSubscriber_Add")
	}
	return false
}

func FSubscriber_Update( dbid int, ptr * FSubscriber) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update fsubscriber set imsi=$1 , plmn=$2 , skey=$3 , opc=$4 , amf=$5 , seq=$6 , accountstatus=$7 , defsst=$8 , defsd=$9 , uplink=$10 , downlink=$11 , dnnid1=$12 , dnnid2=$13 , dnnid3=$14 , createdby=$15 , updatedby=$16 , isactive=$17 , rowversion=$18 , msisdn=$19 , mcc=$20 , mnc=$21 where subscriberid=$22`
		_, err := dbConn.Db.Exec( updateStmt, ptr.IMSI, ptr.PLMN, ptr.SKey, ptr.OPc, ptr.AMF, ptr.SEQ, ptr.AccountStatus, ptr.DefSST, ptr.DefSD, ptr.Uplink, ptr.Downlink, ptr.DnnId1, ptr.DnnId2, ptr.DnnId3, ptr.CreatedBy, ptr.UpdatedBy, ptr.IsActive, ptr.RowVersion, ptr.MSISDN, ptr.MCC, ptr.MNC, ptr.SubscriberId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrFSubscriberUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql FSubscriber_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrFSubscriberUpdateSuccess.Increment();

		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrFSubscriberUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - FSubscriber_Update")
	}
	return false
}

func FSubscriber_Get( dbid int, ptr * FSubscriber) * FSubscriber {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrFSubscriberSelectError.Increment();
		common.DbLogError("DBConnection Allocation failed - FSubscriber_Get");
		return nil;
	}

	selectStmt := `select subscriberid, imsi, plmn, skey, opc, amf, seq, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, createdby, updatedby, isactive, rowversion, msisdn, mcc, mnc from fsubscriber where subscriberid=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.SubscriberId)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrFSubscriberSelectError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-sql FSubscriber_Get  %s %v+", selectStmt, err));
		return nil;
	}

	var obj * FSubscriber = new(FSubscriber)

	var strPLMN sql.NullString
	var strSKey sql.NullString
	var strOPc sql.NullString
	var strAMF sql.NullString
	var strSEQ sql.NullString
	var strDefSD sql.NullString
	var strMSISDN sql.NullString

	for rows.Next() {
		err = rows.Scan( &obj.SubscriberId, &obj.IMSI, &strPLMN, &strSKey, &strOPc, &strAMF, &strSEQ, &obj.AccountStatus, &obj.DefSST, &strDefSD, &obj.Uplink, &obj.Downlink, &obj.DnnId1, &obj.DnnId2, &obj.DnnId3, &obj.CreatedBy, &obj.UpdatedBy, &obj.IsActive, &obj.RowVersion, &strMSISDN, &obj.MCC, &obj.MNC)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in Scan FSubscriber_Get  %v+", err));
		}

		if strPLMN.String != "" {
			obj.PLMN = strPLMN.String;
		}

		if strSKey.String != "" {
			obj.SKey = strSKey.String;
		}

		if strOPc.String != "" {
			obj.OPc = strOPc.String;
		}

		if strAMF.String != "" {
			obj.AMF = strAMF.String;
		}

		if strSEQ.String != "" {
			obj.SEQ = strSEQ.String;
		}

		if strDefSD.String != "" {
			obj.DefSD = strDefSD.String;
		}

		if strMSISDN.String != "" {
			obj.MSISDN = strMSISDN.String;
		}

	}
	rows.Close()

	ctrFSubscriberSelectSuccess.Increment();
	return obj
}

func FSubscriber_GetByPaging( dbid int, pageInfo common.FPageInfo, ptr * FSubscriber) * FSubscriber_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrFSubscriberSelectPageError.Increment();
		common.DbLogError("DBConnection Allocation failed - FSubscriber_GetByPaging");
		return nil;
	}

	//fmt.Printf("ptr=%+v\n", ptr);
	//fmt.Printf("pageInfo=%+v\n", pageInfo);
	
	var slist * FSubscriber_list = new(FSubscriber_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT subscriberid, imsi, plmn, skey, opc, amf, seq, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, createdby, updatedby, isactive, rowversion, msisdn, mcc, mnc FROM fsubscriber  where (imsi=$1 OR $1=0) and isactive=$2 ORDER BY subscriberid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	//fmt.Printf("selQuery=%+v\n", selQuery);
	
	
	rows, err := dbConn.Db.Query( selQuery, ptr.IMSI, ptr.IsActive);

	if err == nil {
		for rows.Next() {

			var obj FSubscriber

			var strPLMN sql.NullString
			var strSKey sql.NullString
			var strOPc sql.NullString
			var strAMF sql.NullString
			var strSEQ sql.NullString
			var strDefSD sql.NullString
			var strMSISDN sql.NullString

			if rerr := rows.Scan( &obj.SubscriberId, &obj.IMSI, &strPLMN, &strSKey, &strOPc, &strAMF, &strSEQ, &obj.AccountStatus, &obj.DefSST, &strDefSD, &obj.Uplink, &obj.Downlink, &obj.DnnId1, &obj.DnnId2, &obj.DnnId3, &obj.CreatedBy, &obj.UpdatedBy, &obj.IsActive, &obj.RowVersion, &strMSISDN, &obj.MCC, &obj.MNC); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-paging-scan FSubscriber_GetByPaging  %v+", rerr));
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			if strSKey.String != "" {
				obj.SKey = strSKey.String;
			}

			if strOPc.String != "" {
				obj.OPc = strOPc.String;
			}

			if strAMF.String != "" {
				obj.AMF = strAMF.String;
			}

			if strSEQ.String != "" {
				obj.SEQ = strSEQ.String;
			}

			if strDefSD.String != "" {
				obj.DefSD = strDefSD.String;
			}

			if strMSISDN.String != "" {
				obj.MSISDN = strMSISDN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		terr := dbConn.Db.QueryRow( "SELECT count(1) id FROM fsubscriber where 1=1  and (imsi=$1 OR $1=0) and isactive=$2", ptr.IMSI, ptr.IsActive).Scan( &slist.TotalCount)
		if terr != nil {
			ptr.ErrorNo = 5006;
			common.DbLogError( fmt.Sprintf( "error in select-paging-count-sql FSubscriber_GetByPaging  %v+", terr));
		}

		if slist.TotalCount == 0 {
			slist.Items = []FSubscriber{}
		}
		ctrFSubscriberSelectPageSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrFSubscriberSelectPageError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-paging-sql FSubscriber_GetByPaging  %s %v+", selQuery, err));
	}

	//fmt.Printf("slist=%+v\n", slist);

	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func FSubscriber_GetAll( dbid int, ptr * FSubscriber) * FSubscriber_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrFSubscriberSelectAllError.Increment();
		common.DbLogError("DBConnection Allocation failed - FSubscriber_GetAll");
		return nil;
	}

	var slist * FSubscriber_list = new(FSubscriber_list)
	slist.TotalCount = 0

	selQuery := fmt.Sprintf("SELECT subscriberid, imsi, plmn, skey, opc, amf, seq, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, createdby, updatedby, isactive, rowversion, msisdn, mcc, mnc FROM fsubscriber  ORDER BY subscriberid")
	rows, err := dbConn.Db.Query( selQuery);

	if err == nil {
		for rows.Next() {

			var obj FSubscriber

			var strPLMN sql.NullString
			var strSKey sql.NullString
			var strOPc sql.NullString
			var strAMF sql.NullString
			var strSEQ sql.NullString
			var strDefSD sql.NullString
			var strMSISDN sql.NullString

			if rerr := rows.Scan( &obj.SubscriberId, &obj.IMSI, &strPLMN, &strSKey, &strOPc, &strAMF, &strSEQ, &obj.AccountStatus, &obj.DefSST, &strDefSD, &obj.Uplink, &obj.Downlink, &obj.DnnId1, &obj.DnnId2, &obj.DnnId3, &obj.CreatedBy, &obj.UpdatedBy, &obj.IsActive, &obj.RowVersion, &strMSISDN, &obj.MCC, &obj.MNC); rerr != nil {
				ptr.ErrorNo = 5005;
				common.DbLogError( fmt.Sprintf( "error in select-all-scan FSubscriber_GetAll  %v+", rerr));
			}

			if strPLMN.String != "" {
				obj.PLMN = strPLMN.String;
			}

			if strSKey.String != "" {
				obj.SKey = strSKey.String;
			}

			if strOPc.String != "" {
				obj.OPc = strOPc.String;
			}

			if strAMF.String != "" {
				obj.AMF = strAMF.String;
			}

			if strSEQ.String != "" {
				obj.SEQ = strSEQ.String;
			}

			if strDefSD.String != "" {
				obj.DefSD = strDefSD.String;
			}

			if strMSISDN.String != "" {
				obj.MSISDN = strMSISDN.String;
			}

			slist.Items = append( slist.Items, obj)
		}
		rows.Close()

		if slist.Items == nil {
			slist.Items = []FSubscriber{};
		}
		ctrFSubscriberSelectAllSuccess.Increment();
	} else {
		ptr.ErrorNo = 5004;
		ctrFSubscriberSelectAllError.Increment();
		common.DbLogError( fmt.Sprintf( "error in select-all-sql FSubscriber_GetAll  %s %v+", selQuery, err));
	}

	slist.TotalCount = len( slist.Items);
	slist.ErrorNo = ptr.ErrorNo;
	return slist
}

func FSubscriber_Delete( dbid int, ptr * FSubscriber) bool {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn != nil {
		updateStmt := `update fsubscriber set IsActive=10 where subscriberid=$1`
		_, err := dbConn.Db.Exec( updateStmt, ptr.SubscriberId)

		if err != nil {
			ptr.ErrorNo = 5003;
			ctrFSubscriberUpdateError.Increment();
			common.DbLogError( fmt.Sprintf( "error in update-sql FSubscriber_Update  %s %v+", updateStmt, err));
			return false;
		}

		ctrFSubscriberUpdateSuccess.Increment();
		return true
	} else {
		ptr.ErrorNo = 5001;
		ctrFSubscriberUpdateError.Increment();
		common.DbLogError( "DBConnection Allocation failed - FSubscriber_Update")
	}
	return false
}

func FSubscriber_SelectByIMSI( dbid int, ptr * FSubscriber) * FSubscriber_list {
	dbConn := postgres.GetDbConnPool("xfgc").GetConn();
	defer postgres.PutPGConnection( dbConn)

	if dbConn == nil {
		ptr.ErrorNo = 5001;
		ctrFSubscriber_SelectByIMSIError.Increment();
		common.DbLogError("DBConnection Allocation failed - FSubscriber_SelectByIMSI");
		return nil;
	}

	selectStmt := `select subscriberid, imsi, plmn, skey, opc, amf, seq, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, createdby, updatedby, isactive, rowversion, msisdn, mcc, mnc from fsubscriber where imsi=$1`
	rows, err := dbConn.Db.Query( selectStmt, ptr.IMSI)

	if err != nil {
		ptr.ErrorNo = 5004;
		ctrFSubscriber_SelectByIMSIError.Increment();
		common.DbLogError( fmt.Sprintf( "error in FSubscriber_SelectByIMSI-sql  %s %v+", selectStmt, err));
		return nil;
	}

	var slist * FSubscriber_list = new(FSubscriber_list)
	slist.TotalCount = 0

	for rows.Next() {
		var obj FSubscriber

		var strPLMN sql.NullString
		var strSKey sql.NullString
		var strOPc sql.NullString
		var strAMF sql.NullString
		var strSEQ sql.NullString
		var strDefSD sql.NullString
		var strMSISDN sql.NullString

		err = rows.Scan( &obj.SubscriberId, &obj.IMSI, &strPLMN, &strSKey, &strOPc, &strAMF, &strSEQ, &obj.AccountStatus, &obj.DefSST, &strDefSD, &obj.Uplink, &obj.Downlink, &obj.DnnId1, &obj.DnnId2, &obj.DnnId3, &obj.CreatedBy, &obj.UpdatedBy, &obj.IsActive, &obj.RowVersion, &strMSISDN, &obj.MCC, &obj.MNC)

		if err != nil {
			ptr.ErrorNo = 5005;
			common.DbLogError( fmt.Sprintf( "error in FSubscriber_SelectByIMSI-Scan  %v+", err));
		}

		if strPLMN.String != "" {
			obj.PLMN = strPLMN.String;
		}

		if strSKey.String != "" {
			obj.SKey = strSKey.String;
		}

		if strOPc.String != "" {
			obj.OPc = strOPc.String;
		}

		if strAMF.String != "" {
			obj.AMF = strAMF.String;
		}

		if strSEQ.String != "" {
			obj.SEQ = strSEQ.String;
		}

		if strDefSD.String != "" {
			obj.DefSD = strDefSD.String;
		}

		if strMSISDN.String != "" {
			obj.MSISDN = strMSISDN.String;
		}

		slist.Items = append( slist.Items, obj)
	}
	rows.Close()

	slist.TotalCount = len(slist.Items)
	ctrFSubscriber_SelectByIMSISuccess.Increment();
	return slist
}



// Table Name : FSubscriber

func addFSubscriber(c * gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.CreateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var new_FSubscriber FSubscriber

	if err := c.BindJSON(&new_FSubscriber); err != nil {
		common.AppLogError("BindJSON failed for addFSubscriber");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if new_FSubscriber.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}

	if new_FSubscriber.SKey == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"SKey is required"})
		return
	}

	if new_FSubscriber.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"OPc is required"})
		return
	}

	if new_FSubscriber.DefSD == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"DefSD is required"})
		return
	}

	if new_FSubscriber.MSISDN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"MSISDN is required"})
		return
	}


	//set custom feild values
	new_FSubscriber.AMF = "02";
	new_FSubscriber.SEQ = "0";
	new_FSubscriber.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05");
	new_FSubscriber.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05");
	new_FSubscriber.CreatedBy = user.UserId;
	new_FSubscriber.UpdatedBy = user.UserId;
	new_FSubscriber.IsActive = 1;
	new_FSubscriber.RowVersion = 1;

	FSubscriber_Add( 0, &new_FSubscriber)
	c.IndentedJSON( http.StatusOK, gin.H{"SubscriberId": new_FSubscriber.SubscriberId, "ErrorNo": new_FSubscriber.ErrorNo, "ErrorName": new_FSubscriber.ErrorName});
}


func updateFSubscriber(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.UpdateAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_FSubscriber FSubscriber

	if err := c.BindJSON( &obj_FSubscriber); err != nil {
		common.AppLogError("BindJSON failed for updateFSubscriber");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	//validation

	if obj_FSubscriber.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"PLMN is required"})
		return
	}

	if obj_FSubscriber.SKey == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"SKey is required"})
		return
	}

	if obj_FSubscriber.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"OPc is required"})
		return
	}

	if obj_FSubscriber.DefSD == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"DefSD is required"})
		return
	}

	if obj_FSubscriber.MSISDN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo":6000,"ErrorName":"MSISDN is required"})
		return
	}


	//set custom feild values
	obj_FSubscriber.AMF = "02";
	obj_FSubscriber.SEQ = "0";
	obj_FSubscriber.CreatedBy = user.UserId;
	obj_FSubscriber.UpdatedBy = user.UserId;
	obj_FSubscriber.IsActive = 1;
	obj_FSubscriber.RowVersion = 1;


	sts := FSubscriber_Update( 0, &obj_FSubscriber)
	c.IndentedJSON( http.StatusOK, gin.H{"SubscriberId": obj_FSubscriber.SubscriberId, "Status": sts, "ErrorNo": obj_FSubscriber.ErrorNo, "ErrorName": obj_FSubscriber.ErrorName});
}


func getFSubscriber(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ViewAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_FSubscriber FSubscriber;
	id, ok := c.GetQuery("id");
	if ok {
		obj_FSubscriber.SubscriberId,_ = strconv.Atoi(id);
	}

	if obj_FSubscriber.SubscriberId == 0 {
		common.AppLogError("id not found in query string - getFSubscriber");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := FSubscriber_Get( 0, &obj_FSubscriber);
	c.IndentedJSON( http.StatusOK, res)
}


func deleteFSubscriber(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.DeleteAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_FSubscriber FSubscriber;
	id, ok := c.GetQuery("id");
	if ok {
		obj_FSubscriber.SubscriberId,_ = strconv.Atoi(id);
	}

	if obj_FSubscriber.SubscriberId == 0 {
		common.AppLogError("id not found in query string - getFSubscriber");
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4999, "ErrorName": "id Not Provided"});
		return
	}


	res := FSubscriber_Delete( 0, &obj_FSubscriber);
	c.IndentedJSON( http.StatusOK, res)
}


func getFSubscriberByPaging(c *gin.Context) {

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
		common.AppLogError("BindJSON failed for getFSubscriberList");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":perr.Error()})
		return
	}

	var obj FSubscriber
	obj.IMSI = common.IntConv(pageInfo.Info.Search);
	obj.IsActive = 1;

	list := FSubscriber_GetByPaging( 0, pageInfo, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getAllFSubscriber(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj FSubscriber

	list := FSubscriber_GetAll( 0, &obj)
	c.IndentedJSON( http.StatusOK, list)
}

func getFSubscriberByIMSI(c *gin.Context) {

	user := common.Session_GetUser(c);
	if user == nil {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"});
		return
	}

	if user.HasAccess( 4, common.ListAccess) == false {
		c.IndentedJSON( http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"});
		return
	}

	var obj_FSubscriber FSubscriber

	if err := c.BindJSON( &obj_FSubscriber); err != nil {
		common.AppLogError("BindJSON failed for updateFSubscriber");
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo":5000,"ErrorName":err.Error()})
		return
	}


	list := FSubscriber_SelectByIMSI( 0, &obj_FSubscriber)
	c.IndentedJSON( http.StatusOK, list)
}





//// Table Name : FSubscriber
func InitModule_FSubscriber( router * gin.Engine) {
	router.POST("/fsubscriber/add", addFSubscriber);
	router.POST("/fsubscriber/upd", updateFSubscriber);
	router.POST("/fsubscriber/getfsubscriberbypaging", getFSubscriberByPaging);
	router.GET("/fsubscriber/getall", getAllFSubscriber);
	router.GET("/fsubscriber/getbyid", getFSubscriber);
	router.GET("/fsubscriber/delete", deleteFSubscriber);
}




