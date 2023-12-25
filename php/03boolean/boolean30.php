<?php
function boolval($nat){
	return ($nat? "true": "false");
}
do{
	$a = mt_rand(1, 50);
	$b = mt_rand(1, 50);
	$c = mt_rand(1, 50);
}
while(($a+$b <= $c) || ($a+$c <= $b) || ($b+$c <= $a));
echo "Длина первой стороны: ".$a;
echo "<br>Длина второй стороны: ".$b;
echo "<br>Длина третьей стороны: ".$c;
$mant = ($a == $b) && ($b == $c);
echo "<br>Равносторонность треугольника:  ".boolval($mant);
?>