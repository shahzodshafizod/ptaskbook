<?php
$mass = array('N', 'S', 'E', 'W');
$n = rand(0, 3);
$far = array(1, -1, 2);
$f1 = rand(0, 2);
$f2 = rand(0, 2);
$qimat = abs($far[$f1] + $far[$f2]);
$qimat = abs($far[$f1] + $far[$f2]);
switch($mass[$n]){
       case 'N':
            switch($qimat){
                   case 0: case 4:
                        $har = 'N';
                        break;
                   case 1:
                        $har = 'W';
                        break;
                   case 2:
                        $har = 'S';
                        break;
                   case 3:
                        $har = 'E';
            }
            break;
       case 'S':
            switch($qimat){
                   case 0: case 4:
                        $har = 'S';
                        break;
                   case 1:
                        $har = 'E';
                        berak;
                   case 2:
                        $har = 'W';
                        break;
                   case 3:
                        $har = 'N';
            }
            break;
       case 'E':
            switch($qimat){
                   case 0: case 4:
                        $har = 'E';
                        break;
                   case 1:
                        $har = 'N';
                        break;
                   case 2:
                        $har = 'S';
                        break;
                   case 3:
                        $har = 'W';
            }
            break;
       case 'W':
            switch($qimat){
                   case 0: case 4:
                        $har = 'W';
                        break;
                   case 1:
                        $har = 'S';
                        break;
                   case 2:
                        $har = 'E';
                        break;
                   case 3:
                        $har = 'N';
            }
}
echo "Исходное направление локатора - $mass[$n].<br>";
echo "Первая команда: $far[$f1]<br>";
echo "Вторая команда: $far[$f2]<br>";
echo "Направление локатора после выполнения команд - $har.";
?>

/*
<?php
$mass = array('N', 'S', 'E', 'W');
$n = rand(0, 3);
$far = array(1, -1, 2);
$f1 = rand(0, 2);
$f2 = rand(0, 2);
$qimat = abs($far[$f1] + $far[$f2]);
$qimat = abs($far[$f1] + $far[$f2]);
switch($mass[$n]){
       case 'N':
            switch($qimat){
                   case 0: case 4:
                        $har = 'N';
                        break;
                   case 1:
                        $har = 'W';
                        break;
                   case 2:
                        $har = 'S';
                        break;
                   case 3:
                        $har = 'E';
            }
            break;
       case 'S':
            switch($qimat){
                   case 0: case 4:
                        $har = 'S';
                        break;
                   case 1:
                        $har = 'E';
                        berak;
                   case 2:
                        $har = 'W';
                        break;
                   case 3:
                        $har = 'N';
            }
            break;
       case 'E':
            switch($qimat){
                   case 0: case 4:
                        $har = 'E';
                        break;
                   case 1:
                        $har = 'N';
                        break;
                   case 2:
                        $har = 'S';
                        break;
                   case 3:
                        $har = 'W';
            }
            break;
       case 'W':
            switch($qimat){
                   case 0: case 4:
                        $har = 'W';
                        break;
                   case 1:
                        $har = 'S';
                        break;
                   case 2:
                        $har = 'E';
                        break;
                   case 3:
                        $har = 'N';
            }
}
echo "Самти аввалаи локатор $mass[$n] аст.<br>";
echo "Фармони якум: $far[$f1] аст<br>";
echo "Фармони дуюм: $far[$f2] аст<br>";
echo "Самти локатор пас ичрои фармонхо $har аст.<br>";
?>
*/