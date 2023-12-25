<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
do{
	$A = mt_rand(-50, 50);
}
while($A == 0);
$B = mt_rand(-50, 50);
$C = mt_rand(-50, 50);
$D = $B * $B - 4 * $A * $C;
$mant = ($D >= 0);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>Дискриминант: ".$D;
echo "<br>У уравнения вещественные корни:  ".boolval($mant);
?>