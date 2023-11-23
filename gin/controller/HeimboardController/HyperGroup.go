package HeimboardController

import (
	"fmt"
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelGroup"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type HyperGroupEdge struct {
	GroupID1   string
	GroupName1 string
	GroupID2   string
	GroupName2 string
	Value1     float64
	Value2     float64
	Affinity   float64
	Gcd        int
}

// CalculateHyperGroup 计算超群数据控制器
// selected_groups(数组) => hyper_group_points([]ModelGroup.Group), hyper_group_edges([]HyperGroupEdge), max_group(int), min_group(int), max_edge(int), min_edge(int), max_gcd(int), min_gcd(int)
func CalculateHyperGroup(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	selectedGroupStr := ctx.PostForm("selected_groups")
	selectedGroups := strings.Split(selectedGroupStr, ",")
	if len(selectedGroups) <= 1 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "选择的群聊个数不能低于2个")
		return
	}
	var groups []ModelGroup.Group
	var hyperGroupEdges []HyperGroupEdge
	maxGroup := 0
	minGroup := 3000
	maxEdge := 0.0
	minEdge := 3000.0
	maxGcd := 0
	minGcd := 3000

	for i := 0; i < len(selectedGroups); i++ {
		var group ModelGroup.Group
		DB.Where("group_id = ?", selectedGroups[i]).First(&group)
		if group.MemberNum > maxGroup {
			maxGroup = group.MemberNum
		}
		if group.MemberNum < minGroup {
			minGroup = group.MemberNum
		}
		groups = append(groups, group)
	}

	for i := 0; i < len(selectedGroups); i++ {
		for j := 0; j < len(selectedGroups) && j < i; j++ {
			// 发出 post 请求
			resp, err := http.Get("http://121.5.167.91/util/gcd.php?group_id1=" + selectedGroups[i] + "&group_id2=" + selectedGroups[j])
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			defer resp.Body.Close()
			content, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
			values := strings.Split(*str, ",")
			value1, _ := strconv.ParseFloat(values[0], 64)
			value2, _ := strconv.ParseFloat(values[1], 64)
			gcd, _ := strconv.Atoi(values[2])
			if gcd > maxGcd {
				maxGcd = gcd
			}
			if gcd < minGcd {
				minGcd = gcd
			}
			affinity := float64(gcd) * value1 * value2
			if affinity > maxEdge {
				maxEdge = affinity
			}
			if affinity < minEdge {
				minEdge = affinity
			}
			var group1 ModelGroup.Group
			var group2 ModelGroup.Group
			DB.Where("group_id = ?", selectedGroups[i]).First(&group1)
			DB.Where("group_id = ?", selectedGroups[j]).First(&group2)
			newEdge := HyperGroupEdge{
				GroupID1:   selectedGroups[i],
				GroupName1: group1.GroupName,
				GroupID2:   selectedGroups[j],
				GroupName2: group2.GroupName,
				Value1:     value1,
				Value2:     value2,
				Affinity:   affinity,
				Gcd:        gcd,
			}
			hyperGroupEdges = append(hyperGroupEdges, newEdge)
		}
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{
		"hyper_group_points": groups,
		"hyper_group_edges":  hyperGroupEdges,
		"max_group":          maxGroup,
		"min_group":          minGroup,
		"max_edge":           maxEdge,
		"min_edge":           minEdge,
		"max_gcd":            maxGcd,
		"min_gcd":            minGcd,
	}, "获取超群计算结果成功")
}
