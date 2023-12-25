<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$x = mt_rand(-50, 50);
$y = mt_rand(-50, 50);
$mant = (($x * $y) > 0);
echo "Координаты точки:";
echo "<br>x = ".$x."<br>y = ".$y;
echo "<br>Принадлежность данной точки первой или третьей четверти:  ".boolval($mant);
?>