<?php
function IncTime(&$H, &$M, &$S, $T){
        $H += intval($T / 3600);
        $M += intval($T % 3600 / 60);
        $S += $T % 3600 % 60;
}
$soat = mt_rand(1, 24);
$daq = mt_rand(1, 60);
$son = mt_rand(1, 60);
$T = mt_rand();
echo "Вақти аввала:";
echo "<br>Cоат: ".$soat;
echo "<br>Дақиқа: ".$daq;
echo "<br>Сония: ".$son;
echo "<br>Вақти додашуда бо сонияҳо: ".$T;
IncTime($soat, $daq, $son, $T);
echo "<br><br>Вақти зиёдшуда ба миқдори ".$T." сония:";
echo "<br>Cоат: ".$soat;
echo "<br>Дақиқа: ".$daq;
echo "<br>Сония: ".$son;
?>