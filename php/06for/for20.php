<?php
$fact = 1;
$sum = 0;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= $N; $i++){
	$fact *= $i;
	$sum += $fact;
}
echo "Cуммаи факториалҳо = ".$sum;
?>