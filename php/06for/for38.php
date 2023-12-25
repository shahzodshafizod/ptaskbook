<?php
$sum = 0;
$N = mt_rand(1, 50);
echo "Mиқдори ададҳо: ".$N."<br>";
for($i = 1; $i <= $N; $i++){
	$dar = 1;
	for($j = $N - $i + 1; $j >= 1; $j--)
			$dar *= $i;
	$sum += $dar;
}
echo "Сумма = ".$sum;
?>