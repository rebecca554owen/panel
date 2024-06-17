package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"github.com/TheTNB/panel/app/http/controllers"
	"github.com/TheTNB/panel/app/http/middleware"
)

func Api() {
	facades.Route().Prefix("api/panel").Group(func(r route.Router) {
		r.Prefix("info").Group(func(r route.Router) {
			infoController := controllers.NewInfoController()
			r.Get("panel", infoController.Panel)
			r.Middleware(middleware.Jwt()).Get("homePlugins", infoController.HomePlugins)
			r.Middleware(middleware.Jwt()).Get("nowMonitor", infoController.NowMonitor)
			r.Middleware(middleware.Jwt()).Get("systemInfo", infoController.SystemInfo)
			r.Middleware(middleware.Jwt()).Get("countInfo", infoController.CountInfo)
			r.Middleware(middleware.Jwt()).Get("installedDbAndPhp", infoController.InstalledDbAndPhp)
			r.Middleware(middleware.Jwt()).Get("checkUpdate", infoController.CheckUpdate)
			r.Middleware(middleware.Jwt()).Get("updateInfo", infoController.UpdateInfo)
			r.Middleware(middleware.Jwt()).Post("update", infoController.Update)
			r.Middleware(middleware.Jwt()).Post("restart", infoController.Restart)
		})
		r.Prefix("user").Group(func(r route.Router) {
			userController := controllers.NewUserController()
			r.Post("login", userController.Login)
			r.Middleware(middleware.Jwt()).Get("info", userController.Info)
		})
		r.Prefix("task").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			taskController := controllers.NewTaskController()
			r.Get("status", taskController.Status)
			r.Get("list", taskController.List)
			r.Get("log", taskController.Log)
			r.Post("delete", taskController.Delete)
		})
		r.Prefix("website").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			websiteController := controllers.NewWebsiteController()
			r.Get("defaultConfig", websiteController.GetDefaultConfig)
			r.Post("defaultConfig", websiteController.SaveDefaultConfig)
			r.Get("backupList", websiteController.BackupList)
			r.Put("uploadBackup", websiteController.UploadBackup)
			r.Delete("deleteBackup", websiteController.DeleteBackup)
		})
		r.Prefix("websites").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			websiteController := controllers.NewWebsiteController()
			r.Get("/", websiteController.List)
			r.Post("/", websiteController.Add)
			r.Delete("{id}", websiteController.Delete)
			r.Get("{id}/config", websiteController.GetConfig)
			r.Post("{id}/config", websiteController.SaveConfig)
			r.Delete("{id}/log", websiteController.ClearLog)
			r.Post("{id}/updateRemark", websiteController.UpdateRemark)
			r.Post("{id}/createBackup", websiteController.CreateBackup)
			r.Post("{id}/restoreBackup", websiteController.RestoreBackup)
			r.Post("{id}/resetConfig", websiteController.ResetConfig)
			r.Post("{id}/status", websiteController.Status)
		})
		r.Prefix("cert").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			certController := controllers.NewCertController()
			r.Get("caProviders", certController.CAProviders)
			r.Get("dnsProviders", certController.DNSProviders)
			r.Get("algorithms", certController.Algorithms)
			r.Get("users", certController.UserList)
			r.Post("users", certController.UserStore)
			r.Put("users/{id}", certController.UserUpdate)
			r.Get("users/{id}", certController.UserShow)
			r.Delete("users/{id}", certController.UserDestroy)
			r.Get("dns", certController.DNSList)
			r.Post("dns", certController.DNSStore)
			r.Put("dns/{id}", certController.DNSUpdate)
			r.Get("dns/{id}", certController.DNSShow)
			r.Delete("dns/{id}", certController.DNSDestroy)
			r.Get("certs", certController.CertList)
			r.Post("certs", certController.CertStore)
			r.Put("certs/{id}", certController.CertUpdate)
			r.Get("certs/{id}", certController.CertShow)
			r.Delete("certs/{id}", certController.CertDestroy)
			r.Post("obtain", certController.Obtain)
			r.Post("renew", certController.Renew)
			r.Post("manualDNS", certController.ManualDNS)
			r.Post("deploy", certController.Deploy)
		})
		r.Prefix("plugin").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			pluginController := controllers.NewPluginController()
			r.Get("list", pluginController.List)
			r.Post("install", pluginController.Install)
			r.Post("uninstall", pluginController.Uninstall)
			r.Post("update", pluginController.Update)
			r.Post("updateShow", pluginController.UpdateShow)
			r.Get("isInstalled", pluginController.IsInstalled)
		})
		r.Prefix("cron").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			cronController := controllers.NewCronController()
			r.Get("list", cronController.List)
			r.Get("{id}", cronController.Script)
			r.Post("add", cronController.Add)
			r.Put("{id}", cronController.Update)
			r.Delete("{id}", cronController.Delete)
			r.Post("status", cronController.Status)
			r.Get("log/{id}", cronController.Log)
		})
		r.Prefix("safe").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			safeController := controllers.NewSafeController()
			r.Get("firewallStatus", safeController.GetFirewallStatus)
			r.Post("firewallStatus", safeController.SetFirewallStatus)
			r.Get("firewallRules", safeController.GetFirewallRules)
			r.Post("firewallRules", safeController.AddFirewallRule)
			r.Delete("firewallRules", safeController.DeleteFirewallRule)
			r.Get("sshStatus", safeController.GetSshStatus)
			r.Post("sshStatus", safeController.SetSshStatus)
			r.Get("sshPort", safeController.GetSshPort)
			r.Post("sshPort", safeController.SetSshPort)
			r.Get("pingStatus", safeController.GetPingStatus)
			r.Post("pingStatus", safeController.SetPingStatus)
		})
		r.Prefix("container").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			containerController := controllers.NewContainerController()
			r.Get("list", containerController.ContainerList)
			r.Get("search", containerController.ContainerSearch)
			r.Post("create", containerController.ContainerCreate)
			r.Post("remove", containerController.ContainerRemove)
			r.Post("start", containerController.ContainerStart)
			r.Post("stop", containerController.ContainerStop)
			r.Post("restart", containerController.ContainerRestart)
			r.Post("pause", containerController.ContainerPause)
			r.Post("unpause", containerController.ContainerUnpause)
			r.Get("inspect", containerController.ContainerInspect)
			r.Post("kill", containerController.ContainerKill)
			r.Post("rename", containerController.ContainerRename)
			r.Get("stats", containerController.ContainerStats)
			r.Get("exist", containerController.ContainerExist)
			r.Get("logs", containerController.ContainerLogs)
			r.Post("prune", containerController.ContainerPrune)

			r.Prefix("network").Group(func(r route.Router) {
				r.Get("list", containerController.NetworkList)
				r.Post("create", containerController.NetworkCreate)
				r.Post("remove", containerController.NetworkRemove)
				r.Get("exist", containerController.NetworkExist)
				r.Get("inspect", containerController.NetworkInspect)
				r.Post("connect", containerController.NetworkConnect)
				r.Post("disconnect", containerController.NetworkDisconnect)
				r.Post("prune", containerController.NetworkPrune)
			})

			r.Prefix("image").Group(func(r route.Router) {
				r.Get("list", containerController.ImageList)
				r.Get("exist", containerController.ImageExist)
				r.Post("pull", containerController.ImagePull)
				r.Post("remove", containerController.ImageRemove)
				r.Get("inspect", containerController.ImageInspect)
				r.Post("prune", containerController.ImagePrune)
			})

			r.Prefix("volume").Group(func(r route.Router) {
				r.Get("list", containerController.VolumeList)
				r.Post("create", containerController.VolumeCreate)
				r.Get("exist", containerController.VolumeExist)
				r.Get("inspect", containerController.VolumeInspect)
				r.Post("remove", containerController.VolumeRemove)
				r.Post("prune", containerController.VolumePrune)
			})
		})
		r.Prefix("file").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			fileController := controllers.NewFileController()
			r.Post("create", fileController.Create)
			r.Get("content", fileController.Content)
			r.Post("save", fileController.Save)
			r.Post("delete", fileController.Delete)
			r.Post("upload", fileController.Upload)
			r.Post("move", fileController.Move)
			r.Post("copy", fileController.Copy)
			r.Get("download", fileController.Download)
			r.Post("remoteDownload", fileController.RemoteDownload)
			r.Get("info", fileController.Info)
			r.Post("permission", fileController.Permission)
			r.Post("archive", fileController.Archive)
			r.Post("unArchive", fileController.UnArchive)
			r.Post("search", fileController.Search)
			r.Get("list", fileController.List)
		})
		r.Prefix("monitor").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			monitorController := controllers.NewMonitorController()
			r.Post("switch", monitorController.Switch)
			r.Post("saveDays", monitorController.SaveDays)
			r.Post("clear", monitorController.Clear)
			r.Get("list", monitorController.List)
			r.Get("switchAndDays", monitorController.SwitchAndDays)
		})
		r.Prefix("ssh").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			sshController := controllers.NewSshController()
			r.Get("info", sshController.GetInfo)
			r.Post("info", sshController.UpdateInfo)
			r.Get("session", sshController.Session)
		})
		r.Prefix("setting").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			settingController := controllers.NewSettingController()
			r.Get("list", settingController.List)
			r.Post("update", settingController.Update)
		})
		r.Prefix("system").Middleware(middleware.Jwt()).Group(func(r route.Router) {
			controller := controllers.NewSystemController()
			r.Get("service/status", controller.ServiceStatus)
			r.Get("service/isEnabled", controller.ServiceIsEnabled)
			r.Post("service/enable", controller.ServiceEnable)
			r.Post("service/disable", controller.ServiceDisable)
			r.Post("service/restart", controller.ServiceRestart)
			r.Post("service/reload", controller.ServiceReload)
			r.Post("service/start", controller.ServiceStart)
			r.Post("service/stop", controller.ServiceStop)
		})
	})

	// 文档
	swaggerController := controllers.NewSwaggerController()
	facades.Route().Get("swagger/*any", swaggerController.Index)

	// 静态文件
	entrance := facades.Config().GetString("http.entrance")
	if entrance == "/" {
		entrance = ""
	}
	assetController := controllers.NewAssetController()
	facades.Route().Get("favicon.png", assetController.Favicon)
	facades.Route().Get("robots.txt", assetController.Robots)
	facades.Route().Get(entrance+"/assets/{any}", assetController.Index)
	facades.Route().Get(entrance+"/loading/{any}", assetController.Index)
	facades.Route().Get(entrance+"/{any}", assetController.Index)
	facades.Route().Get(entrance+"/", assetController.Index)
	facades.Route().Fallback(assetController.NotFound)
}
