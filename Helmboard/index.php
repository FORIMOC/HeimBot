<?php
//session_start();
//if (!isset($_SESSION['unique_id'])){
//    header("location: ../login.php");
//}
//include_once "../php/config.php";
//$sql = mysqli_query($user_conn, "SELECT * FROM users WHERE unique_id = {$_SESSION['unique_id']}");
//if (mysqli_num_rows($sql) > 0){
//    $user_row = mysqli_fetch_assoc($sql);
//}
//echo "欢迎用户".$user_row['name'];
//echo "<img src='../php/icons/".$user_row['img']."' width='50px' height='50px'>";
include_once "php/config.php";
$sql=mysqli_query($conn,"SELECT group_id,count(*) as user_num from Hoshino_users group by group_id");
for($i=0;$i<mysqli_num_rows($sql);$i++){
    $groups[$i]=mysqli_fetch_assoc($sql); // 群0 id: groups[0][group_id], 群0 用户数量: groups[0][user_num]
}
for($i=0;$i<sizeof($groups);$i++){
    $sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE group_id='{$groups[$i]['group_id']}' ORDER BY freq DESC");
    for($j=0;$j<5;$j++){
        $top5users[$i][$j]=mysqli_fetch_assoc($sql);
        if (!isset($top5users[$i][$j])){
            $top5users[$i][$j]['id']=0;
            $top5users[$i][$j]['user_id']='暂无';
            $top5users[$i][$j]['username']='暂无';
            $top5users[$i][$j]['group_id']=$groups[$i]['group_id'];
            $top5users[$i][$j]['freq']=0;
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
    $chart="
<div id='chart".$i."' style='width: 600px; height: 400px'></div>
<script>
var chart=echarts.init(document.getElementById('chart".$i."'));

var option = {
    title: {
        text: '群".$groups[$i]['group_id']."前5活跃用户',
        subtext: '(共".$groups[$i]['user_num']."用户)',
        left: 'center'
    },
    tooltip: {
        trigger: 'item'
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
                { value: ".$top5users[$i][0]['freq'].", name: '".$top5users[$i][0]['username']."(".$top5users[$i][0]['user_id'].")' },
                { value: ".$top5users[$i][1]['freq'].", name: '".$top5users[$i][1]['username']."(".$top5users[$i][1]['user_id'].")' },
                { value: ".$top5users[$i][2]['freq'].", name: '".$top5users[$i][2]['username']."(".$top5users[$i][2]['user_id'].")' },
                { value: ".$top5users[$i][3]['freq'].", name: '".$top5users[$i][3]['username']."(".$top5users[$i][3]['user_id'].")' },
                { value: ".$top5users[$i][4]['freq'].", name: '".$top5users[$i][4]['username']."(".$top5users[$i][4]['user_id'].")' }
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
    echo $chart;
}
?>
</body>
</html>