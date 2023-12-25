<?php
$sh_kg = mt_rand(1, 10);
$sh_narx = mt_rand(10, 100);
$ir_kg = mt_rand(1, 10);
$ir_narx = mt_rand(5, 50);
$sh_narx /= $sh_kg;
$ir_narx /= $ir_kg;
$garon = $sh_narx / $ir_narx;
$sh_narx = number_format($sh_narx, 2);
$ir_narx = number_format($ir_narx, 2);
$garon = number_format($garon, 2);
echo "Шоколад:<br>1 кг = ".$sh_narx." сомони<br>";
echo "Ирис:<br>1 кг = ".$ir_narx." сомони<br>";
echo "Шоколад дороже ирисок в ".$garon." раз.<br>";
?>

/*
<?php
$sh_kg = mt_rand(1, 10);
$sh_narx = mt_rand(10, 100);
$ir_kg = mt_rand(1, 10);
$ir_narx = mt_rand(5, 50);
$sh_narx /= $sh_kg;
$ir_narx /= $ir_kg;
$garon = $sh_narx / $ir_narx;
$sh_narx = number_format($sh_narx, 2);
$ir_narx = number_format($ir_narx, 2);
$garon = number_format($garon, 2);
echo "Шоколад:<br>1 кг = ".$sh_narx." сомонӣ аст<br>";
echo "Ирис:<br>1 кг = ".$ir_narx." сомонӣ аст<br>";
echo "Шоколад аз ирис ".$garon." маротиба қимматтар аст.<br>";
?>
*/