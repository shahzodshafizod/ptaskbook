<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$x = mt_rand(-50, 50);
$y = mt_rand(-50, 50);
$mant = ($x < 0) && ($y != 0);
echo "Координаты точки:";
echo "<br>x = ".$x."<br>y = ".$y;
echo "<br>Принадлежность данной точки второй или третьей четверти:  ".boolval($mant);
?>