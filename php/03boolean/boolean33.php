<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$a = mt_rand(1, 50);
$b = mt_rand(1, 50);
$c = mt_rand(1, 50);
$mant = ($a+$b > $c) && ($a+$c > $b) && ($b+$c > $a);
echo "Длина первой стороны: ".$a;
echo "<br>Длина второй стороны: ".$b;
echo "<br>Длина третьей стороны: ".$c;
echo "<br>Существование треугольника:  ".boolval($mant);
?>