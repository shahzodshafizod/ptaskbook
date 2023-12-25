<?php
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
	$C = mt_rand(1, 100);
}
while(($C >= $A) || ($C >= $B));
$C_dar_A = intval($A / $C);
$C_dar_B = intval($B / $C);
$kvad = $C_dar_A * $C_dar_B;
$ozod = $A * $B - $kvad * $C * $C;
echo "Тарафи якӯми росткунҷа: ".$A;
echo "<br>Тарафи дуюми росткунҷа: ".$B;
echo "<br>Тарафи квадрат: ".$C;
echo "<br>Миқдори квадратҳо: ".$kvad;
echo "<br>Масоҳати бандНАбуда: ".$ozod;
?>