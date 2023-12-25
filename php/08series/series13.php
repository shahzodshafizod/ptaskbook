<?php
$hisob = 0;
$sum = 0;
do{
	$N = mt_rand(-100, 100);
	if($N != 0){
		echo "N".($hisob + 1)." = ".$N."<br>";
		if(($N > 0) && (($N % 2) == 0))
			$sum += $N;
		$hisob++;
	}
}
while($N != 0);
echo "Суммаи ададҳои ҷуфти мусбӣ = ".$sum;
?>