package constants


type PermissionSlug string

const (
	// ➤ User Management
	UserCreate PermissionSlug = "user.create"
	UserUpdate PermissionSlug = "user.update"
	UserDelete PermissionSlug = "user.delete"
	UserView   PermissionSlug = "user.view"

	// ➤ Role/Admin Management
	RoleCreate PermissionSlug = "role.create"
	RoleUpdate PermissionSlug = "role.update"
	RoleDelete PermissionSlug = "role.delete"
	RoleView   PermissionSlug = "role.view"

	// ➤ Product/Inventory (ERP Feature)
	ProductCreate PermissionSlug = "product.create"
	ProductUpdate PermissionSlug = "product.update"
	ProductDelete PermissionSlug = "product.delete"
	ProductView   PermissionSlug = "product.view"

	// ➤ Reports & Analytics
	ReportView   PermissionSlug = "report.view"
	ReportExport PermissionSlug = "report.export"

	// ➤ Settings
	SettingsView PermissionSlug = "settings.view"
	SettingsEdit PermissionSlug = "settings.edit"

	//  ➤ Permission Management
	PermissionAssign PermissionSlug = "permission.assign"
)


var AllPermissions = []struct {
	Name string         
	Slug PermissionSlug 
}{
	
	{Name: "Create User", Slug: UserCreate},
	{Name: "Update User", Slug: UserUpdate},
	{Name: "Delete User", Slug: UserDelete},
	{Name: "View User List", Slug: UserView},

	
	{Name: "Create Role", Slug: RoleCreate},
	{Name: "Update Role", Slug: RoleUpdate},
	{Name: "Delete Role", Slug: RoleDelete},
	{Name: "View Roles", Slug: RoleView},

	
	{Name: "Create Product", Slug: ProductCreate},
	{Name: "Update Product", Slug: ProductUpdate},
	{Name: "Delete Product", Slug: ProductDelete},
	{Name: "View Product List", Slug: ProductView},

	
	{Name: "View Reports", Slug: ReportView},
	{Name: "Export Reports", Slug: ReportExport},

	
	{Name: "View Settings", Slug: SettingsView},
	{Name: "Edit Settings", Slug: SettingsEdit},

	
	{Name: "Assign Permissions", Slug: PermissionAssign},
}