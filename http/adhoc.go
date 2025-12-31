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
		ctx := c.Request().Context()

		c.Logger().Infof("handleGetPrograms called with area_id='%s' (len=%d)", areaID, len(areaID))

		// 過去7日間の番組を取得
		now := time.Now().In(timeutil.JST())
		var allStationsWithDate []stationWithDate

		// エリアIDが空の場合、現在地を自動検出
		actualAreaID := areaID
		if areaID == "" {
			c.Logger().Infof("areaID is empty, detecting area...")
			client, err := radikoutil.NewClient(
				ctx,
				radikoutil.WithPremium(radikoEmail, radikoPassword),
			)
			if err != nil {
				c.Logger().Errorf("failed to create radiko client for area detection: %v", err)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error": "Failed to detect area",
				})
			}
			actualAreaID = client.AreaID()
			c.Logger().Infof("detected area ID: %s", actualAreaID)
		}

		for i := 0; i < 7; i++ {
			targetDate := now.AddDate(0, 0, -i)

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
					c.Logger().Errorf("failed to create radiko client for date %s: %v", targetDate.Format("2006-01-02"), err)
					continue
				}

				stations, err = client.GetStations(ctx, targetDate)
				if err != nil {
					c.Logger().Errorf("failed to get stations for date %s: %v", targetDate.Format("2006-01-02"), err)
					continue
				}

				programCache.Set(cacheKey, stations)
			}

			// 日付情報を付加
			for _, station := range stations {
				allStationsWithDate = append(allStationsWithDate, stationWithDate{
					Station: station,
					Date:    targetDate,
					AreaID:  actualAreaID,
				})
			}
		}

		c.Logger().Infof("fetched %d stations with dates, actualAreaID='%s'", len(allStationsWithDate), actualAreaID)

		acceptHeader := c.Request().Header.Get("Accept")
		if acceptHeader == "application/json" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"stations": allStationsWithDate,
			})
		}

		// HTML fragment を返す（htmx用）
		return renderProgramTable(c, allStationsWithDate, actualAreaID)
	}
}

type stationWithDate struct {
	Station radiko.Station
	Date    time.Time
	AreaID  string
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
			c.Logger().Errorf("failed to bind request: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		c.Logger().Infof("received adhoc request: station_id=%s, from=%s, area_id=%s", req.StationID, req.From, req.AreaID)

		// JSTタイムゾーンでパース
		fromTime, err := time.ParseInLocation("20060102150405", req.From, timeutil.JST())
		if err != nil {
			c.Logger().Errorf("failed to parse time %s: %v", req.From, err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": fmt.Sprintf("invalid time format: %v", err),
			})
		}

		if fromTime.After(time.Now()) {
			c.Logger().Warnf("attempt to record future program: %s", fromTime)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "未来の番組は録音できません。過去に放送された番組のみ録音可能です。",
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
func renderProgramTable(c echo.Context, stationsWithDate []stationWithDate, areaID string) error {
	tmpl := `
<div class="overflow-x-auto">
	<table class="min-w-full bg-white border border-gray-300">
		<thead class="bg-gray-100">
			<tr>
				<th class="px-4 py-2 border">日付</th>
				<th class="px-4 py-2 border">ステーション</th>
				<th class="px-4 py-2 border">番組</th>
				<th class="px-4 py-2 border">時間</th>
				<th class="px-4 py-2 border">操作</th>
			</tr>
		</thead>
		<tbody>
			{{range $stationDate := .Stations}}
				{{range $stationDate.Station.Progs.Progs}}
				<tr class="hover:bg-gray-50 program-row" data-title="{{.Title}}">
					<td class="px-4 py-2 border text-sm">{{formatDate .Ft}}</td>
					<td class="px-4 py-2 border">{{$stationDate.Station.Name}}</td>
					<td class="px-4 py-2 border">
						<div class="font-semibold">{{.Title}}</div>
						<div class="text-sm text-gray-600">{{stripHTML .Desc}}</div>
					</td>
					<td class="px-4 py-2 border text-sm">
						{{formatTime .Ft}} - {{formatTime .To}}
					</td>
					<td class="px-4 py-2 border text-center">
						{{if isPast .Ft}}
						<button
							hx-post="/api/record/adhoc"
							hx-vals='{"station_id": "{{$stationDate.Station.ID}}", "from": "{{.Ft}}", "area_id": "{{$stationDate.AreaID}}"}'
							hx-swap="none"
							hx-on::after-request="htmx.trigger('#task-list', 'taskCreated')"
							class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-3 rounded text-sm">
							録音
						</button>
						{{else}}
						<button
							disabled
							class="bg-gray-300 text-gray-500 font-bold py-1 px-3 rounded text-sm cursor-not-allowed"
							title="未来の番組は録音できません">
							録音不可
						</button>
						{{end}}
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
		"formatDate": func(timeStr string) string {
			// YYYYMMDDhhmmss -> MM/DD
			if len(timeStr) >= 8 {
				month := timeStr[4:6]
				day := timeStr[6:8]
				return month + "/" + day
			}
			return timeStr
		},
		"stripHTML": func(s string) string {
			return strictPolicy.Sanitize(s)
		},
		"isPast": func(timeStr string) bool {
			// YYYYMMDDhhmmss 形式をパースして過去かどうかを判定
			t, err := time.Parse("20060102150405", timeStr)
			if err != nil {
				return false
			}
			return t.Before(time.Now())
		},
	}

	t, err := template.New("program-table").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return err
	}

	type templateData struct {
		Stations []struct {
			Station radiko.Station
			DateStr string
			AreaID  string
		}
	}

	data := templateData{}
	for _, sd := range stationsWithDate {
		data.Stations = append(data.Stations, struct {
			Station radiko.Station
			DateStr string
			AreaID  string
		}{
			Station: sd.Station,
			DateStr: sd.Date.Format("01/02"),
			AreaID:  areaID,
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
<div class="space-y-3">
	{{range .Tasks}}
	<div class="bg-white border-2 {{if eq .Status "completed"}}border-green-200{{else if eq .Status "failed"}}border-red-200{{else if eq .Status "recording"}}border-blue-200{{else}}border-gray-200{{end}} rounded-lg p-4 shadow-sm">
		<div class="flex items-start justify-between mb-2">
			<div class="font-semibold text-gray-900 flex-1">
				{{if ne .ProgramTitle ""}}{{.ProgramTitle}}{{else}}{{.StationID}}{{end}}
			</div>
			{{if eq .Status "pending"}}
			<span class="inline-flex items-center px-2 py-1 text-xs font-medium text-gray-700 bg-gray-100 rounded">待機中</span>
			{{else if eq .Status "recording"}}
			<span class="inline-flex items-center px-2 py-1 text-xs font-medium text-blue-700 bg-blue-100 rounded">
				<svg class="animate-spin -ml-1 mr-2 h-3 w-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
					<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
					<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
				</svg>
				録音中
			</span>
			{{else if eq .Status "completed"}}
			<span class="inline-flex items-center px-2 py-1 text-xs font-medium text-green-700 bg-green-100 rounded">✓ 完了</span>
			{{else if eq .Status "failed"}}
			<span class="inline-flex items-center px-2 py-1 text-xs font-medium text-red-700 bg-red-100 rounded">✗ 失敗</span>
			{{end}}
		</div>
		<div class="text-sm text-gray-600">
			<div>{{.StationID}} • {{.From.Format "01/02 15:04"}}</div>
			{{if ne .Error ""}}
			<div class="text-red-600 mt-1">{{.Error}}</div>
			{{end}}
		</div>
	</div>
	{{else}}
	<div class="text-center text-gray-500 py-8 bg-gray-50 rounded-lg">
		録音タスクはありません
	</div>
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
