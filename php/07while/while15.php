<?php
$S = 1000;
$K = 0;
$P = mt_rand(1, 24);
echo "Фоиз: ".$P;
while($S <= 1100){
	$S += $S * $P / 100;
	$K++;
}
echo "<br>Миқдори моҳҳо = ".$K;
echo "<br>пасандоз [>1100] = ".number_format($S, 2);
?>