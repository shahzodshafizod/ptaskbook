<?php
$A1 = 1;
$A2 = 2;
$A3 = 3;
$N = mt_rand(3, 23);
echo "Mиқдори ададҳои пайдарпайӣ: ".$N."<br>";
echo "A1 = 1<br>A2 = 2<br>A3 = 3<br>";
for($i = 4; $i <= $N; $i++){
	$Ak = $A3 + $A2 - 2 * $A1;
	$A1 = $A2;
	$A2 = $A3;
	$A3 = $Ak;
	echo "A".$i." = ".$Ak."<br>";
}
?>