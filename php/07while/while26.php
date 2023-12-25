<?php
$F1 = $F2 = 1;
do{
	$N = mt_rand(2, 100);
	$Fk = $F1 + $F2;
	$F1 = $F2;
	$F2 = $Fk;
	while($N > $Fk){
		$Fk = $F1 + $F2;
		$F1 = $F2;
		$F2 = $Fk;
	}
}
while($N != $Fk);
echo "N = ".$N;
echo "<br>F[k-1] = ".($Fk - ($F2 - $F1));
echo "<br>\nFk = ".$Fk;
echo "<br>\nF[k+1] = ".($Fk + $F1);
?>