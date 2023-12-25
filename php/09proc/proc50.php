<?php
function TimeToHMS($T, &$H, &$M, &$S){
        $H = intval($T / 3600);
        $M = intval($T % 3600 / 60);
        $S = $T % 3600 % 60;
}
$hisob = 1;
while($hisob <= 5){
	$T = mt_rand();
	TimeToHMS($T, $soat, $daq, $son);
	echo "Вақт: ".$T;
	echo "<br>Соат: ".$soat;
	echo "<br>Дақиқа: ".$daq;
	echo "<br>Сония: ".$son."<br><br>";
	$hisob++;
}
?>