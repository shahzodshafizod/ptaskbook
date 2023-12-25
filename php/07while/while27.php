<?php
do{
	$N = mt_rand(2, 100);
	$F1 = 1;
	$F2 = 1;
	$hisob = 3;
	$Fk = $F1 + $F2;
	$F1 = $F2;
	$F2 = $Fk;
	while($N > $Fk){
		$Fk = $F1 + $F2;
		$F1 = $F2;
		$F2 = $Fk;
		$hisob++;
	}
}
while($N != $Fk);
echo "N = ".$N;
echo "<br>Миқдори ададҳои Фибоначчи = ".$hisob;
?>