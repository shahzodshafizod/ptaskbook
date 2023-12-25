<?php
$F1 = $F2 = 1;
$N = mt_rand(2, 20);
echo "Mиқдори ададҳо пайдарпайии Фибоначчи: ".$N."<br><br>";
echo "F1 = 1<br>F2 = 1<br>";
for($i = 3; $i <= $N; $i++){
	$Fk = $F1 + $F2;
	$F1 = $F2;
	$F2 = $Fk;
	echo "F".$i." = ".$Fk."<br>";
}
?>