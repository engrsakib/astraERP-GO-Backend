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

    // ➤ Permission Management
    PermissionAssign PermissionSlug = "permission.assign"

    // ➤ FAQ Management
    FaqCreate PermissionSlug = "faq.create"
    FaqUpdate PermissionSlug = "faq.update"
    FaqDelete PermissionSlug = "faq.delete"
    FaqView   PermissionSlug = "faq.view"

    // ➤ Group Type Management
    GroupTypeCreate PermissionSlug = "group_type.create"
    GroupTypeUpdate PermissionSlug = "group_type.update"
    GroupTypeDelete PermissionSlug = "group_type.delete"
    GroupTypeView   PermissionSlug = "group_type.view"

    // ➤ Group Management
    GroupCreate PermissionSlug = "group.create"
    GroupUpdate PermissionSlug = "group.update"
    GroupDelete PermissionSlug = "group.delete"
    GroupView   PermissionSlug = "group.view"

    // ➤ Member Management
    MemberCreate PermissionSlug = "member.create"
    MemberUpdate PermissionSlug = "member.update"
    MemberDelete PermissionSlug = "member.delete"
    MemberView   PermissionSlug = "member.view"
)

var AllPermissions = []struct {
    Name string
    Slug PermissionSlug
}{
    // User
    {Name: "Create User", Slug: UserCreate},
    {Name: "Update User", Slug: UserUpdate},
    {Name: "Delete User", Slug: UserDelete},
    {Name: "View User List", Slug: UserView},

    // Role
    {Name: "Create Role", Slug: RoleCreate},
    {Name: "Update Role", Slug: RoleUpdate},
    {Name: "Delete Role", Slug: RoleDelete},
    {Name: "View Roles", Slug: RoleView},

    // Permission Assign
    {Name: "Assign Permissions", Slug: PermissionAssign},

    // FAQ
    {Name: "Create FAQ", Slug: FaqCreate},
    {Name: "Update FAQ", Slug: FaqUpdate},
    {Name: "Delete FAQ", Slug: FaqDelete},
    {Name: "View FAQ List", Slug: FaqView},

    // Group Type
    {Name: "Create Group Type", Slug: GroupTypeCreate},
    {Name: "Update Group Type", Slug: GroupTypeUpdate},
    {Name: "Delete Group Type", Slug: GroupTypeDelete},
    {Name: "View Group Types", Slug: GroupTypeView},

    // Group
    {Name: "Create Group", Slug: GroupCreate},
    {Name: "Update Group", Slug: GroupUpdate},
    {Name: "Delete Group", Slug: GroupDelete},
    {Name: "View Groups", Slug: GroupView},

    // Member
    {Name: "Create Member", Slug: MemberCreate},
    {Name: "Update Member", Slug: MemberUpdate},
    {Name: "Delete Member", Slug: MemberDelete},
    {Name: "View Members", Slug: MemberView},
}
