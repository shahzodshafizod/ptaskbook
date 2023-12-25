<?php
$hisob = 0;
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A <= $B);
echo "A = ".$A;
echo "<br>B = ".$B;
while($A >= $B){
	$A -= $B;
	$hisob++;
}
echo "<br>Mиқдори порчаҳо = ".$hisob;
?>