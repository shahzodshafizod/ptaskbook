<?php
function Calc($A, $B, $Op){
	switch($Op){
		case 1:
				$nat = $A - $B;
				break;
		case 2:
				$nat = $A * $B;
				break;
		case 3:
				$nat = $A / $B;
				break;
		default:
				$nat = $A + $B;
	}
	return $nat;
}
$hisob = 1;
do{
	$A = mt_rand(-100, 100);
	$B = mt_rand(-100, 100);
}
while(($A == 0) || ($B == 0));
echo "A = ".$A;
echo "<br>B = ".$B;
while($hisob <= 3){
	$N = mt_rand(1, 10);
	$natija = Calc($A, $B, $N);
	echo "<br>Рақами амал: ".$N;
	echo "<br>Натиҷаи амал: ".$natija."<br>";
	$hisob++;
}
?>