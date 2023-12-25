<?php
$N = mt_rand(2, 20);
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A >= $B);
echo "Mиқдори порчаҳо: ".$N;
echo "<br>қимати аввали порча: ".$A;
echo "<br>қимати охири порча: ".$B."<br>";
$H = ($B - $A) / $N;
echo "Дарозии ҳар як порча: H = ".$H."<br>";
for($i = 0; $i <= $N; $i++)
	echo "A + ".$i."H = ".($A + $i * $H)."<br>";
?>