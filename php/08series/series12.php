<?php
$hisob = 0;
do{
	$A = mt_rand(-100, 100);
	if($A != 0){
		echo "A".($hisob + 1)." = ".$A."<br>";
		$hisob++;
	}
}
while($A != 0);
echo "Миқдори ададҳои маҷмӯъ = ".($hisob - 1);
?>