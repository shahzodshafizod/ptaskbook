<?php
$N = mt_rand(2, 20);
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A >= $B);
echo "Mиқдори порчаҳо: ".$M;
echo "<br>қимати аввалини порча: ".$A;
echo "<br>қимати охирини порча: ".$B."<br>";
$H = number_format(($B - $A) / $N, 2); 
echo "Дарозии ҳар як порча = ".$H."<br>";
for($i = 0; $i <= $N; $i++){
	$Fx = 1 - sin($A + $i * $H);
	$Fx = number_format($Fx, 2);
	echo "F(".($A + $i * $H).") = ".$Fx."<br>";
}
?>