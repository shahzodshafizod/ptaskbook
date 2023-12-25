<?php
$K = mt_rand(1, 365);
$N = mt_rand(1, 7);
$hafta = ($K + $N - 2) % 7 + 1;
echo "Рӯзи сол: ".$K;
echo "<br>Рақами рӯзи ҳафтаи аввали сол: ".$N;
echo "<br>Рақами рӯзи ҳафта: ".$hafta;
?>