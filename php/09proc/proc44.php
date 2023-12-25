<?php
function Arctg1($x, $e){
        $dar = $x;
		$sum = $x;
        $i = 1;
		$alomat = 1;
        while(abs($dar / $i) <= $e){
                $dar *= $x * $x;
                $i += 2;
                $alomat *= -1;
                $sum += $alomat * $dar / $i;
        }
		$sum = number_format($sum, 2);
        return $sum;
}
$hisob = 1;
$x = mt_rand(-10, 10);
echo "x = ".$x."<br>";
$e = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
while($hisob <= 6){
	$index = mt_rand(0, 8);
	$nat = Arctg1($x, $e[$index]);
	echo "<br>e = ".$e[$index];
	echo "<br>Arctg1(".$x.") = ".$nat."<br><br>";
	$hisob++;
}
?>