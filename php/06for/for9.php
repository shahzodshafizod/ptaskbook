<?php
$kv_sum = 0;
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}

while($A >= $B);
echo "A = ".$A;
echo "<br>B = ".$B."<br>";
for($i = $A; $i <= $B; $i++)
	$kv_sum += $i * $i;
echo "Суммаи квадратҳои ададҳо аз A то B: ".$kv_sum;
?>