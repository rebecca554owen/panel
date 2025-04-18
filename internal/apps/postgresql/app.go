package postgresql

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/tnb-labs/panel/internal/app"
	"github.com/tnb-labs/panel/internal/service"
	"github.com/tnb-labs/panel/pkg/io"
	"github.com/tnb-labs/panel/pkg/shell"
	"github.com/tnb-labs/panel/pkg/systemctl"
	"github.com/tnb-labs/panel/pkg/types"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (s *App) Route(r chi.Router) {
	r.Get("/config", s.GetConfig)
	r.Post("/config", s.UpdateConfig)
	r.Get("/userConfig", s.GetUserConfig)
	r.Post("/userConfig", s.UpdateUserConfig)
	r.Get("/load", s.Load)
	r.Get("/log", s.Log)
	r.Post("/clearLog", s.ClearLog)
}

// GetConfig 获取配置
func (s *App) GetConfig(w http.ResponseWriter, r *http.Request) {
	// 获取配置
	config, err := io.Read(fmt.Sprintf("%s/server/postgresql/data/postgresql.conf", app.Root))
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL配置失败")
		return
	}

	service.Success(w, config)
}

// UpdateConfig 保存配置
func (s *App) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	req, err := service.Bind[UpdateConfig](r)
	if err != nil {
		service.Error(w, http.StatusUnprocessableEntity, "%v", err)
		return
	}

	if err = io.Write(fmt.Sprintf("%s/server/postgresql/data/postgresql.conf", app.Root), req.Config, 0644); err != nil {
		service.Error(w, http.StatusInternalServerError, "写入PostgreSQL配置失败")
		return
	}

	if err = systemctl.Reload("postgresql"); err != nil {
		service.Error(w, http.StatusInternalServerError, "重载服务失败: %v", err)
		return
	}

	service.Success(w, nil)
}

// GetUserConfig 获取用户配置
func (s *App) GetUserConfig(w http.ResponseWriter, r *http.Request) {
	// 获取配置
	config, err := io.Read(fmt.Sprintf("%s/server/postgresql/data/pg_hba.conf", app.Root))
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL配置失败")
		return
	}

	service.Success(w, config)
}

// UpdateUserConfig 保存用户配置
func (s *App) UpdateUserConfig(w http.ResponseWriter, r *http.Request) {
	req, err := service.Bind[UpdateConfig](r)
	if err != nil {
		service.Error(w, http.StatusUnprocessableEntity, "%v", err)
		return
	}

	if err = io.Write(fmt.Sprintf("%s/server/postgresql/data/pg_hba.conf", app.Root), req.Config, 0644); err != nil {
		service.Error(w, http.StatusInternalServerError, "写入PostgreSQL配置失败")
		return
	}

	if err = systemctl.Reload("postgresql"); err != nil {
		service.Error(w, http.StatusInternalServerError, "重载服务失败: %v", err)
		return
	}

	service.Success(w, nil)
}

// Load 获取负载
func (s *App) Load(w http.ResponseWriter, r *http.Request) {
	status, _ := systemctl.Status("postgresql")
	if !status {
		service.Success(w, []types.NV{})
		return
	}

	start, err := shell.Execf(`echo "select pg_postmaster_start_time();" | su - postgres -c "psql" | sed -n 3p | cut -d'.' -f1`)
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL启动时间失败")
		return
	}
	pid, err := shell.Execf(`echo "select pg_backend_pid();" | su - postgres -c "psql" | sed -n 3p`)
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL进程PID失败")
		return
	}
	process, err := shell.Execf(`ps aux | grep postgres | grep -v grep | wc -l`)
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL进程数失败")
		return
	}
	connections, err := shell.Execf(`echo "SELECT count(*) FROM pg_stat_activity WHERE NOT pid=pg_backend_pid();" | su - postgres -c "psql" | sed -n 3p`)
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL连接数失败")
		return
	}
	storage, err := shell.Execf(`echo "select pg_size_pretty(pg_database_size('postgres'));" | su - postgres -c "psql" | sed -n 3p`)
	if err != nil {
		service.Error(w, http.StatusInternalServerError, "获取PostgreSQL空间占用失败")
		return
	}

	data := []types.NV{
		{Name: "启动时间", Value: start},
		{Name: "进程 PID", Value: pid},
		{Name: "进程数", Value: process},
		{Name: "总连接数", Value: connections},
		{Name: "空间占用", Value: storage},
	}

	service.Success(w, data)
}

// Log 获取日志
func (s *App) Log(w http.ResponseWriter, r *http.Request) {
	service.Success(w, fmt.Sprintf("%s/server/postgresql/logs/postgresql-%s.log", app.Root, time.Now().Format(time.DateOnly)))
}

// ClearLog 清空日志
func (s *App) ClearLog(w http.ResponseWriter, r *http.Request) {
	if _, err := shell.Execf("rm -rf %s/server/postgresql/logs/postgresql-*.log", app.Root); err != nil {
		service.Error(w, http.StatusInternalServerError, "%v", err)
		return
	}

	service.Success(w, nil)
}
