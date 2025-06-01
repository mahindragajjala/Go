package common

type ModuleInfo struct {
	ModuleId 	int;
	View 		bool;
	Create 		bool;
	Update 		bool;
	Delete 		bool;
	List 		bool;
}

type UserInfo struct {
	UserId int;
	UserTypeId int;
	RoleId []int;
	OrganizationId int;				//Current Viewing
	ParentOrganizationId int;		//Parent OrgId
	OrganizationIds []int;			//User Having Access to List
	DisplayName string;
	LastLogin string;
	Modules []ModuleInfo
}

type AccessType int

const (
	ViewAccess 		AccessType = 1
	CreateAccess	AccessType = 2
	UpdateAccess	AccessType = 3
	DeleteAccess	AccessType = 4
	ListAccess		AccessType = 5
)

func (user * UserInfo) GetModuleInfo( moduleId int) * ModuleInfo {
	
	/*
	for i := 0; i < len(user.Modules); i++ {
		if user.Modules[i].ModuleId == moduleId {
			return &user.Modules[i];
		}
	}
	*/
	
	return nil;
}

func (user * UserInfo) HasAccess( ModuleId int, accessType AccessType) bool {
	
	return true;
	
	
	module :=  user.GetModuleInfo( ModuleId)
	
	if module != nil {
		switch accessType {
			case ViewAccess:
				return module.View;
			case CreateAccess:
				return module.Create;
			case UpdateAccess:
				return module.Update;
			case DeleteAccess:
				return module.Delete;
			case ListAccess:
				return module.List;
			default:
				return false;
		}
	}
	
	return false;
}




