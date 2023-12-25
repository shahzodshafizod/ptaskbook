<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
$c = mt_rand(-100, 100);
if($a > $b)
	if($b > $c)
			$mid = $b;
	else
			if($a > $c)
					$mid = $c;
			else
					$mid = $a;
else if($b > $a)
	if($a > $c)
			$mid = $a;
	else
			if($b > $c)
					$mid = $c;
			else
					$mid = $b;
else if($a > $c)
	if($c > $b)
			$mid = $c;
	else
			if($a > $b)
					$mid = $b;
			else
					$mid = $a;
echo "Aдади якӯм: ".$a;
echo "<br>Aдади дуюм: ".$b;
echo "<br>Aдади сеюм: ".$c;
echo "<br>Адади мобайнӣ: ".$mid;
?>