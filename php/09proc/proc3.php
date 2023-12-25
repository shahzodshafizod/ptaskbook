<?php
function Mean($X, $Y, &$AMean, &$GMean){
        $AMean = ($X + $Y) / 2;
        $GMean = sqrt($X * $Y);
}
$a = mt_rand(1, 50);
$b = mt_rand(1, 50); 
$c = mt_rand(1, 50);
$d = mt_rand(1, 50);
echo "a = ".$a;
echo "<br>b = ".$b;
echo "<br>c = ".$c;
echo "<br>d = ".$d;
Mean($a, $b, $arif, $geom);
$arif = number_format($arif, 2);
$geom = number_format($geom, 2);
echo "<br>Қимати миёнаи арифметики a ва b:  ".$arif;
echo "<br>Қимати миёнаи геометрии a ва b:  ".$geom."<br>";
Mean($a, $c, $arif, $geom);
$arif = number_format($arif, 2);
$geom = number_format($geom, 2);
echo "<br>Қимати миёнаи арифметики a ва c:  ".$arif;
echo "<br>Қимати миёнаи геометрии a ва c:  ".$geom."<br>";
Mean($a, $d, $arif, $geom);
$arif = number_format($arif, 2);
$geom = number_format($geom, 2);
echo "<br>Қимати миёнаи арифметики a ва d:  ".$arif;
echo "<br>Қимати миёнаи геометрии a ва d:  ".$geom;
?>