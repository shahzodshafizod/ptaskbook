<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$x1 = mt_rand(-50, 50);
$y1 = mt_rand(-50, 50);
$x2 = mt_rand(-50, 50);
$y2 = mt_rand(-50, 50);
$x = mt_rand(-50, 50);
$y = mt_rand(-50, 50);
$mant = ($x1 < $x) && ($x < $x2) && ($y1 > $y) && ($y > $y2);
echo "Координаты левой верхней вершины:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координаты правой нижней вершины:";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
echo "<br>Координаты данной точки:";
echo "<br>x = ".$x."<br>y = ".$y;
echo "<br>Принадлежность данной точки прямоугольнику:  ".boolval($mant);
?>