<?php
function KTU2($A, $B){
        while(($A != 0) && ($B != 0))
                if($A > $B)
                        $A %= $B;
                else
                        $B %= $A;
        return ($A + $B);
}
function Frac1($a, $b, &$p, &$q){
        $ktu = KTU2($a, $b);
        $p = intval($a / $ktu);
        $q = intval($b / $ktu);
}
$a = mt_rand(1, 100);
$b = mt_rand(1, 100);
$c = mt_rand(1, 100);
$d = mt_rand(1, 100);
echo "a = ".$a;
echo "<br>b = ".$b;
echo "<br>c = ".$c;
echo "<br>d = ".$d;
Frac1($a, $b, $p, $q);
$nat1 = $p / $q;
echo "<br>a / b = ".$p."/".$q;
Frac1($c, $d, $p, $q);
$nat2 = $p / $q;
echo "<br>c / d = ".$p."/".$q;
$nat = $nat1 + $nat2;
$nat = number_format($nat, 2);
echo "<br>a / b + c / d = ".$nat;

$e = mt_rand(1, 100);
$f = mt_rand(1, 100);
$g = mt_rand(1, 100);
$h = mt_rand(1, 100);
echo "<br><br>e = ".$e;
echo "<br>f = ".$f;
echo "<br>g = ".$g;
echo "<br>h = ".$h;
Frac1($e, $f, $p, $q);
$nat1 = $p / $q;
echo "<br>e / f = ".$p."/".$q;
Frac1($g, $h, $p, $q);
$nat2 = $p / $q;
echo "<br>g / h = ".$p."/".$q;
$nat = $nat1 + $nat2;
$nat = number_format($nat, 2);
echo "<br>e / f + g / h = ".$nat;
?>