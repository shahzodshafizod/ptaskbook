<?php
function Power4($x, $a, $e){
        $adad = 1;
		$dar = 1;
		$sum = 1;
		$fact = 1;
        $i = 1;
        while(abs($adad * $dar / $fact) > $e){
                $adad *= ($adad - $i + 1);
                $dar *= $x;
                $fact *= $i;
                $sum += $adad * $dar / $fact;
                $i++;
        }
		$sum = number_format($sum, 2);
        return $sum;
}
$x = mt_rand(-10, 10);
echo "x = ".$x."<br>";
$e = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$hisob = 1;
$a = mt_rand(-50, 50);
echo "a = ".$a."<br>";
while($hisob <= 6){
	$index = mt_rand(0, 8);
	$nat = Power4($x, $a, $e[$index]);
	echo "<br>e = ".$e[$index];
	echo "(1+".$x.")^".$a." = ".$nat."<br><br>";
	$hisob++;
}
?>