<?php
$K = mt_rand(1, 365);
$hafta = ($K + 3) % 7;
echo "Рӯзи сол: ".$K;
echo "<br>Рақами рӯзи ҳафта: ".$hafta;
?>