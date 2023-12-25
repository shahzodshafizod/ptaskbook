<?php
$A1 = 1;
$A2 = 2;
$N = mt_rand(2, 20);
echo "Mиқдори ададҳои пайдарпайӣ: ".$N."<br><br>";
echo "A1 = 1<br>A2 = 2<br>";
for($i = 3; $i <= $N; $i++){
	$Ak = ($A1 + 2 * $A2) / 3;
	$A1 = $A2;
	$A2 = $Ak;
	echo "A".$i." = ".number_format($Ak, 2)."<br>";
}
?>