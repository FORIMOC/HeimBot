<?php
session_start();
if (isset($_SESSION['unique_id'])){
    include_once "../php/config.php";
    $sql = mysqli_query($user_conn, "SELECT * FROM users WHERE unique_id = {$_SESSION['unique_id']}");
    if (mysqli_num_rows($sql) > 0){
        $user_row = mysqli_fetch_assoc($sql);
    }
    echo "欢迎用户".$user_row['name'];
    echo "<img src='../php/icons/".$user_row['img']."' width='50px' height='50px'>";
}else{
    echo "欢迎游客Hoshino";
    echo "<img src='images/Hoshino.jpg' width='50px' height='50px'>";
}
include_once "php/config.php";
include_once "php/filter.php";

$sql=mysqli_query($conn,"SELECT group_id,count(*) as user_num from Hoshino_users group by group_id");
for($i=0;$i<mysqli_num_rows($sql);$i++){
    $groups[$i]['group_id']=0;
    $groups[$i]['user_num']=0;
    $groups[$i]=mysqli_fetch_assoc($sql); // 群0 id: groups[0][group_id], 群0 用户数量: groups[0][user_num]
}
// 所有用户数据检索
for($i=0;$i<sizeof($groups);$i++){
    $sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE group_id='{$groups[$i]['group_id']}' ORDER BY freq DESC");
    for($j=0;$j<($groups[$i]['user_num']>5?$groups[$i]['user_num']:5);$j++){ // 循环不小于 5 次
        $users[$i][$j]=mysqli_fetch_assoc($sql); // 第 i 个群的第 j 个用户(已由 freq 字段排序)
        if (!isset($users[$i][$j])){
            $users[$i][$j]['id']=0;
            $users[$i][$j]['user_id']='暂无';
            $users[$i][$j]['username']='暂无';
            $users[$i][$j]['group_id']=$groups[$i]['group_id'];
            $users[$i][$j]['freq']=0;
        }
    }
}
// 关系网络图数据检索
for ($i=0;$i<sizeof($groups);$i++){
    for ($j=0;$j<$groups[$i]['user_num'];$j++){
        for ($k=0;$k<$groups[$i]['user_num'];$k++){
            $graph[$i][$j][$k]['value']=0;
            $graph[$i][$j][$k]['user1name']='null';
            $graph[$i][$j][$k]['user2name']='null';
            if ($j!=$k){
                $sql=mysqli_query($conn,"SELECT * FROM Hoshino_graph WHERE group_id='{$groups[$i]['group_id']}' AND user1_id='{$users[$i][$j]['user_id']}' AND user2_id='{$users[$i][$k]['user_id']}'");
                if (mysqli_num_rows($sql)>0){
                    $graph[$i][$j][$k]=mysqli_fetch_assoc($sql); // 第 i 个群 图[i][j]这条关系记录
                }
            }
        }
    }
// 用户边权总值数据检索
    for ($j=0;$j<$groups[$i]['user_num'];$j++){
        $users_value[$i][$j]=0;
        for ($k=0;$k<$groups[$i]['user_num'];$k++){
            $users_value[$i][$j]+=$graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value'];
        }
    }

// 最大边权值和最大点
    $max_node_value[$i]=0;
    $max_link_value[$i]=0;
    for ($j=0;$j<$groups[$i]['user_num'];$j++){
            $max_node_value[$i]=max($max_node_value[$i], $users_value[$i][$j]);
    }
    for ($j=0;$j<sizeof($graph[$i]);$j++){
        for ($k=0;$k<sizeof($graph[$i]);$k++){
            if (isset($graph[$i][$j][$k])){
                $max_link_value[$i]=max($max_link_value[$i],$graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']);
            }
        }
    }
}
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hoshino | Helmboard</title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="icon" href="images/Hoshino.jpg">
    <script src="files/echarts.js"></script>
    <link rel="stylesheet" href="files/index.css">
</head>
<body>
<?php
for ($i=0;$i<sizeof($groups);$i++){
    // rela_chart 部分
    $rela_chart="
<div id='rela_chart".$i."' style='width: 600px; height: 400px'></div>
<script>
var myChart = echarts.init(document.getElementById('rela_chart".$i."'));
    var categories = [];
    for (var i = 0; i <= 3; i++) {
        categories[i] = {
            name: '类别' + i
        };
    }
    option = {
        // 图的标题
        title: {
            text: '群".$groups[$i]['group_id']."用户社交网络',
            subtext: '(最大点: ".$max_node_value[$i]." | 最大边: ".$max_link_value[$i].")',
            left: 'center'
        },
        // 提示框的配置
        tooltip: {
            formatter: function (x) {
                return x.data.des;
            }
        },
        // 工具箱
        toolbox: {
            // 显示工具箱
            show: true,
            feature: {
                mark: {
                    show: true
                },
                // 还原
                restore: {
                    show: true
                },
                // 保存为图片
                saveAsImage: {
                    show: true
                }
            }
        },
        legend: [{
            // selectedMode: 'single',
            data: categories.map(function (a) {
                return a.name;
            }),
            orient: 'vertical',
            left: 'left'
        }],
        series: [{
            type: 'graph', // 类型:关系图
            layout: 'force', //图的布局，类型为力导图
            symbolSize: 20, // 调整节点的大小
            roam: true, // 是否开启鼠标缩放和平移漫游。默认不开启。如果只想要开启缩放或者平移,可以设置成 'scale' 或者 'move'。设置成 true 为都开启
            edgeSymbolSize: [2, 10],
            emphasis: {
                scale: true,
                focus: 'adjacency',
            },
            edgeLabel: {
                show: true,
                textStyle: {
                    fontSize: 15
                },
                formatter: function (x) {
                    return x.data.value;
                }
            },
            force: {
                repulsion: 5000,
                edgeLength: [30, 80]
            },
            draggable: true,
            lineStyle: {
                width: 2
            },
            label: {
                show: true,
                position: 'top'
            },
            data: [";
    for ($j=0;$j<$groups[$i]['user_num'];$j++){
        $rela_chart.="
            {
                name: '".$users[$i][$j]['username']."',
                des: '".$users[$i][$j]['username']."[".$users_value[$i][$j]."](".$users[$i][$j]['user_id'].")',";
        if ($users_value[$i][$j]<=floor($max_node_value[$i]*0.25)){
            $rela_chart.="category: 3,
            symbolSize: 30
            },";
        }elseif ($users_value[$i][$j]<=floor($max_node_value[$i]*0.5)){
            $rela_chart.="category: 2,
            symbolSize: 40
            },";
        }elseif ($users_value[$i][$j]<=floor($max_node_value[$i]*0.75)){
            $rela_chart.="category: 1,
            symbolSize: 50
            },";
        }else{
            $rela_chart.="category: 0,
            symbolSize: 60
            },";
        }
    }

    $rela_chart.="],
            links: [";
    for ($j=0;$j<sizeof($graph[$i]);$j++){
        for ($k=0;$k<sizeof($graph[$i]);$k++){
            if ($graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']!=0){
                if ($graph[$i][$j][$k]['user1name']!="null"){
                    $rela_chart.="{
                name: '".$graph[$i][$j][$k]['user1name']."-".$graph[$i][$j][$k]['user2name']."',
                source: '".$users[$i][$j]['username']."',
                target: '".$users[$i][$k]['username']."',
                des: '".$graph[$i][$j][$k]['user1name'].": ".$graph[$i][$j][$k]['value']."->|<-".$graph[$i][$k][$j]['value']." :".$graph[$i][$j][$k]['user2name']."',
                value: '".$graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']."',";
                    if ($graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']<=floor($max_link_value[$i]*0.25)){
                        $rela_chart.="lineStyle: {
                    width: 2
                }
            },";
                    }elseif ($graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']<=floor($max_link_value[$i]*0.5)){
                        $rela_chart.="lineStyle: {
                    width: 5
                }
            },";
                    }elseif ($graph[$i][$j][$k]['value']+$graph[$i][$k][$j]['value']<=floor($max_link_value[$i]*0.75)){
                        $rela_chart.="lineStyle: {
                    width: 8
                }
            },";
                    }else{
                        $rela_chart.="lineStyle: {
                    width: 12
                }
            },";
                    }
                }
            }
        }
    }
    $rela_chart.="],
            categories: categories
        }]
    };
    myChart.setOption(option);
</script>
    ";

    // freq_chart 部分
    $freq_chart="
<div id='freq_chart".$i."' style='width: 600px; height: 400px'></div>
<script>
var chart=echarts.init(document.getElementById('freq_chart".$i."'));
var option = {
    title: {
        text: '群".$groups[$i]['group_id']."前5活跃用户',
        subtext: '(共".$groups[$i]['user_num']."用户)',
        left: 'center'
    },
    tooltip: {
        trigger: 'item',
    },
    toolbox: {
            show: true,
            feature: {
                mark: {
                    show: true
                },
                restore: {
                    show: true
                },
                saveAsImage: {
                    show: true
                }
            }
        },
    legend:{
        orient: 'vertical',
        left: 'left'
    },
    series:[
        {
            name: '发言频率',
            type: 'pie',
            radius: '50%',
            data:[
                { value: ".$users[$i][0]['freq'].", name: '".$users[$i][0]['username']."(".$users[$i][0]['user_id'].")' },
                { value: ".$users[$i][1]['freq'].", name: '".$users[$i][1]['username']."(".$users[$i][1]['user_id'].")' },
                { value: ".$users[$i][2]['freq'].", name: '".$users[$i][2]['username']."(".$users[$i][2]['user_id'].")' },
                { value: ".$users[$i][3]['freq'].", name: '".$users[$i][3]['username']."(".$users[$i][3]['user_id'].")' },
                { value: ".$users[$i][4]['freq'].", name: '".$users[$i][4]['username']."(".$users[$i][4]['user_id'].")' }
            ],
            emphasis: {
                itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0,0,0,0.5)'
                }
            }
        }
    ]
}
chart.setOption(option)
</script>";

    // last_words 部分
    $sql=mysqli_query($conn,"SELECT * FROM Hoshino_premessages WHERE group_id='{$groups[$i]['group_id']}'");
    if (mysqli_num_rows($sql)>0){
        for ($j=0;$j<5;$j++){
            $premessages[$j]=mysqli_fetch_assoc($sql);
        }
    }else{
        for ($j=0;$j<5;$j++){
            $premessages[$j]['content']='暂无消息';
            $premessages[$j]['user_id']='暂无';
            $premessages[$j]['username']='暂无';
        }
    }
    $last_words="<div class='last_words'>";
    for ($j=0;$j<5;$j++){
        $last_words.="<p><span class='word_name'>".$premessages[$j]['username']."(".$premessages[$j]['user_id'].")</span> <span class='output_mark'>> </span>".filter($premessages[$j]['content'])."</p>";
    }
    $last_words.="</div>";

    // 输出div
    echo "<div class='container'>";
    echo $rela_chart;
    echo $freq_chart;
    echo "</div>";
    echo "<div class='container'>";
    echo $last_words;
    echo "</div>";
}
?>
</body>
</html>