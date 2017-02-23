package admin

var IndexHeadersHandler = AdminHeadersHandler

var IndexPageHandler = NewAdminPageHandler("page/admin/index")

var ArticlesHeadersHandler = AdminHeadersHandler

var ArticlesPageHandler = NewAdminPageHandler("page/admin/article")

var CategoryHeadersHandler = AdminHeadersHandler

var CategoryPageHandler = NewAdminPageHandler("page/admin/category")

var SiteInfoHeadersHandler = AdminHeadersHandler

var SiteInfoPageHandler = NewAdminPageHandler("page/admin/site-info")

var AccountHeadersHandler = AdminHeadersHandler

var AccountPageHandler = NewAdminPageHandler("page/admin/account")
