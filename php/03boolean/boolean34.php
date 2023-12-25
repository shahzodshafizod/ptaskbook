<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$x = mt_rand(1, 8);
$y = mt_rand(1, 8);
$mant = (($x + $y) % 2 == 1);
echo "Координаты поля шахматной доски:";
echo "<br>x = ".$x."<br>y = ".$y;
echo "<br>Белоцветность данного поля:  ".boolval($mant);
?>