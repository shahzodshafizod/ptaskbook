<?php
$i = $j = 0;
$A = mt_rand(1, 100);
$B = mt_rand(1, 100);
do{
	$C = mt_rand(1, 100);
}
while(($C >= $A) || ($C >= $B));
echo "Tарафҳои росткунҷа:";
echo "<br>A = ".$A;
echo "<br>B = ".$B;
echo "<br>Tарафи квадрат:";
echo "<br>C = ".$C;
while($A >= $C){
	$A -= $C;
	$i++;
}
while($B >= $C){
	$B -= $C;
	$j++;
}
echo "<br>Mиқдори квадратҳо: ".($i * $j);
?>