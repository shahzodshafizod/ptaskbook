<?php
$t = false;
$du = false;
do{
	do{	
		$N = mt_rand(1, 1000);
		for($k = 2; $k <= $N; $k++){
			if(pow(2, $k) == $N)
				$du = true;
		}
	}
	while(!$du);
	echo "Дараҷаи адади 2 = ".$N;
	for($i = 0; $i <= intval($N/2); $i++)
		if(pow(2, $i) == $N){
			$t = true;
			$dar = $i;
		}
}
while(!$t);
echo "<br>K = ".$dar;
?>