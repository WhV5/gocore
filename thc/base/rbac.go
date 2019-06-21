package base

import "github.com/pssauron/gocore/libs"

type BaseDTO struct {
	CreateBy libs.Integer `xorm:"create_by" json:"create_by"` //创建人

	ModifyBy libs.Integer `xorm:"modify_by" json:"modify_by"` //修改人

	CreateAt libs.Time `xorm:"create_at" json:"create_at"` //创建日期

	ModifyAt libs.Time `xorm:"modify_at" json:"modify_at"` //修改日期

}

//RBACSite 站点
type RBACSite struct {
	ID       libs.Integer `xorm:"id" json:"id"`
	SiteName libs.String  `xorm:"site_name" json:"site_name"` //站点名称
	Status   libs.String  `xorm:"status" json:"status"`       //状态：0：停用,1:启用
	Memo     libs.String  `xorm:"memo" json:"memo"`
	BaseDTO
}

//RBACRole 角色表
type RBACRole struct {
	ID       libs.Integer `xorm:"id" json:"id"`               // 主键
	SiteID   libs.Integer `xorm:"site_id" json:"site_id"`     //所属站点
	RoleName libs.String  `xorm:"role_name" json:"role_name"` //角色名称
	Status   libs.String  `xorm:"status" json:"status"`       //状态：0 停用，1,启用
	Memo     libs.String  `xorm:"memo" json:"memo"`           //角色描述
	BaseDTO
}

//RBACMenu 页面
type RBACMenu struct {
	ID           libs.Integer `xorm:"id" json:"id"`                       //
	SiteID       libs.Integer `xorm:"site_id" json:"site_id"`             // 所属站点
	ResourceName libs.String  `xorm:"resource_name" json:"resource_name"` //资源名称
	Status       libs.String  `xorm:"status" json:"status"`               //资源状态 0 停用，1启用
	URL          libs.String  `xorm:"url" json:"url"`                     //地址
	IsFolder     libs.Boolean `xorm:"is_folder" json:"is_folder"`         //是否是目录
	IsCtrl       libs.Boolean `xorm:"is_ctrl" json:"is_ctrl"`             //需要权限
	PID          libs.String  `xorm:"pid" json:"pid"`                     //父目录ID
	BaseDTO
}

//RBACMenu 按钮操作
type RBACOperator struct {
	ID     libs.Integer `xorm:"id" json:"id"`
	MenuID libs.Integer `xorm:"menu_id" json:"menu_id"` //菜单ID
	OpCode libs.String  `xorm:"op_code" json:"op_code"` //操作编码 `query,get,save update other`
	OpName libs.String  `xorm:"op_name" json:"op_name"` //操作名称
	API    libs.String  `xorm:"api" json:"api"`         //操作API //每个资源对于一个API 操作，有 其他操作必须有查询权限
	Memo   libs.String  `xorm:"memo" json:"memo"`       //描述
	BaseDTO
}

type RBACGroupLimit struct {
	ID     libs.Integer `xorm:"id" json:"id"`
	RoleID libs.Integer `xorm:"role_id" json:"role_id"` //角色ID
	OpID   libs.Integer `xorm:"op_id" json:"op_id"`     //操作ID

}
