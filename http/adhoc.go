package http

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/upamune/radicaster/radikoutil"
	"github.com/upamune/radicaster/record"
	"github.com/upamune/radicaster/timeutil"
	"github.com/yyoshiki41/go-radiko"
)

var strictPolicy = bluemonday.StrictPolicy()

// handlePrograms は番組表UIを表示
func handlePrograms(
	radikoEmail, radikoPassword string,
	programCache *programCacheStore,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		areaID := c.QueryParam("area_id")
		dateStr := c.QueryParam("date")

		var targetDate time.Time
		if dateStr == "" {
			targetDate = time.Now().In(timeutil.JST())
		} else {
			var err error
			targetDate, err = time.Parse("2006-01-02", dateStr)
			if err != nil {
				return c.String(http.StatusBadRequest, "Invalid date format")
			}
		}

		data := map[string]interface{}{
			"CurrentDate": targetDate.Format("2006-01-02"),
			"AreaID":      areaID,
			"IsPremium":   radikoEmail != "",
		}

		return c.Render(http.StatusOK, "programs.html.tmpl", data)
	}
}

// handleGetPrograms は番組表データを取得（JSONまたはHTML fragment）
func handleGetPrograms(
	radikoEmail, radikoPassword string,
	programCache *programCacheStore,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		areaID := c.QueryParam("area_id")
		dateStr := c.QueryParam("date")

		var targetDate time.Time
		if dateStr == "" {
			targetDate = time.Now().In(timeutil.JST())
		} else {
			var err error
			targetDate, err = time.Parse("2006-01-02", dateStr)
			if err != nil {
				return c.String(http.StatusBadRequest, "Invalid date format")
			}
		}

		ctx := c.Request().Context()

		// キャッシュキー生成
		cacheKey := fmt.Sprintf("%s:%s", areaID, targetDate.Format("2006-01-02"))
		stations, ok := programCache.Get(cacheKey)

		if !ok {
			// Radikoクライアント初期化
			client, err := radikoutil.NewClient(
				ctx,
				radikoutil.WithAreaID(areaID),
				radikoutil.WithPremium(radikoEmail, radikoPassword),
			)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			stations, err = client.GetStations(ctx, targetDate)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			programCache.Set(cacheKey, stations)
		}

		acceptHeader := c.Request().Header.Get("Accept")
		if acceptHeader == "application/json" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"stations": stations,
			})
		}

		// HTML fragment を返す（htmx用）
		return renderProgramTable(c, stations, areaID, targetDate)
	}
}

// handleAdHocRecord はアドホック録音を開始
func handleAdHocRecord(recorder *record.Recorder) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			StationID string `json:"station_id"`
			From      string `json:"from"`
			AreaID    string `json:"area_id"`
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		fromTime, err := time.Parse("20060102150405", req.From)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": fmt.Sprintf("invalid time format: %v", err),
			})
		}

		taskID, err := recorder.RecordAdHoc(
			c.Request().Context(),
			req.StationID,
			fromTime,
			req.AreaID,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"task_id": taskID,
			"status":  "pending",
		})
	}
}

// handleAdHocStatus は録音状態を取得
func handleAdHocStatus(recorder *record.Recorder) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskIDsParam := c.QueryParam("task_ids")
		var taskIDs []string
		if taskIDsParam != "" {
			taskIDs = strings.Split(taskIDsParam, ",")
		}

		tasks := recorder.GetAdHocTaskStatus(taskIDs)

		acceptHeader := c.Request().Header.Get("Accept")
		if acceptHeader == "application/json" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"tasks": tasks,
			})
		}

		// HTML fragment を返す（htmx用）
		return renderTaskList(c, tasks)
	}
}

// renderProgramTable は番組表のHTML fragmentを生成
func renderProgramTable(c echo.Context, stations radiko.Stations, areaID string, targetDate time.Time) error {
	tmpl := `
<div class="overflow-x-auto">
	<table class="min-w-full bg-white border border-gray-300">
		<thead class="bg-gray-100">
			<tr>
				<th class="px-4 py-2 border">ステーション</th>
				<th class="px-4 py-2 border">番組</th>
				<th class="px-4 py-2 border">時間</th>
				<th class="px-4 py-2 border">操作</th>
			</tr>
		</thead>
		<tbody>
			{{range $station := .Stations}}
				{{range .Progs.Progs}}
				<tr class="hover:bg-gray-50 program-row" data-title="{{.Title}}" data-station="{{$station.StationName}}">
					<td class="px-4 py-2 border">{{$station.StationName}}</td>
					<td class="px-4 py-2 border">
						<div class="font-semibold">{{.Title}}</div>
						<div class="text-sm text-gray-600">{{stripHTML .Desc}}</div>
					</td>
					<td class="px-4 py-2 border text-sm">
						{{formatTime .Ft}} - {{formatTime .To}}
					</td>
					<td class="px-4 py-2 border text-center">
						<button
							hx-post="/api/record/adhoc"
							hx-vals='{"station_id": "{{$station.StationID}}", "from": "{{.Ft}}", "area_id": "{{$station.AreaID}}"}'
							hx-swap="none"
							hx-on::after-request="htmx.trigger('#task-list', 'taskCreated')"
							class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-3 rounded text-sm">
							録音
						</button>
					</td>
				</tr>
				{{end}}
			{{end}}
		</tbody>
	</table>
</div>
`

	funcMap := template.FuncMap{
		"formatTime": func(timeStr string) string {
			// YYYYMMDDhhmmss -> HH:MM
			if len(timeStr) >= 12 {
				hour := timeStr[8:10]
				minute := timeStr[10:12]
				return hour + ":" + minute
			}
			return timeStr
		},
		"stripHTML": func(s string) string {
			return strictPolicy.Sanitize(s)
		},
	}

	t, err := template.New("program-table").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return err
	}

	type templateData struct {
		Stations []struct {
			StationID   string
			StationName string
			Progs       radiko.Progs
			AreaID      string
		}
	}

	data := templateData{}
	for _, station := range stations {
		data.Stations = append(data.Stations, struct {
			StationID   string
			StationName string
			Progs       radiko.Progs
			AreaID      string
		}{
			StationID:   station.ID,
			StationName: station.Name,
			Progs:       station.Progs,
			AreaID:      areaID,
		})
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, buf.String())
}

// renderTaskList はタスク一覧のHTML fragmentを生成
func renderTaskList(c echo.Context, tasks []*record.AdHocTask) error {
	tmpl := `
<div class="space-y-2">
	{{range .Tasks}}
	<div class="bg-white border border-gray-300 rounded-lg p-4">
		<div class="flex justify-between items-start">
			<div class="flex-1">
				<div class="font-semibold">{{.StationID}} - {{.From.Format "2006-01-02 15:04"}}</div>
				<div class="text-sm text-gray-600 mt-1">タスクID: {{.ID}}</div>
				{{if ne .FilePath ""}}
				<div class="text-sm text-gray-600 mt-1">ファイル: {{.FilePath}}</div>
				{{end}}
				{{if ne .Error ""}}
				<div class="text-sm text-red-600 mt-1">エラー: {{.Error}}</div>
				{{end}}
			</div>
			<div class="ml-4">
				{{if eq .Status "pending"}}
				<span class="inline-block px-3 py-1 text-sm font-semibold text-gray-700 bg-gray-200 rounded-full">待機中</span>
				{{else if eq .Status "recording"}}
				<span class="inline-block px-3 py-1 text-sm font-semibold text-blue-700 bg-blue-100 rounded-full">録音中</span>
				{{else if eq .Status "completed"}}
				<span class="inline-block px-3 py-1 text-sm font-semibold text-green-700 bg-green-100 rounded-full">完了</span>
				{{else if eq .Status "failed"}}
				<span class="inline-block px-3 py-1 text-sm font-semibold text-red-700 bg-red-100 rounded-full">失敗</span>
				{{end}}
			</div>
		</div>
	</div>
	{{else}}
	<div class="text-center text-gray-500 py-8">録音タスクはありません</div>
	{{end}}
</div>
`

	t, err := template.New("task-list").Parse(tmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]interface{}{"Tasks": tasks}); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, buf.String())
}

// programCacheStore は番組表のキャッシュ
type programCacheStore struct {
	cache map[string]cacheEntry
	ttl   time.Duration
}

type cacheEntry struct {
	data      radiko.Stations
	expiresAt time.Time
}

func newProgramCacheStore(ttl time.Duration) *programCacheStore {
	return &programCacheStore{
		cache: make(map[string]cacheEntry),
		ttl:   ttl,
	}
}

func (s *programCacheStore) Get(key string) (radiko.Stations, bool) {
	entry, ok := s.cache[key]
	if !ok || time.Now().After(entry.expiresAt) {
		return radiko.Stations{}, false
	}
	return entry.data, true
}

func (s *programCacheStore) Set(key string, data radiko.Stations) {
	s.cache[key] = cacheEntry{
		data:      data,
		expiresAt: time.Now().Add(s.ttl),
	}
}
