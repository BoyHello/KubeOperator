package constant

const (
	LOGIN  = "登录|Login"
	LOGOUT = "退出|Logout"

	// 项目
	CREATE_PROJECT      = "创建项目|Create project"
	UPDATE_PROJECT_INFO = "更新项目信息|Update project information"
	DELETE_PROJECT      = "删除项目|Delete project"

	BIND_CLUSTER_MEMBER            = "绑定集群成员|Binding cluster members"
	UNBIND_CLUSTER_MEMBER          = "解绑集群成员|Unbind cluster members"
	UPDATE_CLUSTER_MEMBER_ROLE     = "更新集群成员权限|Update cluster member permissions"
	BIND_CLUSTER_RESOURCE_PLAN     = "绑定集群资源(部署计划)|Binding cluster resources(plan)"
	BIND_CLUSTER_RESOURCE_BACKUP   = "绑定集群资源(备份账号)|Binding cluster resources(backup_account)"
	BIND_CLUSTER_RESOURCE_HOST     = "绑定集群资源(主机)|Binding cluster resources(host)"
	UNBIND_CLUSTER_RESOURCE_PLAN   = "解绑集群资源(部署计划)|Unbind cluster resources(plan)"
	UNBIND_CLUSTER_RESOURCE_BACKUP = "解绑集群资源(备份账号)|Unbind cluster resources(backup_account)"
	UNBIND_CLUSTER_RESOURCE_HOST   = "解绑集群资源(主机)|Unbind cluster resources(host)"

	BIND_PROJECT_MEMBER            = "绑定项目成员|Binding project members"
	UNBIND_PROJECT_MEMBER          = "解绑项目成员|Unbind project members"
	UPDATE_PROJECT_MEMBER_ROLE     = "更新项目成员权限|Update project member permissions"
	BIND_PROJECT_RESOURCE_PLAN     = "绑定项目资源(部署计划)|Binding project resources(plan)"
	BIND_PROJECT_RESOURCE_BACKUP   = "绑定项目资源(备份账号)|Binding project resources(backup_account)"
	BIND_PROJECT_RESOURCE_HOST     = "绑定项目资源(主机)|Binding project resources(host)"
	UNBIND_PROJECT_RESOURCE_PLAN   = "解绑项目资源(部署计划)|Unbind project resources(plan)"
	UNBIND_PROJECT_RESOURCE_BACKUP = "解绑项目资源(备份账号)|Unbind project resources(backup_account)"
	UNBIND_PROJECT_RESOURCE_HOST   = "解绑项目资源(主机)|Unbind project resources(host)"

	// 集群
	CREATE_CLUSTER  = "添加集群|Create cluster"
	IMPORT_CLUSTER  = "导入集群|Import cluster"
	INIT_CLUSTER    = "初始化集群|Init cluster"
	DELETE_CLUSTER  = "删除集群|Delete cluster"
	UPGRADE_CLUSTER = "集群升级|Upgrade cluster"
	HEALTH_CHECK    = "集群健康检查|Health check"
	HEALTH_RECOVER  = "集群健康恢复|Health recover"

	CREATE_CLUSTER_NODE = "添加集群节点|Create cluster node"
	DELETE_CLUSTER_NODE = "删除集群节点|Delete cluster node"

	CREATE_CLUSTER_STORAGE_SUPPLIER = "添加集群存储供应商|Create cluster storage vendor"
	DELETE_CLUSTER_STORAGE_SUPPLIER = "删除集群存储供应商|Delete cluster storage vendor"
	SYNC_CLUSTER_STORAGE_SUPPLIER   = "同步集群存储供应商|Sync cluster storage vendor"
	CREATE_CLUSTER_PVC              = "添加集群持久卷|Create cluster pvc"
	DELETE_CLUSTER_PVC              = "删除集群持久卷|Delete cluster pvc"

	ENABLE_CLUSTER_NPD  = "启用NPD|Enable cluster NPD"
	DISABLE_CLUSTER_NPD = "关闭NPD|Disable cluster NPD"

	ENABLE_CLUSTER_TOOL   = "启用集群工具|Enable cluster tools"
	UPGRADE_CLUSTER_TOOL  = "升级集群工具|Upgrade cluster tools"
	DISABLE_CLUSTER_TOOL  = "禁用集群工具|Disable cluster tools"
	ENABLE_CLUSTER_ISTIO  = "启用/修改集群 Istio|Enable/Update cluster Istio"
	DISABLE_CLUSTER_ISTIO = "禁用集群 Istio|Disable cluster Istio"

	CREATE_CLUSTER_STORAGE_CLASS   = "添加存储类|Create storage class"
	DELETE_CLUSTER_STORAGE_CLASS   = "删除存储类|Delete storage class"
	CREATE_CLUSTER_NAMESPACE       = "添加命名空间|Create cluster namespace"
	DELETE_CLUSTER_NAMESPACE       = "删除命名空间|Delete cluster namespace"
	CREATE_CLUSTER_BACKUP_STRATEGY = "添加集群备份策略|Create cluster backup strategy"
	START_CLUSTER_BACKUP           = "开始备份|Start cluster backup"
	UPLOAD_LOCAL_RECOVERY_FILE     = "上传本地恢复文件|Upload local recovery file"
	DELETE_RECOVERY_LIST           = "删除备份文件|Delete backup files"
	RECOVER_FROM_RECOVERY          = "从备份列表恢复|Restore from backup list"
	START_CLUSTER_CIS_SCAN         = "开始集群CIS扫描|Start cluster CIS scan"
	DELETE_CLUSTER_CIS_SCAN_RESULT = "删除集群CIS扫描结果|Delete cluster CIS scan results"

	// 主机
	CREATE_HOST    = "添加主机|Create host"
	SYNC_HOST_LIST = "主机同步|Sync host"
	DELETE_HOST    = "删除主机|Delete host"

	// 自动模式
	CREATE_REGION        = "添加区域|Create region"
	DELETE_REGION        = "删除区域|Delete region"
	CREATE_ZONE          = "添加可用区|Create zone"
	UPDATE_ZONE          = "修改可用区信息|Update zone information"
	DELETE_ZONE          = "删除可用区|Delete zone"
	CREATE_PLAN          = "添加部署计划|Create plan"
	DELETE_PLAN          = "删除部署计划|Delete plan"
	UPDATE_PLAN          = "更新部署计划|Update plan"
	CREATE_VM_CONFIG     = "添加虚拟机配置|Create virtual machine configuration"
	UPDATE_VM_CONFIG     = "修改虚拟机配置信息|Update virtual machine configuration information"
	DELETE_VM_CONFIG     = "删除虚拟机配置|Delete virtual machine configuration"
	CREATE_IP_POOL       = "添加IP池|Create Ip pool"
	BACTH_DELETE_IP_POOL = "批量删除IP池|Batch delete IP pool"
	DELETE_IP_POOL       = "删除IP池|Delete IP Pool"
	CREATE_IP            = "添加IP"
	DELETE_IP            = "删除Ip"

	// 用户
	CREATE_USER          = "添加用户|Create user"
	UPDATE_USER          = "修改用户信息|Update user information"
	UPDATE_USER_PASSWORD = "修改用户密码|Delete user password"
	DELETE_USER          = "删除用户|Delete user"
	FORGOT_USER_PASSWORD = "忘记密码|forgot password"

	// 版本
	ENABLE_VERSION  = "启用ko版本|Enable ko version"
	DISABLE_VERSION = "停用ko版本|Disable ko version"

	// 系统设置
	CREATE_CREDENTIALS    = "添加凭证|Create credentials"
	UPDATE_CREDENTIALS    = "修改凭证信息|Update credential information"
	DELETE_CREDENTIALS    = "删除凭证|Delete credentials"
	CREATE_REGISTRY       = "添加仓库信息|Create registry"
	UPDATE_REGISTRY       = "更新仓库信息|Delete registry"
	UPDATE_NEXUS_PASSWORD = "更新 Nexus 仓库密码|Update nexus password"
	DELETE_REGISTRY       = "删除仓库信息|Delete registry"
	CREATE_BACKUP_ACCOUNT = "添加备份账号|Create backup account"
	UPDATE_BACKUP_ACCOUNT = "修改备份账号信息|Update backup account information"
	DELETE_BACKUP_ACCOUNT = "删除备份账号|Delete backup account"
	CREATE_EMAIL          = "设置系统配置|Set system config"
	IMPORT_LICENCE        = "导入许可证书|import licence"

	CREATE_NTP = "添加 NTP|Create NTP server"
	UPDATE_NTP = "修改 NTP|Update NTP server"
	DELETE_NTP = "删除 NTP|DElete NTP server"
)
