<?php
$sum = 0;
$N = mt_rand(1, 50);
$K = mt_rand(1, 50);
echo "Mиқдори ададҳо: ".$N;
echo "<br>дараҷа: ".$K;
for($i = 1; $i <= $N; $i++){
	$dar = 1;
	for($j = 1; $j <= $K; $j++)
			$dar *= $i;
	$sum += $dar;
}
echo "<br>Сумма = ".$sum;
?>