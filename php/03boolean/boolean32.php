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
$mant = ($a*$a == $b*$b+$c*$c) && ($b*$b == $a*$a+$c*$c) && ($c*$c == $a*$a+$b*$b);
echo "Длина первой стороны: ".$a;
echo "<br>Длина второй стороны: ".$b;
echo "<br>Длина третьей стороны: ".$c;
echo "<br>Прямоугольность треугольника:  ".boolval($mant);
?>