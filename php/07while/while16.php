<?php
$S = $nat = 10;
$K = 0;
$P = mt_rand(1, 49);
echo "Фоиз: ".$P;
while($nat <= 200){
	$S += $S * $P / 100;
	$nat += $S;
	$K++;
}
echo "<br>Миқдори моҳҳо = ".$K;
echo "<br>Mасофа = ".number_format($nat, 3);
?>