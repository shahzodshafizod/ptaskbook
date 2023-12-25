<?php
$narx = mt_rand(1, 100) / 4;
$narx = number_format($narx, 2);
echo "Нарх: ".$narx."<br><br>";
for($i = 1.2; $i < 2.2; $i += 0.2)
	echo $i." кг = ".($i * $narx)." сомонӣ<br>";
?>