package application

import (
	"fmt"
	"time"

	"net/http"
	"strconv"

	// "database/sql"

	"xius.com/xfgc/common"
	"xius.com/xfgc/xfgc"

	// "xius.com/xfgc/postgres"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthInfoRequest struct {
	Id int
}

type ModuleInfoResponse struct {
	ModuleId          int
	CreatePermission  int
	ViewPermission    int
	EditPermission    int
	DeletePerimission int
}

type AuthInfoResponse struct {
	UserId           int
	UserName         string
	OrganizationName string
	City             string
	UserTypeId       int
	Modules          []ModuleInfoResponse
	ReadOnly         int
	ErrorNo          int
	ErrorName        string
}

func AppUserDummy() {
	fmt.Printf("%s\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
}

func AppWebLogin(userName string, password string) *xfgc.AppUser {

	var userObj xfgc.AppUser
	userObj.AppUserName = userName

	uList := xfgc.AppUser_SelectByUserName(0, &userObj)

	if uList != nil {
		if uList.TotalCount == 1 {
			if uList.Items[0].Password == password && uList.Items[0].IsAllowed == 1 && uList.Items[0].IsActive == 1 {
				return &uList.Items[0]
			}
		}
	}

	return nil
}
func AuthInfo(c *gin.Context) {
	session := sessions.Default(c)
	x := session.Get("user")

	if x == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Access"})
		return
	}

	user, ok := x.(xfgc.AppUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	var authInfoResponse AuthInfoResponse
	authInfoResponse.UserId = user.AppUserId
	authInfoResponse.UserName = user.AppUserName
	authInfoResponse.UserTypeId = user.UserTypeId
	authInfoResponse.ReadOnly = 0
	authInfoResponse.OrganizationName = ""
	authInfoResponse.City = ""
	authInfoResponse.ErrorNo = 0
	authInfoResponse.ErrorName = ""

	var mod ModuleInfoResponse
	mod.ModuleId = 1
	authInfoResponse.Modules = append(authInfoResponse.Modules, mod)

	c.JSON(http.StatusOK, authInfoResponse)
}

func SaveSubscriber(c *gin.Context) {
	user := common.Session_GetUser(c)
	if user == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"})
		return
	}

	if user.HasAccess(4, common.CreateAccess) == false {
		c.IndentedJSON(http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"})
		return
	}

	var new_FSubscriber xfgc.FSubscriber

	if err := c.BindJSON(&new_FSubscriber); err != nil {
		common.AppLogError("BindJSON failed for addFSubscriber")
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo": 5000, "ErrorName": err.Error()})
		return
	}

	//validation

	if new_FSubscriber.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "PLMN is required"})
		return
	}

	if new_FSubscriber.SKey == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "SKey is required"})
		return
	}

	if new_FSubscriber.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "OPc is required"})
		return
	}

	if new_FSubscriber.DefSD == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "DefSD is required"})
		return
	}

	if new_FSubscriber.MSISDN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "MSISDN is required"})
		return
	}

	uList := xfgc.FSubscriber_SelectByIMSI(0, &new_FSubscriber)

	if uList != nil {
		if uList.TotalCount == 1 {
			c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "IMSI Already Exists"})
			return
		}
	}

	//set custom feild values
	new_FSubscriber.AMF = "02"
	new_FSubscriber.SEQ = "0"
	new_FSubscriber.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
	new_FSubscriber.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
	new_FSubscriber.CreatedBy = user.UserId
	new_FSubscriber.UpdatedBy = user.UserId
	new_FSubscriber.IsActive = 1
	new_FSubscriber.RowVersion = 1

	xfgc.FSubscriber_Add(0, &new_FSubscriber)
	c.IndentedJSON(http.StatusOK, gin.H{"SubscriberId": new_FSubscriber.SubscriberId, "ErrorNo": new_FSubscriber.ErrorNo, "ErrorName": new_FSubscriber.ErrorName})
}

func SaveImsiRange(c *gin.Context) {
	user := common.Session_GetUser(c)
	if user == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"ErrorNo": 4998, "ErrorName": "UnAuthorized"})
		return
	}

	if user.HasAccess(4, common.CreateAccess) == false {
		c.IndentedJSON(http.StatusOK, gin.H{"ErrorNo": 4997, "ErrorName": "No Access"})
		return
	}

	var new_IMSIRange xfgc.IMSIRange

	if err := c.BindJSON(&new_IMSIRange); err != nil {
		common.AppLogError("BindJSON failed for addIMSIRange")
		c.JSON(http.StatusBadRequest, gin.H{"ErrorNo": 5000, "ErrorName": err.Error()})
		return
	}

	//validation

	if new_IMSIRange.SubK == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "SubK is required"})
		return
	}

	if new_IMSIRange.OPc == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "OPc is required"})
		return
	}

	if new_IMSIRange.PLMN == "" {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "PLMN is required"})
		return
	}

	if new_IMSIRange.IMSIBegin > new_IMSIRange.IMSIEnd {
		c.JSON(http.StatusOK, gin.H{"ErrorNo": 6000, "ErrorName": "IMSI Start No should be less than End No"})
		return
	}

	//set custom feild values
	new_IMSIRange.IsActive = 1
	new_IMSIRange.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
	new_IMSIRange.CreatedUserId = user.UserId

	xfgc.IMSIRange_Add(0, &new_IMSIRange)

	var imsiExists bool = false

	for i := new_IMSIRange.IMSIBegin; i <= new_IMSIRange.IMSIEnd; i++ {

		var new_FSubscriber xfgc.FSubscriber

		new_FSubscriber.IMSI = i
		new_FSubscriber.PLMN = new_IMSIRange.PLMN
		new_FSubscriber.SKey = new_IMSIRange.SubK
		new_FSubscriber.OPc = new_IMSIRange.OPc
		new_FSubscriber.AccountStatus = 1
		new_FSubscriber.DefSST = 1
		new_FSubscriber.DefSD = "123456"
		new_FSubscriber.Uplink = 100000
		new_FSubscriber.Downlink = 500000
		new_FSubscriber.DnnId1 = 1
		new_FSubscriber.DnnId2 = 1
		new_FSubscriber.DnnId3 = 1
		new_FSubscriber.MSISDN = strconv.Itoa(i)

		if len(new_FSubscriber.MSISDN) == 15 {
			new_FSubscriber.MSISDN = new_FSubscriber.MSISDN[5:15]
		}

		new_FSubscriber.MCC = new_IMSIRange.MCC
		new_FSubscriber.MNC = new_IMSIRange.MNC

		new_FSubscriber.AMF = "02"
		new_FSubscriber.SEQ = "0"
		new_FSubscriber.CreatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
		new_FSubscriber.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
		new_FSubscriber.CreatedBy = user.UserId
		new_FSubscriber.UpdatedBy = user.UserId
		new_FSubscriber.IsActive = 1
		new_FSubscriber.RowVersion = 1

		uList := xfgc.FSubscriber_SelectByIMSI(0, &new_FSubscriber)

		imsiExists = false

		if uList != nil {
			if uList.TotalCount == 1 {
				imsiExists = true
			}
		}

		fmt.Printf("imsiExists=%+v\n", imsiExists)
		fmt.Printf("uList=%+v\n\n", uList)

		if imsiExists == false {
			xfgc.FSubscriber_Add(0, &new_FSubscriber)
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"IMSIRangeId": new_IMSIRange.IMSIRangeId, "ErrorNo": new_IMSIRange.ErrorNo, "ErrorName": new_IMSIRange.ErrorName})
}

func InitModule_Auth(router *gin.Engine) {
	router.POST("/login/authinfo", AuthInfo)
	router.POST("/application/savesub", SaveSubscriber)
	router.POST("/application/saveir", SaveImsiRange)
}

//rohini-lpu
